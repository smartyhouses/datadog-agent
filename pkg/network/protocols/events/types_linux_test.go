// Code generated by genpost.go; DO NOT EDIT.

package events

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/ebpf/ebpftest"
)

func TestCgoAlignment_Batch(t *testing.T) {
	ebpftest.TestCgoAlignment[Batch](t)
}

func TestCgoAlignment_batchKey(t *testing.T) {
	ebpftest.TestCgoAlignment[batchKey](t)
}
