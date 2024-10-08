// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.33.0
// 	protoc        v4.25.3
// source: mailer/msg_verification.proto

package mailer

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

type SendVerificationRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Code  int32  `protobuf:"varint,3,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *SendVerificationRequest) Reset() {
	*x = SendVerificationRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailer_msg_verification_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationRequest) ProtoMessage() {}

func (x *SendVerificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_msg_verification_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationRequest.ProtoReflect.Descriptor instead.
func (*SendVerificationRequest) Descriptor() ([]byte, []int) {
	return file_mailer_msg_verification_proto_rawDescGZIP(), []int{0}
}

func (x *SendVerificationRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SendVerificationRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *SendVerificationRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type SendVerificationResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Message string `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *SendVerificationResponse) Reset() {
	*x = SendVerificationResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mailer_msg_verification_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SendVerificationResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SendVerificationResponse) ProtoMessage() {}

func (x *SendVerificationResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mailer_msg_verification_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SendVerificationResponse.ProtoReflect.Descriptor instead.
func (*SendVerificationResponse) Descriptor() ([]byte, []int) {
	return file_mailer_msg_verification_proto_rawDescGZIP(), []int{1}
}

func (x *SendVerificationResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_mailer_msg_verification_proto protoreflect.FileDescriptor

var file_mailer_msg_verification_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x2f, 0x6d, 0x73, 0x67, 0x5f, 0x76, 0x65, 0x72,
	0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x06, 0x6d, 0x61, 0x69, 0x6c, 0x65, 0x72, 0x22, 0x57, 0x0a, 0x17, 0x53, 0x65, 0x6e, 0x64, 0x56,
	0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12, 0x12, 0x0a, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x22, 0x34, 0x0a, 0x18, 0x53, 0x65, 0x6e, 0x64, 0x56, 0x65, 0x72, 0x69, 0x66, 0x69, 0x63, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x3b, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x62, 0x69, 0x61, 0x6e, 0x72, 0x77, 0x78, 0x2f, 0x6d,
	0x79, 0x73, 0x68, 0x6f, 0x70, 0x73, 0x74, 0x61, 0x74, 0x69, 0x73, 0x74, 0x69, 0x63, 0x73, 0x2d,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6f, 0x6c, 0x61, 0x6e, 0x67, 0x2f, 0x6d, 0x61, 0x69,
	0x6c, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_mailer_msg_verification_proto_rawDescOnce sync.Once
	file_mailer_msg_verification_proto_rawDescData = file_mailer_msg_verification_proto_rawDesc
)

func file_mailer_msg_verification_proto_rawDescGZIP() []byte {
	file_mailer_msg_verification_proto_rawDescOnce.Do(func() {
		file_mailer_msg_verification_proto_rawDescData = protoimpl.X.CompressGZIP(file_mailer_msg_verification_proto_rawDescData)
	})
	return file_mailer_msg_verification_proto_rawDescData
}

var file_mailer_msg_verification_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_mailer_msg_verification_proto_goTypes = []interface{}{
	(*SendVerificationRequest)(nil),  // 0: mailer.SendVerificationRequest
	(*SendVerificationResponse)(nil), // 1: mailer.SendVerificationResponse
}
var file_mailer_msg_verification_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mailer_msg_verification_proto_init() }
func file_mailer_msg_verification_proto_init() {
	if File_mailer_msg_verification_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mailer_msg_verification_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerificationRequest); i {
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
		file_mailer_msg_verification_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SendVerificationResponse); i {
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
			RawDescriptor: file_mailer_msg_verification_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_mailer_msg_verification_proto_goTypes,
		DependencyIndexes: file_mailer_msg_verification_proto_depIdxs,
		MessageInfos:      file_mailer_msg_verification_proto_msgTypes,
	}.Build()
	File_mailer_msg_verification_proto = out.File
	file_mailer_msg_verification_proto_rawDesc = nil
	file_mailer_msg_verification_proto_goTypes = nil
	file_mailer_msg_verification_proto_depIdxs = nil
}
