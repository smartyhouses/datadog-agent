// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.3
// 	protoc        v5.29.3
// source: datadog/process/workloadmeta_process.proto

package process

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type ProcessStreamResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	EventID       int32                  `protobuf:"varint,1,opt,name=eventID,proto3" json:"eventID,omitempty"`
	SetEvents     []*ProcessEventSet     `protobuf:"bytes,2,rep,name=setEvents,proto3" json:"setEvents,omitempty"`
	UnsetEvents   []*ProcessEventUnset   `protobuf:"bytes,3,rep,name=unsetEvents,proto3" json:"unsetEvents,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessStreamResponse) Reset() {
	*x = ProcessStreamResponse{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessStreamResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStreamResponse) ProtoMessage() {}

func (x *ProcessStreamResponse) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStreamResponse.ProtoReflect.Descriptor instead.
func (*ProcessStreamResponse) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{0}
}

func (x *ProcessStreamResponse) GetEventID() int32 {
	if x != nil {
		return x.EventID
	}
	return 0
}

func (x *ProcessStreamResponse) GetSetEvents() []*ProcessEventSet {
	if x != nil {
		return x.SetEvents
	}
	return nil
}

func (x *ProcessStreamResponse) GetUnsetEvents() []*ProcessEventUnset {
	if x != nil {
		return x.UnsetEvents
	}
	return nil
}

type ProcessEventSet struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pid           int32                  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	Nspid         int32                  `protobuf:"varint,2,opt,name=nspid,proto3" json:"nspid,omitempty"`
	ContainerID   string                 `protobuf:"bytes,3,opt,name=containerID,proto3" json:"containerID,omitempty"`
	CreationTime  int64                  `protobuf:"varint,4,opt,name=creationTime,proto3" json:"creationTime,omitempty"`
	Language      *Language              `protobuf:"bytes,5,opt,name=language,proto3" json:"language,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessEventSet) Reset() {
	*x = ProcessEventSet{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessEventSet) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessEventSet) ProtoMessage() {}

func (x *ProcessEventSet) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessEventSet.ProtoReflect.Descriptor instead.
func (*ProcessEventSet) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{1}
}

func (x *ProcessEventSet) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

func (x *ProcessEventSet) GetNspid() int32 {
	if x != nil {
		return x.Nspid
	}
	return 0
}

func (x *ProcessEventSet) GetContainerID() string {
	if x != nil {
		return x.ContainerID
	}
	return ""
}

func (x *ProcessEventSet) GetCreationTime() int64 {
	if x != nil {
		return x.CreationTime
	}
	return 0
}

func (x *ProcessEventSet) GetLanguage() *Language {
	if x != nil {
		return x.Language
	}
	return nil
}

type ProcessEventUnset struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Pid           int32                  `protobuf:"varint,1,opt,name=pid,proto3" json:"pid,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessEventUnset) Reset() {
	*x = ProcessEventUnset{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessEventUnset) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessEventUnset) ProtoMessage() {}

func (x *ProcessEventUnset) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessEventUnset.ProtoReflect.Descriptor instead.
func (*ProcessEventUnset) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{2}
}

func (x *ProcessEventUnset) GetPid() int32 {
	if x != nil {
		return x.Pid
	}
	return 0
}

type Language struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Name          string                 `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Language) Reset() {
	*x = Language{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Language) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Language) ProtoMessage() {}

func (x *Language) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Language.ProtoReflect.Descriptor instead.
func (*Language) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{3}
}

