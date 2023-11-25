// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: proto/mav_controller_outside/trjxcommon.proto

package trjxmavlink

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

type Result_CodeDef int32

const (
	Result_Complete        Result_CodeDef = 0
	Result_InternalError   Result_CodeDef = 1
	Result_NoExistAircraft Result_CodeDef = 2
)

// Enum value maps for Result_CodeDef.
var (
	Result_CodeDef_name = map[int32]string{
		0: "Complete",
		1: "InternalError",
		2: "NoExistAircraft",
	}
	Result_CodeDef_value = map[string]int32{
		"Complete":        0,
		"InternalError":   1,
		"NoExistAircraft": 2,
	}
)

func (x Result_CodeDef) Enum() *Result_CodeDef {
	p := new(Result_CodeDef)
	*p = x
	return p
}

func (x Result_CodeDef) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Result_CodeDef) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_mav_controller_outside_trjxcommon_proto_enumTypes[0].Descriptor()
}

func (Result_CodeDef) Type() protoreflect.EnumType {
	return &file_proto_mav_controller_outside_trjxcommon_proto_enumTypes[0]
}

func (x Result_CodeDef) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Result_CodeDef.Descriptor instead.
func (Result_CodeDef) EnumDescriptor() ([]byte, []int) {
	return file_proto_mav_controller_outside_trjxcommon_proto_rawDescGZIP(), []int{3, 0}
}

type SystemVersion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Major int32  `protobuf:"varint,1,opt,name=major,proto3" json:"major,omitempty"`
	Minor int32  `protobuf:"varint,2,opt,name=minor,proto3" json:"minor,omitempty"`
	Build string `protobuf:"bytes,3,opt,name=build,proto3" json:"build,omitempty"`
}

func (x *SystemVersion) Reset() {
	*x = SystemVersion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SystemVersion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SystemVersion) ProtoMessage() {}

func (x *SystemVersion) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SystemVersion.ProtoReflect.Descriptor instead.
func (*SystemVersion) Descriptor() ([]byte, []int) {
	return file_proto_mav_controller_outside_trjxcommon_proto_rawDescGZIP(), []int{0}
}

func (x *SystemVersion) GetMajor() int32 {
	if x != nil {
		return x.Major
	}
	return 0
}

func (x *SystemVersion) GetMinor() int32 {
	if x != nil {
		return x.Minor
	}
	return 0
}

func (x *SystemVersion) GetBuild() string {
	if x != nil {
		return x.Build
	}
	return ""
}

type SeriviceAttribute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Version *SystemVersion `protobuf:"bytes,1,opt,name=Version,proto3" json:"Version,omitempty"`
}

func (x *SeriviceAttribute) Reset() {
	*x = SeriviceAttribute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SeriviceAttribute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SeriviceAttribute) ProtoMessage() {}

func (x *SeriviceAttribute) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SeriviceAttribute.ProtoReflect.Descriptor instead.
func (*SeriviceAttribute) Descriptor() ([]byte, []int) {
	return file_proto_mav_controller_outside_trjxcommon_proto_rawDescGZIP(), []int{1}
}

func (x *SeriviceAttribute) GetVersion() *SystemVersion {
	if x != nil {
		return x.Version
	}
	return nil
}

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_mav_controller_outside_trjxcommon_proto_rawDescGZIP(), []int{2}
}

type Result struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   Result_CodeDef `protobuf:"varint,1,opt,name=Code,proto3,enum=trjxmavlink.Result_CodeDef" json:"Code,omitempty"`
	Detail string         `protobuf:"bytes,2,opt,name=Detail,proto3" json:"Detail,omitempty"`
}

func (x *Result) Reset() {
	*x = Result{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Result) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Result) ProtoMessage() {}

func (x *Result) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Result.ProtoReflect.Descriptor instead.
func (*Result) Descriptor() ([]byte, []int) {
	return file_proto_mav_controller_outside_trjxcommon_proto_rawDescGZIP(), []int{3}
}

func (x *Result) GetCode() Result_CodeDef {
	if x != nil {
		return x.Code
	}
	return Result_Complete
}

func (x *Result) GetDetail() string {
	if x != nil {
		return x.Detail
	}
	return ""
}

type AutopilotModelDef struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Model:
	//	*AutopilotModelDef_Name
	//	*AutopilotModelDef_ModelSet
	Model isAutopilotModelDef_Model `protobuf_oneof:"Model"`
}

func (x *AutopilotModelDef) Reset() {
	*x = AutopilotModelDef{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AutopilotModelDef) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AutopilotModelDef) ProtoMessage() {}

