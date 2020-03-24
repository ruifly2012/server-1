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

type C2M_AccountLogon struct {
	RpcId                int32    `protobuf:"varint,90,opt,name=RpcId,proto3" json:"RpcId,omitempty"`
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *C2M_AccountLogon) Reset()         { *m = C2M_AccountLogon{} }
func (m *C2M_AccountLogon) String() string { return proto.CompactTextString(m) }
func (*C2M_AccountLogon) ProtoMessage()    {}
func (*C2M_AccountLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{1}
}

func (m *C2M_AccountLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_C2M_AccountLogon.Unmarshal(m, b)
}
func (m *C2M_AccountLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_C2M_AccountLogon.Marshal(b, m, deterministic)
}
func (m *C2M_AccountLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_C2M_AccountLogon.Merge(m, src)
}
func (m *C2M_AccountLogon) XXX_Size() int {
	return xxx_messageInfo_C2M_AccountLogon.Size(m)
}
func (m *C2M_AccountLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_C2M_AccountLogon.DiscardUnknown(m)
}

var xxx_messageInfo_C2M_AccountLogon proto.InternalMessageInfo

func (m *C2M_AccountLogon) GetRpcId() int32 {
	if m != nil {
		return m.RpcId
	}
	return 0
}

func (m *C2M_AccountLogon) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *C2M_AccountLogon) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

type M2C_AccountLogon struct {
	RpcId                int32    `protobuf:"varint,90,opt,name=RpcId,proto3" json:"RpcId,omitempty"`
	Error                int32    `protobuf:"varint,91,opt,name=Error,proto3" json:"Error,omitempty"`
	Message              string   `protobuf:"bytes,92,opt,name=Message,proto3" json:"Message,omitempty"`
	UserId               int64    `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	AccountId            int64    `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	PlayerId             int64    `protobuf:"varint,3,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	PlayerName           string   `protobuf:"bytes,4,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	PlayerLevel          int32    `protobuf:"varint,5,opt,name=PlayerLevel,proto3" json:"PlayerLevel,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *M2C_AccountLogon) Reset()         { *m = M2C_AccountLogon{} }
func (m *M2C_AccountLogon) String() string { return proto.CompactTextString(m) }
func (*M2C_AccountLogon) ProtoMessage()    {}
func (*M2C_AccountLogon) Descriptor() ([]byte, []int) {
	return fileDescriptor_d66906c5773c9d08, []int{2}
}

func (m *M2C_AccountLogon) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_M2C_AccountLogon.Unmarshal(m, b)
}
func (m *M2C_AccountLogon) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_M2C_AccountLogon.Marshal(b, m, deterministic)
}
func (m *M2C_AccountLogon) XXX_Merge(src proto.Message) {
	xxx_messageInfo_M2C_AccountLogon.Merge(m, src)
}
func (m *M2C_AccountLogon) XXX_Size() int {
	return xxx_messageInfo_M2C_AccountLogon.Size(m)
}
func (m *M2C_AccountLogon) XXX_DiscardUnknown() {
	xxx_messageInfo_M2C_AccountLogon.DiscardUnknown(m)
}

var xxx_messageInfo_M2C_AccountLogon proto.InternalMessageInfo

func (m *M2C_AccountLogon) GetRpcId() int32 {
	if m != nil {
		return m.RpcId
	}
	return 0
}

func (m *M2C_AccountLogon) GetError() int32 {
	if m != nil {
		return m.Error
	}
	return 0
}

func (m *M2C_AccountLogon) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *M2C_AccountLogon) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func (m *M2C_AccountLogon) GetAccountId() int64 {
	if m != nil {
		return m.AccountId
	}
	return 0
}

func (m *M2C_AccountLogon) GetPlayerId() int64 {
	if m != nil {
		return m.PlayerId
	}
	return 0
}

func (m *M2C_AccountLogon) GetPlayerName() string {
	if m != nil {
		return m.PlayerName
	}
	return ""
}

func (m *M2C_AccountLogon) GetPlayerLevel() int32 {
	if m != nil {
		return m.PlayerLevel
	}
	return 0
}

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
	proto.RegisterType((*C2M_AccountLogon)(nil), "yokai_account.C2M_AccountLogon")
	proto.RegisterType((*M2C_AccountLogon)(nil), "yokai_account.M2C_AccountLogon")
	proto.RegisterType((*MC_HeartBeat)(nil), "yokai_account.MC_HeartBeat")
	proto.RegisterType((*MS_HeartBeat)(nil), "yokai_account.MS_HeartBeat")
	proto.RegisterType((*MC_AccountConnected)(nil), "yokai_account.MC_AccountConnected")
	proto.RegisterType((*MC_AccountDisconnect)(nil), "yokai_account.MC_AccountDisconnect")
}

func init() { proto.RegisterFile("account/account.proto", fileDescriptor_d66906c5773c9d08) }

