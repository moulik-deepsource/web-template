// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.15.0
// source: protobuf/os/os.proto

package os

import (
	proto "github.com/golang/protobuf/proto"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

type OS struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name       string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	CreatedAt  *timestamppb.Timestamp `protobuf:"bytes,9000,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt *timestamppb.Timestamp `protobuf:"bytes,9001,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
}

func (x *OS) Reset() {
	*x = OS{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protobuf_os_os_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OS) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OS) ProtoMessage() {}

func (x *OS) ProtoReflect() protoreflect.Message {
	mi := &file_protobuf_os_os_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OS.ProtoReflect.Descriptor instead.
func (*OS) Descriptor() ([]byte, []int) {
	return file_protobuf_os_os_proto_rawDescGZIP(), []int{0}
}

func (x *OS) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *OS) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OS) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *OS) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

var File_protobuf_os_os_proto protoreflect.FileDescriptor

var file_protobuf_os_os_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x6f, 0x73, 0x2f, 0x6f, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x6f, 0x73, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xa2, 0x01, 0x0a, 0x02,
	0x4f, 0x53, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x64, 0x5f, 0x61, 0x74, 0x18, 0xa8, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64,
	0x41, 0x74, 0x12, 0x3c, 0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x61,
	0x74, 0x18, 0xa9, 0x46, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x0a, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x41, 0x74,
	0x42, 0x32, 0x5a, 0x30, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72,
	0x69, 0x63, 0x6b, 0x79, 0x70, 0x61, 0x69, 0x2f, 0x77, 0x65, 0x62, 0x2d, 0x74, 0x65, 0x6d, 0x70,
	0x6c, 0x61, 0x74, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x6f, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_protobuf_os_os_proto_rawDescOnce sync.Once
	file_protobuf_os_os_proto_rawDescData = file_protobuf_os_os_proto_rawDesc
)

func file_protobuf_os_os_proto_rawDescGZIP() []byte {
	file_protobuf_os_os_proto_rawDescOnce.Do(func() {
		file_protobuf_os_os_proto_rawDescData = protoimpl.X.CompressGZIP(file_protobuf_os_os_proto_rawDescData)
	})
	return file_protobuf_os_os_proto_rawDescData
}

var file_protobuf_os_os_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_protobuf_os_os_proto_goTypes = []interface{}{
	(*OS)(nil),                    // 0: os.OS
	(*timestamppb.Timestamp)(nil), // 1: google.protobuf.Timestamp
}
var file_protobuf_os_os_proto_depIdxs = []int32{
	1, // 0: os.OS.created_at:type_name -> google.protobuf.Timestamp
	1, // 1: os.OS.modified_at:type_name -> google.protobuf.Timestamp
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_protobuf_os_os_proto_init() }
func file_protobuf_os_os_proto_init() {
	if File_protobuf_os_os_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protobuf_os_os_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OS); i {
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
			RawDescriptor: file_protobuf_os_os_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_protobuf_os_os_proto_goTypes,
		DependencyIndexes: file_protobuf_os_os_proto_depIdxs,
		MessageInfos:      file_protobuf_os_os_proto_msgTypes,
	}.Build()
	File_protobuf_os_os_proto = out.File
	file_protobuf_os_os_proto_rawDesc = nil
	file_protobuf_os_os_proto_goTypes = nil
	file_protobuf_os_os_proto_depIdxs = nil
}