func (x *Language) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type ProcessStreamEntitiesRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ProcessStreamEntitiesRequest) Reset() {
	*x = ProcessStreamEntitiesRequest{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ProcessStreamEntitiesRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProcessStreamEntitiesRequest) ProtoMessage() {}

func (x *ProcessStreamEntitiesRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProcessStreamEntitiesRequest.ProtoReflect.Descriptor instead.
func (*ProcessStreamEntitiesRequest) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{4}
}

// ParentLanguageAnnotationRequest is sent from the Core-Agent to the Cluster-Agent to notify that
// a language was detected for a given container
type ParentLanguageAnnotationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	PodDetails    []*PodLanguageDetails  `protobuf:"bytes,1,rep,name=podDetails,proto3" json:"podDetails,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ParentLanguageAnnotationRequest) Reset() {
	*x = ParentLanguageAnnotationRequest{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ParentLanguageAnnotationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ParentLanguageAnnotationRequest) ProtoMessage() {}

func (x *ParentLanguageAnnotationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ParentLanguageAnnotationRequest.ProtoReflect.Descriptor instead.
func (*ParentLanguageAnnotationRequest) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{5}
}

func (x *ParentLanguageAnnotationRequest) GetPodDetails() []*PodLanguageDetails {
	if x != nil {
		return x.PodDetails
	}
	return nil
}

// PodLanguageDetails holds the language metadata associated to a given pod
type PodLanguageDetails struct {
	state                protoimpl.MessageState      `protogen:"open.v1"`
	Namespace            string                      `protobuf:"bytes,1,opt,name=namespace,proto3" json:"namespace,omitempty"`
	Name                 string                      `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ContainerDetails     []*ContainerLanguageDetails `protobuf:"bytes,3,rep,name=containerDetails,proto3" json:"containerDetails,omitempty"`
	Ownerref             *KubeOwnerInfo              `protobuf:"bytes,4,opt,name=ownerref,proto3" json:"ownerref,omitempty"`
	InitContainerDetails []*ContainerLanguageDetails `protobuf:"bytes,5,rep,name=initContainerDetails,proto3" json:"initContainerDetails,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *PodLanguageDetails) Reset() {
	*x = PodLanguageDetails{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *PodLanguageDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PodLanguageDetails) ProtoMessage() {}

func (x *PodLanguageDetails) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PodLanguageDetails.ProtoReflect.Descriptor instead.
func (*PodLanguageDetails) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{6}
}

func (x *PodLanguageDetails) GetNamespace() string {
	if x != nil {
		return x.Namespace
	}
	return ""
}

func (x *PodLanguageDetails) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *PodLanguageDetails) GetContainerDetails() []*ContainerLanguageDetails {
	if x != nil {
		return x.ContainerDetails
	}
	return nil
}

func (x *PodLanguageDetails) GetOwnerref() *KubeOwnerInfo {
	if x != nil {
		return x.Ownerref
	}
	return nil
}

func (x *PodLanguageDetails) GetInitContainerDetails() []*ContainerLanguageDetails {
	if x != nil {
		return x.InitContainerDetails
	}
	return nil
}

// ContainerLanguageDetails contains the different languages used in a container
type ContainerLanguageDetails struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ContainerName string                 `protobuf:"bytes,1,opt,name=containerName,proto3" json:"containerName,omitempty"`
	Languages     []*Language            `protobuf:"bytes,2,rep,name=languages,proto3" json:"languages,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ContainerLanguageDetails) Reset() {
	*x = ContainerLanguageDetails{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ContainerLanguageDetails) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ContainerLanguageDetails) ProtoMessage() {}

func (x *ContainerLanguageDetails) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ContainerLanguageDetails.ProtoReflect.Descriptor instead.
func (*ContainerLanguageDetails) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{7}
}

func (x *ContainerLanguageDetails) GetContainerName() string {
	if x != nil {
		return x.ContainerName
	}
	return ""
}

func (x *ContainerLanguageDetails) GetLanguages() []*Language {
	if x != nil {
		return x.Languages
	}
	return nil
}

// KubeOwnerInfo holds metadata about the owner of the pod
type KubeOwnerInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            string                 `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Kind          string                 `protobuf:"bytes,3,opt,name=kind,proto3" json:"kind,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *KubeOwnerInfo) Reset() {
	*x = KubeOwnerInfo{}
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *KubeOwnerInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*KubeOwnerInfo) ProtoMessage() {}

