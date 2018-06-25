// Code generated by protoc-gen-go. DO NOT EDIT.
// source: cmsg/client_msg.proto

/*
Package cmsg is a generated protocol buffer package.

It is generated from these files:
	cmsg/client_msg.proto

It has these top-level messages:
	CReqSyncSystemTime
	CRespSyncSystemTime
	CReqAuth
	CRespAuth
	CNotifyLoginInfo
	CReqLogin
	CRespLogin
	CReqLogout
	CRespLogout
	CNotifyLogout
	CReqNotifyUserData
	CRespNotifyUserData
	CReqUseSkill
	CRespUseSkill
	CNotifyUseSkill
	CNotifyGameStart
	CNotifyGeneralStatus
*/
package cmsg

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import gamedef "github.com/wenxiu2199/gameserver/src/server/gameproto/gamedef"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 同步系统时间
type CReqSyncSystemTime struct {
	Timestamp int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *CReqSyncSystemTime) Reset()                    { *m = CReqSyncSystemTime{} }
func (m *CReqSyncSystemTime) String() string            { return proto.CompactTextString(m) }
func (*CReqSyncSystemTime) ProtoMessage()               {}
func (*CReqSyncSystemTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CReqSyncSystemTime) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type CRespSyncSystemTime struct {
	Timestamp       int64 `protobuf:"varint,1,opt,name=timestamp" json:"timestamp,omitempty"`
	ServerTimestamp int64 `protobuf:"varint,2,opt,name=serverTimestamp" json:"serverTimestamp,omitempty"`
}

func (m *CRespSyncSystemTime) Reset()                    { *m = CRespSyncSystemTime{} }
func (m *CRespSyncSystemTime) String() string            { return proto.CompactTextString(m) }
func (*CRespSyncSystemTime) ProtoMessage()               {}
func (*CRespSyncSystemTime) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CRespSyncSystemTime) GetTimestamp() int64 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

func (m *CRespSyncSystemTime) GetServerTimestamp() int64 {
	if m != nil {
		return m.ServerTimestamp
	}
	return 0
}

// 玩家认证请求
type CReqAuth struct {
	Account  string                    `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Password string                    `protobuf:"bytes,2,opt,name=password" json:"password,omitempty"`
	Extra    *gamedef.ExtraAccountInfo `protobuf:"bytes,3,opt,name=extra" json:"extra,omitempty"`
}

func (m *CReqAuth) Reset()                    { *m = CReqAuth{} }
func (m *CReqAuth) String() string            { return proto.CompactTextString(m) }
func (*CReqAuth) ProtoMessage()               {}
func (*CReqAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *CReqAuth) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CReqAuth) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

func (m *CReqAuth) GetExtra() *gamedef.ExtraAccountInfo {
	if m != nil {
		return m.Extra
	}
	return nil
}

type CRespAuth struct {
	ErrCode    uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg     string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	Account    string `protobuf:"bytes,3,opt,name=account" json:"account,omitempty"`
	UserID     uint64 `protobuf:"varint,4,opt,name=userID" json:"userID,omitempty"`
	Sign       string `protobuf:"bytes,5,opt,name=sign" json:"sign,omitempty"`
	UnlockTime int64  `protobuf:"varint,6,opt,name=unlockTime" json:"unlockTime,omitempty"`
}

func (m *CRespAuth) Reset()                    { *m = CRespAuth{} }
func (m *CRespAuth) String() string            { return proto.CompactTextString(m) }
func (*CRespAuth) ProtoMessage()               {}
func (*CRespAuth) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *CRespAuth) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespAuth) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *CRespAuth) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CRespAuth) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CRespAuth) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

func (m *CRespAuth) GetUnlockTime() int64 {
	if m != nil {
		return m.UnlockTime
	}
	return 0
}

// 账号登录通知
type CNotifyLoginInfo struct {
	Account string `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
	Ip      string `protobuf:"bytes,2,opt,name=ip" json:"ip,omitempty"`
}

