// Code generated by protoc-gen-go. DO NOT EDIT.
// source: game/hero.proto

package game

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

////////////////////////////////////////////////
// Hero
type Hero struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	TypeId               int32    `protobuf:"varint,2,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	Exp                  int64    `protobuf:"varint,3,opt,name=exp,proto3" json:"exp,omitempty"`
	Level                int32    `protobuf:"varint,4,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Hero) Reset()         { *m = Hero{} }
func (m *Hero) String() string { return proto.CompactTextString(m) }
func (*Hero) ProtoMessage()    {}
func (*Hero) Descriptor() ([]byte, []int) {
	return fileDescriptor_1698784dfd4e5521, []int{0}
}

func (m *Hero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Hero.Unmarshal(m, b)
}
func (m *Hero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Hero.Marshal(b, m, deterministic)
}
func (m *Hero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Hero.Merge(m, src)
}
func (m *Hero) XXX_Size() int {
	return xxx_messageInfo_Hero.Size(m)
}
func (m *Hero) XXX_DiscardUnknown() {
	xxx_messageInfo_Hero.DiscardUnknown(m)
}

var xxx_messageInfo_Hero proto.InternalMessageInfo

func (m *Hero) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *Hero) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

func (m *Hero) GetExp() int64 {
	if m != nil {
		return m.Exp
	}
	return 0
}

func (m *Hero) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type MC_AddHero struct {
	TypeId               int32    `protobuf:"varint,1,opt,name=type_id,json=typeId,proto3" json:"type_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_AddHero) Reset()         { *m = MC_AddHero{} }
func (m *MC_AddHero) String() string { return proto.CompactTextString(m) }
func (*MC_AddHero) ProtoMessage()    {}
func (*MC_AddHero) Descriptor() ([]byte, []int) {
	return fileDescriptor_1698784dfd4e5521, []int{1}
}

func (m *MC_AddHero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_AddHero.Unmarshal(m, b)
}
func (m *MC_AddHero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_AddHero.Marshal(b, m, deterministic)
}
func (m *MC_AddHero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_AddHero.Merge(m, src)
}
func (m *MC_AddHero) XXX_Size() int {
	return xxx_messageInfo_MC_AddHero.Size(m)
}
func (m *MC_AddHero) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_AddHero.DiscardUnknown(m)
}

var xxx_messageInfo_MC_AddHero proto.InternalMessageInfo

func (m *MC_AddHero) GetTypeId() int32 {
	if m != nil {
		return m.TypeId
	}
	return 0
}

type MC_DelHero struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_DelHero) Reset()         { *m = MC_DelHero{} }
func (m *MC_DelHero) String() string { return proto.CompactTextString(m) }
func (*MC_DelHero) ProtoMessage()    {}
func (*MC_DelHero) Descriptor() ([]byte, []int) {
	return fileDescriptor_1698784dfd4e5521, []int{2}
}

func (m *MC_DelHero) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_DelHero.Unmarshal(m, b)
}
func (m *MC_DelHero) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_DelHero.Marshal(b, m, deterministic)
}
func (m *MC_DelHero) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_DelHero.Merge(m, src)
}
func (m *MC_DelHero) XXX_Size() int {
	return xxx_messageInfo_MC_DelHero.Size(m)
}
func (m *MC_DelHero) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_DelHero.DiscardUnknown(m)
}

var xxx_messageInfo_MC_DelHero proto.InternalMessageInfo

func (m *MC_DelHero) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

type MC_QueryHeros struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_QueryHeros) Reset()         { *m = MC_QueryHeros{} }
func (m *MC_QueryHeros) String() string { return proto.CompactTextString(m) }
func (*MC_QueryHeros) ProtoMessage()    {}
func (*MC_QueryHeros) Descriptor() ([]byte, []int) {
	return fileDescriptor_1698784dfd4e5521, []int{3}
}

func (m *MC_QueryHeros) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_QueryHeros.Unmarshal(m, b)
}
func (m *MC_QueryHeros) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_QueryHeros.Marshal(b, m, deterministic)
}
func (m *MC_QueryHeros) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_QueryHeros.Merge(m, src)
}
func (m *MC_QueryHeros) XXX_Size() int {
	return xxx_messageInfo_MC_QueryHeros.Size(m)
}
func (m *MC_QueryHeros) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_QueryHeros.DiscardUnknown(m)
}

var xxx_messageInfo_MC_QueryHeros proto.InternalMessageInfo

type MS_HeroList struct {
	Heros                []*Hero  `protobuf:"bytes,1,rep,name=heros,proto3" json:"heros,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_HeroList) Reset()         { *m = MS_HeroList{} }
func (m *MS_HeroList) String() string { return proto.CompactTextString(m) }
func (*MS_HeroList) ProtoMessage()    {}
func (*MS_HeroList) Descriptor() ([]byte, []int) {
	return fileDescriptor_1698784dfd4e5521, []int{4}
}

func (m *MS_HeroList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_HeroList.Unmarshal(m, b)
}
func (m *MS_HeroList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_HeroList.Marshal(b, m, deterministic)
}
func (m *MS_HeroList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_HeroList.Merge(m, src)
}
func (m *MS_HeroList) XXX_Size() int {
	return xxx_messageInfo_MS_HeroList.Size(m)
}
func (m *MS_HeroList) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_HeroList.DiscardUnknown(m)
}

var xxx_messageInfo_MS_HeroList proto.InternalMessageInfo

func (m *MS_HeroList) GetHeros() []*Hero {
	if m != nil {
		return m.Heros
	}
	return nil
}

