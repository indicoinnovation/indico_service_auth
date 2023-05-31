// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: mfa.proto

package mfa

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

type GenerateOTPTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Principal    string `protobuf:"bytes,1,opt,name=principal,proto3" json:"principal,omitempty"`
	ClientSecret string `protobuf:"bytes,2,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *GenerateOTPTokenRequest) Reset() {
	*x = GenerateOTPTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mfa_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateOTPTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateOTPTokenRequest) ProtoMessage() {}

func (x *GenerateOTPTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mfa_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateOTPTokenRequest.ProtoReflect.Descriptor instead.
func (*GenerateOTPTokenRequest) Descriptor() ([]byte, []int) {
	return file_mfa_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateOTPTokenRequest) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *GenerateOTPTokenRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type GenerateOTPTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
}

func (x *GenerateOTPTokenResponse) Reset() {
	*x = GenerateOTPTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mfa_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateOTPTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateOTPTokenResponse) ProtoMessage() {}

func (x *GenerateOTPTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mfa_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateOTPTokenResponse.ProtoReflect.Descriptor instead.
func (*GenerateOTPTokenResponse) Descriptor() ([]byte, []int) {
	return file_mfa_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateOTPTokenResponse) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type ValidateOTPTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token        string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Principal    string `protobuf:"bytes,2,opt,name=principal,proto3" json:"principal,omitempty"`
	ClientSecret string `protobuf:"bytes,3,opt,name=client_secret,json=clientSecret,proto3" json:"client_secret,omitempty"`
}

func (x *ValidateOTPTokenRequest) Reset() {
	*x = ValidateOTPTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mfa_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateOTPTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateOTPTokenRequest) ProtoMessage() {}

func (x *ValidateOTPTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mfa_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateOTPTokenRequest.ProtoReflect.Descriptor instead.
func (*ValidateOTPTokenRequest) Descriptor() ([]byte, []int) {
	return file_mfa_proto_rawDescGZIP(), []int{2}
}

func (x *ValidateOTPTokenRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *ValidateOTPTokenRequest) GetPrincipal() string {
	if x != nil {
		return x.Principal
	}
	return ""
}

func (x *ValidateOTPTokenRequest) GetClientSecret() string {
	if x != nil {
		return x.ClientSecret
	}
	return ""
}

type ValidateOTPTokenResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IsValid bool `protobuf:"varint,1,opt,name=is_valid,json=isValid,proto3" json:"is_valid"`
}

func (x *ValidateOTPTokenResponse) Reset() {
	*x = ValidateOTPTokenResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_mfa_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ValidateOTPTokenResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateOTPTokenResponse) ProtoMessage() {}

func (x *ValidateOTPTokenResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mfa_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateOTPTokenResponse.ProtoReflect.Descriptor instead.
func (*ValidateOTPTokenResponse) Descriptor() ([]byte, []int) {
	return file_mfa_proto_rawDescGZIP(), []int{3}
}

func (x *ValidateOTPTokenResponse) GetIsValid() bool {
	if x != nil {
		return x.IsValid
	}
	return false
}

var File_mfa_proto protoreflect.FileDescriptor

var file_mfa_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6d, 0x66, 0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x5c, 0x0a, 0x17, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69,
	0x70, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x69, 0x6e, 0x63,
	0x69, 0x70, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73,
	0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69,
	0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22, 0x30, 0x0a, 0x18, 0x47, 0x65, 0x6e,
	0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x72, 0x0a, 0x17, 0x56,
	0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1c, 0x0a, 0x09,
	0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x70, 0x72, 0x69, 0x6e, 0x63, 0x69, 0x70, 0x61, 0x6c, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x6c,
	0x69, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0c, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x22,
	0x35, 0x0a, 0x18, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x19, 0x0a, 0x08, 0x69,
	0x73, 0x5f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x69,
	0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x32, 0xa2, 0x01, 0x0a, 0x0a, 0x4d, 0x46, 0x41, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x10, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x4f, 0x54, 0x50, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x4f, 0x54,
	0x50, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x49, 0x0a, 0x10, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x18, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f,
	0x54, 0x50, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19,
	0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65, 0x4f, 0x54, 0x50, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x3d, 0x5a, 0x3b, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x49, 0x4e, 0x44, 0x49, 0x43, 0x4f,
	0x2d, 0x49, 0x4e, 0x4e, 0x4f, 0x56, 0x41, 0x54, 0x49, 0x4f, 0x4e, 0x2f, 0x69, 0x6e, 0x64, 0x69,
	0x63, 0x6f, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x61, 0x75, 0x74, 0x68, 0x2f,
	0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2f, 0x6d, 0x66, 0x61, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_mfa_proto_rawDescOnce sync.Once
	file_mfa_proto_rawDescData = file_mfa_proto_rawDesc
)

func file_mfa_proto_rawDescGZIP() []byte {
	file_mfa_proto_rawDescOnce.Do(func() {
		file_mfa_proto_rawDescData = protoimpl.X.CompressGZIP(file_mfa_proto_rawDescData)
	})
	return file_mfa_proto_rawDescData
}

var file_mfa_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_mfa_proto_goTypes = []interface{}{
	(*GenerateOTPTokenRequest)(nil),  // 0: GenerateOTPTokenRequest
	(*GenerateOTPTokenResponse)(nil), // 1: GenerateOTPTokenResponse
	(*ValidateOTPTokenRequest)(nil),  // 2: ValidateOTPTokenRequest
	(*ValidateOTPTokenResponse)(nil), // 3: ValidateOTPTokenResponse
}
var file_mfa_proto_depIdxs = []int32{
	0, // 0: MFAService.GenerateOTPToken:input_type -> GenerateOTPTokenRequest
	2, // 1: MFAService.ValidateOTPToken:input_type -> ValidateOTPTokenRequest
	1, // 2: MFAService.GenerateOTPToken:output_type -> GenerateOTPTokenResponse
	3, // 3: MFAService.ValidateOTPToken:output_type -> ValidateOTPTokenResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_mfa_proto_init() }
func file_mfa_proto_init() {
	if File_mfa_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_mfa_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateOTPTokenRequest); i {
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
		file_mfa_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateOTPTokenResponse); i {
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
		file_mfa_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateOTPTokenRequest); i {
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
		file_mfa_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ValidateOTPTokenResponse); i {
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
			RawDescriptor: file_mfa_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mfa_proto_goTypes,
		DependencyIndexes: file_mfa_proto_depIdxs,
		MessageInfos:      file_mfa_proto_msgTypes,
	}.Build()
	File_mfa_proto = out.File
	file_mfa_proto_rawDesc = nil
	file_mfa_proto_goTypes = nil
	file_mfa_proto_depIdxs = nil
}
