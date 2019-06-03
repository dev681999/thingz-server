// Code generated by protoc-gen-go. DO NOT EDIT.
// source: user.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type User struct {
	// @inject_tag: bson:"_id"
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty" bson:"_id"`
	Email                string   `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	Password             string   `protobuf:"bytes,4,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func UserFromBytes(data []byte) (*User, error) {
	t := &User{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *User) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *User) Reset()                   { *m = User{} }
func (m *User) String() string           { return proto.CompactTextString(m) }
func (*User) ProtoMessage()              {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type CreateUserRequest struct {
	User                 *User    `protobuf:"bytes,1,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func CreateUserRequestFromBytes(data []byte) (*CreateUserRequest, error) {
	t := &CreateUserRequest{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *CreateUserRequest) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *CreateUserRequest) Reset()                   { *m = CreateUserRequest{} }
func (m *CreateUserRequest) String() string           { return proto.CompactTextString(m) }
func (*CreateUserRequest) ProtoMessage()              {}
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{1}
}

func (m *CreateUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserRequest.Unmarshal(m, b)
}
func (m *CreateUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserRequest.Marshal(b, m, deterministic)
}
func (m *CreateUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserRequest.Merge(m, src)
}
func (m *CreateUserRequest) XXX_Size() int {
	return xxx_messageInfo_CreateUserRequest.Size(m)
}
func (m *CreateUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserRequest proto.InternalMessageInfo

func (m *CreateUserRequest) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

type CreateUserResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	Id                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func CreateUserResponseFromBytes(data []byte) (*CreateUserResponse, error) {
	t := &CreateUserResponse{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *CreateUserResponse) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *CreateUserResponse) Reset()                   { *m = CreateUserResponse{} }
func (m *CreateUserResponse) String() string           { return proto.CompactTextString(m) }
func (*CreateUserResponse) ProtoMessage()              {}
func (*CreateUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{2}
}

func (m *CreateUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateUserResponse.Unmarshal(m, b)
}
func (m *CreateUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateUserResponse.Marshal(b, m, deterministic)
}
func (m *CreateUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateUserResponse.Merge(m, src)
}
func (m *CreateUserResponse) XXX_Size() int {
	return xxx_messageInfo_CreateUserResponse.Size(m)
}
func (m *CreateUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateUserResponse proto.InternalMessageInfo

func (m *CreateUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *CreateUserResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *CreateUserResponse) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type VerifyUserRequest struct {
	Email                string   `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,2,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func VerifyUserRequestFromBytes(data []byte) (*VerifyUserRequest, error) {
	t := &VerifyUserRequest{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *VerifyUserRequest) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *VerifyUserRequest) Reset()                   { *m = VerifyUserRequest{} }
func (m *VerifyUserRequest) String() string           { return proto.CompactTextString(m) }
func (*VerifyUserRequest) ProtoMessage()              {}
func (*VerifyUserRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{3}
}

func (m *VerifyUserRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserRequest.Unmarshal(m, b)
}
func (m *VerifyUserRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserRequest.Marshal(b, m, deterministic)
}
func (m *VerifyUserRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserRequest.Merge(m, src)
}
func (m *VerifyUserRequest) XXX_Size() int {
	return xxx_messageInfo_VerifyUserRequest.Size(m)
}
func (m *VerifyUserRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserRequest.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserRequest proto.InternalMessageInfo

func (m *VerifyUserRequest) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *VerifyUserRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type VerifyUserResponse struct {
	Success              bool     `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Error                string   `protobuf:"bytes,2,opt,name=error,proto3" json:"error,omitempty"`
	User                 *User    `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func VerifyUserResponseFromBytes(data []byte) (*VerifyUserResponse, error) {
	t := &VerifyUserResponse{}
	err := proto.Unmarshal(data, t)
	return t, err
}

func (m *VerifyUserResponse) ToBytes() ([]byte, error) { return proto.Marshal(m) }
func (m *VerifyUserResponse) Reset()                   { *m = VerifyUserResponse{} }
func (m *VerifyUserResponse) String() string           { return proto.CompactTextString(m) }
func (*VerifyUserResponse) ProtoMessage()              {}
func (*VerifyUserResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_116e343673f7ffaf, []int{4}
}

func (m *VerifyUserResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_VerifyUserResponse.Unmarshal(m, b)
}
func (m *VerifyUserResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_VerifyUserResponse.Marshal(b, m, deterministic)
}
func (m *VerifyUserResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifyUserResponse.Merge(m, src)
}
func (m *VerifyUserResponse) XXX_Size() int {
	return xxx_messageInfo_VerifyUserResponse.Size(m)
}
func (m *VerifyUserResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifyUserResponse.DiscardUnknown(m)
}

var xxx_messageInfo_VerifyUserResponse proto.InternalMessageInfo

func (m *VerifyUserResponse) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *VerifyUserResponse) GetError() string {
	if m != nil {
		return m.Error
	}
	return ""
}

func (m *VerifyUserResponse) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*CreateUserRequest)(nil), "proto.CreateUserRequest")
	proto.RegisterType((*CreateUserResponse)(nil), "proto.CreateUserResponse")
	proto.RegisterType((*VerifyUserRequest)(nil), "proto.VerifyUserRequest")
	proto.RegisterType((*VerifyUserResponse)(nil), "proto.VerifyUserResponse")
}

