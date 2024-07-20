// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: shop/shop.proto

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

// Request and Response messages
type RegisterShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"` // Other fields as needed
}

func (x *RegisterShopRequest) Reset() {
	*x = RegisterShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_shop_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterShopRequest) ProtoMessage() {}

func (x *RegisterShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterShopRequest.ProtoReflect.Descriptor instead.
func (*RegisterShopRequest) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{0}
}

func (x *RegisterShopRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type RegisterShopResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *RegisterShopResponse) Reset() {
	*x = RegisterShopResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_shop_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegisterShopResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegisterShopResponse) ProtoMessage() {}

func (x *RegisterShopResponse) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegisterShopResponse.ProtoReflect.Descriptor instead.
func (*RegisterShopResponse) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{1}
}

func (x *RegisterShopResponse) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

type GetShopRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ShopId string `protobuf:"bytes,1,opt,name=shop_id,json=shopId,proto3" json:"shop_id,omitempty"`
}

func (x *GetShopRequest) Reset() {
	*x = GetShopRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_shop_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetShopRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetShopRequest) ProtoMessage() {}

func (x *GetShopRequest) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetShopRequest.ProtoReflect.Descriptor instead.
func (*GetShopRequest) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{2}
}

func (x *GetShopRequest) GetShopId() string {
	if x != nil {
		return x.ShopId
	}
	return ""
}

type Shop struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"` // Other fields as needed
}

func (x *Shop) Reset() {
	*x = Shop{}
	if protoimpl.UnsafeEnabled {
		mi := &file_shop_shop_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Shop) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Shop) ProtoMessage() {}

func (x *Shop) ProtoReflect() protoreflect.Message {
	mi := &file_shop_shop_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Shop.ProtoReflect.Descriptor instead.
func (*Shop) Descriptor() ([]byte, []int) {
	return file_shop_shop_proto_rawDescGZIP(), []int{3}
}

func (x *Shop) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Shop) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_shop_shop_proto protoreflect.FileDescriptor

var file_shop_shop_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x68, 0x6f, 0x70, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x29, 0x0a, 0x13, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x6f,
	0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x14,
	0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x29, 0x0a,
	0x0e, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x17, 0x0a, 0x07, 0x73, 0x68, 0x6f, 0x70, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x68, 0x6f, 0x70, 0x49, 0x64, 0x22, 0x2a, 0x0a, 0x04, 0x53, 0x68, 0x6f, 0x70,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x32, 0x6d, 0x0a, 0x0b, 0x53, 0x68, 0x6f, 0x70, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x3b, 0x0a, 0x0c, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53,
	0x68, 0x6f, 0x70, 0x12, 0x14, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x65, 0x72, 0x53, 0x68,
	0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x65, 0x72, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x21, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x53, 0x68, 0x6f, 0x70, 0x12, 0x0f, 0x2e, 0x47, 0x65,
	0x74, 0x53, 0x68, 0x6f, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x05, 0x2e, 0x53,
	0x68, 0x6f, 0x70, 0x42, 0x39, 0x5a, 0x37, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x66, 0x61, 0x62, 0x69, 0x61, 0x6e, 0x72, 0x77, 0x78, 0x2f, 0x6d, 0x79, 0x73, 0x68,
	0x6f, 0x70, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2d, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_shop_shop_proto_rawDescOnce sync.Once
	file_shop_shop_proto_rawDescData = file_shop_shop_proto_rawDesc
)

func file_shop_shop_proto_rawDescGZIP() []byte {
	file_shop_shop_proto_rawDescOnce.Do(func() {
		file_shop_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_shop_shop_proto_rawDescData)
	})
	return file_shop_shop_proto_rawDescData
}

var file_shop_shop_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_shop_shop_proto_goTypes = []interface{}{
	(*RegisterShopRequest)(nil),  // 0: RegisterShopRequest
	(*RegisterShopResponse)(nil), // 1: RegisterShopResponse
	(*GetShopRequest)(nil),       // 2: GetShopRequest
	(*Shop)(nil),                 // 3: Shop
}
var file_shop_shop_proto_depIdxs = []int32{
	0, // 0: ShopService.RegisterShop:input_type -> RegisterShopRequest
	2, // 1: ShopService.GetShop:input_type -> GetShopRequest
	1, // 2: ShopService.RegisterShop:output_type -> RegisterShopResponse
	3, // 3: ShopService.GetShop:output_type -> Shop
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_shop_shop_proto_init() }
func file_shop_shop_proto_init() {
	if File_shop_shop_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_shop_shop_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterShopRequest); i {
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
		file_shop_shop_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegisterShopResponse); i {
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
		file_shop_shop_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetShopRequest); i {
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
		file_shop_shop_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Shop); i {
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
			RawDescriptor: file_shop_shop_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_shop_shop_proto_goTypes,
		DependencyIndexes: file_shop_shop_proto_depIdxs,
		MessageInfos:      file_shop_shop_proto_msgTypes,
	}.Build()
	File_shop_shop_proto = out.File
	file_shop_shop_proto_rawDesc = nil
	file_shop_shop_proto_goTypes = nil
	file_shop_shop_proto_depIdxs = nil
}