type MS_HeroInfo struct {
	Info                 *Hero    `protobuf:"bytes,1,opt,name=info,proto3" json:"info,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_HeroInfo) Reset()         { *m = MS_HeroInfo{} }
func (m *MS_HeroInfo) String() string { return proto.CompactTextString(m) }
func (*MS_HeroInfo) ProtoMessage()    {}
func (*MS_HeroInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_1698784dfd4e5521, []int{5}
}

func (m *MS_HeroInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_HeroInfo.Unmarshal(m, b)
}
func (m *MS_HeroInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_HeroInfo.Marshal(b, m, deterministic)
}
func (m *MS_HeroInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_HeroInfo.Merge(m, src)
}
func (m *MS_HeroInfo) XXX_Size() int {
	return xxx_messageInfo_MS_HeroInfo.Size(m)
}
func (m *MS_HeroInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_HeroInfo.DiscardUnknown(m)
}

var xxx_messageInfo_MS_HeroInfo proto.InternalMessageInfo

func (m *MS_HeroInfo) GetInfo() *Hero {
	if m != nil {
		return m.Info
	}
	return nil
}

func init() {
	proto.RegisterType((*Hero)(nil), "yokai_game.Hero")
	proto.RegisterType((*MC_AddHero)(nil), "yokai_game.MC_AddHero")
	proto.RegisterType((*MC_DelHero)(nil), "yokai_game.MC_DelHero")
	proto.RegisterType((*MC_QueryHeros)(nil), "yokai_game.MC_QueryHeros")
	proto.RegisterType((*MS_HeroList)(nil), "yokai_game.MS_HeroList")
	proto.RegisterType((*MS_HeroInfo)(nil), "yokai_game.MS_HeroInfo")
}

func init() { proto.RegisterFile("game/hero.proto", fileDescriptor_1698784dfd4e5521) }

var fileDescriptor_1698784dfd4e5521 = []byte{
	// 262 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x90, 0x5f, 0x4b, 0xc3, 0x30,
	0x14, 0xc5, 0x49, 0xff, 0x4c, 0xb8, 0x43, 0x37, 0x82, 0x60, 0x1e, 0x7c, 0x28, 0x41, 0xa5, 0x88,
	0xb4, 0xe0, 0xf0, 0x03, 0xe8, 0x7c, 0x70, 0x60, 0x1f, 0xac, 0x0f, 0x82, 0x2f, 0x61, 0x33, 0x77,
	0x5b, 0xb0, 0x5b, 0x4a, 0xda, 0x0d, 0xfb, 0xed, 0x25, 0xb7, 0xa0, 0x7b, 0xd8, 0x53, 0xcf, 0xed,
	0xf9, 0x9d, 0x93, 0xdc, 0xc0, 0x68, 0x35, 0xdf, 0x60, 0xbe, 0x46, 0x67, 0xb3, 0xda, 0xd9, 0xd6,
	0x72, 0xe8, 0xec, 0xf7, 0xdc, 0x28, 0xff, 0x5b, 0x7e, 0x40, 0xf4, 0x82, 0xce, 0xf2, 0x33, 0x08,
	0x8c, 0x16, 0x2c, 0x61, 0x69, 0x58, 0x06, 0x46, 0xf3, 0x0b, 0x38, 0x69, 0xbb, 0x1a, 0x95, 0xd1,
	0x22, 0x48, 0x58, 0x1a, 0x97, 0x03, 0x3f, 0xce, 0x34, 0x1f, 0x43, 0x88, 0x3f, 0xb5, 0x08, 0x89,
	0xf4, 0x92, 0x9f, 0x43, 0x5c, 0xe1, 0x1e, 0x2b, 0x11, 0x11, 0xd8, 0x0f, 0xf2, 0x1a, 0xa0, 0x98,
	0xaa, 0x47, 0xad, 0xa9, 0xfe, 0xa0, 0x8e, 0x1d, 0xd6, 0xc9, 0x4b, 0xc2, 0x9e, 0xb1, 0x3a, 0x76,
	0x0b, 0x39, 0x82, 0xd3, 0x62, 0xaa, 0xde, 0x76, 0xe8, 0x3a, 0xef, 0x37, 0xf2, 0x01, 0x86, 0xc5,
	0xbb, 0xf2, 0xfa, 0xd5, 0x34, 0x2d, 0xbf, 0x81, 0xd8, 0xef, 0xd5, 0x08, 0x96, 0x84, 0xe9, 0xf0,
	0x7e, 0x9c, 0xfd, 0x6f, 0x96, 0x79, 0xa8, 0xec, 0x6d, 0x39, 0xf9, 0x8b, 0xcd, 0xb6, 0x4b, 0xcb,
	0xaf, 0x20, 0x32, 0xdb, 0xa5, 0xa5, 0x83, 0x8e, 0xa5, 0xc8, 0x7d, 0xba, 0xfb, 0xbc, 0x5d, 0x99,
	0x76, 0xbd, 0x5b, 0x64, 0x5f, 0x76, 0x93, 0x13, 0x63, 0x6c, 0xff, 0x55, 0x0d, 0xba, 0x3d, 0xba,
	0x9c, 0xde, 0x33, 0xf7, 0xc1, 0xc5, 0x80, 0xf4, 0xe4, 0x37, 0x00, 0x00, 0xff, 0xff, 0x38, 0xa8,
	0x8c, 0x13, 0x6e, 0x01, 0x00, 0x00,
}