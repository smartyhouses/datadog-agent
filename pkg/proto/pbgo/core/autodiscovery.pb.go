// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: datadog/autodiscovery/autodiscovery.proto

package core

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ConfigEventType int32

const (
	ConfigEventType_SCHEDULE   ConfigEventType = 0
	ConfigEventType_UNSCHEDULE ConfigEventType = 1
)

// Enum value maps for ConfigEventType.
var (
	ConfigEventType_name = map[int32]string{
		0: "SCHEDULE",
		1: "UNSCHEDULE",
	}
	ConfigEventType_value = map[string]int32{
		"SCHEDULE":   0,
		"UNSCHEDULE": 1,
	}
)

func (x ConfigEventType) Enum() *ConfigEventType {
	p := new(ConfigEventType)
	*p = x
	return p
}

func (x ConfigEventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ConfigEventType) Descriptor() protoreflect.EnumDescriptor {
	return file_datadog_autodiscovery_autodiscovery_proto_enumTypes[0].Descriptor()
}

func (ConfigEventType) Type() protoreflect.EnumType {
	return &file_datadog_autodiscovery_autodiscovery_proto_enumTypes[0]
}

func (x ConfigEventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ConfigEventType.Descriptor instead.
func (ConfigEventType) EnumDescriptor() ([]byte, []int) {
	return file_datadog_autodiscovery_autodiscovery_proto_rawDescGZIP(), []int{0}
}

type KubeNamespacedName struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Namespace     string                 `protobuf:"bytes,2,opt,name=namespace,proto3" json:"namespace,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KubeNamespacedName) Reset() {
	*x = KubeNamespacedName{}
	mi := &file_datadog_autodiscovery_autodiscovery_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KubeNamespacedName) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeNamespacedName) ProtoMessage() {}

func (x *KubeNamespacedName) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_autodiscovery_autodiscovery_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeNamespacedName.ProtoReflect.Descriptor instead.
func (*KubeNamespacedName) Descriptor() ([]byte, []int) {
	return file_datadog_autodiscovery_autodiscovery_proto_rawDescGZIP(), []int{0}
}

func (x *KubeNamespacedName) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubeNamespacedName) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

type AdvancedADIdentifier struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	KubeService   *KubeNamespacedName    `protobuf:"bytes,1,opt,name=kubeService,proto3" json:"kubeService,omitempty"`
	KubeEndpoints *KubeNamespacedName    `protobuf:"bytes,2,opt,name=kubeEndpoints,proto3" json:"kubeEndpoints,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AdvancedADIdentifier) Reset() {
	*x = AdvancedADIdentifier{}
	mi := &file_datadog_autodiscovery_autodiscovery_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AdvancedADIdentifier) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AdvancedADIdentifier) ProtoMessage() {}

func (x *AdvancedADIdentifier) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_autodiscovery_autodiscovery_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AdvancedADIdentifier.ProtoReflect.Descriptor instead.
func (*AdvancedADIdentifier) Descriptor() ([]byte, []int) {
	return file_datadog_autodiscovery_autodiscovery_proto_rawDescGZIP(), []int{1}
}

func (x *AdvancedADIdentifier) GetKubeService() *KubeNamespacedName {
	if x != nil {
		return x.KubeService
	}
	return nil
}

func (x *AdvancedADIdentifier) GetKubeEndpoints() *KubeNamespacedName {
	if x != nil {
		return x.KubeEndpoints
	}
	return nil
}

