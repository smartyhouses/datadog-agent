// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build kubeapiserver

package provider

import (
	"context"
	"fmt"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes/scheme"
	v1core "k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/tools/record"

	workloadmeta "github.com/DataDog/datadog-agent/comp/core/workloadmeta/def"
	"github.com/DataDog/datadog-agent/pkg/aggregator/sender"
	"github.com/DataDog/datadog-agent/pkg/clusteragent/autoscaling"
	"github.com/DataDog/datadog-agent/pkg/clusteragent/autoscaling/workload"
	"github.com/DataDog/datadog-agent/pkg/clusteragent/autoscaling/workload/local"
	"github.com/DataDog/datadog-agent/pkg/clusteragent/autoscaling/workload/model"
	"github.com/DataDog/datadog-agent/pkg/util/kubernetes/apiserver"
	"github.com/DataDog/datadog-agent/pkg/util/kubernetes/apiserver/leaderelection"
)

const maxDatadogPodAutoscalerObjects int = 100

// StartWorkloadAutoscaling starts the workload autoscaling controller
func StartWorkloadAutoscaling(
	ctx context.Context,
	clusterID string,
	apiCl *apiserver.APIClient,
	rcClient workload.RcClient,
	wlm workloadmeta.Component,
	senderManager sender.SenderManager,
) (workload.PodPatcher, error) {
	if apiCl == nil {
		return nil, fmt.Errorf("Impossible to start workload autoscaling without valid APIClient")
	}

	le, err := leaderelection.GetLeaderEngine()
	if err != nil {
		return nil, fmt.Errorf("Unable to start workload autoscaling as LeaderElection failed with: %v", err)
	}

	eventBroadcaster := record.NewBroadcaster()
	eventBroadcaster.StartRecordingToSink(&v1core.EventSinkImpl{Interface: apiCl.Cl.CoreV1().Events("")})
	eventRecorder := eventBroadcaster.NewRecorder(scheme.Scheme, corev1.EventSource{Component: "datadog-workload-autoscaler"})

	store := autoscaling.NewStore[model.PodAutoscalerInternal]()
	podPatcher := workload.NewPodPatcher(store, le.IsLeader, apiCl.DynamicCl, eventRecorder)
	podWatcher := workload.NewPodWatcher(wlm, podPatcher)
	localRecommender := local.NewRecommender(podWatcher, store)

	_, err = workload.NewConfigRetriever(store, le.IsLeader, rcClient)
	if err != nil {
		return nil, fmt.Errorf("Unable to start workload autoscaling config retriever: %w", err)
	}

	// We could consider the sender to be optional, but it's actually required for backend information
	sender, err := senderManager.GetSender("workload_autoscaling")
	sender.DisableDefaultHostname(true)
	if err != nil {
		return nil, fmt.Errorf("Unable to start local telemetry for autoscaling: %w", err)
	}

	limitHeap := autoscaling.NewHashHeap(maxDatadogPodAutoscalerObjects, store)

	controller, err := workload.NewController(clusterID, eventRecorder, apiCl.RESTMapper, apiCl.ScaleCl, apiCl.DynamicInformerCl, apiCl.DynamicInformerFactory, le.IsLeader, store, podWatcher, sender, limitHeap)
	if err != nil {
		return nil, fmt.Errorf("Unable to start workload autoscaling controller: %w", err)
	}

	// Start informers & controllers (informers can be started multiple times)
	apiCl.DynamicInformerFactory.Start(ctx.Done())
	apiCl.InformerFactory.Start(ctx.Done())

	// TODO: Wait POD Watcher sync before running the controller
	go podWatcher.Run(ctx)
	go controller.Run(ctx)
	go localRecommender.Run(ctx)

	return podPatcher, nil
}
