// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.14.0
// source: combat/combat.proto

package combat

import (
	game "e.coding.net/mmstudio/blade/proto/go_out/game"
	proto "github.com/golang/protobuf/proto"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

// unit scene att
type UnitInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UnitTypeId  int32       `protobuf:"varint,1,opt,name=UnitTypeId,proto3" json:"UnitTypeId,omitempty"`
	UnitAttList []*game.Att `protobuf:"bytes,2,rep,name=UnitAttList,proto3" json:"UnitAttList,omitempty"`
}

func (x *UnitInfo) Reset() {
	*x = UnitInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_combat_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UnitInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UnitInfo) ProtoMessage() {}

func (x *UnitInfo) ProtoReflect() protoreflect.Message {
	mi := &file_combat_combat_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UnitInfo.ProtoReflect.Descriptor instead.
func (*UnitInfo) Descriptor() ([]byte, []int) {
	return file_combat_combat_proto_rawDescGZIP(), []int{0}
}

func (x *UnitInfo) GetUnitTypeId() int32 {
	if x != nil {
		return x.UnitTypeId
	}
	return 0
}

func (x *UnitInfo) GetUnitAttList() []*game.Att {
	if x != nil {
		return x.UnitAttList
	}
	return nil
}

type StartStageCombatReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId         int64       `protobuf:"varint,1,opt,name=SceneId,proto3" json:"SceneId,omitempty"` // generated by game service
	SceneType       int32       `protobuf:"varint,2,opt,name=SceneType,proto3" json:"SceneType,omitempty"`
	AttackId        int64       `protobuf:"varint,3,opt,name=AttackId,proto3" json:"AttackId,omitempty"`
	AttackUnitList  []*UnitInfo `protobuf:"bytes,4,rep,name=AttackUnitList,proto3" json:"AttackUnitList,omitempty"`
	DefenceId       int64       `protobuf:"varint,5,opt,name=DefenceId,proto3" json:"DefenceId,omitempty"`
	DefenceUnitList []*UnitInfo `protobuf:"bytes,6,rep,name=DefenceUnitList,proto3" json:"DefenceUnitList,omitempty"`
}

func (x *StartStageCombatReq) Reset() {
	*x = StartStageCombatReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_combat_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartStageCombatReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStageCombatReq) ProtoMessage() {}

func (x *StartStageCombatReq) ProtoReflect() protoreflect.Message {
	mi := &file_combat_combat_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStageCombatReq.ProtoReflect.Descriptor instead.
func (*StartStageCombatReq) Descriptor() ([]byte, []int) {
	return file_combat_combat_proto_rawDescGZIP(), []int{1}
}

func (x *StartStageCombatReq) GetSceneId() int64 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *StartStageCombatReq) GetSceneType() int32 {
	if x != nil {
		return x.SceneType
	}
	return 0
}

func (x *StartStageCombatReq) GetAttackId() int64 {
	if x != nil {
		return x.AttackId
	}
	return 0
}

func (x *StartStageCombatReq) GetAttackUnitList() []*UnitInfo {
	if x != nil {
		return x.AttackUnitList
	}
	return nil
}

func (x *StartStageCombatReq) GetDefenceId() int64 {
	if x != nil {
		return x.DefenceId
	}
	return 0
}

func (x *StartStageCombatReq) GetDefenceUnitList() []*UnitInfo {
	if x != nil {
		return x.DefenceUnitList
	}
	return nil
}

type StartStageCombatReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SceneId   int64 `protobuf:"varint,1,opt,name=SceneId,proto3" json:"SceneId,omitempty"`
	AttackId  int64 `protobuf:"varint,2,opt,name=AttackId,proto3" json:"AttackId,omitempty"`
	DefenceId int64 `protobuf:"varint,3,opt,name=DefenceId,proto3" json:"DefenceId,omitempty"`
	Result    bool  `protobuf:"varint,4,opt,name=Result,proto3" json:"Result,omitempty"` // true == win
}

func (x *StartStageCombatReply) Reset() {
	*x = StartStageCombatReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_combat_combat_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StartStageCombatReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StartStageCombatReply) ProtoMessage() {}

