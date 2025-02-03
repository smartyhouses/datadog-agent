// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2025-present Datadog, Inc.

//go:build kubeapiserver

// Package local provides local recommendations for autoscaling workloads.
package local

import (
	"fmt"

	datadoghq "github.com/DataDog/datadog-operator/api/datadoghq/v1alpha1"

	corev1 "k8s.io/api/core/v1"
)

const (
	watermarkTolerance             = 5
	containerCPUUsageMetricName    = "container.cpu.usage"
	containerMemoryUsageMetricName = "container.memory.usage"
)

var (
	resourceToMetric = map[corev1.ResourceName]string{
		corev1.ResourceCPU:    containerCPUUsageMetricName,
		corev1.ResourceMemory: containerMemoryUsageMetricName,
	}
)

type resourceRecommenderSettings struct {
	metricName    string
	containerName string
	lowWatermark  float64
	highWatermark float64
}

func newResourceRecommenderSettings(target datadoghq.DatadogPodAutoscalerTarget) (*resourceRecommenderSettings, error) {
	if target.Type == datadoghq.DatadogPodAutoscalerContainerResourceTargetType {
		return getOptionsFromContainerResource(target.ContainerResource)
	}
	if target.Type == datadoghq.DatadogPodAutoscalerResourceTargetType {
		return getOptionsFromPodResource(target.PodResource)
	}
	return nil, fmt.Errorf("Invalid target type: %s", target.Type)
}

func getOptionsFromPodResource(target *datadoghq.DatadogPodAutoscalerResourceTarget) (*resourceRecommenderSettings, error) {
	if target == nil {
		return nil, fmt.Errorf("nil target")
	}

	if err := validateTarget(target.Value.Type, target.Name, target.Value); err != nil {
		return nil, err
	}

	recSettings := &resourceRecommenderSettings{
		metricName:    resourceToMetric[target.Name],
		lowWatermark:  float64((*target.Value.Utilization - watermarkTolerance)) / 100.0,
		highWatermark: float64((*target.Value.Utilization + watermarkTolerance)) / 100.0,
	}
	return recSettings, nil
}

func getOptionsFromContainerResource(target *datadoghq.DatadogPodAutoscalerContainerResourceTarget) (*resourceRecommenderSettings, error) {
	if target == nil {
		return nil, fmt.Errorf("nil target")
	}

	if err := validateTarget(target.Value.Type, target.Name, target.Value); err != nil {
		return nil, err
	}

	recSettings := &resourceRecommenderSettings{
		metricName:    resourceToMetric[target.Name],
		lowWatermark:  float64((*target.Value.Utilization - watermarkTolerance)) / 100.0,
		highWatermark: float64((*target.Value.Utilization + watermarkTolerance)) / 100.0,
		containerName: target.Container,
	}
	return recSettings, nil
}

func validateTarget(targetType datadoghq.DatadogPodAutoscalerTargetValueType, name corev1.ResourceName, value datadoghq.DatadogPodAutoscalerTargetValue) error {
	if targetType != datadoghq.DatadogPodAutoscalerUtilizationTargetValueType {
		return fmt.Errorf("invalid value type: %s", targetType)
	}

	_, ok := resourceToMetric[name]
	if !ok {
		return fmt.Errorf("invalid resource name: %s", name)
	}

	if err := validateUtilizationValue(value); err != nil {
		return fmt.Errorf("invalid utilization value: %s", err)
	}

	return nil
}

func validateUtilizationValue(value datadoghq.DatadogPodAutoscalerTargetValue) error {
	if value.Utilization == nil {
		return fmt.Errorf("missing utilization value")
	}
	if *value.Utilization < 1 || *value.Utilization > 100 {
		return fmt.Errorf("utilization value must be between 1 and 100")
	}
	return nil
}