func (m *CNotifyLoginInfo) Reset()                    { *m = CNotifyLoginInfo{} }
func (m *CNotifyLoginInfo) String() string            { return proto.CompactTextString(m) }
func (*CNotifyLoginInfo) ProtoMessage()               {}
func (*CNotifyLoginInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *CNotifyLoginInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *CNotifyLoginInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

// 玩家登录请求
type CReqLogin struct {
	UserID uint64 `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Sign   string `protobuf:"bytes,2,opt,name=sign" json:"sign,omitempty"`
}

func (m *CReqLogin) Reset()                    { *m = CReqLogin{} }
func (m *CReqLogin) String() string            { return proto.CompactTextString(m) }
func (*CReqLogin) ProtoMessage()               {}
func (*CReqLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *CReqLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CReqLogin) GetSign() string {
	if m != nil {
		return m.Sign
	}
	return ""
}

type CRespLogin struct {
	ErrCode     uint32        `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg      string        `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	UserID      uint64        `protobuf:"varint,3,opt,name=userID" json:"userID,omitempty"`
	User        *gamedef.User `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	IsReconnect bool          `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
}

func (m *CRespLogin) Reset()                    { *m = CRespLogin{} }
func (m *CRespLogin) String() string            { return proto.CompactTextString(m) }
func (*CRespLogin) ProtoMessage()               {}
func (*CRespLogin) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *CRespLogin) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespLogin) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *CRespLogin) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CRespLogin) GetUser() *gamedef.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CRespLogin) GetIsReconnect() bool {
	if m != nil {
		return m.IsReconnect
	}
	return false
}

// 玩家登出请求
type CReqLogout struct {
}

func (m *CReqLogout) Reset()                    { *m = CReqLogout{} }
func (m *CReqLogout) String() string            { return proto.CompactTextString(m) }
func (*CReqLogout) ProtoMessage()               {}
func (*CReqLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type CRespLogout struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *CRespLogout) Reset()                    { *m = CRespLogout{} }
func (m *CRespLogout) String() string            { return proto.CompactTextString(m) }
func (*CRespLogout) ProtoMessage()               {}
func (*CRespLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *CRespLogout) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespLogout) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

// 玩家被提出
type CNotifyLogout struct {
	LoginInfo *gamedef.LoginInfo `protobuf:"bytes,1,opt,name=loginInfo" json:"loginInfo,omitempty"`
}

func (m *CNotifyLogout) Reset()                    { *m = CNotifyLogout{} }
func (m *CNotifyLogout) String() string            { return proto.CompactTextString(m) }
func (*CNotifyLogout) ProtoMessage()               {}
func (*CNotifyLogout) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *CNotifyLogout) GetLoginInfo() *gamedef.LoginInfo {
	if m != nil {
		return m.LoginInfo
	}
	return nil
}

// 请求玩家数据
type CReqNotifyUserData struct {
}

func (m *CReqNotifyUserData) Reset()                    { *m = CReqNotifyUserData{} }
func (m *CReqNotifyUserData) String() string            { return proto.CompactTextString(m) }
func (*CReqNotifyUserData) ProtoMessage()               {}
func (*CReqNotifyUserData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

// 返回玩家数据
type CRespNotifyUserData struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *CRespNotifyUserData) Reset()                    { *m = CRespNotifyUserData{} }
func (m *CRespNotifyUserData) String() string            { return proto.CompactTextString(m) }
func (*CRespNotifyUserData) ProtoMessage()               {}
func (*CRespNotifyUserData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *CRespNotifyUserData) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespNotifyUserData) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

// ===============Game数据====================
type CReqUseSkill struct {
	SkillID uint32 `protobuf:"varint,1,opt,name=skillID" json:"skillID,omitempty"`
}

func (m *CReqUseSkill) Reset()                    { *m = CReqUseSkill{} }
func (m *CReqUseSkill) String() string            { return proto.CompactTextString(m) }
func (*CReqUseSkill) ProtoMessage()               {}
func (*CReqUseSkill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CReqUseSkill) GetSkillID() uint32 {
	if m != nil {
		return m.SkillID
	}
	return 0
}

// 返回玩家数据
type CRespUseSkill struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	SkillID uint32 `protobuf:"varint,3,opt,name=skillID" json:"skillID,omitempty"`
}

func (m *CRespUseSkill) Reset()                    { *m = CRespUseSkill{} }
func (m *CRespUseSkill) String() string            { return proto.CompactTextString(m) }
func (*CRespUseSkill) ProtoMessage()               {}
func (*CRespUseSkill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *CRespUseSkill) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespUseSkill) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *CRespUseSkill) GetSkillID() uint32 {
	if m != nil {
		return m.SkillID
	}
	return 0
}

type CNotifyUseSkill struct {
	UserID  uint64 `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	SkillID uint32 `protobuf:"varint,2,opt,name=skillID" json:"skillID,omitempty"`
}

func (m *CNotifyUseSkill) Reset()                    { *m = CNotifyUseSkill{} }
func (m *CNotifyUseSkill) String() string            { return proto.CompactTextString(m) }
func (*CNotifyUseSkill) ProtoMessage()               {}
func (*CNotifyUseSkill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *CNotifyUseSkill) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

func (m *CNotifyUseSkill) GetSkillID() uint32 {
	if m != nil {
		return m.SkillID
	}
	return 0
}

type CNotifyGameStart struct {
	MyGeneral *gamedef.GameGeneral `protobuf:"bytes,1,opt,name=myGeneral" json:"myGeneral,omitempty"`
	OpGeneral *gamedef.GameGeneral `protobuf:"bytes,2,opt,name=opGeneral" json:"opGeneral,omitempty"`
}

func (m *CNotifyGameStart) Reset()                    { *m = CNotifyGameStart{} }
func (m *CNotifyGameStart) String() string            { return proto.CompactTextString(m) }
func (*CNotifyGameStart) ProtoMessage()               {}
func (*CNotifyGameStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *CNotifyGameStart) GetMyGeneral() *gamedef.GameGeneral {
	if m != nil {
		return m.MyGeneral
	}
	return nil
}

func (m *CNotifyGameStart) GetOpGeneral() *gamedef.GameGeneral {
	if m != nil {
		return m.OpGeneral
	}
	return nil
}

type CNotifyGeneralStatus struct {
	GameGeneral *gamedef.GameGeneral `protobuf:"bytes,1,opt,name=gameGeneral" json:"gameGeneral,omitempty"`
}

func (m *CNotifyGeneralStatus) Reset()                    { *m = CNotifyGeneralStatus{} }
func (m *CNotifyGeneralStatus) String() string            { return proto.CompactTextString(m) }
func (*CNotifyGeneralStatus) ProtoMessage()               {}
func (*CNotifyGeneralStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *CNotifyGeneralStatus) GetGameGeneral() *gamedef.GameGeneral {
	if m != nil {
		return m.GameGeneral
	}
	return nil
}

func init() {
	proto.RegisterType((*CReqSyncSystemTime)(nil), "cmsg.CReqSyncSystemTime")
	proto.RegisterType((*CRespSyncSystemTime)(nil), "cmsg.CRespSyncSystemTime")
	proto.RegisterType((*CReqAuth)(nil), "cmsg.CReqAuth")
	proto.RegisterType((*CRespAuth)(nil), "cmsg.CRespAuth")
	proto.RegisterType((*CNotifyLoginInfo)(nil), "cmsg.CNotifyLoginInfo")
	proto.RegisterType((*CReqLogin)(nil), "cmsg.CReqLogin")
	proto.RegisterType((*CRespLogin)(nil), "cmsg.CRespLogin")
	proto.RegisterType((*CReqLogout)(nil), "cmsg.CReqLogout")
	proto.RegisterType((*CRespLogout)(nil), "cmsg.CRespLogout")
	proto.RegisterType((*CNotifyLogout)(nil), "cmsg.CNotifyLogout")
	proto.RegisterType((*CReqNotifyUserData)(nil), "cmsg.CReqNotifyUserData")
	proto.RegisterType((*CRespNotifyUserData)(nil), "cmsg.CRespNotifyUserData")
	proto.RegisterType((*CReqUseSkill)(nil), "cmsg.CReqUseSkill")
	proto.RegisterType((*CRespUseSkill)(nil), "cmsg.CRespUseSkill")
	proto.RegisterType((*CNotifyUseSkill)(nil), "cmsg.CNotifyUseSkill")
	proto.RegisterType((*CNotifyGameStart)(nil), "cmsg.CNotifyGameStart")
	proto.RegisterType((*CNotifyGeneralStatus)(nil), "cmsg.CNotifyGeneralStatus")
}

func init() { proto.RegisterFile("cmsg/client_msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 601 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0x93, 0x34, 0xbf, 0x66, 0xd2, 0xfc, 0x8a, 0x96, 0x52, 0x85, 0x0a, 0xa1, 0xb0, 0xa7,
	0x9c, 0x62, 0x30, 0x12, 0xa8, 0x12, 0x12, 0xaa, 0x5c, 0x54, 0x55, 0x2a, 0x3d, 0x6c, 0xda, 0x13,
	0x42, 0x95, 0xeb, 0x4c, 0xdc, 0x55, 0x6d, 0xaf, 0xbb, 0xbb, 0xa6, 0x0d, 0x5f, 0x85, 0x23, 0x5f,
	0x14, 0xed, 0xfa, 0x6f, 0x10, 0x54, 0x0a, 0xa7, 0xcc, 0x8c, 0xdf, 0xbc, 0x79, 0xf3, 0x76, 0x37,
	0xf0, 0x2c, 0x4c, 0x54, 0xe4, 0x86, 0x31, 0xc7, 0x54, 0x5f, 0x25, 0x2a, 0x9a, 0x65, 0x52, 0x68,
	0x41, 0x7a, 0xa6, 0x7c, 0x70, 0x16, 0x71, 0x7d, 0x93, 0x5f, 0xcf, 0x42, 0x91, 0xb8, 0xf7, 0x98,
	0x3e, 0xf0, 0xdc, 0x7b, 0x73, 0x78, 0xe8, 0x46, 0x41, 0x82, 0x0a, 0xe5, 0x37, 0x94, 0xae, 0x92,
	0xa1, 0x5b, 0x86, 0xa6, 0x6a, 0x9b, 0x6d, 0xb4, 0xc0, 0xa5, 0xfd, 0xbd, 0x5a, 0xe0, 0xb2, 0xe0,
	0xa4, 0x1e, 0x10, 0x9f, 0xe1, 0xdd, 0x7c, 0x95, 0x86, 0xf3, 0x95, 0xd2, 0x98, 0x5c, 0xf0, 0x04,
	0xc9, 0x0b, 0x18, 0x68, 0x9e, 0xa0, 0xd2, 0x41, 0x92, 0x8d, 0x9d, 0x89, 0x33, 0xed, 0xb2, 0xa6,
	0x40, 0xbf, 0xc2, 0x53, 0x9f, 0xa1, 0xca, 0x36, 0x69, 0x22, 0x53, 0xd8, 0x2d, 0x34, 0x5d, 0xd4,
	0x98, 0x8e, 0xc5, 0xfc, 0x5e, 0xa6, 0x77, 0xb0, 0x6d, 0x24, 0x1d, 0xe5, 0xfa, 0x86, 0x8c, 0xe1,
	0xbf, 0x20, 0x0c, 0x45, 0x9e, 0x6a, 0xcb, 0x38, 0x60, 0x55, 0x4a, 0x0e, 0x60, 0x3b, 0x0b, 0x94,
	0xba, 0x17, 0x72, 0x61, 0x89, 0x06, 0xac, 0xce, 0x89, 0x0b, 0x5b, 0xf8, 0xa0, 0x65, 0x30, 0xee,
	0x4e, 0x9c, 0xe9, 0xd0, 0x7b, 0x3e, 0x2b, 0x97, 0x9f, 0x7d, 0x32, 0xd5, 0xa3, 0x82, 0xe1, 0x34,
	0x5d, 0x0a, 0x56, 0xe0, 0xe8, 0x4f, 0x07, 0x06, 0x76, 0xa5, 0x6a, 0x28, 0x4a, 0xe9, 0x8b, 0x05,
	0xda, 0xa1, 0x23, 0x56, 0xa5, 0x64, 0x1f, 0xfa, 0x28, 0xe5, 0x67, 0x15, 0x95, 0x23, 0xcb, 0xac,
	0x2d, 0xb3, 0xbb, 0x2e, 0x73, 0x1f, 0xfa, 0xb9, 0x42, 0x79, 0x7a, 0x3c, 0xee, 0x4d, 0x9c, 0x69,
	0x8f, 0x95, 0x19, 0x21, 0xd0, 0x53, 0x3c, 0x4a, 0xc7, 0x5b, 0x16, 0x6e, 0x63, 0xf2, 0x12, 0x20,
	0x4f, 0x63, 0x11, 0xde, 0x1a, 0x2f, 0xc6, 0x7d, 0xeb, 0x4e, 0xab, 0x42, 0x3f, 0xc0, 0x13, 0xff,
	0x5c, 0x68, 0xbe, 0x5c, 0x9d, 0x89, 0x88, 0xa7, 0x66, 0x81, 0x47, 0x0c, 0xfa, 0x1f, 0x3a, 0x3c,
	0x2b, 0x75, 0x76, 0x78, 0x46, 0xdf, 0xdb, 0x15, 0xef, 0x6c, 0x6b, 0x4b, 0x96, 0xf3, 0x47, 0x59,
	0x9d, 0x46, 0x16, 0xfd, 0xe1, 0x00, 0x58, 0x73, 0x8a, 0xd6, 0xcd, 0xdd, 0x69, 0x86, 0x75, 0xd7,
	0x86, 0xbd, 0x82, 0x9e, 0x89, 0xac, 0x33, 0x43, 0x6f, 0x54, 0x9f, 0xd2, 0xa5, 0x42, 0xc9, 0xec,
	0x27, 0x32, 0x81, 0x21, 0x57, 0x0c, 0x43, 0x91, 0xa6, 0x18, 0x6a, 0xeb, 0xd6, 0x36, 0x6b, 0x97,
	0xe8, 0x8e, 0x15, 0x67, 0xd6, 0x12, 0xb9, 0xa6, 0x1f, 0x61, 0x58, 0x49, 0x15, 0xb9, 0xde, 0x5c,
	0x2b, 0x3d, 0x82, 0x51, 0xe3, 0xb1, 0xa1, 0x78, 0x0d, 0x83, 0xb8, 0x72, 0xdb, 0x92, 0x0c, 0x3d,
	0x52, 0x2b, 0xad, 0xcf, 0x81, 0x35, 0x20, 0xba, 0x57, 0x3c, 0xa9, 0x82, 0xc5, 0xec, 0x72, 0x1c,
	0xe8, 0x80, 0x9e, 0x94, 0x8f, 0x66, 0xbd, 0xfc, 0x0f, 0x0a, 0xa7, 0xb0, 0x63, 0xe8, 0x2f, 0x15,
	0xce, 0x6f, 0x79, 0x1c, 0x1b, 0x06, 0x65, 0x82, 0xf2, 0x2c, 0x47, 0xac, 0x4a, 0xe9, 0x17, 0x18,
	0xd9, 0x91, 0x6d, 0xe8, 0xe6, 0x17, 0xbb, 0x22, 0xef, 0xae, 0x93, 0xfb, 0xb0, 0xeb, 0xd7, 0xbb,
	0x14, 0xf4, 0x7f, 0xbb, 0x54, 0x2d, 0x92, 0xce, 0x3a, 0xc9, 0xf7, 0xfa, 0x46, 0x9f, 0x04, 0x09,
	0xce, 0x75, 0x20, 0x35, 0xf1, 0x60, 0x90, 0xac, 0x4e, 0x30, 0x45, 0x19, 0xc4, 0xa5, 0xe1, 0x7b,
	0xb5, 0xe1, 0x06, 0x56, 0x7e, 0x63, 0x0d, 0xcc, 0xf4, 0x88, 0xac, 0xea, 0xe9, 0x3c, 0xd6, 0x53,
	0xc3, 0xe8, 0x39, 0xec, 0x55, 0xb3, 0x8b, 0xca, 0x5c, 0x07, 0x3a, 0x57, 0xe4, 0x1d, 0x0c, 0xa3,
	0xa6, 0xe3, 0x51, 0x05, 0x6d, 0xe0, 0x75, 0xdf, 0xfe, 0xa1, 0xbe, 0xfd, 0x15, 0x00, 0x00, 0xff,
	0xff, 0x57, 0xde, 0x10, 0xd5, 0xbd, 0x05, 0x00, 0x00,
}