func (x *KubeOwnerInfo) ProtoReflect() protoreflect.Message {
	mi := &file_datadog_process_workloadmeta_process_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use KubeOwnerInfo.ProtoReflect.Descriptor instead.
func (*KubeOwnerInfo) Descriptor() ([]byte, []int) {
	return file_datadog_process_workloadmeta_process_proto_rawDescGZIP(), []int{8}
}

func (x *KubeOwnerInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *KubeOwnerInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *KubeOwnerInfo) GetKind() string {
	if x != nil {
		return x.Kind
	}
	return ""
}

var File_datadog_process_workloadmeta_process_proto protoreflect.FileDescriptor

var file_datadog_process_workloadmeta_process_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x2f, 0x77, 0x6f, 0x72, 0x6b, 0x6c, 0x6f, 0x61, 0x64, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0f, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x22, 0xb7, 0x01,
	0x0a, 0x15, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x49,
	0x44, 0x12, 0x3e, 0x0a, 0x09, 0x73, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x18, 0x02,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x20, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70,
	0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x76,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x52, 0x09, 0x73, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x73, 0x12, 0x44, 0x0a, 0x0b, 0x75, 0x6e, 0x73, 0x65, 0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73,
	0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67,
	0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x45, 0x76, 0x65, 0x6e, 0x74, 0x55, 0x6e, 0x73, 0x65, 0x74, 0x52, 0x0b, 0x75, 0x6e, 0x73, 0x65,
	0x74, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x22, 0xb6, 0x01, 0x0a, 0x0f, 0x50, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x6e, 0x73, 0x70, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6e, 0x73,
	0x70, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72,
	0x49, 0x44, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x49, 0x44, 0x12, 0x22, 0x0a, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x35, 0x0a, 0x08, 0x6c, 0x61, 0x6e,
	0x67, 0x75, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x19, 0x2e, 0x64, 0x61,
	0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x08, 0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65,
	0x22, 0x25, 0x0a, 0x11, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x55, 0x6e, 0x73, 0x65, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x70, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x03, 0x70, 0x69, 0x64, 0x22, 0x1e, 0x0a, 0x08, 0x4c, 0x61, 0x6e, 0x67, 0x75,
	0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1e, 0x0a, 0x1c, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x22, 0x66, 0x0a, 0x1f, 0x50, 0x61, 0x72, 0x65, 0x6e,
	0x74, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x43, 0x0a, 0x0a, 0x70, 0x6f,
	0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x23,
	0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73,
	0x2e, 0x50, 0x6f, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x52, 0x0a, 0x70, 0x6f, 0x64, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22,
	0xb8, 0x02, 0x0a, 0x12, 0x50, 0x6f, 0x64, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73, 0x70,
	0x61, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x6e, 0x61, 0x6d, 0x65, 0x73,
	0x70, 0x61, 0x63, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x55, 0x0a, 0x10, 0x63, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f,
	0x63, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x61,
	0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x10, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12,
	0x3a, 0x0a, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x66, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x1e, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x63,
	0x65, 0x73, 0x73, 0x2e, 0x4b, 0x75, 0x62, 0x65, 0x4f, 0x77, 0x6e, 0x65, 0x72, 0x49, 0x6e, 0x66,
	0x6f, 0x52, 0x08, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x72, 0x65, 0x66, 0x12, 0x5d, 0x0a, 0x14, 0x69,
	0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61,
	0x69, 0x6c, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x64, 0x61, 0x74, 0x61,
	0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x43, 0x6f, 0x6e, 0x74,
	0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x44, 0x65, 0x74,
	0x61, 0x69, 0x6c, 0x73, 0x52, 0x14, 0x69, 0x6e, 0x69, 0x74, 0x43, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x22, 0x79, 0x0a, 0x18, 0x43, 0x6f,
	0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x44,
	0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69,
	0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x63,
	0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x37, 0x0a, 0x09,
	0x6c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x19, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x2e, 0x4c, 0x61, 0x6e, 0x67, 0x75, 0x61, 0x67, 0x65, 0x52, 0x09, 0x6c, 0x61, 0x6e, 0x67,
	0x75, 0x61, 0x67, 0x65, 0x73, 0x22, 0x47, 0x0a, 0x0d, 0x4b, 0x75, 0x62, 0x65, 0x4f, 0x77, 0x6e,
	0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6b, 0x69,
	0x6e, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6b, 0x69, 0x6e, 0x64, 0x32, 0x80,
	0x01, 0x0a, 0x13, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x79,
	0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x12, 0x69, 0x0a, 0x0e, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d,
	0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73, 0x12, 0x2d, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64,
	0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65,
	0x73, 0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x45, 0x6e, 0x74, 0x69, 0x74, 0x69, 0x65, 0x73,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x26, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x64, 0x6f,
	0x67, 0x2e, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x2e, 0x50, 0x72, 0x6f, 0x63, 0x65, 0x73,
	0x73, 0x53, 0x74, 0x72, 0x65, 0x61, 0x6d, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x30,
	0x01, 0x42, 0x18, 0x5a, 0x16, 0x70, 0x6b, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x62, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x63, 0x65, 0x73, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_datadog_process_workloadmeta_process_proto_rawDescOnce sync.Once
	file_datadog_process_workloadmeta_process_proto_rawDescData = file_datadog_process_workloadmeta_process_proto_rawDesc
)

