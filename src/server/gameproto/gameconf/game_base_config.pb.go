// Code generated by protoc-gen-go. DO NOT EDIT.
// source: gameconf/game_base_config.proto

/*
Package gameconf is a generated protocol buffer package.

It is generated from these files:
	gameconf/game_base_config.proto
	gameconf/game_type.proto

It has these top-level messages:
	GameBaseConfig
	SkillEffect
	SkillConfDefine
	SkillEffectConfDefine
	Learn
	GeneralConfDefine
	DoNotUseThis
	DoNotUseThisDefine
*/
package gameconf

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

// Defined in table: GameBaseConfig
type GameBaseConfig struct {
	DoNotUseThis    []*DoNotUseThisDefine    `protobuf:"bytes,1,rep,name=DoNotUseThis" json:"DoNotUseThis,omitempty"`
	SkillConf       []*SkillConfDefine       `protobuf:"bytes,2,rep,name=SkillConf" json:"SkillConf,omitempty"`
	SkillEffectConf []*SkillEffectConfDefine `protobuf:"bytes,3,rep,name=SkillEffectConf" json:"SkillEffectConf,omitempty"`
	GeneralConf     []*GeneralConfDefine     `protobuf:"bytes,4,rep,name=GeneralConf" json:"GeneralConf,omitempty"`
}

func (m *GameBaseConfig) Reset()                    { *m = GameBaseConfig{} }
func (m *GameBaseConfig) String() string            { return proto.CompactTextString(m) }
func (*GameBaseConfig) ProtoMessage()               {}
func (*GameBaseConfig) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GameBaseConfig) GetDoNotUseThis() []*DoNotUseThisDefine {
	if m != nil {
		return m.DoNotUseThis
	}
	return nil
}

func (m *GameBaseConfig) GetSkillConf() []*SkillConfDefine {
	if m != nil {
		return m.SkillConf
	}
	return nil
}

func (m *GameBaseConfig) GetSkillEffectConf() []*SkillEffectConfDefine {
	if m != nil {
		return m.SkillEffectConf
	}
	return nil
}

func (m *GameBaseConfig) GetGeneralConf() []*GeneralConfDefine {
	if m != nil {
		return m.GeneralConf
	}
	return nil
}

// Defined in table: SkillConf
type SkillEffect struct {
	// 效果对象
	Object SkillEffectObjectTyp `protobuf:"varint,1,opt,name=Object,enum=gameconf.SkillEffectObjectTyp" json:"Object,omitempty"`
	// 最短持续时间
	Min int32 `protobuf:"varint,2,opt,name=Min" json:"Min,omitempty"`
	// 最长持续时间
	Max int32 `protobuf:"varint,3,opt,name=Max" json:"Max,omitempty"`
	// 触发概率
	Chance int32 `protobuf:"varint,4,opt,name=Chance" json:"Chance,omitempty"`
	// 等级
	Power int32 `protobuf:"varint,5,opt,name=Power" json:"Power,omitempty"`
	// 效果ID
	Id uint32 `protobuf:"varint,6,opt,name=Id" json:"Id,omitempty"`
}

func (m *SkillEffect) Reset()                    { *m = SkillEffect{} }
func (m *SkillEffect) String() string            { return proto.CompactTextString(m) }
func (*SkillEffect) ProtoMessage()               {}
func (*SkillEffect) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *SkillEffect) GetObject() SkillEffectObjectTyp {
	if m != nil {
		return m.Object
	}
	return SkillEffectObjectTyp_SEOTInvalid
}

func (m *SkillEffect) GetMin() int32 {
	if m != nil {
		return m.Min
	}
	return 0
}

func (m *SkillEffect) GetMax() int32 {
	if m != nil {
		return m.Max
	}
	return 0
}

func (m *SkillEffect) GetChance() int32 {
	if m != nil {
		return m.Chance
	}
	return 0
}

