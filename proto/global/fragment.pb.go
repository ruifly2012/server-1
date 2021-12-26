// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.17.3
// source: fragment.proto

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

type S2C_HeroFragmentsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frags []*Fragment `protobuf:"bytes,1,rep,name=frags,proto3" json:"frags,omitempty"`
}

func (x *S2C_HeroFragmentsList) Reset() {
	*x = S2C_HeroFragmentsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fragment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeroFragmentsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeroFragmentsList) ProtoMessage() {}

func (x *S2C_HeroFragmentsList) ProtoReflect() protoreflect.Message {
	mi := &file_fragment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeroFragmentsList.ProtoReflect.Descriptor instead.
func (*S2C_HeroFragmentsList) Descriptor() ([]byte, []int) {
	return file_fragment_proto_rawDescGZIP(), []int{0}
}

func (x *S2C_HeroFragmentsList) GetFrags() []*Fragment {
	if x != nil {
		return x.Frags
	}
	return nil
}

type S2C_CollectionFragmentsList struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frags []*Fragment `protobuf:"bytes,1,rep,name=frags,proto3" json:"frags,omitempty"`
}

func (x *S2C_CollectionFragmentsList) Reset() {
	*x = S2C_CollectionFragmentsList{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fragment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_CollectionFragmentsList) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_CollectionFragmentsList) ProtoMessage() {}

func (x *S2C_CollectionFragmentsList) ProtoReflect() protoreflect.Message {
	mi := &file_fragment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_CollectionFragmentsList.ProtoReflect.Descriptor instead.
func (*S2C_CollectionFragmentsList) Descriptor() ([]byte, []int) {
	return file_fragment_proto_rawDescGZIP(), []int{1}
}

func (x *S2C_CollectionFragmentsList) GetFrags() []*Fragment {
	if x != nil {
		return x.Frags
	}
	return nil
}

// 英雄碎片更新
type S2C_HeroFragmentsUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frags []*Fragment `protobuf:"bytes,1,rep,name=frags,proto3" json:"frags,omitempty"`
}

func (x *S2C_HeroFragmentsUpdate) Reset() {
	*x = S2C_HeroFragmentsUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fragment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_HeroFragmentsUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_HeroFragmentsUpdate) ProtoMessage() {}

func (x *S2C_HeroFragmentsUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_fragment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_HeroFragmentsUpdate.ProtoReflect.Descriptor instead.
func (*S2C_HeroFragmentsUpdate) Descriptor() ([]byte, []int) {
	return file_fragment_proto_rawDescGZIP(), []int{2}
}

func (x *S2C_HeroFragmentsUpdate) GetFrags() []*Fragment {
	if x != nil {
		return x.Frags
	}
	return nil
}

// 收集品碎片更新
type S2C_CollectionFragmentsUpdate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Frags []*Fragment `protobuf:"bytes,1,rep,name=frags,proto3" json:"frags,omitempty"`
}

func (x *S2C_CollectionFragmentsUpdate) Reset() {
	*x = S2C_CollectionFragmentsUpdate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fragment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *S2C_CollectionFragmentsUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*S2C_CollectionFragmentsUpdate) ProtoMessage() {}

func (x *S2C_CollectionFragmentsUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_fragment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use S2C_CollectionFragmentsUpdate.ProtoReflect.Descriptor instead.
func (*S2C_CollectionFragmentsUpdate) Descriptor() ([]byte, []int) {
	return file_fragment_proto_rawDescGZIP(), []int{3}
}

func (x *S2C_CollectionFragmentsUpdate) GetFrags() []*Fragment {
	if x != nil {
		return x.Frags
	}
	return nil
}

// 英雄碎片合成
type C2S_HeroFragmentsCompose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FragId int32 `protobuf:"varint,1,opt,name=FragId,proto3" json:"FragId,omitempty"`
}

func (x *C2S_HeroFragmentsCompose) Reset() {
	*x = C2S_HeroFragmentsCompose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fragment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_HeroFragmentsCompose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_HeroFragmentsCompose) ProtoMessage() {}

func (x *C2S_HeroFragmentsCompose) ProtoReflect() protoreflect.Message {
	mi := &file_fragment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_HeroFragmentsCompose.ProtoReflect.Descriptor instead.
func (*C2S_HeroFragmentsCompose) Descriptor() ([]byte, []int) {
	return file_fragment_proto_rawDescGZIP(), []int{4}
}

func (x *C2S_HeroFragmentsCompose) GetFragId() int32 {
	if x != nil {
		return x.FragId
	}
	return 0
}

// 收集品碎片合成
type C2S_CollectionFragmentsCompose struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	FragId int32 `protobuf:"varint,1,opt,name=FragId,proto3" json:"FragId,omitempty"`
}

func (x *C2S_CollectionFragmentsCompose) Reset() {
	*x = C2S_CollectionFragmentsCompose{}
	if protoimpl.UnsafeEnabled {
		mi := &file_fragment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *C2S_CollectionFragmentsCompose) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*C2S_CollectionFragmentsCompose) ProtoMessage() {}

