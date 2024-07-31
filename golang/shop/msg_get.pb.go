// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: shop/msg_get.proto

package shop

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

type GetByIdRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *GetByIdRequest) Reset() {
	*x = GetByIdRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_msg_get_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetByIdRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetByIdRequest) ProtoMessage() {}

func (x *GetByIdRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_msg_get_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetByIdRequest.ProtoReflect.Descriptor instead.
func (*GetByIdRequest) Descriptor() ([]byte, []int) {
	return file_shop_msg_get_proto_rawDescGZIP(), []int{0}
}

func (x *GetByIdRequest) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

var File_shop_msg_get_proto protoreflect.FileDescriptor

var file_shop_msg_get_proto_rawDesc = []byte{
	0x0a, 0x12, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x67, 0x65, 0x74, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x73, 0x68, 0x6f, 0x70, 0x22, 0x29, 0x0a, 0x0e, 0x47, 0x65,
	0x74, 0x42, 0x79, 0x49, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07,
	0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73,
	0x68, 0x6f, 0x70, 0x49, 0x64, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x62, 0x69, 0x61, 0x6e, 0x72, 0x77, 0x78, 0x2f, 0x6d, 0x79,
	0x73, 0x68, 0x6f, 0x70, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x70,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_msg_get_proto_rawDescOnce sync.Once
	file_shop_msg_get_proto_rawDescData = file_shop_msg_get_proto_rawDesc
)

func file_shop_msg_get_proto_rawDescGZIP() []byte {
	file_shop_msg_get_proto_rawDescOnce.Do(func() {
		file_shop_msg_get_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_msg_get_proto_rawDescData)
	})
	return file_shop_msg_get_proto_rawDescData
}

var file_shop_msg_get_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_shop_msg_get_proto_goTypes = []interface{}{
	(*GetByIdRequest)(nil), // 0: shop.GetByIdRequest
}
var file_shop_msg_get_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shop_msg_get_proto_init() }
func file_shop_msg_get_proto_init() {
	if File_shop_msg_get_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_msg_get_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetByIdRequest); i {
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
			RawDescriptor: file_shop_msg_get_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_shop_msg_get_proto_goTypes,
		DependencyIndexes: file_shop_msg_get_proto_depIdxs,
		MessageInfos:      file_shop_msg_get_proto_msgTypes,
	}.Build()
	File_shop_msg_get_proto = out.File
	file_shop_msg_get_proto_rawDesc = nil
	file_shop_msg_get_proto_goTypes = nil
	file_shop_msg_get_proto_depIdxs = nil
}
