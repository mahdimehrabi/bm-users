// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: src/apps/grpc/proto/auth/auth.proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_src_apps_grpc_proto_auth_auth_proto_rawDescGZIP(), []int{0}
}

type IDReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID int64 `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
}

func (x *IDReq) Reset() {
	*x = IDReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IDReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IDReq) ProtoMessage() {}

func (x *IDReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IDReq.ProtoReflect.Descriptor instead.
func (*IDReq) Descriptor() ([]byte, []int) {
	return file_src_apps_grpc_proto_auth_auth_proto_rawDescGZIP(), []int{1}
}

func (x *IDReq) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

type TokenReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token string `protobuf:"bytes,1,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *TokenReq) Reset() {
	*x = TokenReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenReq) ProtoMessage() {}

func (x *TokenReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenReq.ProtoReflect.Descriptor instead.
func (*TokenReq) Descriptor() ([]byte, []int) {
	return file_src_apps_grpc_proto_auth_auth_proto_rawDescGZIP(), []int{2}
}

func (x *TokenReq) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type Credits struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ID       int64  `protobuf:"varint,1,opt,name=ID,proto3" json:"ID,omitempty"`
	Email    string `protobuf:"bytes,2,opt,name=Email,proto3" json:"Email,omitempty"`
	Fullname string `protobuf:"bytes,3,opt,name=Fullname,proto3" json:"Fullname,omitempty"`
}

func (x *Credits) Reset() {
	*x = Credits{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Credits) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Credits) ProtoMessage() {}

func (x *Credits) ProtoReflect() protoreflect.Message {
	mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Credits.ProtoReflect.Descriptor instead.
func (*Credits) Descriptor() ([]byte, []int) {
	return file_src_apps_grpc_proto_auth_auth_proto_rawDescGZIP(), []int{3}
}

func (x *Credits) GetID() int64 {
	if x != nil {
		return x.ID
	}
	return 0
}

func (x *Credits) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *Credits) GetFullname() string {
	if x != nil {
		return x.Fullname
	}
	return ""
}

type LoginReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=Username,proto3" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
}

func (x *LoginReq) Reset() {
	*x = LoginReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LoginReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LoginReq) ProtoMessage() {}

func (x *LoginReq) ProtoReflect() protoreflect.Message {
	mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LoginReq.ProtoReflect.Descriptor instead.
func (*LoginReq) Descriptor() ([]byte, []int) {
	return file_src_apps_grpc_proto_auth_auth_proto_rawDescGZIP(), []int{4}
}

func (x *LoginReq) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *LoginReq) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type TokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken     string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
	RefreshToken    string `protobuf:"bytes,2,opt,name=RefreshToken,proto3" json:"RefreshToken,omitempty"`
	AccessTokenExp  int64  `protobuf:"varint,3,opt,name=AccessTokenExp,proto3" json:"AccessTokenExp,omitempty"`
	RefreshTokenExp int64  `protobuf:"varint,4,opt,name=RefreshTokenExp,proto3" json:"RefreshTokenExp,omitempty"`
}

func (x *TokensResponse) Reset() {
	*x = TokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokensResponse) ProtoMessage() {}