func (m *SkillEffect) GetPower() int32 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *SkillEffect) GetId() uint32 {
	if m != nil {
		return m.Id
	}
	return 0
}

// Defined in table: SkillConf
type SkillConfDefine struct {
	SkillID         uint32         `protobuf:"varint,1,opt,name=SkillID" json:"SkillID,omitempty"`
	SkillName       string         `protobuf:"bytes,2,opt,name=SkillName" json:"SkillName,omitempty"`
	SkillType       SkillTyp       `protobuf:"varint,3,opt,name=SkillType,enum=gameconf.SkillTyp" json:"SkillType,omitempty"`
	SkillAttackType SkillAttackTyp `protobuf:"varint,4,opt,name=SkillAttackType,enum=gameconf.SkillAttackTyp" json:"SkillAttackType,omitempty"`
	Power           uint32         `protobuf:"varint,5,opt,name=Power" json:"Power,omitempty"`
	Effect          []*SkillEffect `protobuf:"bytes,6,rep,name=Effect" json:"Effect,omitempty"`
}

func (m *SkillConfDefine) Reset()                    { *m = SkillConfDefine{} }
func (m *SkillConfDefine) String() string            { return proto.CompactTextString(m) }
func (*SkillConfDefine) ProtoMessage()               {}
func (*SkillConfDefine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *SkillConfDefine) GetSkillID() uint32 {
	if m != nil {
		return m.SkillID
	}
	return 0
}

func (m *SkillConfDefine) GetSkillName() string {
	if m != nil {
		return m.SkillName
	}
	return ""
}

func (m *SkillConfDefine) GetSkillType() SkillTyp {
	if m != nil {
		return m.SkillType
	}
	return SkillTyp_STInvalid
}

func (m *SkillConfDefine) GetSkillAttackType() SkillAttackTyp {
	if m != nil {
		return m.SkillAttackType
	}
	return SkillAttackTyp_SATInvliad
}

func (m *SkillConfDefine) GetPower() uint32 {
	if m != nil {
		return m.Power
	}
	return 0
}

func (m *SkillConfDefine) GetEffect() []*SkillEffect {
	if m != nil {
		return m.Effect
	}
	return nil
}

// Defined in table: SkillEffectConf
type SkillEffectConfDefine struct {
	SkillEffectID   uint32         `protobuf:"varint,1,opt,name=SkillEffectID" json:"SkillEffectID,omitempty"`
	SkillName       string         `protobuf:"bytes,2,opt,name=SkillName" json:"SkillName,omitempty"`
	SkillEffectType SkillEffectTyp `protobuf:"varint,3,opt,name=SkillEffectType,enum=gameconf.SkillEffectTyp" json:"SkillEffectType,omitempty"`
	LevelUp         int32          `protobuf:"varint,4,opt,name=LevelUp" json:"LevelUp,omitempty"`
}

func (m *SkillEffectConfDefine) Reset()                    { *m = SkillEffectConfDefine{} }
func (m *SkillEffectConfDefine) String() string            { return proto.CompactTextString(m) }
func (*SkillEffectConfDefine) ProtoMessage()               {}
func (*SkillEffectConfDefine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *SkillEffectConfDefine) GetSkillEffectID() uint32 {
	if m != nil {
		return m.SkillEffectID
	}
	return 0
}

func (m *SkillEffectConfDefine) GetSkillName() string {
	if m != nil {
		return m.SkillName
	}
	return ""
}

func (m *SkillEffectConfDefine) GetSkillEffectType() SkillEffectTyp {
	if m != nil {
		return m.SkillEffectType
	}
	return SkillEffectTyp_SETInvliad
}

func (m *SkillEffectConfDefine) GetLevelUp() int32 {
	if m != nil {
		return m.LevelUp
	}
	return 0
}

// Defined in table: GeneralConf
type Learn struct {
	// 等级
	Level uint32 `protobuf:"varint,1,opt,name=Level" json:"Level,omitempty"`
	// 技能
	SkillID uint32 `protobuf:"varint,2,opt,name=SkillID" json:"SkillID,omitempty"`
}

