// Code generated by protoc-gen-go. DO NOT EDIT.
// source: account/account.proto

package account

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

type LiteAccount struct {
	Id                   int64    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Level                int32    `protobuf:"varint,3,opt,name=level,proto3" json:"level,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LiteAccount) Reset()         { *m = LiteAccount{} }
func (m *LiteAccount) String() string { return proto.CompactTextString(m) }
func (*LiteAccount) ProtoMessage()    {}
func (*LiteAccount) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{0}
}

func (m *LiteAccount) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LiteAccount.Unmarshal(m, b)
}
func (m *LiteAccount) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LiteAccount.Marshal(b, m, deterministic)
}
func (m *LiteAccount) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LiteAccount.Merge(m, src)
}
func (m *LiteAccount) XXX_Size() int {
	return xxx_messageInfo_LiteAccount.Size(m)
}
func (m *LiteAccount) XXX_DiscardUnknown() {
	xxx_messageInfo_LiteAccount.DiscardUnknown(m)
}

var xxx_messageInfo_LiteAccount proto.InternalMessageInfo

func (m *LiteAccount) GetId() int64 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *LiteAccount) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *LiteAccount) GetLevel() int32 {
	if m != nil {
		return m.Level
	}
	return 0
}

type MC_AccountLogon struct {
	UserId               int64    `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_AccountLogon) Reset()         { *m = MC_AccountLogon{} }
func (m *MC_AccountLogon) String() string { return proto.CompactTextString(m) }
func (*MC_AccountLogon) ProtoMessage()    {}
func (*MC_AccountLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{1}
}

func (m *MC_AccountLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_AccountLogon.Unmarshal(m, b)
}
func (m *MC_AccountLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_AccountLogon.Marshal(b, m, deterministic)
}
func (m *MC_AccountLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_AccountLogon.Merge(m, src)
}
func (m *MC_AccountLogon) XXX_Size() int {
	return xxx_messageInfo_MC_AccountLogon.Size(m)
}
func (m *MC_AccountLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_AccountLogon.DiscardUnknown(m)
}

var xxx_messageInfo_MC_AccountLogon proto.InternalMessageInfo

func (m *MC_AccountLogon) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *MC_AccountLogon) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type MS_AccountLogon struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_AccountLogon) Reset()         { *m = MS_AccountLogon{} }
func (m *MS_AccountLogon) String() string { return proto.CompactTextString(m) }
func (*MS_AccountLogon) ProtoMessage()    {}
func (*MS_AccountLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{2}
}

func (m *MS_AccountLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_AccountLogon.Unmarshal(m, b)
}
func (m *MS_AccountLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_AccountLogon.Marshal(b, m, deterministic)
}
func (m *MS_AccountLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_AccountLogon.Merge(m, src)
}
func (m *MS_AccountLogon) XXX_Size() int {
	return xxx_messageInfo_MS_AccountLogon.Size(m)
}
func (m *MS_AccountLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_AccountLogon.DiscardUnknown(m)
}

var xxx_messageInfo_MS_AccountLogon proto.InternalMessageInfo

type MC_HeartBeat struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_HeartBeat) Reset()         { *m = MC_HeartBeat{} }
func (m *MC_HeartBeat) String() string { return proto.CompactTextString(m) }
func (*MC_HeartBeat) ProtoMessage()    {}
func (*MC_HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{3}
}

func (m *MC_HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_HeartBeat.Unmarshal(m, b)
}
func (m *MC_HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_HeartBeat.Marshal(b, m, deterministic)
}
func (m *MC_HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_HeartBeat.Merge(m, src)
}
func (m *MC_HeartBeat) XXX_Size() int {
	return xxx_messageInfo_MC_HeartBeat.Size(m)
}
func (m *MC_HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_MC_HeartBeat proto.InternalMessageInfo

type MS_HeartBeat struct {
	Timestamp            uint32   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MS_HeartBeat) Reset()         { *m = MS_HeartBeat{} }
func (m *MS_HeartBeat) String() string { return proto.CompactTextString(m) }
func (*MS_HeartBeat) ProtoMessage()    {}
func (*MS_HeartBeat) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{4}
}

