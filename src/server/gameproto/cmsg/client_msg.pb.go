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
	CNotifyDataChange
	CReqUseSkill
	CRespUseSkill
	CNotifyUseSkill
	CNotifyGameStart
	CNotifyGeneralStatus
	CNotifyGameStage
	CNotifyGameResult
	CReqUserInit
	CRespUserInit
	CReqStageFight
	CRespStageFight
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

type CNotifyDataChangeType int32

const (
	CNotifyDataChange_DCTInvalid CNotifyDataChangeType = 0
	CNotifyDataChange_DCTUser    CNotifyDataChangeType = 1
	CNotifyDataChange_DCTGeneral CNotifyDataChangeType = 2
)

var CNotifyDataChangeType_name = map[int32]string{
	0: "DCTInvalid",
	1: "DCTUser",
	2: "DCTGeneral",
}
var CNotifyDataChangeType_value = map[string]int32{
	"DCTInvalid": 0,
	"DCTUser":    1,
	"DCTGeneral": 2,
}

func (x CNotifyDataChangeType) String() string {
	return proto.EnumName(CNotifyDataChangeType_name, int32(x))
}
func (CNotifyDataChangeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{12, 0} }

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
	ErrCode     uint32             `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg      string             `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	UserID      uint64             `protobuf:"varint,3,opt,name=userID" json:"userID,omitempty"`
	User        *gamedef.User      `protobuf:"bytes,4,opt,name=user" json:"user,omitempty"`
	IsReconnect bool               `protobuf:"varint,5,opt,name=isReconnect" json:"isReconnect,omitempty"`
	Generals    []*gamedef.General `protobuf:"bytes,6,rep,name=generals" json:"generals,omitempty"`
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