func (x *AutopilotModelDef) ProtoReflect() protoreflect.Message {
	mi := &file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AutopilotModelDef.ProtoReflect.Descriptor instead.
func (*AutopilotModelDef) Descriptor() ([]byte, []int) {
	return file_proto_mav_controller_outside_trjxcommon_proto_rawDescGZIP(), []int{4}
}

func (m *AutopilotModelDef) GetModel() isAutopilotModelDef_Model {
	if m != nil {
		return m.Model
	}
	return nil
}

func (x *AutopilotModelDef) GetName() string {
	if x, ok := x.GetModel().(*AutopilotModelDef_Name); ok {
		return x.Name
	}
	return ""
}

func (x *AutopilotModelDef) GetModelSet() string {
	if x, ok := x.GetModel().(*AutopilotModelDef_ModelSet); ok {
		return x.ModelSet
	}
	return ""
}

type isAutopilotModelDef_Model interface {
	isAutopilotModelDef_Model()
}

type AutopilotModelDef_Name struct {
	Name string `protobuf:"bytes,1,opt,name=Name,proto3,oneof"` // MS model.jsonモデル名
}

type AutopilotModelDef_ModelSet struct {
	ModelSet string `protobuf:"bytes,2,opt,name=ModelSet,proto3,oneof"` // json
}

func (*AutopilotModelDef_Name) isAutopilotModelDef_Model() {}

func (*AutopilotModelDef_ModelSet) isAutopilotModelDef_Model() {}

var File_proto_mav_controller_outside_trjxcommon_proto protoreflect.FileDescriptor

var file_proto_mav_controller_outside_trjxcommon_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x76, 0x5f, 0x63, 0x6f, 0x6e, 0x74,
	0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f, 0x6f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x2f, 0x74,
	0x72, 0x6a, 0x78, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0b, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x51, 0x0a, 0x0d,
	0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x6d, 0x61, 0x6a, 0x6f, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x6d, 0x61,
	0x6a, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x6d, 0x69, 0x6e, 0x6f, 0x72, 0x12, 0x14, 0x0a, 0x05, 0x62, 0x75, 0x69,
	0x6c, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x62, 0x75, 0x69, 0x6c, 0x64, 0x22,
	0x49, 0x0a, 0x11, 0x53, 0x65, 0x72, 0x69, 0x76, 0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69,
	0x62, 0x75, 0x74, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c,
	0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f,
	0x6e, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d,
	0x70, 0x74, 0x79, 0x22, 0x92, 0x01, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x2f,
	0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1b, 0x2e, 0x74,
	0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x44, 0x65, 0x66, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x16, 0x0a, 0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x22, 0x3f, 0x0a, 0x07, 0x43, 0x6f, 0x64, 0x65, 0x44,
	0x65, 0x66, 0x12, 0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x00,
	0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x10, 0x01, 0x12, 0x13, 0x0a, 0x0f, 0x4e, 0x6f, 0x45, 0x78, 0x69, 0x73, 0x74, 0x41, 0x69,
	0x72, 0x63, 0x72, 0x61, 0x66, 0x74, 0x10, 0x02, 0x22, 0x50, 0x0a, 0x11, 0x41, 0x75, 0x74, 0x6f,
	0x70, 0x69, 0x6c, 0x6f, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66, 0x12, 0x14, 0x0a,
	0x04, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x04, 0x4e,
	0x61, 0x6d, 0x65, 0x12, 0x1c, 0x0a, 0x08, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65, 0x74, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x08, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x53, 0x65,
	0x74, 0x42, 0x07, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x42, 0x16, 0x5a, 0x14, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f, 0x2f, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69,
	0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_mav_controller_outside_trjxcommon_proto_rawDescOnce sync.Once
	file_proto_mav_controller_outside_trjxcommon_proto_rawDescData = file_proto_mav_controller_outside_trjxcommon_proto_rawDesc
)

func file_proto_mav_controller_outside_trjxcommon_proto_rawDescGZIP() []byte {
	file_proto_mav_controller_outside_trjxcommon_proto_rawDescOnce.Do(func() {
		file_proto_mav_controller_outside_trjxcommon_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_mav_controller_outside_trjxcommon_proto_rawDescData)
	})
	return file_proto_mav_controller_outside_trjxcommon_proto_rawDescData
}

var file_proto_mav_controller_outside_trjxcommon_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_mav_controller_outside_trjxcommon_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_proto_mav_controller_outside_trjxcommon_proto_goTypes = []interface{}{
	(Result_CodeDef)(0),       // 0: trjxmavlink.Result.CodeDef
	(*SystemVersion)(nil),     // 1: trjxmavlink.SystemVersion
	(*SeriviceAttribute)(nil), // 2: trjxmavlink.SeriviceAttribute
	(*Empty)(nil),             // 3: trjxmavlink.Empty
	(*Result)(nil),            // 4: trjxmavlink.Result
	(*AutopilotModelDef)(nil), // 5: trjxmavlink.AutopilotModelDef
}
var file_proto_mav_controller_outside_trjxcommon_proto_depIdxs = []int32{
	1, // 0: trjxmavlink.SeriviceAttribute.Version:type_name -> trjxmavlink.SystemVersion
	0, // 1: trjxmavlink.Result.Code:type_name -> trjxmavlink.Result.CodeDef
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_mav_controller_outside_trjxcommon_proto_init() }
func file_proto_mav_controller_outside_trjxcommon_proto_init() {
	if File_proto_mav_controller_outside_trjxcommon_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SystemVersion); i {
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
		file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SeriviceAttribute); i {
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
		file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Result); i {
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
		file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AutopilotModelDef); i {
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
	file_proto_mav_controller_outside_trjxcommon_proto_msgTypes[4].OneofWrappers = []interface{}{
		(*AutopilotModelDef_Name)(nil),
		(*AutopilotModelDef_ModelSet)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_mav_controller_outside_trjxcommon_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_mav_controller_outside_trjxcommon_proto_goTypes,
		DependencyIndexes: file_proto_mav_controller_outside_trjxcommon_proto_depIdxs,
		EnumInfos:         file_proto_mav_controller_outside_trjxcommon_proto_enumTypes,
		MessageInfos:      file_proto_mav_controller_outside_trjxcommon_proto_msgTypes,
	}.Build()
	File_proto_mav_controller_outside_trjxcommon_proto = out.File
	file_proto_mav_controller_outside_trjxcommon_proto_rawDesc = nil
	file_proto_mav_controller_outside_trjxcommon_proto_goTypes = nil
	file_proto_mav_controller_outside_trjxcommon_proto_depIdxs = nil
}
