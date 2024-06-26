// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.1
// source: jwt.proto

package jwtpb

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type GenerateTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId   int32  `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Username string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *GenerateTokensRequest) Reset() {
	*x = GenerateTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokensRequest) ProtoMessage() {}

func (x *GenerateTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokensRequest.ProtoReflect.Descriptor instead.
func (*GenerateTokensRequest) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{0}
}

func (x *GenerateTokensRequest) GetUserId() int32 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *GenerateTokensRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type GenerateTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair *TokenPair `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
}

func (x *GenerateTokensResponse) Reset() {
	*x = GenerateTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenerateTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenerateTokensResponse) ProtoMessage() {}

func (x *GenerateTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenerateTokensResponse.ProtoReflect.Descriptor instead.
func (*GenerateTokensResponse) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{1}
}

func (x *GenerateTokensResponse) GetPair() *TokenPair {
	if x != nil {
		return x.Pair
	}
	return nil
}

type RefreshTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
	Username     string `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
}

func (x *RefreshTokensRequest) Reset() {
	*x = RefreshTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokensRequest) ProtoMessage() {}

func (x *RefreshTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokensRequest.ProtoReflect.Descriptor instead.
func (*RefreshTokensRequest) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{2}
}

func (x *RefreshTokensRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

func (x *RefreshTokensRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

type RefreshTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pair *TokenPair `protobuf:"bytes,1,opt,name=pair,proto3" json:"pair,omitempty"`
}

func (x *RefreshTokensResponse) Reset() {
	*x = RefreshTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RefreshTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RefreshTokensResponse) ProtoMessage() {}

func (x *RefreshTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RefreshTokensResponse.ProtoReflect.Descriptor instead.
func (*RefreshTokensResponse) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{3}
}

func (x *RefreshTokensResponse) GetPair() *TokenPair {
	if x != nil {
		return x.Pair
	}
	return nil
}

type RemoveAllTokensForUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
}

func (x *RemoveAllTokensForUserRequest) Reset() {
	*x = RemoveAllTokensForUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RemoveAllTokensForUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RemoveAllTokensForUserRequest) ProtoMessage() {}

func (x *RemoveAllTokensForUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RemoveAllTokensForUserRequest.ProtoReflect.Descriptor instead.
func (*RemoveAllTokensForUserRequest) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{4}
}

func (x *RemoveAllTokensForUserRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type IsValidAccessTokenRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken string `protobuf:"bytes,1,opt,name=AccessToken,proto3" json:"AccessToken,omitempty"`
}

func (x *IsValidAccessTokenRequest) Reset() {
	*x = IsValidAccessTokenRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *IsValidAccessTokenRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*IsValidAccessTokenRequest) ProtoMessage() {}

func (x *IsValidAccessTokenRequest) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use IsValidAccessTokenRequest.ProtoReflect.Descriptor instead.
func (*IsValidAccessTokenRequest) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{5}
}

func (x *IsValidAccessTokenRequest) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

type TokenPair struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=access_token,json=accessToken,proto3" json:"access_token,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refresh_token,json=refreshToken,proto3" json:"refresh_token,omitempty"`
}

func (x *TokenPair) Reset() {
	*x = TokenPair{}
	if protoimpl.UnsafeEnabled {
		mi := &file_jwt_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TokenPair) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TokenPair) ProtoMessage() {}

func (x *TokenPair) ProtoReflect() protoreflect.Message {
	mi := &file_jwt_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TokenPair.ProtoReflect.Descriptor instead.
func (*TokenPair) Descriptor() ([]byte, []int) {
	return file_jwt_proto_rawDescGZIP(), []int{6}
}

func (x *TokenPair) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *TokenPair) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

var File_jwt_proto protoreflect.FileDescriptor