func (m *MS_HeartBeat) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MS_HeartBeat.Unmarshal(m, b)
}
func (m *MS_HeartBeat) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MS_HeartBeat.Marshal(b, m, deterministic)
}
func (m *MS_HeartBeat) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MS_HeartBeat.Merge(m, src)
}
func (m *MS_HeartBeat) XXX_Size() int {
	return xxx_messageInfo_MS_HeartBeat.Size(m)
}
func (m *MS_HeartBeat) XXX_DiscardUnknown() {
	xxx_messageInfo_MS_HeartBeat.DiscardUnknown(m)
}

var xxx_messageInfo_MS_HeartBeat proto.InternalMessageInfo

func (m *MS_HeartBeat) GetTimestamp() uint32 {
	if m != nil {
		return m.Timestamp
	}
	return 0
}

type MC_AccountConnected struct {
	AccountId            int64    `protobuf:"varint,1,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_AccountConnected) Reset()         { *m = MC_AccountConnected{} }
func (m *MC_AccountConnected) String() string { return proto.CompactTextString(m) }
func (*MC_AccountConnected) ProtoMessage()    {}
func (*MC_AccountConnected) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{5}
}

func (m *MC_AccountConnected) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_AccountConnected.Unmarshal(m, b)
}
func (m *MC_AccountConnected) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_AccountConnected.Marshal(b, m, deterministic)
}
func (m *MC_AccountConnected) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_AccountConnected.Merge(m, src)
}
func (m *MC_AccountConnected) XXX_Size() int {
	return xxx_messageInfo_MC_AccountConnected.Size(m)
}
func (m *MC_AccountConnected) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_AccountConnected.DiscardUnknown(m)
}

var xxx_messageInfo_MC_AccountConnected proto.InternalMessageInfo

func (m *MC_AccountConnected) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *MC_AccountConnected) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type MC_AccountDisconnect struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *MC_AccountDisconnect) Reset()         { *m = MC_AccountDisconnect{} }
func (m *MC_AccountDisconnect) String() string { return proto.CompactTextString(m) }
func (*MC_AccountDisconnect) ProtoMessage()    {}
func (*MC_AccountDisconnect) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{6}
}

func (m *MC_AccountDisconnect) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MC_AccountDisconnect.Unmarshal(m, b)
}
func (m *MC_AccountDisconnect) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MC_AccountDisconnect.Marshal(b, m, deterministic)
}
func (m *MC_AccountDisconnect) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MC_AccountDisconnect.Merge(m, src)
}
func (m *MC_AccountDisconnect) XXX_Size() int {
	return xxx_messageInfo_MC_AccountDisconnect.Size(m)
}
func (m *MC_AccountDisconnect) XXX_DiscardUnknown() {
	xxx_messageInfo_MC_AccountDisconnect.DiscardUnknown(m)
}

var xxx_messageInfo_MC_AccountDisconnect proto.InternalMessageInfo

func init() {
	proto.RegisterType((*LiteAccount)(nil), "yokai_account.LiteAccount")
	proto.RegisterType((*MC_AccountLogon)(nil), "yokai_account.MC_AccountLogon")
	proto.RegisterType((*MS_AccountLogon)(nil), "yokai_account.MS_AccountLogon")
	proto.RegisterType((*MC_HeartBeat)(nil), "yokai_account.MC_HeartBeat")
	proto.RegisterType((*MS_HeartBeat)(nil), "yokai_account.MS_HeartBeat")
	proto.RegisterType((*MC_AccountConnected)(nil), "yokai_account.MC_AccountConnected")
	proto.RegisterType((*MC_AccountDisconnect)(nil), "yokai_account.MC_AccountDisconnect")
}

func init() { proto.RegisterFile("account/account.proto", fileDescriptor_d66906c5773c9d08) }

var fileDescriptor_d66906c5773c9d08 = []byte{
	// 276 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4e, 0xeb, 0x30,
	0x10, 0x85, 0x95, 0xf4, 0xb6, 0x57, 0x19, 0xfa, 0x23, 0x4c, 0x81, 0x2c, 0x40, 0x8a, 0xbc, 0xca,
	0x02, 0x9a, 0x05, 0x4f, 0x40, 0x83, 0x44, 0x23, 0x35, 0x9b, 0x74, 0xc7, 0x26, 0x72, 0x9d, 0x51,
	0xb1, 0x68, 0xec, 0xca, 0x71, 0x2a, 0xf1, 0xf6, 0xa8, 0x8e, 0x21, 0xa2, 0x62, 0x35, 0x3e, 0x67,
	0xa4, 0x6f, 0x8e, 0x67, 0xe0, 0x9a, 0x71, 0xae, 0x5a, 0x69, 0x12, 0x57, 0x17, 0x07, 0xad, 0x8c,
	0x22, 0x93, 0x4f, 0xf5, 0xc1, 0x44, 0xe9, 0x4c, 0xfa, 0x0a, 0x17, 0x6b, 0x61, 0xf0, 0xb9, 0x93,
	0x64, 0x0a, 0xbe, 0xa8, 0x42, 0x2f, 0xf2, 0xe2, 0x41, 0xe1, 0x8b, 0x8a, 0x10, 0xf8, 0x27, 0x59,
	0x8d, 0xa1, 0x1f, 0x79, 0x71, 0x50, 0xd8, 0x37, 0x99, 0xc3, 0x70, 0x8f, 0x47, 0xdc, 0x87, 0x83,
	0xc8, 0x8b, 0x87, 0x45, 0x27, 0x68, 0x06, 0xb3, 0x3c, 0x2d, 0x1d, 0x67, 0xad, 0x76, 0x4a, 0x92,
	0x5b, 0xf8, 0xdf, 0x36, 0xa8, 0xcb, 0x1f, 0xe2, 0xe8, 0x24, 0xb3, 0x8a, 0xdc, 0x03, 0xb8, 0xf9,
	0xa7, 0x9e, 0x6f, 0x7b, 0x81, 0x73, 0xb2, 0x8a, 0x5e, 0xc2, 0x2c, 0xdf, 0xfc, 0x42, 0xd1, 0x29,
	0x8c, 0xf3, 0xb4, 0x5c, 0x21, 0xd3, 0x66, 0x89, 0xcc, 0xd0, 0x07, 0x18, 0xe7, 0x9b, 0x5e, 0x93,
	0x3b, 0x08, 0x8c, 0xa8, 0xb1, 0x31, 0xac, 0x3e, 0xd8, 0x61, 0x93, 0xa2, 0x37, 0xe8, 0x0a, 0xae,
	0xfa, 0x6c, 0xa9, 0x92, 0x12, 0xb9, 0xc1, 0xf3, 0x18, 0xde, 0x59, 0x8c, 0xbf, 0xfe, 0x4e, 0x6f,
	0x60, 0xde, 0x93, 0x5e, 0x44, 0xc3, 0x3b, 0xd8, 0x32, 0x79, 0x7b, 0xdc, 0x09, 0xf3, 0xde, 0x6e,
	0x17, 0x5c, 0xd5, 0x89, 0x5d, 0xb1, 0x50, 0x5d, 0x2d, 0x1b, 0xd4, 0x47, 0xd4, 0x89, 0x5d, 0xff,
	0xf7, 0x31, 0xb6, 0x23, 0x2b, 0x9f, 0xbe, 0x02, 0x00, 0x00, 0xff, 0xff, 0xa8, 0x0b, 0x88, 0x73,
	0xa6, 0x01, 0x00, 0x00,
}
