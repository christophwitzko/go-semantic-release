// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.12.3
// source: pkg/updater/updater.proto

package updater

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

type FilesUpdaterForFiles struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FilesUpdaterForFiles) Reset() {
	*x = FilesUpdaterForFiles{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_updater_updater_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesUpdaterForFiles) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesUpdaterForFiles) ProtoMessage() {}

func (x *FilesUpdaterForFiles) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_updater_updater_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesUpdaterForFiles.ProtoReflect.Descriptor instead.
func (*FilesUpdaterForFiles) Descriptor() ([]byte, []int) {
	return file_pkg_updater_updater_proto_rawDescGZIP(), []int{0}
}

type FilesUpdaterApply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FilesUpdaterApply) Reset() {
	*x = FilesUpdaterApply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_updater_updater_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesUpdaterApply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesUpdaterApply) ProtoMessage() {}

func (x *FilesUpdaterApply) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_updater_updater_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesUpdaterApply.ProtoReflect.Descriptor instead.
func (*FilesUpdaterApply) Descriptor() ([]byte, []int) {
	return file_pkg_updater_updater_proto_rawDescGZIP(), []int{1}
}

type FilesUpdaterForFiles_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FilesUpdaterForFiles_Request) Reset() {
	*x = FilesUpdaterForFiles_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_updater_updater_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesUpdaterForFiles_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesUpdaterForFiles_Request) ProtoMessage() {}

func (x *FilesUpdaterForFiles_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_updater_updater_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesUpdaterForFiles_Request.ProtoReflect.Descriptor instead.
func (*FilesUpdaterForFiles_Request) Descriptor() ([]byte, []int) {
	return file_pkg_updater_updater_proto_rawDescGZIP(), []int{0, 0}
}

type FilesUpdaterForFiles_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Files string `protobuf:"bytes,1,opt,name=Files,proto3" json:"Files,omitempty"`
}

func (x *FilesUpdaterForFiles_Response) Reset() {
	*x = FilesUpdaterForFiles_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_updater_updater_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesUpdaterForFiles_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesUpdaterForFiles_Response) ProtoMessage() {}

func (x *FilesUpdaterForFiles_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_updater_updater_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesUpdaterForFiles_Response.ProtoReflect.Descriptor instead.
func (*FilesUpdaterForFiles_Response) Descriptor() ([]byte, []int) {
	return file_pkg_updater_updater_proto_rawDescGZIP(), []int{0, 1}
}

func (x *FilesUpdaterForFiles_Response) GetFiles() string {
	if x != nil {
		return x.Files
	}
	return ""
}

type FilesUpdaterApply_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	File       string `protobuf:"bytes,1,opt,name=File,proto3" json:"File,omitempty"`
	NewVersion string `protobuf:"bytes,2,opt,name=NewVersion,proto3" json:"NewVersion,omitempty"`
}

func (x *FilesUpdaterApply_Request) Reset() {
	*x = FilesUpdaterApply_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_updater_updater_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesUpdaterApply_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesUpdaterApply_Request) ProtoMessage() {}

func (x *FilesUpdaterApply_Request) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_updater_updater_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesUpdaterApply_Request.ProtoReflect.Descriptor instead.
func (*FilesUpdaterApply_Request) Descriptor() ([]byte, []int) {
	return file_pkg_updater_updater_proto_rawDescGZIP(), []int{1, 0}
}

func (x *FilesUpdaterApply_Request) GetFile() string {
	if x != nil {
		return x.File
	}
	return ""
}

func (x *FilesUpdaterApply_Request) GetNewVersion() string {
	if x != nil {
		return x.NewVersion
	}
	return ""
}

type FilesUpdaterApply_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Error string `protobuf:"bytes,1,opt,name=Error,proto3" json:"Error,omitempty"`
}

func (x *FilesUpdaterApply_Response) Reset() {
	*x = FilesUpdaterApply_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_pkg_updater_updater_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *FilesUpdaterApply_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FilesUpdaterApply_Response) ProtoMessage() {}