type Config struct {
	state                   protoimpl.MessageState  `protogen:"open.v1"`
	Name                    string                  `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Instances               [][]byte                `protobuf:"bytes,2,rep,name=instances,proto3" json:"instances,omitempty"`
	InitConfig              []byte                  `protobuf:"bytes,3,opt,name=initConfig,proto3" json:"initConfig,omitempty"`
	MetricConfig            []byte                  `protobuf:"bytes,4,opt,name=metricConfig,proto3" json:"metricConfig,omitempty"`
	LogsConfig              []byte                  `protobuf:"bytes,5,opt,name=logsConfig,proto3" json:"logsConfig,omitempty"`
	AdIdentifiers           []string                `protobuf:"bytes,6,rep,name=adIdentifiers,proto3" json:"adIdentifiers,omitempty"`
	AdvancedAdIdentifiers   []*AdvancedADIdentifier `protobuf:"bytes,7,rep,name=advancedAdIdentifiers,proto3" json:"advancedAdIdentifiers,omitempty"`
	Provider                string                  `protobuf:"bytes,8,opt,name=provider,proto3" json:"provider,omitempty"`
	ServiceId               string                  `protobuf:"bytes,9,opt,name=serviceId,proto3" json:"serviceId,omitempty"`
	TaggerEntity            string                  `protobuf:"bytes,10,opt,name=taggerEntity,proto3" json:"taggerEntity,omitempty"`
	ClusterCheck            bool                    `protobuf:"varint,11,opt,name=clusterCheck,proto3" json:"clusterCheck,omitempty"`
	NodeName                string                  `protobuf:"bytes,12,opt,name=nodeName,proto3" json:"nodeName,omitempty"`
	Source                  string                  `protobuf:"bytes,13,opt,name=source,proto3" json:"source,omitempty"`
	IgnoreAutodiscoveryTags bool                    `protobuf:"varint,14,opt,name=ignoreAutodiscoveryTags,proto3" json:"ignoreAutodiscoveryTags,omitempty"`
	MetricsExcluded         bool                    `protobuf:"varint,15,opt,name=metricsExcluded,proto3" json:"metricsExcluded,omitempty"`
	LogsExcluded            bool                    `protobuf:"varint,16,opt,name=logsExcluded,proto3" json:"logsExcluded,omitempty"`
	EventType               ConfigEventType         `protobuf:"varint,17,opt,name=eventType,proto3,enum=datadog.autodiscovery.ConfigEventType" json:"eventType,omitempty"`
	unknownFields           protoimpl.UnknownFields
	sizeCache               protoimpl.SizeCache
}

func (x *Config) Reset() {
	*x = Config{}
	mi := &file_datadog_autodiscovery_autodiscovery_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Config) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Config) ProtoMessage() {}

func (x *Config) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_autodiscovery_autodiscovery_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Config.ProtoReflect.Descriptor instead.
func (*Config) Descriptor() ([]byte, []int) {
	return file_datadog_autodiscovery_autodiscovery_proto_rawDescGZIP(), []int{2}
}

func (x *Config) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Config) GetInstances() [][]byte {
	if x != nil {
		return x.Instances
	}
	return nil
}

func (x *Config) GetInitConfig() []byte {
	if x != nil {
		return x.InitConfig
	}
	return nil
}

func (x *Config) GetMetricConfig() []byte {
	if x != nil {
		return x.MetricConfig
	}
	return nil
}

func (x *Config) GetLogsConfig() []byte {
	if x != nil {
		return x.LogsConfig
	}
	return nil
}

func (x *Config) GetAdIdentifiers() []string {
	if x != nil {
		return x.AdIdentifiers
	}
	return nil
}

func (x *Config) GetAdvancedAdIdentifiers() []*AdvancedADIdentifier {
	if x != nil {
		return x.AdvancedAdIdentifiers
	}
	return nil
}

func (x *Config) GetProvider() string {
	if x != nil {
		return x.Provider
	}
	return ""
}

func (x *Config) GetServiceId() string {
	if x != nil {
		return x.ServiceId
	}
	return ""
}

func (x *Config) GetTaggerEntity() string {
	if x != nil {
		return x.TaggerEntity
	}
	return ""
}

func (x *Config) GetClusterCheck() bool {
	if x != nil {
		return x.ClusterCheck
	}
	return false
}

func (x *Config) GetNodeName() string {
	if x != nil {
		return x.NodeName
	}
	return ""
}

func (x *Config) GetSource() string {
	if x != nil {
		return x.Source
	}
	return ""
}

func (x *Config) GetIgnoreAutodiscoveryTags() bool {
	if x != nil {
		return x.IgnoreAutodiscoveryTags
	}
	return false
}

func (x *Config) GetMetricsExcluded() bool {
	if x != nil {
		return x.MetricsExcluded
	}
	return false
}

func (x *Config) GetLogsExcluded() bool {
	if x != nil {
		return x.LogsExcluded
	}
	return false
}

func (x *Config) GetEventType() ConfigEventType {
	if x != nil {
		return x.EventType
	}
	return ConfigEventType_SCHEDULE
}

type AutodiscoveryStreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Configs       []*Config              `protobuf:"bytes,1,rep,name=configs,proto3" json:"configs,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AutodiscoveryStreamResponse) Reset() {
	*x = AutodiscoveryStreamResponse{}
	mi := &file_datadog_autodiscovery_autodiscovery_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AutodiscoveryStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutodiscoveryStreamResponse) ProtoMessage() {}

