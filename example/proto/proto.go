// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto.proto

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
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	Name                 *string  `protobuf:"bytes,2,req,name=name" json:"name,omitempty"`
	Head                 *string  `protobuf:"bytes,3,req,name=head" json:"head,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{0}
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
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *User) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *User) GetHead() string {
	if m != nil && m.Head != nil {
		return *m.Head
	}
	return ""
}

type CsLogin struct {
	Id                   *string  `protobuf:"bytes,1,req,name=id" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CsLogin) Reset()         { *m = CsLogin{} }
func (m *CsLogin) String() string { return proto.CompactTextString(m) }
func (*CsLogin) ProtoMessage()    {}
func (*CsLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{1}
}

func (m *CsLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CsLogin.Unmarshal(m, b)
}
func (m *CsLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CsLogin.Marshal(b, m, deterministic)
}
func (m *CsLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CsLogin.Merge(m, src)
}
func (m *CsLogin) XXX_Size() int {
	return xxx_messageInfo_CsLogin.Size(m)
}
func (m *CsLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_CsLogin.DiscardUnknown(m)
}

var xxx_messageInfo_CsLogin proto.InternalMessageInfo

func (m *CsLogin) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

type ScLogin struct {
	User                 *User    `protobuf:"bytes,1,req,name=user" json:"user,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScLogin) Reset()         { *m = ScLogin{} }
func (m *ScLogin) String() string { return proto.CompactTextString(m) }
func (*ScLogin) ProtoMessage()    {}
func (*ScLogin) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fcc84b9998d60d8, []int{2}
}

func (m *ScLogin) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScLogin.Unmarshal(m, b)
}
func (m *ScLogin) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScLogin.Marshal(b, m, deterministic)
}
func (m *ScLogin) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScLogin.Merge(m, src)
}
func (m *ScLogin) XXX_Size() int {
	return xxx_messageInfo_ScLogin.Size(m)
}
func (m *ScLogin) XXX_DiscardUnknown() {
	xxx_messageInfo_ScLogin.DiscardUnknown(m)
}

var xxx_messageInfo_ScLogin proto.InternalMessageInfo

func (m *ScLogin) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func init() {
	proto.RegisterType((*User)(nil), "proto.User")
	proto.RegisterType((*CsLogin)(nil), "proto.cs_login")
	proto.RegisterType((*ScLogin)(nil), "proto.sc_login")
}

func init() { proto.RegisterFile("proto.proto", fileDescriptor_2fcc84b9998d60d8) }

var fileDescriptor_2fcc84b9998d60d8 = []byte{
	// 129 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x03, 0x93, 0x42, 0xac, 0x60, 0x4a, 0xc9, 0x8e, 0x8b, 0x25, 0xb4, 0x38, 0xb5, 0x48,
	0x88, 0x8f, 0x8b, 0x29, 0x33, 0x45, 0x82, 0x51, 0x81, 0x49, 0x83, 0x33, 0x88, 0x29, 0x33, 0x45,
	0x48, 0x88, 0x8b, 0x25, 0x2f, 0x31, 0x37, 0x55, 0x82, 0x09, 0x2c, 0x02, 0x66, 0x83, 0xc4, 0x32,
	0x52, 0x13, 0x53, 0x24, 0x98, 0x21, 0x62, 0x20, 0xb6, 0x92, 0x14, 0x17, 0x47, 0x72, 0x71, 0x7c,
	0x4e, 0x7e, 0x7a, 0x66, 0x1e, 0xba, 0x19, 0x4a, 0xda, 0x5c, 0x1c, 0xc5, 0xc9, 0x50, 0x39, 0x79,
	0x2e, 0x96, 0xd2, 0xe2, 0xd4, 0x22, 0xb0, 0x2c, 0xb7, 0x11, 0x37, 0xc4, 0x11, 0x7a, 0x20, 0xab,
	0x83, 0xc0, 0x12, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0d, 0x1c, 0xac, 0x21, 0x9d, 0x00, 0x00,
	0x00,
}
