// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: logon.proto

package global

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// 客户端账号登陆
type C2S_AccountLogon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      string `protobuf:"bytes,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	AccountId   int64  `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	AccountName string `protobuf:"bytes,3,opt,name=AccountName,proto3" json:"AccountName,omitempty"`
}

func (x *C2S_AccountLogon) Reset() {
	*x = C2S_AccountLogon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_AccountLogon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_AccountLogon) ProtoMessage() {}

func (x *C2S_AccountLogon) ProtoReflect() protoreflect.Message {
	mi := &file_logon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_AccountLogon.ProtoReflect.Descriptor instead.
func (*C2S_AccountLogon) Descriptor() ([]byte, []int) {
	return file_logon_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_AccountLogon) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *C2S_AccountLogon) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *C2S_AccountLogon) GetAccountName() string {
	if x != nil {
		return x.AccountName
	}
	return ""
}

type S2C_AccountLogon struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId      int64  `protobuf:"varint,1,opt,name=UserId,proto3" json:"UserId,omitempty"`
	AccountId   int64  `protobuf:"varint,2,opt,name=AccountId,proto3" json:"AccountId,omitempty"`
	PlayerId    int64  `protobuf:"varint,3,opt,name=PlayerId,proto3" json:"PlayerId,omitempty"`
	PlayerName  string `protobuf:"bytes,4,opt,name=PlayerName,proto3" json:"PlayerName,omitempty"`
	PlayerLevel int32  `protobuf:"varint,5,opt,name=PlayerLevel,proto3" json:"PlayerLevel,omitempty"`
}

func (x *S2C_AccountLogon) Reset() {
	*x = S2C_AccountLogon{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_AccountLogon) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_AccountLogon) ProtoMessage() {}

func (x *S2C_AccountLogon) ProtoReflect() protoreflect.Message {
	mi := &file_logon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_AccountLogon.ProtoReflect.Descriptor instead.
func (*S2C_AccountLogon) Descriptor() ([]byte, []int) {
	return file_logon_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_AccountLogon) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *S2C_AccountLogon) GetAccountId() int64 {
	if x != nil {
		return x.AccountId
	}
	return 0
}

func (x *S2C_AccountLogon) GetPlayerId() int64 {
	if x != nil {
		return x.PlayerId
	}
	return 0
}

func (x *S2C_AccountLogon) GetPlayerName() string {
	if x != nil {
		return x.PlayerName
	}
	return ""
}

func (x *S2C_AccountLogon) GetPlayerLevel() int32 {
	if x != nil {
		return x.PlayerLevel
	}
	return 0
}

// 客户端心跳包
type C2S_HeartBeat struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2S_HeartBeat) Reset() {
	*x = C2S_HeartBeat{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_HeartBeat) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_HeartBeat) ProtoMessage() {}