func file_datadog_process_workloadmeta_process_proto_rawDescGZIP() []byte {
	file_datadog_process_workloadmeta_process_proto_rawDescOnce.Do(func() {
		file_datadog_process_workloadmeta_process_proto_rawDescData = protoimpl.X.CompressGZIP(file_datadog_process_workloadmeta_process_proto_rawDescData)
	})
	return file_datadog_process_workloadmeta_process_proto_rawDescData
}

var file_datadog_process_workloadmeta_process_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_datadog_process_workloadmeta_process_proto_goTypes = []any{
	(*ProcessStreamResponse)(nil),           // 0: datadog.process.ProcessStreamResponse
	(*ProcessEventSet)(nil),                 // 1: datadog.process.ProcessEventSet
	(*ProcessEventUnset)(nil),               // 2: datadog.process.ProcessEventUnset
	(*Language)(nil),                        // 3: datadog.process.Language
	(*ProcessStreamEntitiesRequest)(nil),    // 4: datadog.process.ProcessStreamEntitiesRequest
	(*ParentLanguageAnnotationRequest)(nil), // 5: datadog.process.ParentLanguageAnnotationRequest
	(*PodLanguageDetails)(nil),              // 6: datadog.process.PodLanguageDetails
	(*ContainerLanguageDetails)(nil),        // 7: datadog.process.ContainerLanguageDetails
	(*KubeOwnerInfo)(nil),                   // 8: datadog.process.KubeOwnerInfo
}
var file_datadog_process_workloadmeta_process_proto_depIdxs = []int32{
	1, // 0: datadog.process.ProcessStreamResponse.setEvents:type_name -> datadog.process.ProcessEventSet
	2, // 1: datadog.process.ProcessStreamResponse.unsetEvents:type_name -> datadog.process.ProcessEventUnset
	3, // 2: datadog.process.ProcessEventSet.language:type_name -> datadog.process.Language
	6, // 3: datadog.process.ParentLanguageAnnotationRequest.podDetails:type_name -> datadog.process.PodLanguageDetails
	7, // 4: datadog.process.PodLanguageDetails.containerDetails:type_name -> datadog.process.ContainerLanguageDetails
	8, // 5: datadog.process.PodLanguageDetails.ownerref:type_name -> datadog.process.KubeOwnerInfo
	7, // 6: datadog.process.PodLanguageDetails.initContainerDetails:type_name -> datadog.process.ContainerLanguageDetails
	3, // 7: datadog.process.ContainerLanguageDetails.languages:type_name -> datadog.process.Language
	4, // 8: datadog.process.ProcessEntityStream.StreamEntities:input_type -> datadog.process.ProcessStreamEntitiesRequest
	0, // 9: datadog.process.ProcessEntityStream.StreamEntities:output_type -> datadog.process.ProcessStreamResponse
	9, // [9:10] is the sub-list for method output_type
	8, // [8:9] is the sub-list for method input_type
	8, // [8:8] is the sub-list for extension type_name
	8, // [8:8] is the sub-list for extension extendee
	0, // [0:8] is the sub-list for field type_name
}

