// Code generated by genpost.go; DO NOT EDIT.

package ebpf

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/ebpf/ebpftest"
)

func TestCgoAlignment_LockRange(t *testing.T) {
	ebpftest.TestCgoAlignment[LockRange](t)
}

func TestCgoAlignment_ContentionData(t *testing.T) {
	ebpftest.TestCgoAlignment[ContentionData](t)
}