func (x *C2S_CollectionFragmentsCompose) ProtoReflect() protoreflect.Message {
	mi := &file_fragment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use C2S_CollectionFragmentsCompose.ProtoReflect.Descriptor instead.
func (*C2S_CollectionFragmentsCompose) Descriptor() ([]byte, []int) {
	return file_fragment_proto_rawDescGZIP(), []int{5}
}

func (x *C2S_CollectionFragmentsCompose) GetFragId() int32 {
	if x != nil {
		return x.FragId
	}
	return 0
}

var File_fragment_proto protoreflect.FileDescriptor

var file_fragment_proto_rawDesc = []byte{
	0x0a, 0x0e, 0x66, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x0c, 0x64, 0x65, 0x66, 0x69, 0x6e, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x3e, 0x0a, 0x15, 0x53, 0x32, 0x43, 0x5f, 0x48, 0x65, 0x72,
	0x6f, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x25,
	0x0a, 0x05, 0x66, 0x72, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x66, 0x72, 0x61, 0x67, 0x73, 0x22, 0x44, 0x0a, 0x1b, 0x53, 0x32, 0x43, 0x5f, 0x43, 0x6f, 0x6c,
	0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x4c, 0x69, 0x73, 0x74, 0x12, 0x25, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x67,
	0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x66, 0x72, 0x61, 0x67, 0x73, 0x22, 0x40, 0x0a, 0x17, 0x53,
	0x32, 0x43, 0x5f, 0x48, 0x65, 0x72, 0x6f, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73,
	0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25, 0x0a, 0x05, 0x66, 0x72, 0x61, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72,
	0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05, 0x66, 0x72, 0x61, 0x67, 0x73, 0x22, 0x46, 0x0a,
	0x1d, 0x53, 0x32, 0x43, 0x5f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46,
	0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x12, 0x25,
	0x0a, 0x05, 0x66, 0x72, 0x61, 0x67, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0f, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2e, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x05,
	0x66, 0x72, 0x61, 0x67, 0x73, 0x22, 0x32, 0x0a, 0x18, 0x43, 0x32, 0x53, 0x5f, 0x48, 0x65, 0x72,
	0x6f, 0x46, 0x72, 0x61, 0x67, 0x6d, 0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x46, 0x72, 0x61, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x06, 0x46, 0x72, 0x61, 0x67, 0x49, 0x64, 0x22, 0x38, 0x0a, 0x1e, 0x43, 0x32, 0x53,
	0x5f, 0x43, 0x6f, 0x6c, 0x6c, 0x65, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x46, 0x72, 0x61, 0x67, 0x6d,
	0x65, 0x6e, 0x74, 0x73, 0x43, 0x6f, 0x6d, 0x70, 0x6f, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x46,
	0x72, 0x61, 0x67, 0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x06, 0x46, 0x72, 0x61,
	0x67, 0x49, 0x64, 0x42, 0x32, 0x5a, 0x28, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x65, 0x61, 0x73, 0x74, 0x2d, 0x65, 0x64, 0x65, 0x6e, 0x2f, 0x73, 0x65, 0x72, 0x76,
	0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x6c, 0x6f, 0x62, 0x61, 0x6c, 0xaa,
	0x02, 0x05, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_fragment_proto_rawDescOnce sync.Once
	file_fragment_proto_rawDescData = file_fragment_proto_rawDesc
)

func file_fragment_proto_rawDescGZIP() []byte {
	file_fragment_proto_rawDescOnce.Do(func() {
		file_fragment_proto_rawDescData = protoimpl.X.CompressGZIP(file_fragment_proto_rawDescData)
	})
	return file_fragment_proto_rawDescData
}

var file_fragment_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_fragment_proto_goTypes = []interface{}{
	(*S2C_HeroFragmentsList)(nil),          // 0: proto.S2C_HeroFragmentsList
	(*S2C_CollectionFragmentsList)(nil),    // 1: proto.S2C_CollectionFragmentsList
	(*S2C_HeroFragmentsUpdate)(nil),        // 2: proto.S2C_HeroFragmentsUpdate
	(*S2C_CollectionFragmentsUpdate)(nil),  // 3: proto.S2C_CollectionFragmentsUpdate
	(*C2S_HeroFragmentsCompose)(nil),       // 4: proto.C2S_HeroFragmentsCompose
	(*C2S_CollectionFragmentsCompose)(nil), // 5: proto.C2S_CollectionFragmentsCompose
	(*Fragment)(nil),                       // 6: proto.Fragment
}
var file_fragment_proto_depIdxs = []int32{
	6, // 0: proto.S2C_HeroFragmentsList.frags:type_name -> proto.Fragment
	6, // 1: proto.S2C_CollectionFragmentsList.frags:type_name -> proto.Fragment
	6, // 2: proto.S2C_HeroFragmentsUpdate.frags:type_name -> proto.Fragment
	6, // 3: proto.S2C_CollectionFragmentsUpdate.frags:type_name -> proto.Fragment
	4, // [4:4] is the sub-list for method output_type
	4, // [4:4] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_fragment_proto_init() }
func file_fragment_proto_init() {
	if File_fragment_proto != nil {
		return
	}
	file_define_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_fragment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeroFragmentsList); i {
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
		file_fragment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_CollectionFragmentsList); i {
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
		file_fragment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_HeroFragmentsUpdate); i {
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
		file_fragment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*S2C_CollectionFragmentsUpdate); i {
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
		file_fragment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_HeroFragmentsCompose); i {
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
		file_fragment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*C2S_CollectionFragmentsCompose); i {
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
			RawDescriptor: file_fragment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_fragment_proto_goTypes,
		DependencyIndexes: file_fragment_proto_depIdxs,
		MessageInfos:      file_fragment_proto_msgTypes,
	}.Build()
	File_fragment_proto = out.File
	file_fragment_proto_rawDesc = nil
	file_fragment_proto_goTypes = nil
	file_fragment_proto_depIdxs = nil
}
