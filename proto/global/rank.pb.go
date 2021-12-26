// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: rank.proto

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

// 查询排行信息
type C2S_QueryRank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId int32 `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"` // 排行榜id
}

func (x *C2S_QueryRank) Reset() {
	*x = C2S_QueryRank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_QueryRank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_QueryRank) ProtoMessage() {}

func (x *C2S_QueryRank) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_QueryRank.ProtoReflect.Descriptor instead.
func (*C2S_QueryRank) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{0}
}

func (x *C2S_QueryRank) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

// 排行信息
type S2C_QueryRank struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RankId    int32         `protobuf:"varint,1,opt,name=RankId,proto3" json:"RankId,omitempty"`
	RankIndex int32         `protobuf:"varint,2,opt,name=RankIndex,proto3" json:"RankIndex,omitempty"`
	Metadata  *RankMetadata `protobuf:"bytes,3,opt,name=Metadata,proto3" json:"Metadata,omitempty"`
}

func (x *S2C_QueryRank) Reset() {
	*x = S2C_QueryRank{}
	if protoimpl.UnsafeEnabled {
		mi := &file_rank_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_QueryRank) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_QueryRank) ProtoMessage() {}

func (x *S2C_QueryRank) ProtoReflect() protoreflect.Message {
	mi := &file_rank_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_QueryRank.ProtoReflect.Descriptor instead.
func (*S2C_QueryRank) Descriptor() ([]byte, []int) {
	return file_rank_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_QueryRank) GetRankId() int32 {
	if x != nil {
		return x.RankId
	}
	return 0
}

func (x *S2C_QueryRank) GetRankIndex() int32 {
	if x != nil {
		return x.RankIndex
	}
	return 0
}

func (x *S2C_QueryRank) GetMetadata() *RankMetadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

var File_rank_proto protoreflect.FileDescriptor

var file_rank_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x72, 0x61, 0x6e, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x0d, 0x43, 0x32, 0x53, 0x5f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61,
	0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x64, 0x22, 0x76, 0x0a, 0x0d, 0x53, 0x32,
	0x43, 0x5f, 0x51, 0x75, 0x65, 0x72, 0x79, 0x52, 0x61, 0x6e, 0x6b, 0x12, 0x16, 0x0a, 0x06, 0x52,
	0x61, 0x6e, 0x6b, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x52, 0x61, 0x6e,
	0x6b, 0x49, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65, 0x78,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x52, 0x61, 0x6e, 0x6b, 0x49, 0x6e, 0x64, 0x65,
	0x78, 0x12, 0x20, 0x0a, 0x03, 0x52, 0x61, 0x77, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0e,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x52, 0x61, 0x6e, 0x6b, 0x52, 0x61, 0x77, 0x52, 0x03,
	0x52, 0x61, 0x77, 0x42, 0x32, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0xaa,
	0x02, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_rank_proto_rawDescOnce sync.Once
	file_rank_proto_rawDescData = file_rank_proto_rawDesc
)

func file_rank_proto_rawDescGZIP() []byte {
	file_rank_proto_rawDescOnce.Do(func() {
		file_rank_proto_rawDescData = protoimpl.X.CompressGZIP(file_rank_proto_rawDescData)
	})
	return file_rank_proto_rawDescData
}

var file_rank_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_rank_proto_goTypes = []interface{}{
	(*C2S_QueryRank)(nil), // 0: proto.C2S_QueryRank
	(*S2C_QueryRank)(nil), // 1: proto.S2C_QueryRank
	(*RankMetadata)(nil),  // 2: proto.RankMetadata
}
var file_rank_proto_depIdxs = []int32{
	2, // 0: proto.S2C_QueryRank.Metadata:type_name -> proto.RankMetadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_rank_proto_init() }
func file_rank_proto_init() {
	if File_rank_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_rank_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_QueryRank); i {
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
		file_rank_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_QueryRank); i {
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
			RawDescriptor: file_rank_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_rank_proto_goTypes,
		DependencyIndexes: file_rank_proto_depIdxs,
		MessageInfos:      file_rank_proto_msgTypes,
	}.Build()
	File_rank_proto = out.File
	file_rank_proto_rawDesc = nil
	file_rank_proto_goTypes = nil
	file_rank_proto_depIdxs = nil
}
