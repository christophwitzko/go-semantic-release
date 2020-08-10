// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: pkg/condition/ci_condition.proto

package condition

import (
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Name struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Name) Reset() {
	*x = Name{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name) ProtoMessage() {}

func (x *Name) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name.ProtoReflect.Descriptor instead.
func (*Name) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{0}
}

type RunCondition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *RunCondition) Reset() {
	*x = RunCondition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCondition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCondition) ProtoMessage() {}

func (x *RunCondition) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCondition.ProtoReflect.Descriptor instead.
func (*RunCondition) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{1}
}

type GetCurrentBranch struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCurrentBranch) Reset() {
	*x = GetCurrentBranch{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentBranch) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentBranch) ProtoMessage() {}

func (x *GetCurrentBranch) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentBranch.ProtoReflect.Descriptor instead.
func (*GetCurrentBranch) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{2}
}

type GetCurrentSHA struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCurrentSHA) Reset() {
	*x = GetCurrentSHA{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentSHA) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentSHA) ProtoMessage() {}

func (x *GetCurrentSHA) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentSHA.ProtoReflect.Descriptor instead.
func (*GetCurrentSHA) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{3}
}

type Name_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Name_Request) Reset() {
	*x = Name_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name_Request) ProtoMessage() {}

func (x *Name_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name_Request.ProtoReflect.Descriptor instead.
func (*Name_Request) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{0, 0}
}

type Name_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *Name_Response) Reset() {
	*x = Name_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Name_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Name_Response) ProtoMessage() {}

func (x *Name_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Name_Response.ProtoReflect.Descriptor instead.
func (*Name_Response) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{0, 1}
}

