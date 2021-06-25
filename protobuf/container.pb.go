// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.12.3
// source: hedwig/protobuf/container.proto

package protobuf

import (
	any "github.com/golang/protobuf/ptypes/any"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type MetadataV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Message publisher service
	Publisher string `protobuf:"bytes,1,opt,name=publisher,proto3" json:"publisher,omitempty"`
	// Publish timestamp in epoch milliseconds (integer)
	Timestamp *timestamp.Timestamp `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	// Custom headers associated with the message
	Headers map[string]string `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MetadataV1) Reset() {
	*x = MetadataV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hedwig_protobuf_container_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetadataV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetadataV1) ProtoMessage() {}

func (x *MetadataV1) ProtoReflect() protoreflect.Message {
	mi := &file_hedwig_protobuf_container_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetadataV1.ProtoReflect.Descriptor instead.
func (*MetadataV1) Descriptor() ([]byte, []int) {
	return file_hedwig_protobuf_container_proto_rawDescGZIP(), []int{0}
}

func (x *MetadataV1) GetPublisher() string {
	if x != nil {
		return x.Publisher
	}
	return ""
}

func (x *MetadataV1) GetTimestamp() *timestamp.Timestamp {
	if x != nil {
		return x.Timestamp
	}
	return nil
}

func (x *MetadataV1) GetHeaders() map[string]string {
	if x != nil {
		return x.Headers
	}
	return nil
}

type PayloadV1 struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Format version for the message
	FormatVersion string `protobuf:"bytes,1,opt,name=format_version,json=formatVersion,proto3" json:"format_version,omitempty"`
	// Message identifier
	Id string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// Metadata associated with the message
	Metadata *MetadataV1 `protobuf:"bytes,3,opt,name=metadata,proto3" json:"metadata,omitempty"`
	// Schema to validate the data object with - format: <message type>/<message data version>, e.g. TripCreated/1.0
	Schema string `protobuf:"bytes,4,opt,name=schema,proto3" json:"schema,omitempty"`
	// Message data
	Data *any.Any `protobuf:"bytes,5,opt,name=data,proto3" json:"data,omitempty"`
}

func (x *PayloadV1) Reset() {
	*x = PayloadV1{}
	if protoimpl.UnsafeEnabled {
		mi := &file_hedwig_protobuf_container_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PayloadV1) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PayloadV1) ProtoMessage() {}

func (x *PayloadV1) ProtoReflect() protoreflect.Message {
	mi := &file_hedwig_protobuf_container_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PayloadV1.ProtoReflect.Descriptor instead.
func (*PayloadV1) Descriptor() ([]byte, []int) {
	return file_hedwig_protobuf_container_proto_rawDescGZIP(), []int{1}
}

func (x *PayloadV1) GetFormatVersion() string {
	if x != nil {
		return x.FormatVersion
	}
	return ""
}

func (x *PayloadV1) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *PayloadV1) GetMetadata() *MetadataV1 {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *PayloadV1) GetSchema() string {
	if x != nil {
		return x.Schema
	}
	return ""
}

func (x *PayloadV1) GetData() *any.Any {
	if x != nil {
		return x.Data
	}
	return nil
}

var File_hedwig_protobuf_container_proto protoreflect.FileDescriptor

var file_hedwig_protobuf_container_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x68, 0x65, 0x64, 0x77, 0x69, 0x67, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x61, 0x69, 0x6e, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x06, 0x68, 0x65, 0x64, 0x77, 0x69, 0x67, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xdb, 0x01, 0x0a, 0x0a, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61,
	0x74, 0x61, 0x56, 0x31, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68,
	0x65, 0x72, 0x12, 0x38, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x52, 0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x39, 0x0a, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1f, 0x2e,
	0x68, 0x65, 0x64, 0x77, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56,
	0x31, 0x2e, 0x48, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x07,
	0x68, 0x65, 0x61, 0x64, 0x65, 0x72, 0x73, 0x1a, 0x3a, 0x0a, 0x0c, 0x48, 0x65, 0x61, 0x64, 0x65,
	0x72, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0xb4, 0x01, 0x0a, 0x09, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x56,
	0x31, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73,
	0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x6d, 0x65, 0x74, 0x61,
	0x64, 0x61, 0x74, 0x61, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x68, 0x65, 0x64,
	0x77, 0x69, 0x67, 0x2e, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x56, 0x31, 0x52, 0x08,
	0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x63, 0x68, 0x65,
	0x6d, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x63, 0x68, 0x65, 0x6d, 0x61,
	0x12, 0x28, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66,
	0x2e, 0x41, 0x6e, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x42, 0x43, 0x0a, 0x16, 0x63, 0x6f,
	0x6d, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x68, 0x61, 0x63, 0x68, 0x6f, 0x2e, 0x68, 0x65,
	0x64, 0x77, 0x69, 0x67, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d,
	0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x63, 0x68, 0x61, 0x63, 0x68, 0x6f, 0x2f, 0x68, 0x65, 0x64,
	0x77, 0x69, 0x67, 0x2d, 0x67, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_hedwig_protobuf_container_proto_rawDescOnce sync.Once
	file_hedwig_protobuf_container_proto_rawDescData = file_hedwig_protobuf_container_proto_rawDesc
)

func file_hedwig_protobuf_container_proto_rawDescGZIP() []byte {
	file_hedwig_protobuf_container_proto_rawDescOnce.Do(func() {
		file_hedwig_protobuf_container_proto_rawDescData = protoimpl.X.CompressGZIP(file_hedwig_protobuf_container_proto_rawDescData)
	})
	return file_hedwig_protobuf_container_proto_rawDescData
}

var file_hedwig_protobuf_container_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_hedwig_protobuf_container_proto_goTypes = []interface{}{
	(*MetadataV1)(nil),          // 0: hedwig.MetadataV1
	(*PayloadV1)(nil),           // 1: hedwig.PayloadV1
	nil,                         // 2: hedwig.MetadataV1.HeadersEntry
	(*timestamp.Timestamp)(nil), // 3: google.protobuf.Timestamp
	(*any.Any)(nil),             // 4: google.protobuf.Any
}
var file_hedwig_protobuf_container_proto_depIdxs = []int32{
	3, // 0: hedwig.MetadataV1.timestamp:type_name -> google.protobuf.Timestamp
	2, // 1: hedwig.MetadataV1.headers:type_name -> hedwig.MetadataV1.HeadersEntry
	0, // 2: hedwig.PayloadV1.metadata:type_name -> hedwig.MetadataV1
	4, // 3: hedwig.PayloadV1.data:type_name -> google.protobuf.Any
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_hedwig_protobuf_container_proto_init() }
func file_hedwig_protobuf_container_proto_init() {
	if File_hedwig_protobuf_container_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_hedwig_protobuf_container_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetadataV1); i {
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
		file_hedwig_protobuf_container_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PayloadV1); i {
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
			RawDescriptor: file_hedwig_protobuf_container_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_hedwig_protobuf_container_proto_goTypes,
		DependencyIndexes: file_hedwig_protobuf_container_proto_depIdxs,
		MessageInfos:      file_hedwig_protobuf_container_proto_msgTypes,
	}.Build()
	File_hedwig_protobuf_container_proto = out.File
	file_hedwig_protobuf_container_proto_rawDesc = nil
	file_hedwig_protobuf_container_proto_goTypes = nil
	file_hedwig_protobuf_container_proto_depIdxs = nil
}
