// Code generated by genpost.go; DO NOT EDIT.

package ebpf

import (
	"testing"

	"github.com/DataDog/datadog-agent/pkg/ebpf/ebpftest"
)

func TestCgoAlignment_ConnTuple(t *testing.T) {
	ebpftest.TestCgoAlignment[ConnTuple](t)
}

func TestCgoAlignment_TCPStats(t *testing.T) {
	ebpftest.TestCgoAlignment[TCPStats](t)
}

func TestCgoAlignment_ConnStats(t *testing.T) {
	ebpftest.TestCgoAlignment[ConnStats](t)
}

func TestCgoAlignment_Conn(t *testing.T) {
	ebpftest.TestCgoAlignment[Conn](t)
}

func TestCgoAlignment_SkpConn(t *testing.T) {
	ebpftest.TestCgoAlignment[SkpConn](t)
}

func TestCgoAlignment_PidTs(t *testing.T) {
	ebpftest.TestCgoAlignment[PidTs](t)
}

func TestCgoAlignment_Batch(t *testing.T) {
	ebpftest.TestCgoAlignment[Batch](t)
}

func TestCgoAlignment_Telemetry(t *testing.T) {
	ebpftest.TestCgoAlignment[Telemetry](t)
}

func TestCgoAlignment_PortBinding(t *testing.T) {
	ebpftest.TestCgoAlignment[PortBinding](t)
}

func TestCgoAlignment_PIDFD(t *testing.T) {
	ebpftest.TestCgoAlignment[PIDFD](t)
}

func TestCgoAlignment_UDPRecvSock(t *testing.T) {
	ebpftest.TestCgoAlignment[UDPRecvSock](t)
}

func TestCgoAlignment_BindSyscallArgs(t *testing.T) {
	ebpftest.TestCgoAlignment[BindSyscallArgs](t)
}

func TestCgoAlignment_ProtocolStack(t *testing.T) {
	ebpftest.TestCgoAlignment[ProtocolStack](t)
}

func TestCgoAlignment_ProtocolStackWrapper(t *testing.T) {
	ebpftest.TestCgoAlignment[ProtocolStackWrapper](t)
}

func TestCgoAlignment_TLSTags(t *testing.T) {
	ebpftest.TestCgoAlignment[TLSTags](t)
}

func TestCgoAlignment_TLSTagsWrapper(t *testing.T) {
	ebpftest.TestCgoAlignment[TLSTagsWrapper](t)
}
