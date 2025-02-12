// Code generated by cmd/cgo -godefs; DO NOT EDIT.
// cgo -godefs -- -I ../../ebpf/c -I ../../../ebpf/c -fsigned-char types.go

package kafka

const (
	TopicNameBuckets                     = 0xa
	TopicNameMaxSize                     = 0x50
	MaxSupportedProduceRequestApiVersion = 0xa
	MaxSupportedFetchRequestApiVersion   = 0xc
)

type ConnTuple struct {
	Saddr_h  uint64
	Saddr_l  uint64
	Daddr_h  uint64
	Daddr_l  uint64
	Sport    uint16
	Dport    uint16
	Netns    uint32
	Pid      uint32
	Metadata uint32
}

type EbpfTx struct {
	Tup         ConnTuple
	Transaction KafkaTransaction
}

type KafkaTransactionKey struct {
	Tuple     ConnTuple
	Id        int32
	Pad_cgo_0 [4]byte
}
type KafkaTransaction struct {
	Request_started     uint64
	Response_last_seen  uint64
	Records_count       uint32
	Request_api_key     uint8
	Request_api_version uint8
	Topic_name_size     uint8
	Tags                uint8
	Topic_name          [80]byte
	Error_code          int8
	Pad_cgo_0           [7]byte
}

type KafkaResponseContext struct {
	Transaction                 KafkaTransaction
	Remainder_buf               [4]int8
	Record_batches_num_bytes    int32
	Record_batch_length         int32
	Expected_tcp_seq            uint32
	Carry_over_offset           int32
	Partitions_count            uint32
	Varint_value                uint32
	Record_batches_arrays_idx   uint32
	Record_batches_arrays_count uint32
	State                       uint8
	Remainder                   uint8
	Varint_position             uint8
	Partition_error_code        int8
	Partition_state             uint8
	Pad_cgo_0                   [7]byte
}

type RawKernelTelemetry struct {
	Topic_name_size_buckets  [10]uint64
	Produce_no_required_acks uint64
}