func (x *Name_Response) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type RunCondition_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value map[string]string `protobuf:"bytes,1,rep,name=Value,proto3" json:"Value,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *RunCondition_Request) Reset() {
	*x = RunCondition_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCondition_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCondition_Request) ProtoMessage() {}

func (x *RunCondition_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCondition_Request.ProtoReflect.Descriptor instead.
func (*RunCondition_Request) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{1, 0}
}

func (x *RunCondition_Request) GetValue() map[string]string {
	if x != nil {
		return x.Value
	}
	return nil
}

type RunCondition_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *RunCondition_Response) Reset() {
	*x = RunCondition_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RunCondition_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RunCondition_Response) ProtoMessage() {}

func (x *RunCondition_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RunCondition_Response.ProtoReflect.Descriptor instead.
func (*RunCondition_Response) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{1, 1}
}

func (x *RunCondition_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

type GetCurrentBranch_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCurrentBranch_Request) Reset() {
	*x = GetCurrentBranch_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentBranch_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentBranch_Request) ProtoMessage() {}

func (x *GetCurrentBranch_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentBranch_Request.ProtoReflect.Descriptor instead.
func (*GetCurrentBranch_Request) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{2, 0}
}

type GetCurrentBranch_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *GetCurrentBranch_Response) Reset() {
	*x = GetCurrentBranch_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentBranch_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentBranch_Response) ProtoMessage() {}

func (x *GetCurrentBranch_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentBranch_Response.ProtoReflect.Descriptor instead.
func (*GetCurrentBranch_Response) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{2, 1}
}

func (x *GetCurrentBranch_Response) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

type GetCurrentSHA_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *GetCurrentSHA_Request) Reset() {
	*x = GetCurrentSHA_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentSHA_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentSHA_Request) ProtoMessage() {}

func (x *GetCurrentSHA_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentSHA_Request.ProtoReflect.Descriptor instead.
func (*GetCurrentSHA_Request) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{3, 0}
}

type GetCurrentSHA_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Value string `protobuf:"bytes,1,opt,name=Value,proto3" json:"Value,omitempty"`
}

func (x *GetCurrentSHA_Response) Reset() {
	*x = GetCurrentSHA_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_condition_ci_condition_proto_msgTypes[12]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetCurrentSHA_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetCurrentSHA_Response) ProtoMessage() {}

func (x *GetCurrentSHA_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_condition_ci_condition_proto_msgTypes[12]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetCurrentSHA_Response.ProtoReflect.Descriptor instead.
func (*GetCurrentSHA_Response) Descriptor() ([]byte, []int) {
	return file_pkg_condition_ci_condition_proto_rawDescGZIP(), []int{3, 1}
}

func (x *GetCurrentSHA_Response) GetValue() string {
	if x != nil {
		return x.Value
	}
	return ""
}

var File_pkg_condition_ci_condition_proto protoreflect.FileDescriptor

var file_pkg_condition_ci_condition_proto_rawDesc = []byte{
	0x0a, 0x20, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x63, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x22, 0x33, 0x0a, 0x04, 0x4e, 0x61, 0x6d, 0x65, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0xad, 0x01, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x43,
	0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x7b, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x36, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x45,
	0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x1a, 0x38, 0x0a, 0x0a, 0x56,
	0x61, 0x6c, 0x75, 0x65, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76,
	0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75,
	0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x22, 0x3f, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x75,
	0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x1a, 0x09, 0x0a, 0x07, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x22, 0x3c, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x43,
	0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x48, 0x41, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x56, 0x61, 0x6c, 0x75, 0x65, 0x32, 0x86, 0x02, 0x0a, 0x11, 0x43, 0x49, 0x43, 0x6f, 0x6e,
	0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x25, 0x0a, 0x04,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x0d, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x0e, 0x2e, 0x4e, 0x61, 0x6d, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x3d, 0x0a, 0x0c, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74,
	0x69, 0x6f, 0x6e, 0x12, 0x15, 0x2e, 0x52, 0x75, 0x6e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69,
	0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x52, 0x75, 0x6e,
	0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x12, 0x19, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x42, 0x72, 0x61, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1a, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x42, 0x72,
	0x61, 0x6e, 0x63, 0x68, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a,
	0x0d, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x48, 0x41, 0x12, 0x16,
	0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53, 0x48, 0x41, 0x2e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x43, 0x75, 0x72, 0x72,
	0x65, 0x6e, 0x74, 0x53, 0x48, 0x41, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42,
	0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f,
	0x2d, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73,
	0x65, 0x2f, 0x73, 0x65, 0x6d, 0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61,
	0x73, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_pkg_condition_ci_condition_proto_rawDescOnce sync.Once
	file_pkg_condition_ci_condition_proto_rawDescData = file_pkg_condition_ci_condition_proto_rawDesc
)

func file_pkg_condition_ci_condition_proto_rawDescGZIP() []byte {
	file_pkg_condition_ci_condition_proto_rawDescOnce.Do(func() {
		file_pkg_condition_ci_condition_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_condition_ci_condition_proto_rawDescData)
	})
	return file_pkg_condition_ci_condition_proto_rawDescData
}

var file_pkg_condition_ci_condition_proto_msgTypes = make([]protoimpl.MessageInfo, 13)
var file_pkg_condition_ci_condition_proto_goTypes = []interface{}{
	(*Name)(nil),                      // 0: Name
	(*RunCondition)(nil),              // 1: RunCondition
	(*GetCurrentBranch)(nil),          // 2: GetCurrentBranch
	(*GetCurrentSHA)(nil),             // 3: GetCurrentSHA
	(*Name_Request)(nil),              // 4: Name.Request
	(*Name_Response)(nil),             // 5: Name.Response
	(*RunCondition_Request)(nil),      // 6: RunCondition.Request
	(*RunCondition_Response)(nil),     // 7: RunCondition.Response
	nil,                               // 8: RunCondition.Request.ValueEntry
	(*GetCurrentBranch_Request)(nil),  // 9: GetCurrentBranch.Request
	(*GetCurrentBranch_Response)(nil), // 10: GetCurrentBranch.Response
	(*GetCurrentSHA_Request)(nil),     // 11: GetCurrentSHA.Request
	(*GetCurrentSHA_Response)(nil),    // 12: GetCurrentSHA.Response
}
var file_pkg_condition_ci_condition_proto_depIdxs = []int32{
	8,  // 0: RunCondition.Request.Value:type_name -> RunCondition.Request.ValueEntry
	4,  // 1: CIConditionPlugin.Name:input_type -> Name.Request
	6,  // 2: CIConditionPlugin.RunCondition:input_type -> RunCondition.Request
	9,  // 3: CIConditionPlugin.GetCurrentBranch:input_type -> GetCurrentBranch.Request
	11, // 4: CIConditionPlugin.GetCurrentSHA:input_type -> GetCurrentSHA.Request
	5,  // 5: CIConditionPlugin.Name:output_type -> Name.Response
	7,  // 6: CIConditionPlugin.RunCondition:output_type -> RunCondition.Response
	10, // 7: CIConditionPlugin.GetCurrentBranch:output_type -> GetCurrentBranch.Response
	12, // 8: CIConditionPlugin.GetCurrentSHA:output_type -> GetCurrentSHA.Response
	5,  // [5:9] is the sub-list for method output_type
	1,  // [1:5] is the sub-list for method input_type
	1,  // [1:1] is the sub-list for extension type_name
	1,  // [1:1] is the sub-list for extension extendee
	0,  // [0:1] is the sub-list for field type_name
}

func init() { file_pkg_condition_ci_condition_proto_init() }
func file_pkg_condition_ci_condition_proto_init() {
	if File_pkg_condition_ci_condition_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_condition_ci_condition_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCondition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentBranch); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentSHA); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Name_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCondition_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RunCondition_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentBranch_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentBranch_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentSHA_Request); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_pkg_condition_ci_condition_proto_msgTypes[12].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetCurrentSHA_Response); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pkg_condition_ci_condition_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   13,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_condition_ci_condition_proto_goTypes,
		DependencyIndexes: file_pkg_condition_ci_condition_proto_depIdxs,
		MessageInfos:      file_pkg_condition_ci_condition_proto_msgTypes,
	}.Build()
	File_pkg_condition_ci_condition_proto = out.File
	file_pkg_condition_ci_condition_proto_rawDesc = nil
	file_pkg_condition_ci_condition_proto_goTypes = nil
	file_pkg_condition_ci_condition_proto_depIdxs = nil
}