func init() { file_datadog_process_workloadmeta_process_proto_init() }
func file_datadog_process_workloadmeta_process_proto_init() {
	if File_datadog_process_workloadmeta_process_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_datadog_process_workloadmeta_process_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_datadog_process_workloadmeta_process_proto_goTypes,
		DependencyIndexes: file_datadog_process_workloadmeta_process_proto_depIdxs,
		MessageInfos:      file_datadog_process_workloadmeta_process_proto_msgTypes,
	}.Build()
	File_datadog_process_workloadmeta_process_proto = out.File
	file_datadog_process_workloadmeta_process_proto_rawDesc = nil
	file_datadog_process_workloadmeta_process_proto_goTypes = nil
	file_datadog_process_workloadmeta_process_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProcessEntityStreamClient is the client API for ProcessEntityStream service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProcessEntityStreamClient interface {
	StreamEntities(ctx context.Context, in *ProcessStreamEntitiesRequest, opts ...grpc.CallOption) (ProcessEntityStream_StreamEntitiesClient, error)
}

type processEntityStreamClient struct {
	cc grpc.ClientConnInterface
}

func NewProcessEntityStreamClient(cc grpc.ClientConnInterface) ProcessEntityStreamClient {
	return &processEntityStreamClient{cc}
}

func (c *processEntityStreamClient) StreamEntities(ctx context.Context, in *ProcessStreamEntitiesRequest, opts ...grpc.CallOption) (ProcessEntityStream_StreamEntitiesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_ProcessEntityStream_serviceDesc.Streams[0], "/datadog.process.ProcessEntityStream/StreamEntities", opts...)
	if err != nil {
		return nil, err
	}
	x := &processEntityStreamStreamEntitiesClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type ProcessEntityStream_StreamEntitiesClient interface {
	Recv() (*ProcessStreamResponse, error)
	grpc.ClientStream
}

type processEntityStreamStreamEntitiesClient struct {
	grpc.ClientStream
}

func (x *processEntityStreamStreamEntitiesClient) Recv() (*ProcessStreamResponse, error) {
	m := new(ProcessStreamResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// ProcessEntityStreamServer is the server API for ProcessEntityStream service.
type ProcessEntityStreamServer interface {
	StreamEntities(*ProcessStreamEntitiesRequest, ProcessEntityStream_StreamEntitiesServer) error
}

// UnimplementedProcessEntityStreamServer can be embedded to have forward compatible implementations.
type UnimplementedProcessEntityStreamServer struct {
}

func (*UnimplementedProcessEntityStreamServer) StreamEntities(*ProcessStreamEntitiesRequest, ProcessEntityStream_StreamEntitiesServer) error {
	return status.Errorf(codes.Unimplemented, "method StreamEntities not implemented")
}

func RegisterProcessEntityStreamServer(s *grpc.Server, srv ProcessEntityStreamServer) {
	s.RegisterService(&_ProcessEntityStream_serviceDesc, srv)
}

func _ProcessEntityStream_StreamEntities_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(ProcessStreamEntitiesRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(ProcessEntityStreamServer).StreamEntities(m, &processEntityStreamStreamEntitiesServer{stream})
}

type ProcessEntityStream_StreamEntitiesServer interface {
	Send(*ProcessStreamResponse) error
	grpc.ServerStream
}

type processEntityStreamStreamEntitiesServer struct {
	grpc.ServerStream
}

func (x *processEntityStreamStreamEntitiesServer) Send(m *ProcessStreamResponse) error {
	return x.ServerStream.SendMsg(m)
}

var _ProcessEntityStream_serviceDesc = grpc.ServiceDesc{
	ServiceName: "datadog.process.ProcessEntityStream",
	HandlerType: (*ProcessEntityStreamServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamEntities",
			Handler:       _ProcessEntityStream_StreamEntities_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "datadog/process/workloadmeta_process.proto",
}