func (m *Learn) Reset()                    { *m = Learn{} }
func (m *Learn) String() string            { return proto.CompactTextString(m) }
func (*Learn) ProtoMessage()               {}
func (*Learn) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *Learn) GetLevel() uint32 {
	if m != nil {
		return m.Level
	}
	return 0
}

func (m *Learn) GetSkillID() uint32 {
	if m != nil {
		return m.SkillID
	}
	return 0
}

// Defined in table: GeneralConf
type GeneralConfDefine struct {
	GeneralID   uint32   `protobuf:"varint,1,opt,name=GeneralID" json:"GeneralID,omitempty"`
	GeneralName string   `protobuf:"bytes,2,opt,name=GeneralName" json:"GeneralName,omitempty"`
	Hp          int32    `protobuf:"varint,3,opt,name=Hp" json:"Hp,omitempty"`
	Atk         int32    `protobuf:"varint,4,opt,name=Atk" json:"Atk,omitempty"`
	Def         int32    `protobuf:"varint,5,opt,name=Def" json:"Def,omitempty"`
	Satk        int32    `protobuf:"varint,6,opt,name=Satk" json:"Satk,omitempty"`
	Sdef        int32    `protobuf:"varint,7,opt,name=Sdef" json:"Sdef,omitempty"`
	Spd         int32    `protobuf:"varint,8,opt,name=Spd" json:"Spd,omitempty"`
	LearnSkills []*Learn `protobuf:"bytes,9,rep,name=LearnSkills" json:"LearnSkills,omitempty"`
	BaseSkills  []uint32 `protobuf:"varint,10,rep,packed,name=BaseSkills" json:"BaseSkills,omitempty"`
}

func (m *GeneralConfDefine) Reset()                    { *m = GeneralConfDefine{} }
func (m *GeneralConfDefine) String() string            { return proto.CompactTextString(m) }
func (*GeneralConfDefine) ProtoMessage()               {}
func (*GeneralConfDefine) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GeneralConfDefine) GetGeneralID() uint32 {
	if m != nil {
		return m.GeneralID
	}
	return 0
}

func (m *GeneralConfDefine) GetGeneralName() string {
	if m != nil {
		return m.GeneralName
	}
	return ""
}

func (m *GeneralConfDefine) GetHp() int32 {
	if m != nil {
		return m.Hp
	}
	return 0
}

func (m *GeneralConfDefine) GetAtk() int32 {
	if m != nil {
		return m.Atk
	}
	return 0
}

func (m *GeneralConfDefine) GetDef() int32 {
	if m != nil {
		return m.Def
	}
	return 0
}

func (m *GeneralConfDefine) GetSatk() int32 {
	if m != nil {
		return m.Satk
	}
	return 0
}

func (m *GeneralConfDefine) GetSdef() int32 {
	if m != nil {
		return m.Sdef
	}
	return 0
}

func (m *GeneralConfDefine) GetSpd() int32 {
	if m != nil {
		return m.Spd
	}
	return 0
}

func (m *GeneralConfDefine) GetLearnSkills() []*Learn {
	if m != nil {
		return m.LearnSkills
	}
	return nil
}

func (m *GeneralConfDefine) GetBaseSkills() []uint32 {
	if m != nil {
		return m.BaseSkills
	}
	return nil
}

func init() {
	proto.RegisterType((*GameBaseConfig)(nil), "gameconf.GameBaseConfig")
	proto.RegisterType((*SkillEffect)(nil), "gameconf.SkillEffect")
	proto.RegisterType((*SkillConfDefine)(nil), "gameconf.SkillConfDefine")
	proto.RegisterType((*SkillEffectConfDefine)(nil), "gameconf.SkillEffectConfDefine")
	proto.RegisterType((*Learn)(nil), "gameconf.Learn")
	proto.RegisterType((*GeneralConfDefine)(nil), "gameconf.GeneralConfDefine")
}

