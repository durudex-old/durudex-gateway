// Copyright © 2021 Durudex

// This file is part of Durudex: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.

// Durudex is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Affero General Public License for more details.

// You should have received a copy of the GNU Affero General Public License
// along with Durudex. If not, see <https://www.gnu.org/licenses/>.

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: internal/delivery/grpc/protobuf/auth.proto

package protobuf

import (
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

type UserSignUpRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string                 `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Name     string                 `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Email    string                 `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Password string                 `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	Birthday *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=birthday,proto3" json:"birthday,omitempty"`
	Sex      int32                  `protobuf:"varint,6,opt,name=sex,proto3" json:"sex,omitempty"`
}

func (x *UserSignUpRequest) Reset() {
	*x = UserSignUpRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignUpRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignUpRequest) ProtoMessage() {}

func (x *UserSignUpRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignUpRequest.ProtoReflect.Descriptor instead.
func (*UserSignUpRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{0}
}

func (x *UserSignUpRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserSignUpRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserSignUpRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserSignUpRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *UserSignUpRequest) GetBirthday() *timestamppb.Timestamp {
	if x != nil {
		return x.Birthday
	}
	return nil
}

func (x *UserSignUpRequest) GetSex() int32 {
	if x != nil {
		return x.Sex
	}
	return 0
}

type UserSignUpResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserSignUpResponse) Reset() {
	*x = UserSignUpResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignUpResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignUpResponse) ProtoMessage() {}

func (x *UserSignUpResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignUpResponse.ProtoReflect.Descriptor instead.
func (*UserSignUpResponse) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{1}
}

func (x *UserSignUpResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type UserSignInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Username string `protobuf:"bytes,1,opt,name=username,proto3" json:"username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
}

func (x *UserSignInRequest) Reset() {
	*x = UserSignInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignInRequest) ProtoMessage() {}

func (x *UserSignInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignInRequest.ProtoReflect.Descriptor instead.
func (*UserSignInRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{2}
}

func (x *UserSignInRequest) GetUsername() string {
	if x != nil {
		return x.Username
	}
	return ""
}

func (x *UserSignInRequest) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

type UserSignInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *UserSignInResponse) Reset() {
	*x = UserSignInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserSignInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserSignInResponse) ProtoMessage() {}

func (x *UserSignInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserSignInResponse.ProtoReflect.Descriptor instead.
func (*UserSignInResponse) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{3}
}

func (x *UserSignInResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UserSignInResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type UserRefreshTokensRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RefreshToken string `protobuf:"bytes,1,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *UserRefreshTokensRequest) Reset() {
	*x = UserRefreshTokensRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRefreshTokensRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRefreshTokensRequest) ProtoMessage() {}

func (x *UserRefreshTokensRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRefreshTokensRequest.ProtoReflect.Descriptor instead.
func (*UserRefreshTokensRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{4}
}

func (x *UserRefreshTokensRequest) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type UserRefreshTokensResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AccessToken  string `protobuf:"bytes,1,opt,name=accessToken,proto3" json:"accessToken,omitempty"`
	RefreshToken string `protobuf:"bytes,2,opt,name=refreshToken,proto3" json:"refreshToken,omitempty"`
}

func (x *UserRefreshTokensResponse) Reset() {
	*x = UserRefreshTokensResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserRefreshTokensResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserRefreshTokensResponse) ProtoMessage() {}

func (x *UserRefreshTokensResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserRefreshTokensResponse.ProtoReflect.Descriptor instead.
func (*UserRefreshTokensResponse) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{5}
}

func (x *UserRefreshTokensResponse) GetAccessToken() string {
	if x != nil {
		return x.AccessToken
	}
	return ""
}

func (x *UserRefreshTokensResponse) GetRefreshToken() string {
	if x != nil {
		return x.RefreshToken
	}
	return ""
}

type VerifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Code int32  `protobuf:"varint,2,opt,name=code,proto3" json:"code,omitempty"`
}

func (x *VerifyRequest) Reset() {
	*x = VerifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyRequest) ProtoMessage() {}

func (x *VerifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyRequest.ProtoReflect.Descriptor instead.
func (*VerifyRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *VerifyRequest) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

type Status struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status bool `protobuf:"varint,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *Status) Reset() {
	*x = Status{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Status) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Status) ProtoMessage() {}

func (x *Status) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Status.ProtoReflect.Descriptor instead.
func (*Status) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{7}
}

func (x *Status) GetStatus() bool {
	if x != nil {
		return x.Status
	}
	return false
}

type GetVerifyCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *GetVerifyCodeRequest) Reset() {
	*x = GetVerifyCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetVerifyCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetVerifyCodeRequest) ProtoMessage() {}

func (x *GetVerifyCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetVerifyCodeRequest.ProtoReflect.Descriptor instead.
func (*GetVerifyCodeRequest) Descriptor() ([]byte, []int) {
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP(), []int{8}
}

func (x *GetVerifyCodeRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *GetVerifyCodeRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_internal_delivery_grpc_protobuf_auth_proto protoreflect.FileDescriptor

var file_internal_delivery_grpc_protobuf_auth_proto_rawDesc = []byte{
	0x0a, 0x2a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x64, 0x65, 0x6c, 0x69, 0x76,
	0x65, 0x72, 0x79, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x75, 0x74, 0x68, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbf, 0x01, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x12, 0x36,
	0x0a, 0x08, 0x62, 0x69, 0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62,
	0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52, 0x08, 0x62, 0x69,
	0x72, 0x74, 0x68, 0x64, 0x61, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x78, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x03, 0x73, 0x65, 0x78, 0x22, 0x24, 0x0a, 0x12, 0x55, 0x73, 0x65, 0x72,
	0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x4b,
	0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x70, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72, 0x64, 0x22, 0x5a, 0x0a, 0x12, 0x55,
	0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x3e, 0x0a, 0x18, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65,
	0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x61, 0x0a, 0x19, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73, 0x73, 0x54, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x61, 0x63, 0x63, 0x65, 0x73,
	0x73, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x65, 0x66, 0x72, 0x65, 0x73,
	0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x65,
	0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x33, 0x0a, 0x0d, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x22,
	0x20, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x22, 0x40, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f,
	0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61,
	0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x32, 0xdf, 0x02, 0x0a, 0x0b, 0x41, 0x75, 0x74, 0x68, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x12, 0x1a, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e,
	0x55, 0x70, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x55, 0x70, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x41, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x49, 0x6e,
	0x12, 0x1a, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53,
	0x69, 0x67, 0x6e, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1b, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x53, 0x69, 0x67, 0x6e, 0x49,
	0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x31, 0x0a, 0x06, 0x56, 0x65, 0x72,
	0x69, 0x66, 0x79, 0x12, 0x16, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x56, 0x65,
	0x72, 0x69, 0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x3f, 0x0a, 0x0d,
	0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x47, 0x65, 0x74, 0x56, 0x65, 0x72, 0x69, 0x66,
	0x79, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x67,
	0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x56, 0x0a,
	0x0d, 0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x12, 0x21,
	0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x66,
	0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x22, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x52, 0x65, 0x66, 0x72, 0x65, 0x73, 0x68, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x2d, 0x5a, 0x2b, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x44, 0x75, 0x72, 0x75, 0x64, 0x65, 0x78, 0x2f, 0x64, 0x75, 0x72, 0x75,
	0x64, 0x65, 0x78, 0x2d, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x3b, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_delivery_grpc_protobuf_auth_proto_rawDescOnce sync.Once
	file_internal_delivery_grpc_protobuf_auth_proto_rawDescData = file_internal_delivery_grpc_protobuf_auth_proto_rawDesc
)

func file_internal_delivery_grpc_protobuf_auth_proto_rawDescGZIP() []byte {
	file_internal_delivery_grpc_protobuf_auth_proto_rawDescOnce.Do(func() {
		file_internal_delivery_grpc_protobuf_auth_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_delivery_grpc_protobuf_auth_proto_rawDescData)
	})
	return file_internal_delivery_grpc_protobuf_auth_proto_rawDescData
}

var file_internal_delivery_grpc_protobuf_auth_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_internal_delivery_grpc_protobuf_auth_proto_goTypes = []interface{}{
	(*UserSignUpRequest)(nil),         // 0: gateway.UserSignUpRequest
	(*UserSignUpResponse)(nil),        // 1: gateway.UserSignUpResponse
	(*UserSignInRequest)(nil),         // 2: gateway.UserSignInRequest
	(*UserSignInResponse)(nil),        // 3: gateway.UserSignInResponse
	(*UserRefreshTokensRequest)(nil),  // 4: gateway.UserRefreshTokensRequest
	(*UserRefreshTokensResponse)(nil), // 5: gateway.UserRefreshTokensResponse
	(*VerifyRequest)(nil),             // 6: gateway.VerifyRequest
	(*Status)(nil),                    // 7: gateway.Status
	(*GetVerifyCodeRequest)(nil),      // 8: gateway.GetVerifyCodeRequest
	(*timestamppb.Timestamp)(nil),     // 9: google.protobuf.Timestamp
}
var file_internal_delivery_grpc_protobuf_auth_proto_depIdxs = []int32{
	9, // 0: gateway.UserSignUpRequest.birthday:type_name -> google.protobuf.Timestamp
	0, // 1: gateway.AuthService.SignUp:input_type -> gateway.UserSignUpRequest
	2, // 2: gateway.AuthService.SignIn:input_type -> gateway.UserSignInRequest
	6, // 3: gateway.AuthService.Verify:input_type -> gateway.VerifyRequest
	8, // 4: gateway.AuthService.GetVerifyCode:input_type -> gateway.GetVerifyCodeRequest
	4, // 5: gateway.AuthService.RefreshTokens:input_type -> gateway.UserRefreshTokensRequest
	1, // 6: gateway.AuthService.SignUp:output_type -> gateway.UserSignUpResponse
	3, // 7: gateway.AuthService.SignIn:output_type -> gateway.UserSignInResponse
	7, // 8: gateway.AuthService.Verify:output_type -> gateway.Status
	7, // 9: gateway.AuthService.GetVerifyCode:output_type -> gateway.Status
	5, // 10: gateway.AuthService.RefreshTokens:output_type -> gateway.UserRefreshTokensResponse
	6, // [6:11] is the sub-list for method output_type
	1, // [1:6] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_internal_delivery_grpc_protobuf_auth_proto_init() }
func file_internal_delivery_grpc_protobuf_auth_proto_init() {
	if File_internal_delivery_grpc_protobuf_auth_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignUpRequest); i {
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
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignUpResponse); i {
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
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignInRequest); i {
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
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserSignInResponse); i {
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
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRefreshTokensRequest); i {
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
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserRefreshTokensResponse); i {
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
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyRequest); i {
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
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Status); i {
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
		file_internal_delivery_grpc_protobuf_auth_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetVerifyCodeRequest); i {
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
			RawDescriptor: file_internal_delivery_grpc_protobuf_auth_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_delivery_grpc_protobuf_auth_proto_goTypes,
		DependencyIndexes: file_internal_delivery_grpc_protobuf_auth_proto_depIdxs,
		MessageInfos:      file_internal_delivery_grpc_protobuf_auth_proto_msgTypes,
	}.Build()
	File_internal_delivery_grpc_protobuf_auth_proto = out.File
	file_internal_delivery_grpc_protobuf_auth_proto_rawDesc = nil
	file_internal_delivery_grpc_protobuf_auth_proto_goTypes = nil
	file_internal_delivery_grpc_protobuf_auth_proto_depIdxs = nil
}