func (x *FilesUpdaterApply_Response) ProtoReflect() protoreflect.Message {
	mi := &file_pkg_updater_updater_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FilesUpdaterApply_Response.ProtoReflect.Descriptor instead.
func (*FilesUpdaterApply_Response) Descriptor() ([]byte, []int) {
	return file_pkg_updater_updater_proto_rawDescGZIP(), []int{1, 1}
}

func (x *FilesUpdaterApply_Response) GetError() string {
	if x != nil {
		return x.Error
	}
	return ""
}

var File_pkg_updater_updater_proto protoreflect.FileDescriptor

var file_pkg_updater_updater_proto_rawDesc = []byte{
	0x0a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x2f, 0x75, 0x70,
	0x64, 0x61, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x43, 0x0a, 0x14, 0x46,
	0x69, 0x6c, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x1a, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x46, 0x69,
	0x6c, 0x65, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x22, 0x74, 0x0a, 0x11, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x1a, 0x3d, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x46, 0x69, 0x6c, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x46, 0x69, 0x6c, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4e, 0x65, 0x77, 0x56, 0x65, 0x72, 0x73, 0x69,
	0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x4e, 0x65, 0x77, 0x56, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x1a, 0x20, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x14, 0x0a, 0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x05, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x32, 0xa1, 0x01, 0x0a, 0x12, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x50, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x12, 0x49, 0x0a,
	0x08, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x12, 0x1d, 0x2e, 0x46, 0x69, 0x6c, 0x65,
	0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x46, 0x6f, 0x72, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x2e,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x40, 0x0a, 0x05, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x12, 0x1a, 0x2e, 0x46, 0x69, 0x6c, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72,
	0x41, 0x70, 0x70, 0x6c, 0x79, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e,
	0x46, 0x69, 0x6c, 0x65, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x41, 0x70, 0x70, 0x6c,
	0x79, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x3d, 0x5a, 0x3b, 0x67, 0x69,
	0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x67, 0x6f, 0x2d, 0x73, 0x65, 0x6d, 0x61,
	0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x73, 0x65, 0x6d,
	0x61, 0x6e, 0x74, 0x69, 0x63, 0x2d, 0x72, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x2f, 0x70, 0x6b,
	0x67, 0x2f, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_pkg_updater_updater_proto_rawDescOnce sync.Once
	file_pkg_updater_updater_proto_rawDescData = file_pkg_updater_updater_proto_rawDesc
)

func file_pkg_updater_updater_proto_rawDescGZIP() []byte {
	file_pkg_updater_updater_proto_rawDescOnce.Do(func() {
		file_pkg_updater_updater_proto_rawDescData = protoimpl.X.CompressGZIP(file_pkg_updater_updater_proto_rawDescData)
	})
	return file_pkg_updater_updater_proto_rawDescData
}

var file_pkg_updater_updater_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_pkg_updater_updater_proto_goTypes = []interface{}{
	(*FilesUpdaterForFiles)(nil),          // 0: FilesUpdaterForFiles
	(*FilesUpdaterApply)(nil),             // 1: FilesUpdaterApply
	(*FilesUpdaterForFiles_Request)(nil),  // 2: FilesUpdaterForFiles.Request
	(*FilesUpdaterForFiles_Response)(nil), // 3: FilesUpdaterForFiles.Response
	(*FilesUpdaterApply_Request)(nil),     // 4: FilesUpdaterApply.Request
	(*FilesUpdaterApply_Response)(nil),    // 5: FilesUpdaterApply.Response
}
var file_pkg_updater_updater_proto_depIdxs = []int32{
	2, // 0: FilesUpdaterPlugin.ForFiles:input_type -> FilesUpdaterForFiles.Request
	4, // 1: FilesUpdaterPlugin.Apply:input_type -> FilesUpdaterApply.Request
	3, // 2: FilesUpdaterPlugin.ForFiles:output_type -> FilesUpdaterForFiles.Response
	5, // 3: FilesUpdaterPlugin.Apply:output_type -> FilesUpdaterApply.Response
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pkg_updater_updater_proto_init() }
func file_pkg_updater_updater_proto_init() {
	if File_pkg_updater_updater_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_pkg_updater_updater_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesUpdaterForFiles); i {
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
		file_pkg_updater_updater_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesUpdaterApply); i {
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
		file_pkg_updater_updater_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesUpdaterForFiles_Request); i {
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
		file_pkg_updater_updater_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesUpdaterForFiles_Response); i {
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
		file_pkg_updater_updater_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesUpdaterApply_Request); i {
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
		file_pkg_updater_updater_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*FilesUpdaterApply_Response); i {
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
			RawDescriptor: file_pkg_updater_updater_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_pkg_updater_updater_proto_goTypes,
		DependencyIndexes: file_pkg_updater_updater_proto_depIdxs,
		MessageInfos:      file_pkg_updater_updater_proto_msgTypes,
	}.Build()
	File_pkg_updater_updater_proto = out.File
	file_pkg_updater_updater_proto_rawDesc = nil
	file_pkg_updater_updater_proto_goTypes = nil
	file_pkg_updater_updater_proto_depIdxs = nil
}
