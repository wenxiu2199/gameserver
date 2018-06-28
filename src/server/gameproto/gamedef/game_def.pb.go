// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gamedef/game_def.proto

/*
Package gamedef is a generated protocol buffer package.

It is generated from these files:
	gamedef/game_def.proto

It has these top-level messages:
	LoginInfo
	Account
	ExtraAccountInfo
	User
	Individual
	General
	Item
	UserData
	Buff
	GameGeneral
	Chess
*/
package gamedef

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import gameconf "github.com/wenxiu2199/gameserver/src/server/gameproto/gameconf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GameStageTyp int32

const (
	GameStageTyp_GSTInvalid GameStageTyp = 0
	GameStageTyp_GSTChoose  GameStageTyp = 1
	GameStageTyp_GSTAction  GameStageTyp = 2
)

var GameStageTyp_name = map[int32]string{
	0: "GSTInvalid",
	1: "GSTChoose",
	2: "GSTAction",
}
var GameStageTyp_value = map[string]int32{
	"GSTInvalid": 0,
	"GSTChoose":  1,
	"GSTAction":  2,
}

func (x GameStageTyp) String() string {
	return proto.EnumName(GameStageTyp_name, int32(x))
}
func (GameStageTyp) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// ===================================
type ChessTyp int32

const (
	ChessTyp_CTInvalid ChessTyp = 0
	ChessTyp_CTBlank   ChessTyp = 1
	ChessTyp_CTMine    ChessTyp = 2
)

var ChessTyp_name = map[int32]string{
	0: "CTInvalid",
	1: "CTBlank",
	2: "CTMine",
}
var ChessTyp_value = map[string]int32{
	"CTInvalid": 0,
	"CTBlank":   1,
	"CTMine":    2,
}

func (x ChessTyp) String() string {
	return proto.EnumName(ChessTyp_name, int32(x))
}
func (ChessTyp) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

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