func init() { proto.RegisterFile("gameconf/game_base_config.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 566 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0x95, 0x9d, 0xc4, 0x6d, 0x27, 0x24, 0x85, 0x15, 0x45, 0x0b, 0x54, 0xad, 0x65, 0x71, 0xc8,
	0x85, 0x00, 0x41, 0xa2, 0x27, 0x24, 0x9a, 0x06, 0xb5, 0x91, 0x4a, 0x41, 0x9b, 0xf4, 0x5c, 0x6d,
	0xe2, 0x71, 0x1b, 0x92, 0xd8, 0x56, 0x62, 0x41, 0x73, 0xe7, 0x13, 0xe0, 0x1b, 0xf8, 0x06, 0xfe,
	0x0e, 0xed, 0x78, 0x93, 0x5d, 0x9b, 0x48, 0xdc, 0x76, 0xde, 0xbc, 0x37, 0x99, 0x37, 0x33, 0x31,
	0x1c, 0xdf, 0xca, 0x39, 0x8e, 0x93, 0x38, 0x7a, 0xa5, 0x1e, 0x37, 0x23, 0xb9, 0xc4, 0x1b, 0x15,
	0x4e, 0x6e, 0xdb, 0xe9, 0x22, 0xc9, 0x12, 0xb6, 0xbb, 0x26, 0x3c, 0xe3, 0x45, 0x6a, 0xb6, 0x4a,
	0x31, 0xe7, 0x04, 0xbf, 0x5c, 0x68, 0x9e, 0xcb, 0x39, 0x76, 0xe5, 0x12, 0xcf, 0x48, 0xcc, 0x3e,
	0xc0, 0x83, 0x5e, 0x72, 0x95, 0x64, 0xd7, 0x4b, 0x1c, 0xde, 0x4d, 0x96, 0xdc, 0xf1, 0x2b, 0xad,
	0x7a, 0xe7, 0xb0, 0xbd, 0xae, 0xd1, 0xb6, 0xb3, 0x3d, 0x8c, 0x26, 0x31, 0x8a, 0x82, 0x82, 0x9d,
	0xc0, 0xde, 0x60, 0x3a, 0x99, 0xcd, 0x54, 0x41, 0xee, 0x92, 0xfc, 0xa9, 0x91, 0x6f, 0x52, 0x5a,
	0x6b, 0xb8, 0xac, 0x0f, 0xfb, 0x14, 0x7c, 0x8c, 0x22, 0x1c, 0x67, 0x24, 0xaf, 0x90, 0xfc, 0xb8,
	0x24, 0x37, 0x04, 0x5d, 0xa4, 0xac, 0x63, 0xef, 0xa1, 0x7e, 0x8e, 0x31, 0x2e, 0x64, 0xde, 0x45,
	0x95, 0xca, 0x3c, 0x37, 0x65, 0xac, 0xa4, 0x2e, 0x61, 0xf3, 0x83, 0xdf, 0x0e, 0xd4, 0xad, 0x92,
	0xec, 0x1d, 0x78, 0x9f, 0x47, 0x5f, 0x71, 0x9c, 0x71, 0xc7, 0x77, 0x5a, 0xcd, 0xce, 0xd1, 0xd6,
	0x86, 0x72, 0xca, 0x70, 0x95, 0x0a, 0xcd, 0x66, 0x0f, 0xa1, 0xf2, 0x69, 0x12, 0x73, 0xd7, 0x77,
	0x5a, 0x35, 0xa1, 0x9e, 0x84, 0xc8, 0x7b, 0x5e, 0xd1, 0x88, 0xbc, 0x67, 0x4f, 0xc0, 0x3b, 0xbb,
	0x93, 0xf1, 0x18, 0x79, 0x95, 0x40, 0x1d, 0xb1, 0xc7, 0x50, 0xfb, 0x92, 0x7c, 0xc7, 0x05, 0xaf,
	0x11, 0x9c, 0x07, 0xac, 0x09, 0x6e, 0x3f, 0xe4, 0x9e, 0xef, 0xb4, 0x1a, 0xc2, 0xed, 0x87, 0xc1,
	0x0f, 0x57, 0x0f, 0xcd, 0x58, 0x61, 0x1c, 0x76, 0x08, 0xea, 0xf7, 0xa8, 0xdd, 0x86, 0x58, 0x87,
	0xec, 0x50, 0xaf, 0xe6, 0x4a, 0xce, 0x91, 0xba, 0xda, 0x13, 0x06, 0x60, 0xaf, 0x75, 0x76, 0xb8,
	0x4a, 0x91, 0x3a, 0x6c, 0x76, 0x58, 0xc9, 0xa8, 0x32, 0x67, 0x48, 0xac, 0xab, 0x7f, 0xfc, 0x34,
	0xcb, 0xe4, 0x78, 0x4a, 0xba, 0x2a, 0xe9, 0x78, 0x49, 0xb7, 0x21, 0x88, 0xb2, 0xa0, 0xe8, 0xb3,
	0xb1, 0xf6, 0xf9, 0x12, 0xbc, 0x7c, 0xa8, 0xdc, 0xa3, 0xdd, 0x1d, 0x6c, 0x9d, 0xb8, 0xd0, 0xa4,
	0xe0, 0x8f, 0x03, 0x07, 0x5b, 0x4f, 0x83, 0xbd, 0x80, 0x86, 0x95, 0xd8, 0x8c, 0xa4, 0x08, 0xfe,
	0x67, 0x30, 0xdd, 0xc2, 0x61, 0x5a, 0xe3, 0xe1, 0x5b, 0xbb, 0x32, 0x36, 0x8d, 0x40, 0x2d, 0xe5,
	0x12, 0xbf, 0xe1, 0xec, 0x3a, 0xd5, 0x7b, 0x5e, 0x87, 0xc1, 0x09, 0xd4, 0x2e, 0x51, 0x2e, 0x62,
	0x35, 0x09, 0xc2, 0x74, 0x8b, 0x79, 0x60, 0x6f, 0xd3, 0x2d, 0x6c, 0x33, 0xf8, 0xe9, 0xc2, 0xa3,
	0x7f, 0x0e, 0x59, 0x59, 0xd1, 0xe0, 0xc6, 0xac, 0x01, 0x98, 0xbf, 0xf9, 0x63, 0x58, 0x56, 0x6d,
	0x48, 0x5d, 0xd8, 0x45, 0xaa, 0x0f, 0xd4, 0xbd, 0x48, 0xd5, 0xc5, 0x9e, 0x66, 0x53, 0xdd, 0xb4,
	0x7a, 0x2a, 0xa4, 0x87, 0x91, 0xbe, 0x4b, 0xf5, 0x64, 0x0c, 0xaa, 0x03, 0x99, 0x4d, 0xe9, 0x2e,
	0x6b, 0x82, 0xde, 0x84, 0x85, 0x18, 0xf1, 0x1d, 0x8d, 0x85, 0x18, 0x29, 0xe5, 0x20, 0x0d, 0xf9,
	0x6e, 0xae, 0x1c, 0xa4, 0x21, 0x7b, 0x03, 0x75, 0x32, 0x4f, 0x9e, 0x96, 0x7c, 0x8f, 0x96, 0xbd,
	0x6f, 0xc6, 0x4a, 0x49, 0x61, 0x73, 0xd8, 0x11, 0x80, 0xfa, 0x5e, 0x69, 0x05, 0xf8, 0x95, 0x56,
	0x43, 0x58, 0xc8, 0xc8, 0xa3, 0x6f, 0xdb, 0xdb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x82, 0x15,
	0x9a, 0xd9, 0x22, 0x05, 0x00, 0x00,
}