func (x *C2S_HeartBeat) ProtoReflect() protoreflect.Message {
	mi := &file_logon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_HeartBeat.ProtoReflect.Descriptor instead.
func (*C2S_HeartBeat) Descriptor() ([]byte, []int) {
	return file_logon_proto_rawDescGZIP(), []int{2}
}

type S2C_ServerTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp uint32 `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
}

func (x *S2C_ServerTime) Reset() {
	*x = S2C_ServerTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_ServerTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_ServerTime) ProtoMessage() {}

func (x *S2C_ServerTime) ProtoReflect() protoreflect.Message {
	mi := &file_logon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_ServerTime.ProtoReflect.Descriptor instead.
func (*S2C_ServerTime) Descriptor() ([]byte, []int) {
	return file_logon_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_ServerTime) GetTimestamp() uint32 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

// 客户端主动断开连接
type C2S_AccountDisconnect struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *C2S_AccountDisconnect) Reset() {
	*x = C2S_AccountDisconnect{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_AccountDisconnect) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_AccountDisconnect) ProtoMessage() {}

func (x *C2S_AccountDisconnect) ProtoReflect() protoreflect.Message {
	mi := &file_logon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_AccountDisconnect.ProtoReflect.Descriptor instead.
func (*C2S_AccountDisconnect) Descriptor() ([]byte, []int) {
	return file_logon_proto_rawDescGZIP(), []int{4}
}

// ping
type C2S_Ping struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Ping int32 `protobuf:"varint,1,opt,name=Ping,proto3" json:"Ping,omitempty"`
}

func (x *C2S_Ping) Reset() {
	*x = C2S_Ping{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logon_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_Ping) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_Ping) ProtoMessage() {}

func (x *C2S_Ping) ProtoReflect() protoreflect.Message {
	mi := &file_logon_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_Ping.ProtoReflect.Descriptor instead.
func (*C2S_Ping) Descriptor() ([]byte, []int) {
	return file_logon_proto_rawDescGZIP(), []int{5}
}

func (x *C2S_Ping) GetPing() int32 {
	if x != nil {
		return x.Ping
	}
	return 0
}

type S2C_Pong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Pong int32 `protobuf:"varint,1,opt,name=Pong,proto3" json:"Pong,omitempty"`
}

func (x *S2C_Pong) Reset() {
	*x = S2C_Pong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_logon_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_Pong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_Pong) ProtoMessage() {}

func (x *S2C_Pong) ProtoReflect() protoreflect.Message {
	mi := &file_logon_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_Pong.ProtoReflect.Descriptor instead.
func (*S2C_Pong) Descriptor() ([]byte, []int) {
	return file_logon_proto_rawDescGZIP(), []int{6}
}

func (x *S2C_Pong) GetPong() int32 {
	if x != nil {
		return x.Pong
	}
	return 0
}

var File_logon_proto protoreflect.FileDescriptor

var file_logon_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6c, 0x6f, 0x67, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x22, 0x6a, 0x0a, 0x10, 0x43, 0x32, 0x53, 0x5f, 0x41, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64,
	0x12, 0x1c, 0x0a, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x20,
	0x0a, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0b, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x4e, 0x61, 0x6d, 0x65,
	0x22, 0xa6, 0x01, 0x0a, 0x10, 0x53, 0x32, 0x43, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x4c, 0x6f, 0x67, 0x6f, 0x6e, 0x12, 0x16, 0x0a, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x55, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c, 0x0a,
	0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x09, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x50,
	0x6c, 0x61, 0x79, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x50, 0x6c, 0x61,
	0x79, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x50, 0x6c, 0x61, 0x79, 0x65,
	0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x50, 0x6c,
	0x61, 0x79, 0x65, 0x72, 0x4c, 0x65, 0x76, 0x65, 0x6c, 0x22, 0x0f, 0x0a, 0x0d, 0x43, 0x32, 0x53,
	0x5f, 0x48, 0x65, 0x61, 0x72, 0x74, 0x42, 0x65, 0x61, 0x74, 0x22, 0x2e, 0x0a, 0x0e, 0x53, 0x32,
	0x43, 0x5f, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0d, 0x52,
	0x09, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x22, 0x17, 0x0a, 0x15, 0x43, 0x32,
	0x53, 0x5f, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x44, 0x69, 0x73, 0x63, 0x6f, 0x6e, 0x6e,
	0x65, 0x63, 0x74, 0x22, 0x1e, 0x0a, 0x08, 0x43, 0x32, 0x53, 0x5f, 0x50, 0x69, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x69, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50,
	0x69, 0x6e, 0x67, 0x22, 0x1e, 0x0a, 0x08, 0x53, 0x32, 0x43, 0x5f, 0x50, 0x6f, 0x6e, 0x67, 0x12,
	0x12, 0x0a, 0x04, 0x50, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x50,
	0x6f, 0x6e, 0x67, 0x42, 0x32, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0xaa,
	0x02, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_logon_proto_rawDescOnce sync.Once
	file_logon_proto_rawDescData = file_logon_proto_rawDesc
)

func file_logon_proto_rawDescGZIP() []byte {
	file_logon_proto_rawDescOnce.Do(func() {
		file_logon_proto_rawDescData = protoimpl.X.CompressGZIP(file_logon_proto_rawDescData)
	})
	return file_logon_proto_rawDescData
}

var file_logon_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_logon_proto_goTypes = []interface{}{
	(*C2S_AccountLogon)(nil),      // 0: proto.C2S_AccountLogon
	(*S2C_AccountLogon)(nil),      // 1: proto.S2C_AccountLogon
	(*C2S_HeartBeat)(nil),         // 2: proto.C2S_HeartBeat
	(*S2C_ServerTime)(nil),        // 3: proto.S2C_ServerTime
	(*C2S_AccountDisconnect)(nil), // 4: proto.C2S_AccountDisconnect
	(*C2S_Ping)(nil),              // 5: proto.C2S_Ping
	(*S2C_Pong)(nil),              // 6: proto.S2C_Pong
}
var file_logon_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_logon_proto_init() }
func file_logon_proto_init() {
	if File_logon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_logon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_AccountLogon); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_AccountLogon); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_HeartBeat); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_ServerTime); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_AccountDisconnect); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logon_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_Ping); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_logon_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_Pong); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_logon_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_logon_proto_goTypes,
		DependencyIndexes: file_logon_proto_depIdxs,
		MessageInfos:      file_logon_proto_msgTypes,
	}.Build()
	File_logon_proto = out.File
	file_logon_proto_rawDesc = nil
	file_logon_proto_goTypes = nil
	file_logon_proto_depIdxs = nil
}