type Account struct {
	Account string `protobuf:"bytes,1,opt,name=Account" json:"Account,omitempty"`
	UserID  uint64 `protobuf:"varint,2,opt,name=userID" json:"userID,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Account) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Account) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

// 账号额外信息
type ExtraAccountInfo struct {
	AccountType uint32 `protobuf:"varint,1,opt,name=accountType" json:"accountType,omitempty"`
	Channel     uint32 `protobuf:"varint,2,opt,name=channel" json:"channel,omitempty"`
	Platform    uint32 `protobuf:"varint,3,opt,name=platform" json:"platform,omitempty"`
	Ip          string `protobuf:"bytes,4,opt,name=ip" json:"ip,omitempty"`
	Account     string `protobuf:"bytes,5,opt,name=account" json:"account,omitempty"`
}

func (m *ExtraAccountInfo) Reset()                    { *m = ExtraAccountInfo{} }
func (m *ExtraAccountInfo) String() string            { return proto.CompactTextString(m) }
func (*ExtraAccountInfo) ProtoMessage()               {}
func (*ExtraAccountInfo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

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

func (m *ExtraAccountInfo) GetIp() string {
	if m != nil {
		return m.Ip
	}
	return ""
}

func (m *ExtraAccountInfo) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// 用户信息
type User struct {
	UserID         uint64 `protobuf:"varint,1,opt,name=userID" json:"userID,omitempty"`
	Nickname       string `protobuf:"bytes,2,opt,name=nickname" json:"nickname,omitempty"`
	SexID          uint32 `protobuf:"varint,3,opt,name=sexID" json:"sexID,omitempty"`
	FightGeneralID uint32 `protobuf:"varint,4,opt,name=fightGeneralID" json:"fightGeneralID,omitempty"`
}

func (m *User) Reset()                    { *m = User{} }
func (m *User) String() string            { return proto.CompactTextString(m) }
func (*User) ProtoMessage()               {}
func (*User) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

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

func (m *User) GetFightGeneralID() uint32 {
	if m != nil {
		return m.FightGeneralID
	}
	return 0
}

type Individual struct {
	Hp        int32 `protobuf:"varint,1,opt,name=hp" json:"hp,omitempty"`
	Attack    int32 `protobuf:"varint,2,opt,name=attack" json:"attack,omitempty"`
	Defense   int32 `protobuf:"varint,3,opt,name=defense" json:"defense,omitempty"`
	SpAttack  int32 `protobuf:"varint,4,opt,name=spAttack" json:"spAttack,omitempty"`
	SpDefense int32 `protobuf:"varint,5,opt,name=spDefense" json:"spDefense,omitempty"`
	Speed     int32 `protobuf:"varint,6,opt,name=speed" json:"speed,omitempty"`
}

func (m *Individual) Reset()                    { *m = Individual{} }
func (m *Individual) String() string            { return proto.CompactTextString(m) }
func (*Individual) ProtoMessage()               {}
func (*Individual) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Individual) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *Individual) GetAttack() int32 {
	if m != nil {
		return m.Attack
	}
	return 0
}

func (m *Individual) GetDefense() int32 {
	if m != nil {
		return m.Defense
	}
	return 0
}

func (m *Individual) GetSpAttack() int32 {
	if m != nil {
		return m.SpAttack
	}
	return 0
}

func (m *Individual) GetSpDefense() int32 {
	if m != nil {
		return m.SpDefense
	}
	return 0
}

func (m *Individual) GetSpeed() int32 {
	if m != nil {
		return m.Speed
	}
	return 0
}

// 武将信息
type General struct {
	PkID       uint64      `protobuf:"varint,1,opt,name=pkID" json:"pkID,omitempty"`
	GeneralID  uint32      `protobuf:"varint,2,opt,name=generalID" json:"generalID,omitempty"`
	Individual *Individual `protobuf:"bytes,3,opt,name=individual" json:"individual,omitempty"`
	Effort     *Individual `protobuf:"bytes,4,opt,name=effort" json:"effort,omitempty"`
	Skills     []uint32    `protobuf:"varint,5,rep,packed,name=skills" json:"skills,omitempty"`
	Level      uint32      `protobuf:"varint,6,opt,name=level" json:"level,omitempty"`
}

func (m *General) Reset()                    { *m = General{} }
func (m *General) String() string            { return proto.CompactTextString(m) }
func (*General) ProtoMessage()               {}
func (*General) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *General) GetPkID() uint64 {
	if m != nil {
		return m.PkID
	}
	return 0
}

func (m *General) GetGeneralID() uint32 {
	if m != nil {
		return m.GeneralID
	}
	return 0
}

func (m *General) GetIndividual() *Individual {
	if m != nil {
		return m.Individual
	}
	return nil
}

func (m *General) GetEffort() *Individual {
	if m != nil {
		return m.Effort
	}
	return nil
}

func (m *General) GetSkills() []uint32 {
	if m != nil {
		return m.Skills
	}
	return nil
}

func (m *General) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type Item struct {
	PkID   uint64 `protobuf:"varint,1,opt,name=pkID" json:"pkID,omitempty"`
	ItemID uint32 `protobuf:"varint,2,opt,name=itemID" json:"itemID,omitempty"`
	Num    uint32 `protobuf:"varint,3,opt,name=num" json:"num,omitempty"`
}

func (m *Item) Reset()                    { *m = Item{} }
func (m *Item) String() string            { return proto.CompactTextString(m) }
func (*Item) ProtoMessage()               {}
func (*Item) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Item) GetPkID() uint64 {
	if m != nil {
		return m.PkID
	}
	return 0
}

func (m *Item) GetItemID() uint32 {
	if m != nil {
		return m.ItemID
	}
	return 0
}

func (m *Item) GetNum() uint32 {
	if m != nil {
		return m.Num
	}
	return 0
}

type UserData struct {
	User     *User      `protobuf:"bytes,1,opt,name=user" json:"user,omitempty"`
	Generals []*General `protobuf:"bytes,2,rep,name=generals" json:"generals,omitempty"`
	Items    []*Item    `protobuf:"bytes,3,rep,name=items" json:"items,omitempty"`
}

func (m *UserData) Reset()                    { *m = UserData{} }
func (m *UserData) String() string            { return proto.CompactTextString(m) }
func (*UserData) ProtoMessage()               {}
func (*UserData) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *UserData) GetUser() *User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *UserData) GetGenerals() []*General {
	if m != nil {
		return m.Generals
	}
	return nil
}

func (m *UserData) GetItems() []*Item {
	if m != nil {
		return m.Items
	}
	return nil
}

// ===============Game数据====================
type Buff struct {
	BuffType gameconf.SkillEffectTyp `protobuf:"varint,1,opt,name=buffType,enum=gameconf.SkillEffectTyp" json:"buffType,omitempty"`
	Last     int32                   `protobuf:"varint,2,opt,name=last" json:"last,omitempty"`
	Level    int32                   `protobuf:"varint,3,opt,name=level" json:"level,omitempty"`
}

func (m *Buff) Reset()                    { *m = Buff{} }
func (m *Buff) String() string            { return proto.CompactTextString(m) }
func (*Buff) ProtoMessage()               {}
func (*Buff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *Buff) GetBuffType() gameconf.SkillEffectTyp {
	if m != nil {
		return m.BuffType
	}
	return gameconf.SkillEffectTyp_SETInvliad
}

func (m *Buff) GetLast() int32 {
	if m != nil {
		return m.Last
	}
	return 0
}

func (m *Buff) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type GameGeneral struct {
	GeneralID     uint32   `protobuf:"varint,1,opt,name=generalID" json:"generalID,omitempty"`
	Level         uint32   `protobuf:"varint,2,opt,name=level" json:"level,omitempty"`
	BaseHp        int32    `protobuf:"varint,3,opt,name=baseHp" json:"baseHp,omitempty"`
	BaseAttack    int32    `protobuf:"varint,4,opt,name=baseAttack" json:"baseAttack,omitempty"`
	BaseDefense   int32    `protobuf:"varint,5,opt,name=baseDefense" json:"baseDefense,omitempty"`
	BaseSpAttack  int32    `protobuf:"varint,6,opt,name=baseSpAttack" json:"baseSpAttack,omitempty"`
	BaseSpDefense int32    `protobuf:"varint,7,opt,name=baseSpDefense" json:"baseSpDefense,omitempty"`
	BaseSpeed     int32    `protobuf:"varint,8,opt,name=baseSpeed" json:"baseSpeed,omitempty"`
	CurHP         int32    `protobuf:"varint,9,opt,name=curHP" json:"curHP,omitempty"`
	Buffs         []*Buff  `protobuf:"bytes,10,rep,name=buffs" json:"buffs,omitempty"`
	Skills        []uint32 `protobuf:"varint,11,rep,packed,name=skills" json:"skills,omitempty"`
	UserID        uint64   `protobuf:"varint,12,opt,name=userID" json:"userID,omitempty"`
}

func (m *GameGeneral) Reset()                    { *m = GameGeneral{} }
func (m *GameGeneral) String() string            { return proto.CompactTextString(m) }
func (*GameGeneral) ProtoMessage()               {}
func (*GameGeneral) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *GameGeneral) GetGeneralID() uint32 {
	if m != nil {
		return m.GeneralID
	}
	return 0
}

func (m *GameGeneral) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *GameGeneral) GetBaseHp() int32 {
	if m != nil {
		return m.BaseHp
	}
	return 0
}

func (m *GameGeneral) GetBaseAttack() int32 {
	if m != nil {
		return m.BaseAttack
	}
	return 0
}

func (m *GameGeneral) GetBaseDefense() int32 {
	if m != nil {
		return m.BaseDefense
	}
	return 0
}

func (m *GameGeneral) GetBaseSpAttack() int32 {
	if m != nil {
		return m.BaseSpAttack
	}
	return 0
}

func (m *GameGeneral) GetBaseSpDefense() int32 {
	if m != nil {
		return m.BaseSpDefense
	}
	return 0
}

func (m *GameGeneral) GetBaseSpeed() int32 {
	if m != nil {
		return m.BaseSpeed
	}
	return 0
}

func (m *GameGeneral) GetCurHP() int32 {
	if m != nil {
		return m.CurHP
	}
	return 0
}

func (m *GameGeneral) GetBuffs() []*Buff {
	if m != nil {
		return m.Buffs
	}
	return nil
}

func (m *GameGeneral) GetSkills() []uint32 {
	if m != nil {
		return m.Skills
	}
	return nil
}

func (m *GameGeneral) GetUserID() uint64 {
	if m != nil {
		return m.UserID
	}
	return 0
}

type Chess struct {
	ChessType ChessTyp `protobuf:"varint,1,opt,name=chessType,enum=gamedef.ChessTyp" json:"chessType,omitempty"`
	Visible   bool     `protobuf:"varint,2,opt,name=visible" json:"visible,omitempty"`
	MineNum   uint32   `protobuf:"varint,3,opt,name=mineNum" json:"mineNum,omitempty"`
}

func (m *Chess) Reset()                    { *m = Chess{} }
func (m *Chess) String() string            { return proto.CompactTextString(m) }
func (*Chess) ProtoMessage()               {}
func (*Chess) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *Chess) GetChessType() ChessTyp {
	if m != nil {
		return m.ChessType
	}
	return ChessTyp_CTInvalid
}

func (m *Chess) GetVisible() bool {
	if m != nil {
		return m.Visible
	}
	return false
}

func (m *Chess) GetMineNum() uint32 {
	if m != nil {
		return m.MineNum
	}
	return 0
}

func init() {
	proto.RegisterType((*LoginInfo)(nil), "gamedef.LoginInfo")
	proto.RegisterType((*Account)(nil), "gamedef.Account")
	proto.RegisterType((*ExtraAccountInfo)(nil), "gamedef.ExtraAccountInfo")
	proto.RegisterType((*User)(nil), "gamedef.User")
	proto.RegisterType((*Individual)(nil), "gamedef.Individual")
	proto.RegisterType((*General)(nil), "gamedef.General")
	proto.RegisterType((*Item)(nil), "gamedef.Item")
	proto.RegisterType((*UserData)(nil), "gamedef.UserData")
	proto.RegisterType((*Buff)(nil), "gamedef.Buff")
	proto.RegisterType((*GameGeneral)(nil), "gamedef.GameGeneral")
	proto.RegisterType((*Chess)(nil), "gamedef.Chess")
	proto.RegisterEnum("gamedef.GameStageTyp", GameStageTyp_name, GameStageTyp_value)
	proto.RegisterEnum("gamedef.ChessTyp", ChessTyp_name, ChessTyp_value)
}

func init() { proto.RegisterFile("gamedef/game_def.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 840 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x55, 0xcd, 0x8e, 0xe3, 0x44,
	0x10, 0xc6, 0x4e, 0x9c, 0x9f, 0xca, 0x64, 0x14, 0x1a, 0x34, 0xb2, 0x56, 0x08, 0x0d, 0x06, 0xa1,
	0xd1, 0x82, 0x12, 0x91, 0xe1, 0xb2, 0x82, 0xcb, 0x6c, 0xb2, 0x9a, 0x8d, 0x04, 0x2b, 0xd4, 0x09,
	0x67, 0xd4, 0x71, 0xda, 0x49, 0x2b, 0x76, 0xdb, 0x72, 0xdb, 0x61, 0xe6, 0xc6, 0x43, 0x70, 0xe6,
	0x85, 0x78, 0x17, 0x9e, 0x01, 0x55, 0x77, 0xbb, 0xe3, 0x8c, 0xe0, 0x94, 0xfa, 0xea, 0xff, 0xab,
	0x2a, 0x77, 0xe0, 0x66, 0xcf, 0x32, 0xbe, 0xe3, 0xc9, 0x0c, 0x7f, 0x7f, 0xdb, 0xf1, 0x64, 0x5a,
	0x94, 0x79, 0x95, 0x93, 0xbe, 0xd5, 0xbf, 0xfa, 0xb0, 0x17, 0xd5, 0xa1, 0xde, 0x4e, 0xe3, 0x3c,
	0x9b, 0xfd, 0xce, 0xe5, 0x93, 0xa8, 0xe7, 0xdf, 0xbd, 0x79, 0xa3, 0xdd, 0x15, 0x2f, 0x4f, 0xbc,
	0x9c, 0xa9, 0x32, 0x9e, 0x59, 0x11, 0xb5, 0x3a, 0x5e, 0x4b, 0x71, 0x2e, 0x6d, 0xde, 0xea, 0xb9,
	0xe0, 0x26, 0x71, 0x74, 0x0f, 0xc3, 0x9f, 0xf2, 0xbd, 0x90, 0x2b, 0x99, 0xe4, 0xe4, 0x1a, 0x7c,
	0x51, 0x84, 0xde, 0xad, 0x77, 0x37, 0xa4, 0xbe, 0x28, 0xc8, 0x0d, 0xf4, 0x4a, 0xbe, 0x17, 0xb9,
	0x0c, 0x7d, 0xad, 0xb3, 0x28, 0xfa, 0x01, 0xfa, 0x0f, 0x71, 0x9c, 0xd7, 0xb2, 0x22, 0xa1, 0x13,
	0x6d, 0x9c, 0xb3, 0xdc, 0x40, 0xaf, 0x56, 0xbc, 0x5c, 0x2d, 0x75, 0x70, 0x97, 0x5a, 0x14, 0xfd,
	0xe9, 0xc1, 0xe4, 0xdd, 0x53, 0x55, 0x32, 0xeb, 0xa8, 0x2b, 0xdf, 0xc2, 0x88, 0x19, 0xb8, 0x79,
	0x2e, 0xb8, 0x4e, 0x35, 0xa6, 0x6d, 0x15, 0x16, 0x8a, 0x0f, 0x4c, 0x4a, 0x9e, 0xea, 0x7c, 0x63,
	0xda, 0x40, 0xf2, 0x0a, 0x06, 0x45, 0xca, 0xaa, 0x24, 0x2f, 0xb3, 0xb0, 0xa3, 0x4d, 0x0e, 0x5b,
	0x46, 0x5d, 0xc7, 0x28, 0x84, 0xbe, 0x4d, 0x1a, 0x06, 0xa6, 0x5d, 0x0b, 0xa3, 0x27, 0xe8, 0xfe,
	0xaa, 0x78, 0xd9, 0x6a, 0xdb, 0x6b, 0xb7, 0x8d, 0x55, 0xa4, 0x88, 0x8f, 0x92, 0x65, 0xdc, 0x4e,
	0xc3, 0x61, 0xf2, 0x29, 0x04, 0x8a, 0x3f, 0xad, 0x96, 0xb6, 0xbc, 0x01, 0xe4, 0x6b, 0xb8, 0x4e,
	0xc4, 0xfe, 0x50, 0x3d, 0x72, 0xc9, 0x4b, 0x96, 0xae, 0x96, 0xba, 0x8f, 0x31, 0x7d, 0xa1, 0x8d,
	0xfe, 0xf2, 0x00, 0x56, 0x72, 0x27, 0x4e, 0x62, 0x57, 0xb3, 0x14, 0x5b, 0x3e, 0x98, 0x25, 0x04,
	0xd4, 0x3f, 0xe8, 0x25, 0xb0, 0xaa, 0x62, 0xf1, 0x51, 0x97, 0x0d, 0xa8, 0x45, 0x48, 0x65, 0xc7,
	0x13, 0x2e, 0x15, 0xd7, 0x65, 0x03, 0xda, 0x40, 0x6c, 0x55, 0x15, 0x0f, 0x26, 0xa6, 0xab, 0x4d,
	0x0e, 0x93, 0xcf, 0x60, 0xa8, 0x8a, 0xa5, 0x8d, 0x0b, 0xb4, 0xf1, 0xac, 0xd0, 0x44, 0x0a, 0xce,
	0x77, 0x61, 0x4f, 0x5b, 0x0c, 0x88, 0xfe, 0xf6, 0xa0, 0x6f, 0xdb, 0x25, 0x04, 0xba, 0xc5, 0xd1,
	0x0d, 0x47, 0xcb, 0x98, 0x73, 0xef, 0x38, 0x9a, 0xe5, 0x9c, 0x15, 0xe4, 0x1e, 0x40, 0x38, 0x76,
	0xba, 0xd5, 0xd1, 0xfc, 0x93, 0xa9, 0xbd, 0xe7, 0xe9, 0x99, 0x38, 0x6d, 0xb9, 0x91, 0x6f, 0xa0,
	0xc7, 0x93, 0x24, 0x2f, 0x2b, 0x4d, 0xe0, 0x7f, 0x02, 0xac, 0x0b, 0x4e, 0x48, 0x1d, 0x45, 0x9a,
	0xaa, 0x30, 0xb8, 0xed, 0xdc, 0x8d, 0xa9, 0x45, 0xc8, 0x26, 0xe5, 0x27, 0x9e, 0x6a, 0x36, 0x63,
	0x6a, 0x40, 0xb4, 0x84, 0xee, 0xaa, 0xe2, 0xd9, 0x7f, 0x32, 0xb9, 0x81, 0x9e, 0xa8, 0x78, 0xe6,
	0x68, 0x58, 0x44, 0x26, 0xd0, 0x91, 0x75, 0x73, 0x5d, 0x28, 0x46, 0x7f, 0x78, 0x30, 0xc0, 0x7b,
	0x59, 0xb2, 0x8a, 0x91, 0x2f, 0xa0, 0x8b, 0x57, 0xa2, 0x53, 0x8d, 0xe6, 0x63, 0xd7, 0x2b, 0x3a,
	0x50, 0x6d, 0x22, 0xdf, 0xc2, 0xc0, 0x8e, 0x44, 0x85, 0xfe, 0x6d, 0xe7, 0x6e, 0x34, 0x9f, 0x38,
	0x37, 0x3b, 0x5b, 0xea, 0x3c, 0xc8, 0x97, 0x10, 0x60, 0x65, 0x15, 0x76, 0xb4, 0xeb, 0x39, 0x23,
	0x76, 0x4e, 0x8d, 0x2d, 0x4a, 0xa0, 0xfb, 0xb6, 0x4e, 0x12, 0xf2, 0x3d, 0x0c, 0xb6, 0x75, 0x92,
	0xb8, 0x0f, 0xe7, 0x7a, 0x1e, 0x4e, 0x9b, 0xef, 0x7d, 0xba, 0xc6, 0x51, 0xbc, 0x4b, 0x12, 0x1e,
	0xe3, 0x67, 0x44, 0x9d, 0x27, 0xd2, 0x4f, 0x99, 0xaa, 0xec, 0x51, 0x69, 0xf9, 0x3c, 0x30, 0x73,
	0x50, 0x76, 0x60, 0xff, 0xf8, 0x30, 0x7a, 0x64, 0x19, 0x6f, 0x4e, 0xe0, 0x62, 0xdd, 0xde, 0xcb,
	0x75, 0xbb, 0x1c, 0x7e, 0x6b, 0xe8, 0x38, 0xd8, 0x2d, 0x53, 0xfc, 0x7d, 0x61, 0x53, 0x5b, 0x44,
	0x3e, 0x07, 0x40, 0xe9, 0xe2, 0x58, 0x5b, 0x1a, 0x7c, 0x17, 0x10, 0x5d, 0x1e, 0x6c, 0x5b, 0x45,
	0x22, 0xb8, 0x42, 0xb8, 0x6e, 0x0e, 0xde, 0x5c, 0xee, 0x85, 0x8e, 0x7c, 0x05, 0x63, 0x83, 0x9b,
	0x3c, 0x7d, 0xed, 0x74, 0xa9, 0x44, 0x5e, 0x46, 0x81, 0x1f, 0xc0, 0xc0, 0x7c, 0x1a, 0x4e, 0x81,
	0xbc, 0xe2, 0xba, 0x7c, 0xff, 0x4b, 0x38, 0x34, 0xb3, 0xd1, 0x00, 0x17, 0x85, 0x13, 0x55, 0x21,
	0xbc, 0x58, 0x14, 0x6e, 0x86, 0x1a, 0x5b, 0xeb, 0x3e, 0x47, 0x17, 0xf7, 0x79, 0x7e, 0x6a, 0xae,
	0x2e, 0x5e, 0xc8, 0x14, 0x82, 0xc5, 0x81, 0x2b, 0x45, 0x66, 0x30, 0x8c, 0x51, 0x68, 0xad, 0xf6,
	0x63, 0x57, 0x61, 0x61, 0x2d, 0xf4, 0xec, 0x83, 0x6f, 0xc2, 0x49, 0x28, 0xb1, 0x4d, 0xcd, 0x1b,
	0x35, 0xa0, 0x0d, 0x44, 0x4b, 0x26, 0x24, 0xff, 0xe0, 0xae, 0xb8, 0x81, 0xaf, 0x7f, 0x84, 0x2b,
	0xdc, 0xee, 0xba, 0x62, 0x7b, 0xbe, 0x79, 0x2e, 0xc8, 0x35, 0xc0, 0xe3, 0x7a, 0xb3, 0x92, 0x27,
	0x96, 0x8a, 0xdd, 0xe4, 0x23, 0x32, 0x86, 0xe1, 0xe3, 0x7a, 0xb3, 0x38, 0xe4, 0xb9, 0xe2, 0x13,
	0xcf, 0xc2, 0x87, 0xb8, 0x12, 0xb9, 0x9c, 0xf8, 0xaf, 0xe7, 0x30, 0x68, 0x1a, 0x41, 0xd3, 0xa2,
	0x15, 0x38, 0x82, 0xfe, 0x62, 0xf3, 0x36, 0x65, 0xf2, 0x38, 0xf1, 0x08, 0x40, 0x6f, 0xb1, 0xf9,
	0x59, 0x48, 0x3e, 0xf1, 0xb7, 0x3d, 0xfd, 0xd7, 0x73, 0xff, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff,
	0xae, 0xd8, 0x5d, 0x18, 0xed, 0x06, 0x00, 0x00,
}
