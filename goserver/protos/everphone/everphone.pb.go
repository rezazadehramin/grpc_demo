// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.19.4
// source: protos/everphone.proto

package everphone

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

type EverphoneRandomTextInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *EverphoneRandomTextInput) Reset() {
	*x = EverphoneRandomTextInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_everphone_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EverphoneRandomTextInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EverphoneRandomTextInput) ProtoMessage() {}

func (x *EverphoneRandomTextInput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_everphone_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EverphoneRandomTextInput.ProtoReflect.Descriptor instead.
func (*EverphoneRandomTextInput) Descriptor() ([]byte, []int) {
	return file_protos_everphone_proto_rawDescGZIP(), []int{0}
}

func (x *EverphoneRandomTextInput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

type EverphoneRandomTextOutput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Text string `protobuf:"bytes,1,opt,name=Text,proto3" json:"Text,omitempty"`
}

func (x *EverphoneRandomTextOutput) Reset() {
	*x = EverphoneRandomTextOutput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_protos_everphone_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EverphoneRandomTextOutput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EverphoneRandomTextOutput) ProtoMessage() {}

func (x *EverphoneRandomTextOutput) ProtoReflect() protoreflect.Message {
	mi := &file_protos_everphone_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EverphoneRandomTextOutput.ProtoReflect.Descriptor instead.
func (*EverphoneRandomTextOutput) Descriptor() ([]byte, []int) {
	return file_protos_everphone_proto_rawDescGZIP(), []int{1}
}

func (x *EverphoneRandomTextOutput) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

var File_protos_everphone_proto protoreflect.FileDescriptor

var file_protos_everphone_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x65, 0x76, 0x65, 0x72, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2e, 0x0a, 0x18, 0x45, 0x76, 0x65, 0x72,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x49,
	0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x22, 0x2f, 0x0a, 0x19, 0x45, 0x76, 0x65, 0x72,
	0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x4f,
	0x75, 0x74, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x54, 0x65, 0x78, 0x74, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x54, 0x65, 0x78, 0x74, 0x32, 0x50, 0x0a, 0x09, 0x45, 0x76, 0x65,
	0x72, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d,
	0x54, 0x65, 0x78, 0x74, 0x12, 0x19, 0x2e, 0x45, 0x76, 0x65, 0x72, 0x70, 0x68, 0x6f, 0x6e, 0x65,
	0x52, 0x61, 0x6e, 0x64, 0x6f, 0x6d, 0x54, 0x65, 0x78, 0x74, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x1a,
	0x1a, 0x2e, 0x45, 0x76, 0x65, 0x72, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x52, 0x61, 0x6e, 0x64, 0x6f,
	0x6d, 0x54, 0x65, 0x78, 0x74, 0x4f, 0x75, 0x74, 0x70, 0x75, 0x74, 0x42, 0x0d, 0x5a, 0x0b, 0x2e,
	0x2f, 0x65, 0x76, 0x65, 0x72, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_protos_everphone_proto_rawDescOnce sync.Once
	file_protos_everphone_proto_rawDescData = file_protos_everphone_proto_rawDesc
)

func file_protos_everphone_proto_rawDescGZIP() []byte {
	file_protos_everphone_proto_rawDescOnce.Do(func() {
		file_protos_everphone_proto_rawDescData = protoimpl.X.CompressGZIP(file_protos_everphone_proto_rawDescData)
	})
	return file_protos_everphone_proto_rawDescData
}

var file_protos_everphone_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_protos_everphone_proto_goTypes = []interface{}{
	(*EverphoneRandomTextInput)(nil),  // 0: EverphoneRandomTextInput
	(*EverphoneRandomTextOutput)(nil), // 1: EverphoneRandomTextOutput
}
var file_protos_everphone_proto_depIdxs = []int32{
	0, // 0: Everphone.RandomText:input_type -> EverphoneRandomTextInput
	1, // 1: Everphone.RandomText:output_type -> EverphoneRandomTextOutput
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_protos_everphone_proto_init() }
func file_protos_everphone_proto_init() {
	if File_protos_everphone_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_protos_everphone_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EverphoneRandomTextInput); i {
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
		file_protos_everphone_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EverphoneRandomTextOutput); i {
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
			RawDescriptor: file_protos_everphone_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_protos_everphone_proto_goTypes,
		DependencyIndexes: file_protos_everphone_proto_depIdxs,
		MessageInfos:      file_protos_everphone_proto_msgTypes,
	}.Build()
	File_protos_everphone_proto = out.File
	file_protos_everphone_proto_rawDesc = nil
	file_protos_everphone_proto_goTypes = nil
	file_protos_everphone_proto_depIdxs = nil
}
