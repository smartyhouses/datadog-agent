// Code generated by genpost.go; DO NOT EDIT.

package offsetguess

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/ebpf/ebpftest"
)

func TestCgoAlignment_Proc(t *testing.T) {
	ebpftest.TestCgoAlignment[Proc](t)
}

func TestCgoAlignment_TracerStatus(t *testing.T) {
	ebpftest.TestCgoAlignment[TracerStatus](t)
}

func TestCgoAlignment_ConntrackStatus(t *testing.T) {
	ebpftest.TestCgoAlignment[ConntrackStatus](t)
}
