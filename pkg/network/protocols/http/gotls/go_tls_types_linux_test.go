// Code generated by genpost.go; DO NOT EDIT.

package gotls

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/ebpf/ebpftest"
)

func TestCgoAlignment_Location(t *testing.T) {
	ebpftest.TestCgoAlignment[Location](t)
}

func TestCgoAlignment_SliceLocation(t *testing.T) {
	ebpftest.TestCgoAlignment[SliceLocation](t)
}

func TestCgoAlignment_GoroutineIDMetadata(t *testing.T) {
	ebpftest.TestCgoAlignment[GoroutineIDMetadata](t)
}

func TestCgoAlignment_TlsBinaryId(t *testing.T) {
	ebpftest.TestCgoAlignment[TlsBinaryId](t)
}

func TestCgoAlignment_TlsConnLayout(t *testing.T) {
	ebpftest.TestCgoAlignment[TlsConnLayout](t)
}

func TestCgoAlignment_TlsOffsetsData(t *testing.T) {
	ebpftest.TestCgoAlignment[TlsOffsetsData](t)
}
