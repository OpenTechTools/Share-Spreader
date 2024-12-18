// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.0
// 	protoc        v3.19.6
// source: internal/auth/auth.proto

package auth

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

// Example request message
type ServiceRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Authorized    bool                   `protobuf:"varint,1,opt,name=authorized,proto3" json:"authorized,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceRequest) Reset() {
	*x = ServiceRequest{}
	mi := &file_internal_auth_auth_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceRequest) ProtoMessage() {}

func (x *ServiceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_auth_auth_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceRequest.ProtoReflect.Descriptor instead.
func (*ServiceRequest) Descriptor() ([]byte, []int) {
	return file_internal_auth_auth_proto_rawDescGZIP(), []int{0}
}

func (x *ServiceRequest) GetAuthorized() bool {
	if x != nil {
		return x.Authorized
	}
	return false
}

// Example response message
type ServiceResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Access        string                 `protobuf:"bytes,1,opt,name=access,proto3" json:"access,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ServiceResponse) Reset() {
	*x = ServiceResponse{}
	mi := &file_internal_auth_auth_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ServiceResponse) ProtoMessage() {}

func (x *ServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_auth_auth_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ServiceResponse.ProtoReflect.Descriptor instead.
func (*ServiceResponse) Descriptor() ([]byte, []int) {
	return file_internal_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *ServiceResponse) GetAccess() string {
	if x != nil {
		return x.Access
	}
	return ""
}

var File_internal_auth_auth_proto protoreflect.FileDescriptor

var file_internal_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x18, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68,
	0x22, 0x30, 0x0a, 0x0e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x1e, 0x0a, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a, 0x65, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x0a, 0x61, 0x75, 0x74, 0x68, 0x6f, 0x72, 0x69, 0x7a,
	0x65, 0x64, 0x22, 0x29, 0x0a, 0x0f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x32, 0x4c, 0x0a,
	0x04, 0x41, 0x75, 0x74, 0x68, 0x12, 0x44, 0x0a, 0x15, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x53,
	0x70, 0x72, 0x65, 0x61, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x14,
	0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x07, 0x5a, 0x05, 0x2f,
	0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_auth_auth_proto_rawDescOnce sync.Once
	file_internal_auth_auth_proto_rawDescData = file_internal_auth_auth_proto_rawDesc
)

func file_internal_auth_auth_proto_rawDescGZIP() []byte {
	file_internal_auth_auth_proto_rawDescOnce.Do(func() {
		file_internal_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_auth_auth_proto_rawDescData)
	})
	return file_internal_auth_auth_proto_rawDescData
}

var file_internal_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_internal_auth_auth_proto_goTypes = []any{
	(*ServiceRequest)(nil),  // 0: auth.ServiceRequest
	(*ServiceResponse)(nil), // 1: auth.ServiceResponse
}
var file_internal_auth_auth_proto_depIdxs = []int32{
	0, // 0: auth.Auth.AccessSpreaderService:input_type -> auth.ServiceRequest
	1, // 1: auth.Auth.AccessSpreaderService:output_type -> auth.ServiceResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_auth_auth_proto_init() }
func file_internal_auth_auth_proto_init() {
	if File_internal_auth_auth_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_auth_auth_proto_goTypes,
		DependencyIndexes: file_internal_auth_auth_proto_depIdxs,
		MessageInfos:      file_internal_auth_auth_proto_msgTypes,
	}.Build()
	File_internal_auth_auth_proto = out.File
	file_internal_auth_auth_proto_rawDesc = nil
	file_internal_auth_auth_proto_goTypes = nil
	file_internal_auth_auth_proto_depIdxs = nil
}