var fileDescriptor_d66906c5773c9d08 = []byte{
	// 363 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xcf, 0x6e, 0xe2, 0x30,
	0x10, 0xc6, 0x15, 0x42, 0xd8, 0x65, 0xf8, 0x23, 0xe4, 0x65, 0x51, 0xb4, 0xda, 0x56, 0x51, 0x4e,
	0x1c, 0x5a, 0x22, 0xd1, 0x27, 0x28, 0x69, 0x55, 0x90, 0x48, 0x55, 0xa5, 0xea, 0x85, 0x56, 0x8d,
	0x4c, 0x62, 0x51, 0xab, 0x24, 0x46, 0x8e, 0x41, 0xe2, 0x9d, 0xfb, 0x10, 0x15, 0xe3, 0x90, 0x70,
	0xe8, 0xa1, 0xea, 0xc9, 0xfe, 0x7d, 0xb6, 0x67, 0x3e, 0xcf, 0x0c, 0xfc, 0xa5, 0x71, 0x2c, 0xb6,
	0x99, 0xf2, 0x8a, 0x75, 0xb4, 0x91, 0x42, 0x09, 0xd2, 0xd9, 0x8b, 0x77, 0xca, 0xa3, 0x42, 0x74,
	0xef, 0xa0, 0x35, 0xe7, 0x8a, 0x5d, 0x6b, 0x24, 0x5d, 0xa8, 0xf1, 0xc4, 0x36, 0x1c, 0x63, 0x68,
	0x86, 0x35, 0x9e, 0x10, 0x02, 0xf5, 0x8c, 0xa6, 0xcc, 0xae, 0x39, 0xc6, 0xb0, 0x19, 0xe2, 0x9e,
	0xf4, 0xc1, 0x5a, 0xb3, 0x1d, 0x5b, 0xdb, 0xa6, 0x63, 0x0c, 0xad, 0x50, 0x83, 0xfb, 0x0a, 0x3d,
	0x7f, 0x1c, 0x44, 0x45, 0xa0, 0xb9, 0x58, 0x89, 0xec, 0x70, 0x33, 0xdc, 0xc4, 0xb3, 0xc4, 0x5e,
	0xe8, 0x9b, 0x08, 0x64, 0x00, 0x8d, 0xa7, 0x9c, 0xc9, 0xd9, 0x31, 0x4f, 0x41, 0xe4, 0x3f, 0x34,
	0x8b, 0xd7, 0xb3, 0x04, 0x13, 0x9a, 0x61, 0x25, 0xb8, 0x1f, 0x06, 0xf4, 0x82, 0xb1, 0xff, 0x9d,
	0x04, 0x7d, 0xb0, 0x6e, 0xa5, 0x14, 0xd2, 0x7e, 0xd6, 0x2a, 0x02, 0xb1, 0xe1, 0x57, 0xc0, 0xf2,
	0x9c, 0xae, 0x98, 0xfd, 0x82, 0xbf, 0x39, 0xe2, 0xcf, 0x0c, 0x91, 0x7f, 0xf0, 0xfb, 0x61, 0x4d,
	0xf7, 0xf8, 0xce, 0xc4, 0xc3, 0x92, 0xc9, 0x39, 0x80, 0xde, 0xdf, 0x1f, 0x8a, 0x57, 0xc7, 0x74,
	0x27, 0x0a, 0x71, 0xa0, 0xa5, 0x69, 0x8e, 0x85, 0xb4, 0xd0, 0xe7, 0xa9, 0xe4, 0x76, 0xa1, 0x1d,
	0xf8, 0xd1, 0x94, 0x51, 0xa9, 0x26, 0x8c, 0x2a, 0xf7, 0x02, 0xda, 0xc1, 0x63, 0xc5, 0x07, 0x6f,
	0x8a, 0xa7, 0x2c, 0x57, 0x34, 0xdd, 0xa0, 0xed, 0x4e, 0x58, 0x09, 0xee, 0x14, 0xfe, 0x04, 0x65,
	0xa9, 0x7c, 0x91, 0x65, 0x2c, 0x56, 0x2c, 0x21, 0x67, 0x00, 0x45, 0xdf, 0xa3, 0xb2, 0xcb, 0x4d,
	0x5a, 0xfe, 0xe8, 0x8b, 0x66, 0xbb, 0x03, 0xe8, 0x57, 0x91, 0x6e, 0x78, 0x1e, 0xeb, 0x60, 0x13,
	0x6f, 0x71, 0xb9, 0xe2, 0xea, 0x6d, 0xbb, 0x1c, 0xc5, 0x22, 0xf5, 0x70, 0xa6, 0xb8, 0xd0, 0x6b,
	0x94, 0x33, 0xb9, 0x63, 0xd2, 0xc3, 0x79, 0x3b, 0x4e, 0xdf, 0xb2, 0x81, 0x78, 0xf5, 0x19, 0x00,
	0x00, 0xff, 0xff, 0x6f, 0xd5, 0x41, 0x10, 0x97, 0x02, 0x00, 0x00,
}
