// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gamedef/game_def.proto

/*
Package gamedef is a generated protocol buffer package.

It is generated from these files:
	gamedef/game_def.proto

It has these top-level messages:
	LoginInfo
	ExtraAccountInfo
	User
*/
package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 用户登录信息
type LoginInfo struct {
	Ip     string `protobuf:"bytes,1,opt,name=ip" json:"ip,omitempty"`
	Region string `protobuf:"bytes,2,opt,name=region" json:"region,omitempty"`
}

func (m *LoginInfo) Reset()                    { *m = LoginInfo{} }
func (m *LoginInfo) String() string            { return proto.CompactTextString(m) }
func (*LoginInfo) ProtoMessage()               {}
func (*LoginInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *LoginInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *LoginInfo) GetRegion() string {
	if m != nil {
		return m.Region
	}
	return ""
}

// 账号额外信息
type ExtraAccountInfo struct {
	AccountType uint32 `protobuf:"varint,1,opt,name=accountType" json:"accountType,omitempty"`
	Channel     uint32 `protobuf:"varint,2,opt,name=channel" json:"channel,omitempty"`
	Platform    uint32 `protobuf:"varint,3,opt,name=platform" json:"platform,omitempty"`
}

func (m *ExtraAccountInfo) Reset()                    { *m = ExtraAccountInfo{} }
func (m *ExtraAccountInfo) String() string            { return proto.CompactTextString(m) }
func (*ExtraAccountInfo) ProtoMessage()               {}
func (*ExtraAccountInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ExtraAccountInfo) GetAccountType() uint32 {
	if m != nil {
		return m.AccountType
	}
	return 0
}

func (m *ExtraAccountInfo) GetChannel() uint32 {
	if m != nil {
		return m.Channel
	}
	return 0
}

func (m *ExtraAccountInfo) GetPlatform() uint32 {
	if m != nil {
		return m.Platform
	}
	return 0
}

// 用户信息
type User struct {
	UserID   uint64 `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Nickname string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	SexID    uint32 `protobuf:"varint,3,opt,name=sexID" json:"sexID,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *User) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *User) GetNickname() string {
	if m != nil {
		return m.Nickname
	}
	return ""
}

func (m *User) GetSexID() uint32 {
	if m != nil {
		return m.SexID
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginInfo)(nil), "gamedef.LoginInfo")
	proto.RegisterType((*ExtraAccountInfo)(nil), "gamedef.ExtraAccountInfo")
	proto.RegisterType((*User)(nil), "gamedef.User")
}

func init() { proto.RegisterFile("gamedef/game_def.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 214 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0xd9, 0xba, 0xee, 0xba, 0x23, 0x2b, 0x12, 0x64, 0x09, 0x9e, 0x4a, 0x4f, 0x9e, 0xf4,
	0xd0, 0x5f, 0x20, 0xd4, 0x43, 0xc1, 0x83, 0x04, 0x3d, 0x4b, 0x4c, 0x27, 0x35, 0xda, 0x4e, 0x42,
	0x9a, 0x42, 0xfd, 0xf7, 0xd2, 0x34, 0x16, 0x4f, 0xe1, 0x7b, 0xe1, 0x7d, 0x3c, 0x06, 0x4e, 0xad,
	0xec, 0xb1, 0x41, 0xfd, 0x30, 0xbf, 0xef, 0x0d, 0xea, 0x7b, 0xe7, 0x6d, 0xb0, 0x6c, 0x9f, 0xf2,
	0xa2, 0x84, 0xc3, 0xb3, 0x6d, 0x0d, 0xd5, 0xa4, 0x2d, 0xbb, 0x82, 0xcc, 0x38, 0xbe, 0xc9, 0x37,
	0x77, 0x07, 0x91, 0x19, 0xc7, 0x4e, 0xb0, 0xf3, 0xd8, 0x1a, 0x4b, 0x3c, 0x8b, 0x59, 0xa2, 0xe2,
	0x0b, 0xae, 0x9f, 0xa6, 0xe0, 0xe5, 0xa3, 0x52, 0x76, 0xa4, 0x10, 0xbb, 0x39, 0x5c, 0xca, 0x05,
	0x5f, 0x7f, 0x1c, 0x46, 0xc9, 0x51, 0xfc, 0x8f, 0x18, 0x87, 0xbd, 0xfa, 0x94, 0x44, 0xd8, 0x45,
	0xdd, 0x51, 0xfc, 0x21, 0xbb, 0x85, 0x0b, 0xd7, 0xc9, 0xa0, 0xad, 0xef, 0xf9, 0x59, 0xfc, 0x5a,
	0xb9, 0x78, 0x81, 0xed, 0xdb, 0x80, 0x7e, 0xde, 0x32, 0x0e, 0xe8, 0xeb, 0x2a, 0xaa, 0xb7, 0x22,
	0xd1, 0xdc, 0x25, 0xa3, 0xbe, 0x49, 0xf6, 0x98, 0x56, 0xae, 0xcc, 0x6e, 0xe0, 0x7c, 0xc0, 0xa9,
	0xae, 0x92, 0x74, 0x81, 0x8f, 0x5d, 0x3c, 0x41, 0xf9, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x93, 0x3d,
	0x37, 0x47, 0x1c, 0x01, 0x00, 0x00,
}