func init() { proto.RegisterFile("user.proto", fileDescriptor_116e343673f7ffaf) }

var fileDescriptor_116e343673f7ffaf = []byte{
	// 228 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0x69, 0x9a, 0xd5, 0x75, 0x16, 0x84, 0x1d, 0x3c, 0x04, 0x2f, 0x4a, 0x4e, 0x9e, 0xf6,
	0xa0, 0xbe, 0x81, 0xf8, 0x02, 0x45, 0x3d, 0x79, 0x89, 0xed, 0x08, 0x01, 0xdb, 0xd4, 0x99, 0x16,
	0xf1, 0xed, 0xa5, 0x53, 0x5b, 0x82, 0xe0, 0x65, 0x4f, 0xc9, 0xff, 0x27, 0x33, 0xf3, 0x7f, 0x03,
	0x30, 0x0a, 0xf1, 0xa1, 0xe7, 0x34, 0x24, 0xdc, 0xe8, 0xe1, 0x5f, 0xc1, 0x3e, 0x0b, 0x31, 0x9e,
	0x83, 0x89, 0x8d, 0x2b, 0xae, 0x8b, 0x9b, 0xb3, 0xca, 0xc4, 0x06, 0x2f, 0x60, 0x43, 0x6d, 0x88,
	0x1f, 0xce, 0xa8, 0x35, 0x0b, 0x44, 0xb0, 0x5d, 0x68, 0xc9, 0x95, 0x6a, 0xea, 0x1d, 0x2f, 0x61,
	0xdb, 0x07, 0x91, 0xaf, 0xc4, 0x8d, 0xb3, 0xea, 0xaf, 0xda, 0xdf, 0xc3, 0xfe, 0x81, 0x29, 0x0c,
	0x34, 0xcd, 0xa8, 0xe8, 0x73, 0x24, 0x19, 0xf0, 0x0a, 0xec, 0x94, 0x43, 0x87, 0xed, 0x6e, 0x77,
	0x73, 0x9e, 0x83, 0xfe, 0xd0, 0x07, 0xff, 0x04, 0x98, 0x57, 0x49, 0x9f, 0x3a, 0x21, 0x74, 0x70,
	0x2a, 0x63, 0x5d, 0x93, 0x88, 0x56, 0x6e, 0xab, 0x45, 0x6a, 0x56, 0xe6, 0xc4, 0x6b, 0xd6, 0x49,
	0xfc, 0x12, 0x95, 0x0b, 0x91, 0x7f, 0x84, 0xfd, 0x0b, 0x71, 0x7c, 0xff, 0xce, 0xb3, 0xac, 0x98,
	0x45, 0x8e, 0x99, 0x23, 0x99, 0x3f, 0x48, 0x04, 0x98, 0xb7, 0x39, 0x32, 0xdc, 0xb2, 0x83, 0xf2,
	0x9f, 0x1d, 0xbc, 0x9d, 0xa8, 0x73, 0xf7, 0x13, 0x00, 0x00, 0xff, 0xff, 0x3e, 0xbc, 0xaa, 0x8b,
	0xb3, 0x01, 0x00, 0x00,
}