var file_jwt_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6a, 0x77, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x03, 0x6a, 0x77, 0x74,
	0x1a, 0x1b, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a,
	0x15, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3c, 0x0a, 0x16, 0x47,
	0x65, 0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x70, 0x61, 0x69, 0x72, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50,
	0x61, 0x69, 0x72, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x22, 0x57, 0x0a, 0x14, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3b, 0x0a, 0x15, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x22, 0x0a, 0x04, 0x70,
	0x61, 0x69, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6a, 0x77, 0x74, 0x2e,
	0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x69, 0x72, 0x52, 0x04, 0x70, 0x61, 0x69, 0x72, 0x22,
	0x41, 0x0a, 0x1d, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b,
	0x65, 0x6e, 0x22, 0x3d, 0x0a, 0x19, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x20, 0x0a, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x22, 0x53, 0x0a, 0x09, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x50, 0x61, 0x69, 0x72, 0x12, 0x21,
	0x0a, 0x0c, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x5f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65,
	0x6e, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x5f, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x32, 0xc3, 0x02, 0x0a, 0x0a, 0x4a, 0x77, 0x74, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x74,
	0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x1a, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x47, 0x65,
	0x6e, 0x65, 0x72, 0x61, 0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x65, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x46, 0x0a, 0x0d, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x73, 0x12, 0x19, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x6a,
	0x77, 0x74, 0x2e, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x54, 0x0a, 0x16, 0x52, 0x65, 0x6d, 0x6f,
	0x76, 0x65, 0x41, 0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73,
	0x65, 0x72, 0x12, 0x22, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x52, 0x65, 0x6d, 0x6f, 0x76, 0x65, 0x41,
	0x6c, 0x6c, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x46, 0x6f, 0x72, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x12, 0x4c,
	0x0a, 0x12, 0x49, 0x73, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x1e, 0x2e, 0x6a, 0x77, 0x74, 0x2e, 0x49, 0x73, 0x56, 0x61, 0x6c,
	0x69, 0x64, 0x41, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x42, 0x08, 0x5a, 0x06,
	0x2f, 0x6a, 0x77, 0x74, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_jwt_proto_rawDescOnce sync.Once
	file_jwt_proto_rawDescData = file_jwt_proto_rawDesc
)

func file_jwt_proto_rawDescGZIP() []byte {
	file_jwt_proto_rawDescOnce.Do(func() {
		file_jwt_proto_rawDescData = protoimpl.X.CompressGZIP(file_jwt_proto_rawDescData)
	})
	return file_jwt_proto_rawDescData
}

var file_jwt_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_jwt_proto_goTypes = []interface{}{
	(*GenerateTokensRequest)(nil),         // 0: jwt.GenerateTokensRequest
	(*GenerateTokensResponse)(nil),        // 1: jwt.GenerateTokensResponse
	(*RefreshTokensRequest)(nil),          // 2: jwt.RefreshTokensRequest
	(*RefreshTokensResponse)(nil),         // 3: jwt.RefreshTokensResponse
	(*RemoveAllTokensForUserRequest)(nil), // 4: jwt.RemoveAllTokensForUserRequest
	(*IsValidAccessTokenRequest)(nil),     // 5: jwt.IsValidAccessTokenRequest
	(*TokenPair)(nil),                     // 6: jwt.TokenPair
	(*emptypb.Empty)(nil),                 // 7: google.protobuf.Empty
}
var file_jwt_proto_depIdxs = []int32{
	6, // 0: jwt.GenerateTokensResponse.pair:type_name -> jwt.TokenPair
	6, // 1: jwt.RefreshTokensResponse.pair:type_name -> jwt.TokenPair
	0, // 2: jwt.JwtService.GenerateTokens:input_type -> jwt.GenerateTokensRequest
	2, // 3: jwt.JwtService.RefreshTokens:input_type -> jwt.RefreshTokensRequest
	4, // 4: jwt.JwtService.RemoveAllTokensForUser:input_type -> jwt.RemoveAllTokensForUserRequest
	5, // 5: jwt.JwtService.IsValidAccessToken:input_type -> jwt.IsValidAccessTokenRequest
	1, // 6: jwt.JwtService.GenerateTokens:output_type -> jwt.GenerateTokensResponse
	3, // 7: jwt.JwtService.RefreshTokens:output_type -> jwt.RefreshTokensResponse
	7, // 8: jwt.JwtService.RemoveAllTokensForUser:output_type -> google.protobuf.Empty
	7, // 9: jwt.JwtService.IsValidAccessToken:output_type -> google.protobuf.Empty
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_jwt_proto_init() }
func file_jwt_proto_init() {
	if File_jwt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_jwt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokensRequest); i {
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
		file_jwt_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenerateTokensResponse); i {
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
		file_jwt_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokensRequest); i {
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
		file_jwt_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RefreshTokensResponse); i {
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
		file_jwt_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RemoveAllTokensForUserRequest); i {
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
		file_jwt_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*IsValidAccessTokenRequest); i {
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
		file_jwt_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TokenPair); i {
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
			RawDescriptor: file_jwt_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_jwt_proto_goTypes,
		DependencyIndexes: file_jwt_proto_depIdxs,
		MessageInfos:      file_jwt_proto_msgTypes,
	}.Build()
	File_jwt_proto = out.File
	file_jwt_proto_rawDesc = nil
	file_jwt_proto_goTypes = nil
	file_jwt_proto_depIdxs = nil
}