func (x *TokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_src_apps_grpc_proto_auth_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokensResponse.ProtoReflect.Descriptor instead.
func (*TokensResponse) Descriptor() ([]byte, []int) {
	return file_src_apps_grpc_proto_auth_auth_proto_rawDescGZIP(), []int{5}
}

func (x *TokensResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokensResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *TokensResponse) GetAccessTokenExp() int64 {
	if x != nil {
		return x.AccessTokenExp
	}
	return 0
}

func (x *TokensResponse) GetRefreshTokenExp() int64 {
	if x != nil {
		return x.RefreshTokenExp
	}
	return 0
}

var File_src_apps_grpc_proto_auth_auth_proto protoreflect.FileDescriptor

var file_src_apps_grpc_proto_auth_auth_proto_rawDesc = []byte{
	0x0a, 0x23, 0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x61, 0x75, 0x74, 0x68, 0x22, 0x07, 0x0a, 0x05, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x17, 0x0a, 0x05, 0x49, 0x44, 0x52, 0x65, 0x71, 0x12, 0x0e, 0x0a,
	0x02, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x22, 0x20, 0x0a,
	0x08, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22,
	0x4b, 0x0a, 0x07, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x49, 0x44, 0x12, 0x14, 0x0a, 0x05, 0x45, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x45, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x1a, 0x0a, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x46, 0x75, 0x6c, 0x6c, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x42, 0x0a, 0x08,
	0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x12, 0x1a, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64,
	0x22, 0xa8, 0x01, 0x0a, 0x0e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78,
	0x70, 0x12, 0x28, 0x0a, 0x0f, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x45, 0x78, 0x70, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0f, 0x52, 0x65, 0x66, 0x72,
	0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x45, 0x78, 0x70, 0x32, 0xc5, 0x01, 0x0a, 0x0e,
	0x41, 0x75, 0x74, 0x68, 0x65, 0x6e, 0x74, 0x69, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d,
	0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x4c,
	0x6f, 0x67, 0x69, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2b, 0x0a,
	0x0a, 0x47, 0x65, 0x74, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x0e, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x0d, 0x2e, 0x61, 0x75,
	0x74, 0x68, 0x2e, 0x43, 0x72, 0x65, 0x64, 0x69, 0x74, 0x73, 0x12, 0x32, 0x0a, 0x0a, 0x52, 0x65,
	0x6e, 0x65, 0x77, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x0e, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x1a, 0x14, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x07, 0x42, 0x61, 0x6e, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68,
	0x2e, 0x49, 0x44, 0x52, 0x65, 0x71, 0x1a, 0x0b, 0x2e, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x42, 0x23, 0x5a, 0x21, 0x62, 0x6d, 0x2d, 0x75, 0x73, 0x65, 0x72, 0x73, 0x2f,
	0x73, 0x72, 0x63, 0x2f, 0x61, 0x70, 0x70, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_apps_grpc_proto_auth_auth_proto_rawDescOnce sync.Once
	file_src_apps_grpc_proto_auth_auth_proto_rawDescData = file_src_apps_grpc_proto_auth_auth_proto_rawDesc
)

func file_src_apps_grpc_proto_auth_auth_proto_rawDescGZIP() []byte {
	file_src_apps_grpc_proto_auth_auth_proto_rawDescOnce.Do(func() {
		file_src_apps_grpc_proto_auth_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_apps_grpc_proto_auth_auth_proto_rawDescData)
	})
	return file_src_apps_grpc_proto_auth_auth_proto_rawDescData
}

var file_src_apps_grpc_proto_auth_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_src_apps_grpc_proto_auth_auth_proto_goTypes = []interface{}{
	(*Empty)(nil),          // 0: auth.Empty
	(*IDReq)(nil),          // 1: auth.IDReq
	(*TokenReq)(nil),       // 2: auth.TokenReq
	(*Credits)(nil),        // 3: auth.Credits
	(*LoginReq)(nil),       // 4: auth.LoginReq
	(*TokensResponse)(nil), // 5: auth.TokensResponse
}
var file_src_apps_grpc_proto_auth_auth_proto_depIdxs = []int32{
	4, // 0: auth.Authentication.Login:input_type -> auth.LoginReq
	2, // 1: auth.Authentication.GetCredits:input_type -> auth.TokenReq
	2, // 2: auth.Authentication.RenewToken:input_type -> auth.TokenReq
	1, // 3: auth.Authentication.BanUser:input_type -> auth.IDReq
	5, // 4: auth.Authentication.Login:output_type -> auth.TokensResponse
	3, // 5: auth.Authentication.GetCredits:output_type -> auth.Credits
	5, // 6: auth.Authentication.RenewToken:output_type -> auth.TokensResponse
	0, // 7: auth.Authentication.BanUser:output_type -> auth.Empty
	4, // [4:8] is the sub-list for method output_type
	0, // [0:4] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_src_apps_grpc_proto_auth_auth_proto_init() }
func file_src_apps_grpc_proto_auth_auth_proto_init() {
	if File_src_apps_grpc_proto_auth_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_apps_grpc_proto_auth_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_src_apps_grpc_proto_auth_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IDReq); i {
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
		file_src_apps_grpc_proto_auth_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenReq); i {
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
		file_src_apps_grpc_proto_auth_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Credits); i {
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
		file_src_apps_grpc_proto_auth_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LoginReq); i {
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
		file_src_apps_grpc_proto_auth_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokensResponse); i {
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
			RawDescriptor: file_src_apps_grpc_proto_auth_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_src_apps_grpc_proto_auth_auth_proto_goTypes,
		DependencyIndexes: file_src_apps_grpc_proto_auth_auth_proto_depIdxs,
		MessageInfos:      file_src_apps_grpc_proto_auth_auth_proto_msgTypes,
	}.Build()
	File_src_apps_grpc_proto_auth_auth_proto = out.File
	file_src_apps_grpc_proto_auth_auth_proto_rawDesc = nil
	file_src_apps_grpc_proto_auth_auth_proto_goTypes = nil
	file_src_apps_grpc_proto_auth_auth_proto_depIdxs = nil
}
