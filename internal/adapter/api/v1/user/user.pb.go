// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: internal/adapter/api/v1/user/user.proto

package user

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	proto "github.com/patrickkoss/grpc-gateway-example/internal/adapter/api/domain/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type UserId struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UserId) Reset() {
	*x = UserId{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserId) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserId) ProtoMessage() {}

func (x *UserId) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserId.ProtoReflect.Descriptor instead.
func (*UserId) Descriptor() ([]byte, []int) {
	return file_internal_adapter_api_v1_user_user_proto_rawDescGZIP(), []int{0}
}

func (x *UserId) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type Aggregation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Eq   *string `protobuf:"bytes,1,opt,name=eq,proto3,oneof" json:"eq,omitempty"`
	Like *string `protobuf:"bytes,2,opt,name=like,proto3,oneof" json:"like,omitempty"`
}

func (x *Aggregation) Reset() {
	*x = Aggregation{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Aggregation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Aggregation) ProtoMessage() {}

func (x *Aggregation) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Aggregation.ProtoReflect.Descriptor instead.
func (*Aggregation) Descriptor() ([]byte, []int) {
	return file_internal_adapter_api_v1_user_user_proto_rawDescGZIP(), []int{1}
}

func (x *Aggregation) GetEq() string {
	if x != nil && x.Eq != nil {
		return *x.Eq
	}
	return ""
}

func (x *Aggregation) GetLike() string {
	if x != nil && x.Like != nil {
		return *x.Like
	}
	return ""
}

type Page struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Size *string `protobuf:"bytes,1,opt,name=size,proto3,oneof" json:"size,omitempty"`
	Eq   *string `protobuf:"bytes,2,opt,name=eq,proto3,oneof" json:"eq,omitempty"`
}

func (x *Page) Reset() {
	*x = Page{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Page) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Page) ProtoMessage() {}

func (x *Page) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Page.ProtoReflect.Descriptor instead.
func (*Page) Descriptor() ([]byte, []int) {
	return file_internal_adapter_api_v1_user_user_proto_rawDescGZIP(), []int{2}
}

func (x *Page) GetSize() string {
	if x != nil && x.Size != nil {
		return *x.Size
	}
	return ""
}

func (x *Page) GetEq() string {
	if x != nil && x.Eq != nil {
		return *x.Eq
	}
	return ""
}

type UserParams struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Page *Page        `protobuf:"bytes,1,opt,name=page,proto3,oneof" json:"page,omitempty"`
	Name *Aggregation `protobuf:"bytes,3,opt,name=name,proto3,oneof" json:"name,omitempty"`
}

func (x *UserParams) Reset() {
	*x = UserParams{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserParams) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserParams) ProtoMessage() {}

func (x *UserParams) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserParams.ProtoReflect.Descriptor instead.
func (*UserParams) Descriptor() ([]byte, []int) {
	return file_internal_adapter_api_v1_user_user_proto_rawDescGZIP(), []int{3}
}

func (x *UserParams) GetPage() *Page {
	if x != nil {
		return x.Page
	}
	return nil
}

func (x *UserParams) GetName() *Aggregation {
	if x != nil {
		return x.Name
	}
	return nil
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id      string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name    string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,3,opt,name=surname,proto3" json:"surname,omitempty"`
	Email   string `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_internal_adapter_api_v1_user_user_proto_rawDescGZIP(), []int{4}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateUserBody struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name    string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Surname string `protobuf:"bytes,2,opt,name=surname,proto3" json:"surname,omitempty"`
	Email   string `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateUserBody) Reset() {
	*x = CreateUserBody{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserBody) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserBody) ProtoMessage() {}

func (x *CreateUserBody) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserBody.ProtoReflect.Descriptor instead.
func (*CreateUserBody) Descriptor() ([]byte, []int) {
	return file_internal_adapter_api_v1_user_user_proto_rawDescGZIP(), []int{5}
}

func (x *CreateUserBody) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateUserBody) GetSurname() string {
	if x != nil {
		return x.Surname
	}
	return ""
}