func (m *CRespLogin) GetGenerals() []*gamedef.General {
	if m != nil {
		return m.Generals
	}
	return nil
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
	ErrCode  uint32             `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg   string             `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	Generals []*gamedef.General `protobuf:"bytes,3,rep,name=generals" json:"generals,omitempty"`
	Items    []*gamedef.Item    `protobuf:"bytes,4,rep,name=Items" json:"Items,omitempty"`
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

func (m *CRespNotifyUserData) GetGenerals() []*gamedef.General {
	if m != nil {
		return m.Generals
	}
	return nil
}

func (m *CRespNotifyUserData) GetItems() []*gamedef.Item {
	if m != nil {
		return m.Items
	}
	return nil
}

type CNotifyDataChange struct {
	Changes  []CNotifyDataChangeType `protobuf:"varint,1,rep,packed,name=changes,enum=cmsg.CNotifyDataChangeType" json:"changes,omitempty"`
	User     *gamedef.User           `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Generals []*gamedef.General      `protobuf:"bytes,3,rep,name=generals" json:"generals,omitempty"`
}

func (m *CNotifyDataChange) Reset()                    { *m = CNotifyDataChange{} }
func (m *CNotifyDataChange) String() string            { return proto.CompactTextString(m) }
func (*CNotifyDataChange) ProtoMessage()               {}
func (*CNotifyDataChange) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *CNotifyDataChange) GetChanges() []CNotifyDataChangeType {
	if m != nil {
		return m.Changes
	}
	return nil
}

func (m *CNotifyDataChange) GetUser() *gamedef.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *CNotifyDataChange) GetGenerals() []*gamedef.General {
	if m != nil {
		return m.Generals
	}
	return nil
}

// ===============战斗数据====================
type CReqUseSkill struct {
	SkillID uint32 `protobuf:"varint,1,opt,name=skillID" json:"skillID,omitempty"`
}

func (m *CReqUseSkill) Reset()                    { *m = CReqUseSkill{} }
func (m *CReqUseSkill) String() string            { return proto.CompactTextString(m) }
func (*CReqUseSkill) ProtoMessage()               {}
func (*CReqUseSkill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

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
func (*CRespUseSkill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

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
func (*CNotifyUseSkill) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

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
	Generals []*gamedef.GameGeneral `protobuf:"bytes,1,rep,name=generals" json:"generals,omitempty"`
}

func (m *CNotifyGameStart) Reset()                    { *m = CNotifyGameStart{} }
func (m *CNotifyGameStart) String() string            { return proto.CompactTextString(m) }
func (*CNotifyGameStart) ProtoMessage()               {}
func (*CNotifyGameStart) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *CNotifyGameStart) GetGenerals() []*gamedef.GameGeneral {
	if m != nil {
		return m.Generals
	}
	return nil
}

type CNotifyGeneralStatus struct {
	GameGeneral *gamedef.GameGeneral `protobuf:"bytes,1,opt,name=gameGeneral" json:"gameGeneral,omitempty"`
}

func (m *CNotifyGeneralStatus) Reset()                    { *m = CNotifyGeneralStatus{} }
func (m *CNotifyGeneralStatus) String() string            { return proto.CompactTextString(m) }
func (*CNotifyGeneralStatus) ProtoMessage()               {}
func (*CNotifyGeneralStatus) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{17} }

func (m *CNotifyGeneralStatus) GetGameGeneral() *gamedef.GameGeneral {
	if m != nil {
		return m.GameGeneral
	}
	return nil
}

type CNotifyGameStage struct {
	Stage    gamedef.GameStageTyp `protobuf:"varint,1,opt,name=stage,enum=gamedef.GameStageTyp" json:"stage,omitempty"`
	LastTime int32                `protobuf:"varint,2,opt,name=lastTime" json:"lastTime,omitempty"`
}

func (m *CNotifyGameStage) Reset()                    { *m = CNotifyGameStage{} }
func (m *CNotifyGameStage) String() string            { return proto.CompactTextString(m) }
func (*CNotifyGameStage) ProtoMessage()               {}
func (*CNotifyGameStage) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{18} }

func (m *CNotifyGameStage) GetStage() gamedef.GameStageTyp {
	if m != nil {
		return m.Stage
	}
	return gamedef.GameStageTyp_GSTInvalid
}

func (m *CNotifyGameStage) GetLastTime() int32 {
	if m != nil {
		return m.LastTime
	}
	return 0
}

type CNotifyGameResult struct {
	Winner uint64 `protobuf:"varint,1,opt,name=winner" json:"winner,omitempty"`
	Exp    int32  `protobuf:"varint,2,opt,name=exp" json:"exp,omitempty"`
}

func (m *CNotifyGameResult) Reset()                    { *m = CNotifyGameResult{} }
func (m *CNotifyGameResult) String() string            { return proto.CompactTextString(m) }
func (*CNotifyGameResult) ProtoMessage()               {}
func (*CNotifyGameResult) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{19} }

func (m *CNotifyGameResult) GetWinner() uint64 {
	if m != nil {
		return m.Winner
	}
	return 0
}

func (m *CNotifyGameResult) GetExp() int32 {
	if m != nil {
		return m.Exp
	}
	return 0
}

// ===============Game数据====================
// 创建角色
type CReqUserInit struct {
	NickName     string `protobuf:"bytes,1,opt,name=nickName" json:"nickName,omitempty"`
	FirstGeneral uint32 `protobuf:"varint,2,opt,name=firstGeneral" json:"firstGeneral,omitempty"`
}

func (m *CReqUserInit) Reset()                    { *m = CReqUserInit{} }
func (m *CReqUserInit) String() string            { return proto.CompactTextString(m) }
func (*CReqUserInit) ProtoMessage()               {}
func (*CReqUserInit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{20} }

func (m *CReqUserInit) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CReqUserInit) GetFirstGeneral() uint32 {
	if m != nil {
		return m.FirstGeneral
	}
	return 0
}

type CRespUserInit struct {
	ErrCode      uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg       string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
	NickName     string `protobuf:"bytes,3,opt,name=nickName" json:"nickName,omitempty"`
	FirstGeneral uint32 `protobuf:"varint,4,opt,name=firstGeneral" json:"firstGeneral,omitempty"`
}

func (m *CRespUserInit) Reset()                    { *m = CRespUserInit{} }
func (m *CRespUserInit) String() string            { return proto.CompactTextString(m) }
func (*CRespUserInit) ProtoMessage()               {}
func (*CRespUserInit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{21} }

func (m *CRespUserInit) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespUserInit) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
}

func (m *CRespUserInit) GetNickName() string {
	if m != nil {
		return m.NickName
	}
	return ""
}

func (m *CRespUserInit) GetFirstGeneral() uint32 {
	if m != nil {
		return m.FirstGeneral
	}
	return 0
}

type CReqStageFight struct {
	StageID uint32 `protobuf:"varint,1,opt,name=stageID" json:"stageID,omitempty"`
}

func (m *CReqStageFight) Reset()                    { *m = CReqStageFight{} }
func (m *CReqStageFight) String() string            { return proto.CompactTextString(m) }
func (*CReqStageFight) ProtoMessage()               {}
func (*CReqStageFight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{22} }

func (m *CReqStageFight) GetStageID() uint32 {
	if m != nil {
		return m.StageID
	}
	return 0
}

type CRespStageFight struct {
	ErrCode uint32 `protobuf:"varint,1,opt,name=errCode" json:"errCode,omitempty"`
	ErrMsg  string `protobuf:"bytes,2,opt,name=errMsg" json:"errMsg,omitempty"`
}

func (m *CRespStageFight) Reset()                    { *m = CRespStageFight{} }
func (m *CRespStageFight) String() string            { return proto.CompactTextString(m) }
func (*CRespStageFight) ProtoMessage()               {}
func (*CRespStageFight) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{23} }

func (m *CRespStageFight) GetErrCode() uint32 {
	if m != nil {
		return m.ErrCode
	}
	return 0
}

func (m *CRespStageFight) GetErrMsg() string {
	if m != nil {
		return m.ErrMsg
	}
	return ""
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
	proto.RegisterType((*CNotifyDataChange)(nil), "cmsg.CNotifyDataChange")
	proto.RegisterType((*CReqUseSkill)(nil), "cmsg.CReqUseSkill")
	proto.RegisterType((*CRespUseSkill)(nil), "cmsg.CRespUseSkill")
	proto.RegisterType((*CNotifyUseSkill)(nil), "cmsg.CNotifyUseSkill")
	proto.RegisterType((*CNotifyGameStart)(nil), "cmsg.CNotifyGameStart")
	proto.RegisterType((*CNotifyGeneralStatus)(nil), "cmsg.CNotifyGeneralStatus")
	proto.RegisterType((*CNotifyGameStage)(nil), "cmsg.CNotifyGameStage")
	proto.RegisterType((*CNotifyGameResult)(nil), "cmsg.CNotifyGameResult")
	proto.RegisterType((*CReqUserInit)(nil), "cmsg.CReqUserInit")
	proto.RegisterType((*CRespUserInit)(nil), "cmsg.CRespUserInit")
	proto.RegisterType((*CReqStageFight)(nil), "cmsg.CReqStageFight")
	proto.RegisterType((*CRespStageFight)(nil), "cmsg.CRespStageFight")
	proto.RegisterEnum("cmsg.CNotifyDataChangeType", CNotifyDataChangeType_name, CNotifyDataChangeType_value)
}

func init() { proto.RegisterFile("cmsg/client_msg.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 865 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x56, 0xcd, 0x8e, 0x1b, 0x45,
	0x10, 0x66, 0x3c, 0xb6, 0x77, 0x5d, 0x5e, 0x3b, 0xa6, 0xd9, 0x44, 0x43, 0x14, 0x21, 0xd3, 0x5c,
	0x2c, 0x40, 0x36, 0x38, 0x52, 0x50, 0x24, 0x10, 0x5a, 0x8d, 0x21, 0xb2, 0x14, 0xf6, 0xd0, 0x76,
	0x4e, 0x11, 0x8a, 0x26, 0xe3, 0xf6, 0xb8, 0xb5, 0xf3, 0xb7, 0xdd, 0x3d, 0xd9, 0xf5, 0x91, 0xd7,
	0xe0, 0xca, 0xd3, 0x70, 0xe4, 0x8d, 0x50, 0xd7, 0xfc, 0x3a, 0xd9, 0xb0, 0x32, 0x27, 0x57, 0x55,
	0x57, 0x7f, 0xfd, 0xd5, 0x57, 0xd5, 0x3d, 0x86, 0x87, 0x7e, 0xa4, 0x82, 0x99, 0x1f, 0x0a, 0x1e,
	0xeb, 0x37, 0x91, 0x0a, 0xa6, 0xa9, 0x4c, 0x74, 0x42, 0xda, 0x26, 0xfc, 0xf8, 0x65, 0x20, 0xf4,
	0x2e, 0x7b, 0x3b, 0xf5, 0x93, 0x68, 0x76, 0xc3, 0xe3, 0x5b, 0x91, 0xcd, 0xbf, 0x7f, 0xfe, 0x7c,
	0x16, 0x78, 0x11, 0x57, 0x5c, 0xbe, 0xe3, 0x72, 0xa6, 0xa4, 0x3f, 0x2b, 0x4c, 0x13, 0xc5, 0xcd,
	0x68, 0x6d, 0xf8, 0x16, 0x7f, 0xdf, 0x6c, 0xf8, 0x36, 0xc7, 0xa4, 0x73, 0x20, 0x2e, 0xe3, 0xd7,
	0xab, 0x7d, 0xec, 0xaf, 0xf6, 0x4a, 0xf3, 0x68, 0x2d, 0x22, 0x4e, 0x9e, 0x40, 0x4f, 0x8b, 0x88,
	0x2b, 0xed, 0x45, 0xa9, 0x63, 0x8d, 0xad, 0x89, 0xcd, 0xea, 0x00, 0xfd, 0x1d, 0x3e, 0x73, 0x19,
	0x57, 0xe9, 0x31, 0x9b, 0xc8, 0x04, 0x1e, 0xe4, 0x9c, 0xd6, 0x55, 0x4e, 0x0b, 0x73, 0xde, 0x0f,
	0xd3, 0x6b, 0x38, 0x35, 0x94, 0x2e, 0x32, 0xbd, 0x23, 0x0e, 0x9c, 0x78, 0xbe, 0x9f, 0x64, 0xb1,
	0x46, 0xc4, 0x1e, 0x2b, 0x5d, 0xf2, 0x18, 0x4e, 0x53, 0x4f, 0xa9, 0x9b, 0x44, 0x6e, 0x10, 0xa8,
	0xc7, 0x2a, 0x9f, 0xcc, 0xa0, 0xc3, 0x6f, 0xb5, 0xf4, 0x1c, 0x7b, 0x6c, 0x4d, 0xfa, 0xf3, 0xcf,
	0xa7, 0x45, 0xf1, 0xd3, 0x5f, 0x4c, 0xf4, 0x22, 0x47, 0x58, 0xc6, 0xdb, 0x84, 0xe5, 0x79, 0xf4,
	0x2f, 0x0b, 0x7a, 0x58, 0x52, 0x79, 0x28, 0x97, 0xd2, 0x4d, 0x36, 0x1c, 0x0f, 0x1d, 0xb0, 0xd2,
	0x25, 0x8f, 0xa0, 0xcb, 0xa5, 0xfc, 0x4d, 0x05, 0xc5, 0x91, 0x85, 0xd7, 0xa4, 0x69, 0x1f, 0xd2,
	0x7c, 0x04, 0xdd, 0x4c, 0x71, 0xb9, 0x5c, 0x38, 0xed, 0xb1, 0x35, 0x69, 0xb3, 0xc2, 0x23, 0x04,
	0xda, 0x4a, 0x04, 0xb1, 0xd3, 0xc1, 0x74, 0xb4, 0xc9, 0x17, 0x00, 0x59, 0x1c, 0x26, 0xfe, 0x95,
	0xd1, 0xc2, 0xe9, 0xa2, 0x3a, 0x8d, 0x08, 0xfd, 0x11, 0x46, 0xee, 0x65, 0xa2, 0xc5, 0x76, 0xff,
	0x32, 0x09, 0x44, 0x6c, 0x0a, 0xf8, 0x0f, 0x81, 0x86, 0xd0, 0x12, 0x69, 0xc1, 0xb3, 0x25, 0x52,
	0xfa, 0x03, 0x96, 0x78, 0x8d, 0x5b, 0x1b, 0xb4, 0xac, 0x3b, 0x69, 0xb5, 0x6a, 0x5a, 0xf4, 0x6f,
	0x0b, 0x00, 0xc5, 0xc9, 0xb7, 0x1e, 0xaf, 0x4e, 0x7d, 0x98, 0x7d, 0x70, 0xd8, 0x97, 0xd0, 0x36,
	0x16, 0x2a, 0xd3, 0x9f, 0x0f, 0xaa, 0x2e, 0xbd, 0x52, 0x5c, 0x32, 0x5c, 0x22, 0x63, 0xe8, 0x0b,
	0xc5, 0xb8, 0x9f, 0xc4, 0x31, 0xf7, 0x35, 0xaa, 0x75, 0xca, 0x9a, 0x21, 0xf2, 0x2d, 0x9c, 0x06,
	0x3c, 0xe6, 0xd2, 0x0b, 0x95, 0xd3, 0x1d, 0xdb, 0x93, 0xfe, 0x7c, 0x54, 0x01, 0xbd, 0xc8, 0x17,
	0x58, 0x95, 0x41, 0xcf, 0xb0, 0x14, 0x23, 0x42, 0x92, 0x69, 0xfa, 0x33, 0xf4, 0xcb, 0xc2, 0x92,
	0x4c, 0x1f, 0x5f, 0x19, 0xbd, 0x80, 0x41, 0xdd, 0x11, 0x03, 0xf1, 0x1d, 0xf4, 0xc2, 0xb2, 0x37,
	0x08, 0xd2, 0x9f, 0x93, 0x8a, 0x4e, 0xd5, 0x35, 0x56, 0x27, 0xd1, 0xf3, 0xfc, 0x02, 0xe6, 0x28,
	0xa6, 0xf2, 0x85, 0xa7, 0x3d, 0xfa, 0xa7, 0x55, 0xdc, 0xb1, 0xc3, 0xf8, 0xff, 0x10, 0xbf, 0xa9,
	0x8f, 0x7d, 0x9f, 0x3e, 0xe4, 0x2b, 0xe8, 0x2c, 0x35, 0x8f, 0x94, 0xd3, 0xc6, 0xd4, 0xba, 0x27,
	0x26, 0xca, 0xf2, 0x35, 0xfa, 0x8f, 0x05, 0x9f, 0x16, 0x65, 0x1b, 0x52, 0xee, 0xce, 0x8b, 0x03,
	0x4e, 0x9e, 0xc1, 0x89, 0x8f, 0x96, 0x72, 0xac, 0xb1, 0x3d, 0x19, 0xce, 0x9f, 0x4c, 0xcd, 0x7b,
	0x35, 0xfd, 0x20, 0x73, 0xaa, 0xf7, 0x29, 0x67, 0x65, 0x72, 0x35, 0x05, 0xad, 0x8f, 0x4f, 0xc1,
	0x51, 0x35, 0xd0, 0xa7, 0xd0, 0x36, 0x27, 0x90, 0x21, 0xc0, 0xc2, 0x5d, 0x2f, 0xe3, 0x77, 0x5e,
	0x28, 0x36, 0xa3, 0x4f, 0x48, 0x1f, 0x4e, 0x16, 0xee, 0xda, 0xc0, 0x8e, 0xac, 0x62, 0xb1, 0xd8,
	0x3c, 0x6a, 0xd1, 0x09, 0x9c, 0x99, 0x36, 0xbc, 0x52, 0x7c, 0x75, 0x25, 0xc2, 0xd0, 0x08, 0xad,
	0x8c, 0x51, 0xdc, 0x90, 0x01, 0x2b, 0x5d, 0xfa, 0x1a, 0x06, 0xd8, 0x99, 0x66, 0xea, 0xf1, 0xcf,
	0x45, 0x09, 0x6e, 0x1f, 0x82, 0xbb, 0xf0, 0xc0, 0xad, 0x5a, 0x9e, 0xc3, 0x7f, 0xec, 0xaa, 0x36,
	0x40, 0x5a, 0x87, 0x20, 0x8b, 0xea, 0x9d, 0x78, 0xe1, 0x45, 0x7c, 0xa5, 0x3d, 0x69, 0x06, 0xb3,
	0x96, 0xd0, 0x42, 0x09, 0xcf, 0x6b, 0x09, 0xbd, 0x88, 0x7f, 0x28, 0xe3, 0x25, 0x9c, 0x97, 0x28,
	0x79, 0x68, 0xa5, 0x3d, 0x9d, 0x29, 0xf2, 0x0c, 0xfa, 0x41, 0xbd, 0xa1, 0x18, 0xf2, 0xbb, 0xc1,
	0x9a, 0x89, 0xf4, 0xf5, 0xfb, 0xac, 0x02, 0x4e, 0xbe, 0x81, 0x8e, 0x32, 0x06, 0xa2, 0x0c, 0xe7,
	0x0f, 0x0f, 0x50, 0x30, 0x65, 0xbd, 0x4f, 0x59, 0x9e, 0x63, 0x5e, 0xfc, 0xd0, 0x53, 0x1a, 0x1f,
	0x47, 0x53, 0x71, 0x87, 0x55, 0x3e, 0xfd, 0xa9, 0x9a, 0x48, 0xb3, 0x93, 0x71, 0x95, 0x85, 0xf8,
	0xf6, 0xde, 0x88, 0x38, 0xe6, 0xb2, 0x54, 0x2e, 0xf7, 0xc8, 0x08, 0x6c, 0x7e, 0x9b, 0x16, 0x18,
	0xc6, 0xa4, 0x97, 0x55, 0xf7, 0xe5, 0x32, 0x16, 0xf8, 0x71, 0x89, 0x85, 0x7f, 0x75, 0xe9, 0x45,
	0xbc, 0x78, 0x56, 0x2b, 0x9f, 0x50, 0x38, 0xdb, 0x0a, 0xa9, 0x74, 0x29, 0x40, 0x2e, 0xfe, 0x41,
	0x8c, 0xfe, 0x61, 0xd5, 0x43, 0x92, 0x23, 0x1e, 0x3f, 0x24, 0x4d, 0x0e, 0xf6, 0x3d, 0x1c, 0xda,
	0x77, 0x70, 0xf8, 0x1a, 0x86, 0xf8, 0x65, 0x37, 0xda, 0xfd, 0x2a, 0x82, 0x1d, 0x72, 0x40, 0x25,
	0x1b, 0x33, 0x9d, 0xbb, 0x38, 0x76, 0xf8, 0x45, 0x3f, 0x48, 0x3e, 0x8e, 0xf0, 0xdb, 0x2e, 0xfe,
	0xa3, 0x78, 0xfa, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xd1, 0x8a, 0xec, 0x84, 0xbe, 0x08, 0x00,
	0x00,
}