func (x *AutodiscoveryStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_autodiscovery_autodiscovery_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutodiscoveryStreamResponse.ProtoReflect.Descriptor instead.
func (*AutodiscoveryStreamResponse) Descriptor() ([]byte, []int) {
	return file_datadog_autodiscovery_autodiscovery_proto_rawDescGZIP(), []int{3}
}

func (x *AutodiscoveryStreamResponse) GetConfigs() []*Config {
	if x != nil {
		return x.Configs
	}
	return nil
}

var File_datadog_autodiscovery_autodiscovery_proto protoreflect.FileDescriptor

var file_datadog_autodiscovery_autodiscovery_proto_rawDesc = []byte{
	0x0a, 0x29, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2f, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63,
	0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x15, 0x64, 0x61, 0x74,
	0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65,
	0x72, 0x79, 0x22, 0x46, 0x0a, 0x12, 0x4b, 0x75, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x64, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x22, 0xb4, 0x01, 0x0a, 0x14, 0x41,
	0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x41, 0x44, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66,
	0x69, 0x65, 0x72, 0x12, 0x4b, 0x0a, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64,
	0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x4e,
	0x61, 0x6d, 0x65, 0x52, 0x0b, 0x6b, 0x75, 0x62, 0x65, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x4f, 0x0a, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x4b, 0x75, 0x62, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x73, 0x70, 0x61, 0x63, 0x65, 0x64, 0x4e, 0x61,
	0x6d, 0x65, 0x52, 0x0d, 0x6b, 0x75, 0x62, 0x65, 0x45, 0x6e, 0x64, 0x70, 0x6f, 0x69, 0x6e, 0x74,
	0x73, 0x22, 0xab, 0x05, 0x0a, 0x06, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x18, 0x02, 0x20,
	0x03, 0x28, 0x0c, 0x52, 0x09, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e, 0x63, 0x65, 0x73, 0x12, 0x1e,
	0x0a, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0a, 0x69, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x12, 0x22,
	0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x1e, 0x0a, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x6c, 0x6f, 0x67, 0x73, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69,
	0x65, 0x72, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x49, 0x64, 0x65,
	0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x61, 0x0a, 0x15, 0x61, 0x64, 0x76, 0x61,
	0x6e, 0x63, 0x65, 0x64, 0x41, 0x64, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72,
	0x73, 0x18, 0x07, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2b, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e,
	0x41, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x41, 0x44, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x66, 0x69, 0x65, 0x72, 0x52, 0x15, 0x61, 0x64, 0x76, 0x61, 0x6e, 0x63, 0x65, 0x64, 0x41, 0x64,
	0x49, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70,
	0x72, 0x6f, 0x76, 0x69, 0x64, 0x65, 0x72, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x49, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x49, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x74, 0x61, 0x67, 0x67, 0x65, 0x72, 0x45,
	0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x74, 0x61, 0x67,
	0x67, 0x65, 0x72, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x6c, 0x75,
	0x73, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x0c, 0x63, 0x6c, 0x75, 0x73, 0x74, 0x65, 0x72, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x12, 0x1a, 0x0a,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x08, 0x6e, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x6f, 0x75, 0x72, 0x63,
	0x65, 0x12, 0x38, 0x0a, 0x17, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x64,
	0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x18, 0x0e, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x17, 0x69, 0x67, 0x6e, 0x6f, 0x72, 0x65, 0x41, 0x75, 0x74, 0x6f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x54, 0x61, 0x67, 0x73, 0x12, 0x28, 0x0a, 0x0f, 0x6d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x0f,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x0f, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x22, 0x0a, 0x0c, 0x6c, 0x6f, 0x67, 0x73, 0x45, 0x78, 0x63,
	0x6c, 0x75, 0x64, 0x65, 0x64, 0x18, 0x10, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0c, 0x6c, 0x6f, 0x67,
	0x73, 0x45, 0x78, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x64, 0x12, 0x44, 0x0a, 0x09, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x11, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x26, 0x2e, 0x64,
	0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f,
	0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x22,
	0x56, 0x0a, 0x1b, 0x41, 0x75, 0x74, 0x6f, 0x64, 0x69, 0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x37,
	0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x61, 0x75, 0x74, 0x6f, 0x64, 0x69,
	0x73, 0x63, 0x6f, 0x76, 0x65, 0x72, 0x79, 0x2e, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x52, 0x07,
	0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x73, 0x2a, 0x2f, 0x0a, 0x0f, 0x43, 0x6f, 0x6e, 0x66, 0x69,
	0x67, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0c, 0x0a, 0x08, 0x53, 0x43,
	0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x10, 0x00, 0x12, 0x0e, 0x0a, 0x0a, 0x55, 0x4e, 0x53, 0x43,
	0x48, 0x45, 0x44, 0x55, 0x4c, 0x45, 0x10, 0x01, 0x42, 0x15, 0x5a, 0x13, 0x70, 0x6b, 0x67, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x62, 0x67, 0x6f, 0x2f, 0x63, 0x6f, 0x72, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_datadog_autodiscovery_autodiscovery_proto_rawDescOnce sync.Once
	file_datadog_autodiscovery_autodiscovery_proto_rawDescData = file_datadog_autodiscovery_autodiscovery_proto_rawDesc
)

func file_datadog_autodiscovery_autodiscovery_proto_rawDescGZIP() []byte {
	file_datadog_autodiscovery_autodiscovery_proto_rawDescOnce.Do(func() {
		file_datadog_autodiscovery_autodiscovery_proto_rawDescData = protoimpl.X.CompressGZIP(file_datadog_autodiscovery_autodiscovery_proto_rawDescData)
	})
	return file_datadog_autodiscovery_autodiscovery_proto_rawDescData
}

var file_datadog_autodiscovery_autodiscovery_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_datadog_autodiscovery_autodiscovery_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_datadog_autodiscovery_autodiscovery_proto_goTypes = []any{
	(ConfigEventType)(0),                // 0: datadog.autodiscovery.ConfigEventType
	(*KubeNamespacedName)(nil),          // 1: datadog.autodiscovery.KubeNamespacedName
	(*AdvancedADIdentifier)(nil),        // 2: datadog.autodiscovery.AdvancedADIdentifier
	(*Config)(nil),                      // 3: datadog.autodiscovery.Config
	(*AutodiscoveryStreamResponse)(nil), // 4: datadog.autodiscovery.AutodiscoveryStreamResponse
}
var file_datadog_autodiscovery_autodiscovery_proto_depIdxs = []int32{
	1, // 0: datadog.autodiscovery.AdvancedADIdentifier.kubeService:type_name -> datadog.autodiscovery.KubeNamespacedName
	1, // 1: datadog.autodiscovery.AdvancedADIdentifier.kubeEndpoints:type_name -> datadog.autodiscovery.KubeNamespacedName
	2, // 2: datadog.autodiscovery.Config.advancedAdIdentifiers:type_name -> datadog.autodiscovery.AdvancedADIdentifier
	0, // 3: datadog.autodiscovery.Config.eventType:type_name -> datadog.autodiscovery.ConfigEventType
	3, // 4: datadog.autodiscovery.AutodiscoveryStreamResponse.configs:type_name -> datadog.autodiscovery.Config
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_datadog_autodiscovery_autodiscovery_proto_init() }
func file_datadog_autodiscovery_autodiscovery_proto_init() {
	if File_datadog_autodiscovery_autodiscovery_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datadog_autodiscovery_autodiscovery_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_datadog_autodiscovery_autodiscovery_proto_goTypes,
		DependencyIndexes: file_datadog_autodiscovery_autodiscovery_proto_depIdxs,
		EnumInfos:         file_datadog_autodiscovery_autodiscovery_proto_enumTypes,
		MessageInfos:      file_datadog_autodiscovery_autodiscovery_proto_msgTypes,
	}.Build()
	File_datadog_autodiscovery_autodiscovery_proto = out.File
	file_datadog_autodiscovery_autodiscovery_proto_rawDesc = nil
	file_datadog_autodiscovery_autodiscovery_proto_goTypes = nil
	file_datadog_autodiscovery_autodiscovery_proto_depIdxs = nil
}