func (x *CreateUserBody) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type ListUsersResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Users []*User `protobuf:"bytes,1,rep,name=users,proto3" json:"users,omitempty"`
}

func (x *ListUsersResponse) Reset() {
	*x = ListUsersResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ListUsersResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ListUsersResponse) ProtoMessage() {}

func (x *ListUsersResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ListUsersResponse.ProtoReflect.Descriptor instead.
func (*ListUsersResponse) Descriptor() ([]byte, []int) {
	return file_internal_adapter_api_v1_user_user_proto_rawDescGZIP(), []int{6}
}

func (x *ListUsersResponse) GetUsers() []*User {
	if x != nil {
		return x.Users
	}
	return nil
}

type GetUserResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	User *User `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
}

func (x *GetUserResponse) Reset() {
	*x = GetUserResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetUserResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetUserResponse) ProtoMessage() {}

func (x *GetUserResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_adapter_api_v1_user_user_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetUserResponse.ProtoReflect.Descriptor instead.
func (*GetUserResponse) Descriptor() ([]byte, []int) {
	return file_internal_adapter_api_v1_user_user_proto_rawDescGZIP(), []int{7}
}

func (x *GetUserResponse) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

var File_internal_adapter_api_v1_user_user_proto protoreflect.FileDescriptor

var file_internal_adapter_api_v1_user_user_proto_rawDesc = []byte{
	0x0a, 0x27, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74,
	0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x31, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x1a, 0x2f, 0x69, 0x6e,
	0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x61, 0x70, 0x69, 0x73,
	0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x43, 0x69,
	0x6e, 0x63, 0x6c, 0x75, 0x64, 0x65, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d, 0x67, 0x65, 0x6e, 0x2d, 0x6f,
	0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x1a, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x64, 0x61,
	0x70, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x22, 0x29, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0f, 0x92, 0x41, 0x0c, 0xca, 0x3e,
	0x09, 0xfa, 0x02, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x52, 0x02, 0x69, 0x64, 0x22, 0x74,
	0x0a, 0x0b, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x2d, 0x0a,
	0x02, 0x65, 0x71, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x18, 0x92, 0x41, 0x15, 0x32, 0x0b,
	0x65, 0x78, 0x61, 0x63, 0x74, 0x20, 0x6d, 0x61, 0x74, 0x63, 0x68, 0xca, 0x3e, 0x05, 0xfa, 0x02,
	0x02, 0x65, 0x71, 0x48, 0x00, 0x52, 0x02, 0x65, 0x71, 0x88, 0x01, 0x01, 0x12, 0x26, 0x0a, 0x04,
	0x6c, 0x69, 0x6b, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0xca,
	0x3e, 0x07, 0xfa, 0x02, 0x04, 0x6c, 0x69, 0x6b, 0x65, 0x48, 0x01, 0x52, 0x04, 0x6c, 0x69, 0x6b,
	0x65, 0x88, 0x01, 0x01, 0x42, 0x05, 0x0a, 0x03, 0x5f, 0x65, 0x71, 0x42, 0x07, 0x0a, 0x05, 0x5f,
	0x6c, 0x69, 0x6b, 0x65, 0x22, 0x60, 0x0a, 0x04, 0x50, 0x61, 0x67, 0x65, 0x12, 0x26, 0x0a, 0x04,
	0x73, 0x69, 0x7a, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x0d, 0x92, 0x41, 0x0a, 0xca,
	0x3e, 0x07, 0xfa, 0x02, 0x04, 0x73, 0x69, 0x7a, 0x65, 0x48, 0x00, 0x52, 0x04, 0x73, 0x69, 0x7a,
	0x65, 0x88, 0x01, 0x01, 0x12, 0x20, 0x0a, 0x02, 0x65, 0x71, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x42, 0x0b, 0x92, 0x41, 0x08, 0xca, 0x3e, 0x05, 0xfa, 0x02, 0x02, 0x65, 0x71, 0x48, 0x01, 0x52,
	0x02, 0x65, 0x71, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x42,
	0x05, 0x0a, 0x03, 0x5f, 0x65, 0x71, 0x22, 0x8e, 0x02, 0x0a, 0x0a, 0x55, 0x73, 0x65, 0x72, 0x50,
	0x61, 0x72, 0x61, 0x6d, 0x73, 0x12, 0x6b, 0x0a, 0x04, 0x70, 0x61, 0x67, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x37, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77,
	0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72,
	0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e,
	0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x50, 0x61, 0x67, 0x65, 0x42, 0x19, 0x92, 0x41,
	0x16, 0x32, 0x04, 0x70, 0x61, 0x67, 0x65, 0xca, 0x3e, 0x0d, 0xfa, 0x02, 0x0a, 0x70, 0x61, 0x67,
	0x65, 0x5b, 0x73, 0x69, 0x7a, 0x65, 0x5d, 0x48, 0x00, 0x52, 0x04, 0x70, 0x61, 0x67, 0x65, 0x88,
	0x01, 0x01, 0x12, 0x80, 0x01, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x3e, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x41, 0x67, 0x67, 0x72, 0x65, 0x67, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x42, 0x27, 0x92, 0x41, 0x24, 0x32, 0x14, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x20, 0x74,
	0x68, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0xca, 0x3e, 0x0b, 0xfa,
	0x02, 0x08, 0x6e, 0x61, 0x6d, 0x65, 0x5b, 0x65, 0x71, 0x5d, 0x48, 0x01, 0x52, 0x04, 0x6e, 0x61,
	0x6d, 0x65, 0x88, 0x01, 0x01, 0x42, 0x07, 0x0a, 0x05, 0x5f, 0x70, 0x61, 0x67, 0x65, 0x42, 0x07,
	0x0a, 0x05, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x93, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x3a, 0xb6, 0x01, 0x92, 0x41, 0xb2, 0x01, 0x0a, 0x3c, 0x2a, 0x04, 0x55,
	0x73, 0x65, 0x72, 0x32, 0x12, 0x55, 0x73, 0x65, 0x72, 0x20, 0x69, 0x6e, 0x20, 0x74, 0x68, 0x65,
	0x20, 0x73, 0x79, 0x73, 0x74, 0x65, 0x6d, 0xd2, 0x01, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01,
	0x02, 0x69, 0x64, 0xd2, 0x01, 0x0b, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f,
	0x6e, 0xd2, 0x01, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x32, 0x72, 0x7b, 0x22, 0x6e, 0x61, 0x6d,
	0x65, 0x22, 0x3a, 0x20, 0x22, 0x50, 0x61, 0x74, 0x72, 0x69, 0x63, 0x6b, 0x22, 0x2c, 0x20, 0x22,
	0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4b, 0x6f, 0x73, 0x73, 0x22,
	0x2c, 0x20, 0x22, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3a, 0x20, 0x22, 0x70, 0x61, 0x74, 0x72,
	0x69, 0x63, 0x6b, 0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x2c, 0x20,
	0x22, 0x69, 0x64, 0x22, 0x3a, 0x20, 0x22, 0x63, 0x63, 0x37, 0x62, 0x66, 0x35, 0x37, 0x35, 0x2d,
	0x33, 0x64, 0x37, 0x34, 0x2d, 0x34, 0x33, 0x33, 0x37, 0x2d, 0x61, 0x34, 0x31, 0x64, 0x2d, 0x62,
	0x66, 0x62, 0x35, 0x66, 0x62, 0x35, 0x66, 0x33, 0x32, 0x38, 0x34, 0x22, 0x7d, 0x22, 0xdc, 0x01,
	0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79,
	0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x73, 0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14,
	0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65,
	0x6d, 0x61, 0x69, 0x6c, 0x3a, 0x85, 0x01, 0x92, 0x41, 0x81, 0x01, 0x0a, 0x39, 0x2a, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x32, 0x0e, 0x55,
	0x73, 0x65, 0x72, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0xd2, 0x01, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0xd2, 0x01, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0xd2, 0x01, 0x07, 0x73,
	0x75, 0x72, 0x6e, 0x61, 0x6d, 0x65, 0x32, 0x44, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x20, 0x22, 0x50, 0x61, 0x74, 0x72, 0x69, 0x63, 0x6b, 0x22, 0x2c, 0x20, 0x22, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4b, 0x6f, 0x73, 0x73, 0x22, 0x2c, 0x20, 0x22,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3a, 0x20, 0x22, 0x70, 0x61, 0x74, 0x72, 0x69, 0x63, 0x6b,
	0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x7d, 0x22, 0x9d, 0x02, 0x0a,
	0x11, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x12, 0x4d, 0x0a, 0x05, 0x75, 0x73, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x37, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79,
	0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31,
	0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x05, 0x75, 0x73, 0x65, 0x72,
	0x73, 0x3a, 0xb8, 0x01, 0x92, 0x41, 0xb4, 0x01, 0x0a, 0x31, 0x2a, 0x11, 0x4c, 0x69, 0x73, 0x74,
	0x20, 0x61, 0x6c, 0x6c, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x11, 0x4c,
	0x69, 0x73, 0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73,
	0xd2, 0x01, 0x08, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x73, 0x32, 0x7f, 0x7b, 0x22, 0x75,
	0x73, 0x65, 0x72, 0x73, 0x22, 0x3a, 0x20, 0x5b, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a,
	0x20, 0x22, 0x50, 0x61, 0x74, 0x72, 0x69, 0x63, 0x6b, 0x22, 0x2c, 0x20, 0x22, 0x73, 0x75, 0x72,
	0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4b, 0x6f, 0x73, 0x73, 0x22, 0x2c, 0x20, 0x22,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x3a, 0x20, 0x22, 0x70, 0x61, 0x74, 0x72, 0x69, 0x63, 0x6b,
	0x40, 0x67, 0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x2c, 0x20, 0x22, 0x69, 0x64,
	0x22, 0x3a, 0x20, 0x22, 0x63, 0x63, 0x37, 0x62, 0x66, 0x35, 0x37, 0x35, 0x2d, 0x33, 0x64, 0x37,
	0x34, 0x2d, 0x34, 0x33, 0x33, 0x37, 0x2d, 0x61, 0x34, 0x31, 0x64, 0x2d, 0x62, 0x66, 0x62, 0x35,
	0x66, 0x62, 0x35, 0x66, 0x33, 0x32, 0x38, 0x34, 0x22, 0x7d, 0x5d, 0x7d, 0x22, 0x8a, 0x02, 0x0a,
	0x0f, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x4b, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x37,
	0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78,
	0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61,
	0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x3a, 0xa9, 0x01,
	0x92, 0x41, 0xa5, 0x01, 0x0a, 0x25, 0x2a, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x73, 0x69, 0x6e, 0x67,
	0x6c, 0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x32, 0x0b, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0xd2, 0x01, 0x04, 0x75, 0x73, 0x65, 0x72, 0x32, 0x7c, 0x7b, 0x22, 0x75,
	0x73, 0x65, 0x72, 0x22, 0x3a, 0x20, 0x7b, 0x22, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22,
	0x50, 0x61, 0x74, 0x72, 0x69, 0x63, 0x6b, 0x22, 0x2c, 0x20, 0x22, 0x73, 0x75, 0x72, 0x6e, 0x61,
	0x6d, 0x65, 0x22, 0x3a, 0x20, 0x22, 0x4b, 0x6f, 0x73, 0x73, 0x22, 0x2c, 0x20, 0x22, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x22, 0x3a, 0x20, 0x22, 0x70, 0x61, 0x74, 0x72, 0x69, 0x63, 0x6b, 0x40, 0x67,
	0x6d, 0x61, 0x69, 0x6c, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x2c, 0x20, 0x22, 0x69, 0x64, 0x22, 0x3a,
	0x20, 0x22, 0x63, 0x63, 0x37, 0x62, 0x66, 0x35, 0x37, 0x35, 0x2d, 0x33, 0x64, 0x37, 0x34, 0x2d,
	0x34, 0x33, 0x33, 0x37, 0x2d, 0x61, 0x34, 0x31, 0x64, 0x2d, 0x62, 0x66, 0x62, 0x35, 0x66, 0x62,
	0x35, 0x66, 0x33, 0x32, 0x38, 0x34, 0x22, 0x7d, 0x7d, 0x32, 0xb1, 0x0b, 0x0a, 0x0b, 0x55, 0x73,
	0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x9c, 0x03, 0x0a, 0x0a, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x41, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e,
	0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e,
	0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72,
	0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x42, 0x6f, 0x64, 0x79, 0x1a, 0x39, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8f, 0x02, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x0e, 0x22,
	0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x3a, 0x01, 0x2a, 0x92, 0x41, 0xf7,
	0x01, 0x4a, 0x38, 0x0a, 0x03, 0x32, 0x30, 0x31, 0x12, 0x31, 0x0a, 0x0e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x20, 0x61, 0x6e, 0x20, 0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x1d, 0x1a, 0x1b,
	0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6f,
	0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x3c, 0x0a, 0x03, 0x34,
	0x30, 0x30, 0x12, 0x35, 0x0a, 0x0d, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x20, 0x69, 0x6e,
	0x70, 0x75, 0x74, 0x12, 0x24, 0x0a, 0x22, 0x1a, 0x20, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x3d, 0x0a, 0x03, 0x34, 0x30, 0x39,
	0x12, 0x36, 0x0a, 0x0e, 0x41, 0x6c, 0x72, 0x65, 0x61, 0x64, 0x79, 0x20, 0x65, 0x78, 0x69, 0x73,
	0x74, 0x73, 0x12, 0x24, 0x0a, 0x22, 0x1a, 0x20, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x3e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61,
	0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x10, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65,
	0x64, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x1d, 0x1a, 0x1b, 0x23, 0x2f, 0x64,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69,
	0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0xa9, 0x02, 0x0a, 0x08, 0x47, 0x65, 0x74,
	0x55, 0x73, 0x65, 0x72, 0x73, 0x12, 0x3d, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x70,
	0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x50, 0x61,
	0x72, 0x61, 0x6d, 0x73, 0x1a, 0x44, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65,
	0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69,
	0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x97, 0x01, 0x82, 0xd3, 0xe4,
	0x93, 0x02, 0x0b, 0x12, 0x09, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x73, 0x92, 0x41,
	0x82, 0x01, 0x4a, 0x40, 0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x39, 0x0a, 0x0e, 0x4c, 0x69, 0x73,
	0x74, 0x20, 0x61, 0x6c, 0x6c, 0x20, 0x75, 0x73, 0x65, 0x72, 0x73, 0x12, 0x27, 0x0a, 0x25, 0x1a,
	0x23, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x75,
	0x73, 0x65, 0x72, 0x4c, 0x69, 0x73, 0x74, 0x55, 0x73, 0x65, 0x72, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x3e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x33, 0x0a, 0x10, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x1d, 0x1a, 0x1b, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0xe5, 0x02, 0x0a, 0x07, 0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72,
	0x12, 0x39, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e,
	0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c,
	0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e,
	0x75, 0x73, 0x65, 0x72, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x42, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x47, 0x65, 0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22,
	0xda, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x12, 0x0e, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0xc0, 0x01, 0x4a, 0x3f, 0x0a, 0x03,
	0x32, 0x30, 0x30, 0x12, 0x38, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x20, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x20, 0x75, 0x73, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x23, 0x1a, 0x21, 0x23, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x47, 0x65,
	0x74, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x4a, 0x3d, 0x0a,
	0x03, 0x34, 0x30, 0x34, 0x12, 0x36, 0x0a, 0x0e, 0x55, 0x73, 0x65, 0x72, 0x20, 0x6e, 0x6f, 0x74,
	0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x24, 0x0a, 0x22, 0x1a, 0x20, 0x23, 0x2f, 0x64, 0x65,
	0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e,
	0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x3e, 0x0a, 0x07,
	0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12, 0x33, 0x0a, 0x10, 0x75, 0x6e, 0x65, 0x78, 0x70,
	0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x1d, 0x1a,
	0x1b, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64,
	0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0xd5, 0x02, 0x0a,
	0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x12, 0x39, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70,
	0x6c, 0x65, 0x2e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70,
	0x74, 0x65, 0x72, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x76, 0x31, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x1a, 0x39, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x67, 0x61,
	0x74, 0x65, 0x77, 0x61, 0x79, 0x2e, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x69, 0x6e,
	0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2e, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x22, 0xd0, 0x01, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x2a, 0x0e, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x73, 0x2f, 0x7b, 0x69, 0x64, 0x7d, 0x92, 0x41, 0xb6, 0x01, 0x4a, 0x35,
	0x0a, 0x03, 0x32, 0x30, 0x30, 0x12, 0x2e, 0x0a, 0x0b, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x20,
	0x75, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x0a, 0x1d, 0x1a, 0x1b, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69,
	0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x4a, 0x3d, 0x0a, 0x03, 0x34, 0x30, 0x34, 0x12, 0x36, 0x0a, 0x0e,
	0x55, 0x73, 0x65, 0x72, 0x20, 0x6e, 0x6f, 0x74, 0x20, 0x66, 0x6f, 0x75, 0x6e, 0x64, 0x12, 0x24,
	0x0a, 0x22, 0x1a, 0x20, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x4a, 0x3e, 0x0a, 0x07, 0x64, 0x65, 0x66, 0x61, 0x75, 0x6c, 0x74, 0x12,
	0x33, 0x0a, 0x10, 0x75, 0x6e, 0x65, 0x78, 0x70, 0x65, 0x63, 0x74, 0x65, 0x64, 0x20, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x12, 0x1f, 0x0a, 0x1d, 0x1a, 0x1b, 0x23, 0x2f, 0x64, 0x65, 0x66, 0x69, 0x6e,
	0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x1a, 0x17, 0x92, 0x41, 0x14, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12,
	0x0c, 0x55, 0x73, 0x65, 0x72, 0x20, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x42, 0x4a, 0x5a,
	0x48, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x70, 0x61, 0x74, 0x72,
	0x69, 0x63, 0x6b, 0x6b, 0x6f, 0x73, 0x73, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2d, 0x67, 0x61, 0x74,
	0x65, 0x77, 0x61, 0x79, 0x2d, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2f, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x61, 0x64, 0x61, 0x70, 0x74, 0x65, 0x72, 0x2f, 0x61, 0x70,
	0x69, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x33,
}

var (
	file_internal_adapter_api_v1_user_user_proto_rawDescOnce sync.Once
	file_internal_adapter_api_v1_user_user_proto_rawDescData = file_internal_adapter_api_v1_user_user_proto_rawDesc
)

func file_internal_adapter_api_v1_user_user_proto_rawDescGZIP() []byte {
	file_internal_adapter_api_v1_user_user_proto_rawDescOnce.Do(func() {
		file_internal_adapter_api_v1_user_user_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_adapter_api_v1_user_user_proto_rawDescData)
	})
	return file_internal_adapter_api_v1_user_user_proto_rawDescData
}

var file_internal_adapter_api_v1_user_user_proto_msgTypes = make([]protoimpl.MessageInfo, 8)
var file_internal_adapter_api_v1_user_user_proto_goTypes = []interface{}{
	(*UserId)(nil),            // 0: grpc.gateway.example.internal.adapter.api.v1.user.UserId
	(*Aggregation)(nil),       // 1: grpc.gateway.example.internal.adapter.api.v1.user.Aggregation
	(*Page)(nil),              // 2: grpc.gateway.example.internal.adapter.api.v1.user.Page
	(*UserParams)(nil),        // 3: grpc.gateway.example.internal.adapter.api.v1.user.UserParams
	(*User)(nil),              // 4: grpc.gateway.example.internal.adapter.api.v1.user.User
	(*CreateUserBody)(nil),    // 5: grpc.gateway.example.internal.adapter.api.v1.user.CreateUserBody
	(*ListUsersResponse)(nil), // 6: grpc.gateway.example.internal.adapter.api.v1.user.ListUsersResponse
	(*GetUserResponse)(nil),   // 7: grpc.gateway.example.internal.adapter.api.v1.user.GetUserResponse
	(*proto.Message)(nil),     // 8: grpc.gateway.example.internal.adapter.api.domain.Message
}
var file_internal_adapter_api_v1_user_user_proto_depIdxs = []int32{
	2, // 0: grpc.gateway.example.internal.adapter.api.v1.user.UserParams.page:type_name -> grpc.gateway.example.internal.adapter.api.v1.user.Page
	1, // 1: grpc.gateway.example.internal.adapter.api.v1.user.UserParams.name:type_name -> grpc.gateway.example.internal.adapter.api.v1.user.Aggregation
	4, // 2: grpc.gateway.example.internal.adapter.api.v1.user.ListUsersResponse.users:type_name -> grpc.gateway.example.internal.adapter.api.v1.user.User
	4, // 3: grpc.gateway.example.internal.adapter.api.v1.user.GetUserResponse.user:type_name -> grpc.gateway.example.internal.adapter.api.v1.user.User
	5, // 4: grpc.gateway.example.internal.adapter.api.v1.user.UserService.CreateUser:input_type -> grpc.gateway.example.internal.adapter.api.v1.user.CreateUserBody
	3, // 5: grpc.gateway.example.internal.adapter.api.v1.user.UserService.GetUsers:input_type -> grpc.gateway.example.internal.adapter.api.v1.user.UserParams
	0, // 6: grpc.gateway.example.internal.adapter.api.v1.user.UserService.GetUser:input_type -> grpc.gateway.example.internal.adapter.api.v1.user.UserId
	0, // 7: grpc.gateway.example.internal.adapter.api.v1.user.UserService.DeleteUser:input_type -> grpc.gateway.example.internal.adapter.api.v1.user.UserId
	8, // 8: grpc.gateway.example.internal.adapter.api.v1.user.UserService.CreateUser:output_type -> grpc.gateway.example.internal.adapter.api.domain.Message
	6, // 9: grpc.gateway.example.internal.adapter.api.v1.user.UserService.GetUsers:output_type -> grpc.gateway.example.internal.adapter.api.v1.user.ListUsersResponse
	7, // 10: grpc.gateway.example.internal.adapter.api.v1.user.UserService.GetUser:output_type -> grpc.gateway.example.internal.adapter.api.v1.user.GetUserResponse
	8, // 11: grpc.gateway.example.internal.adapter.api.v1.user.UserService.DeleteUser:output_type -> grpc.gateway.example.internal.adapter.api.domain.Message
	8, // [8:12] is the sub-list for method output_type
	4, // [4:8] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_internal_adapter_api_v1_user_user_proto_init() }
func file_internal_adapter_api_v1_user_user_proto_init() {
	if File_internal_adapter_api_v1_user_user_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_adapter_api_v1_user_user_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserId); i {
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
		file_internal_adapter_api_v1_user_user_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Aggregation); i {
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
		file_internal_adapter_api_v1_user_user_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Page); i {
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
		file_internal_adapter_api_v1_user_user_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserParams); i {
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
		file_internal_adapter_api_v1_user_user_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_internal_adapter_api_v1_user_user_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserBody); i {
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
		file_internal_adapter_api_v1_user_user_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ListUsersResponse); i {
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
		file_internal_adapter_api_v1_user_user_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetUserResponse); i {
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
	file_internal_adapter_api_v1_user_user_proto_msgTypes[1].OneofWrappers = []interface{}{}
	file_internal_adapter_api_v1_user_user_proto_msgTypes[2].OneofWrappers = []interface{}{}
	file_internal_adapter_api_v1_user_user_proto_msgTypes[3].OneofWrappers = []interface{}{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_internal_adapter_api_v1_user_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   8,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_adapter_api_v1_user_user_proto_goTypes,
		DependencyIndexes: file_internal_adapter_api_v1_user_user_proto_depIdxs,
		MessageInfos:      file_internal_adapter_api_v1_user_user_proto_msgTypes,
	}.Build()
	File_internal_adapter_api_v1_user_user_proto = out.File
	file_internal_adapter_api_v1_user_user_proto_rawDesc = nil
	file_internal_adapter_api_v1_user_user_proto_goTypes = nil
	file_internal_adapter_api_v1_user_user_proto_depIdxs = nil
}