func (x *StartStageCombatReply) ProtoReflect() protoreflect.Message {
	mi := &file_combat_combat_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StartStageCombatReply.ProtoReflect.Descriptor instead.
func (*StartStageCombatReply) Descriptor() ([]byte, []int) {
	return file_combat_combat_proto_rawDescGZIP(), []int{2}
}

func (x *StartStageCombatReply) GetSceneId() int64 {
	if x != nil {
		return x.SceneId
	}
	return 0
}

func (x *StartStageCombatReply) GetAttackId() int64 {
	if x != nil {
		return x.AttackId
	}
	return 0
}

func (x *StartStageCombatReply) GetDefenceId() int64 {
	if x != nil {
		return x.DefenceId
	}
	return 0
}

func (x *StartStageCombatReply) GetResult() bool {
	if x != nil {
		return x.Result
	}
	return false
}

var File_combat_combat_proto protoreflect.FileDescriptor

var file_combat_combat_proto_rawDesc = []byte{
	0x0a, 0x13, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2f, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x1a, 0x0e, 0x67,
	0x61, 0x6d, 0x65, 0x2f, 0x61, 0x74, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x57, 0x0a,
	0x08, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1e, 0x0a, 0x0a, 0x55, 0x6e, 0x69,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x55,
	0x6e, 0x69, 0x74, 0x54, 0x79, 0x70, 0x65, 0x49, 0x64, 0x12, 0x2b, 0x0a, 0x0b, 0x55, 0x6e, 0x69,
	0x74, 0x41, 0x74, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09,
	0x2e, 0x67, 0x61, 0x6d, 0x65, 0x2e, 0x41, 0x74, 0x74, 0x52, 0x0b, 0x55, 0x6e, 0x69, 0x74, 0x41,
	0x74, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0xfd, 0x01, 0x0a, 0x13, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x71, 0x12, 0x18,
	0x0a, 0x07, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x07, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x53, 0x63, 0x65, 0x6e,
	0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x53, 0x63, 0x65,
	0x6e, 0x65, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b,
	0x49, 0x64, 0x12, 0x38, 0x0a, 0x0e, 0x41, 0x74, 0x74, 0x61, 0x63, 0x6b, 0x55, 0x6e, 0x69, 0x74,
	0x4c, 0x69, 0x73, 0x74, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x62, 0x61, 0x74, 0x2e, 0x55, 0x6e, 0x69, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0e, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x55, 0x6e, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09,
	0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x09, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x49, 0x64, 0x12, 0x3a, 0x0a, 0x0f, 0x44, 0x65,
	0x66, 0x65, 0x6e, 0x63, 0x65, 0x55, 0x6e, 0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x06, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x55, 0x6e, 0x69,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52, 0x0f, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63, 0x65, 0x55, 0x6e,
	0x69, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x22, 0x83, 0x01, 0x0a, 0x15, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79,
	0x12, 0x18, 0x0a, 0x07, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x53, 0x63, 0x65, 0x6e, 0x65, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x41, 0x74,
	0x74, 0x61, 0x63, 0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x44, 0x65, 0x66, 0x65, 0x6e, 0x63,
	0x65, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x44, 0x65, 0x66, 0x65, 0x6e,
	0x63, 0x65, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x61, 0x0a, 0x0d,
	0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x50, 0x0a,
	0x10, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61,
	0x74, 0x12, 0x1b, 0x2e, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74,
	0x53, 0x74, 0x61, 0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x71, 0x1a, 0x1d,
	0x2e, 0x63, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x2e, 0x53, 0x74, 0x61, 0x72, 0x74, 0x53, 0x74, 0x61,
	0x67, 0x65, 0x43, 0x6f, 0x6d, 0x62, 0x61, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42,
	0x31, 0x5a, 0x2f, 0x65, 0x2e, 0x63, 0x6f, 0x64, 0x69, 0x6e, 0x67, 0x2e, 0x6e, 0x65, 0x74, 0x2f,
	0x6d, 0x6d, 0x73, 0x74, 0x75, 0x64, 0x69, 0x6f, 0x2f, 0x62, 0x6c, 0x61, 0x64, 0x65, 0x2f, 0x73,
	0x65, 0x72, 0x76, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x62,
	0x61, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_combat_combat_proto_rawDescOnce sync.Once
	file_combat_combat_proto_rawDescData = file_combat_combat_proto_rawDesc
)

func file_combat_combat_proto_rawDescGZIP() []byte {
	file_combat_combat_proto_rawDescOnce.Do(func() {
		file_combat_combat_proto_rawDescData = protoimpl.X.CompressGZIP(file_combat_combat_proto_rawDescData)
	})
	return file_combat_combat_proto_rawDescData
}

var file_combat_combat_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_combat_combat_proto_goTypes = []interface{}{
	(*UnitInfo)(nil),              // 0: combat.UnitInfo
	(*StartStageCombatReq)(nil),   // 1: combat.StartStageCombatReq
	(*StartStageCombatReply)(nil), // 2: combat.StartStageCombatReply
	(*game.Att)(nil),              // 3: game.Att
}
var file_combat_combat_proto_depIdxs = []int32{
	3, // 0: combat.UnitInfo.UnitAttList:type_name -> game.Att
	0, // 1: combat.StartStageCombatReq.AttackUnitList:type_name -> combat.UnitInfo
	0, // 2: combat.StartStageCombatReq.DefenceUnitList:type_name -> combat.UnitInfo
	1, // 3: combat.CombatService.StartStageCombat:input_type -> combat.StartStageCombatReq
	2, // 4: combat.CombatService.StartStageCombat:output_type -> combat.StartStageCombatReply
	4, // [4:5] is the sub-list for method output_type
	3, // [3:4] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_combat_combat_proto_init() }
func file_combat_combat_proto_init() {
	if File_combat_combat_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_combat_combat_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UnitInfo); i {
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
		file_combat_combat_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartStageCombatReq); i {
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
		file_combat_combat_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StartStageCombatReply); i {
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
			RawDescriptor: file_combat_combat_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_combat_combat_proto_goTypes,
		DependencyIndexes: file_combat_combat_proto_depIdxs,
		MessageInfos:      file_combat_combat_proto_msgTypes,
	}.Build()
	File_combat_combat_proto = out.File
	file_combat_combat_proto_rawDesc = nil
	file_combat_combat_proto_goTypes = nil
	file_combat_combat_proto_depIdxs = nil
}
