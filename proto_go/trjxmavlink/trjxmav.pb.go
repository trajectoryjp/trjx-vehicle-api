// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.19.4
// source: github.com/trajectoryjp/trjx-vehicle-api/proto/mav_controller_outside/trjxmav.proto

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

type Token_Code int32

const (
	Token_Complete      Token_Code = 0 // ログイン成功
	Token_Accepted      Token_Code = 1 // ユーザによって承認。パスワードリセット
	Token_Deny          Token_Code = 2 // ユーザによって拒否。（初期パスワード）
	Token_Invalid       Token_Code = 3 // パスワードが無効
	Token_InternalError Token_Code = 10
)

// Enum value maps for Token_Code.
var (
	Token_Code_name = map[int32]string{
		0:  "Complete",
		1:  "Accepted",
		2:  "Deny",
		3:  "Invalid",
		10: "InternalError",
	}
	Token_Code_value = map[string]int32{
		"Complete":      0,
		"Accepted":      1,
		"Deny":          2,
		"Invalid":       3,
		"InternalError": 10,
	}
)

func (x Token_Code) Enum() *Token_Code {
	p := new(Token_Code)
	*p = x
	return p
}

func (x Token_Code) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Token_Code) Descriptor() protoreflect.EnumDescriptor {
	return file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_enumTypes[0].Descriptor()
}

func (Token_Code) Type() protoreflect.EnumType {
	return &file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_enumTypes[0]
}

func (x Token_Code) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Token_Code.Descriptor instead.
func (Token_Code) EnumDescriptor() ([]byte, []int) {
	return file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescGZIP(), []int{0, 0}
}

type Token struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Result   Token_Code `protobuf:"varint,1,opt,name=Result,proto3,enum=trjxmavlink.Token_Code" json:"Result,omitempty"`
	Password string     `protobuf:"bytes,2,opt,name=Password,proto3" json:"Password,omitempty"`
	Token    string     `protobuf:"bytes,3,opt,name=Token,proto3" json:"Token,omitempty"`
}

func (x *Token) Reset() {
	*x = Token{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Token) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Token) ProtoMessage() {}

func (x *Token) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Token.ProtoReflect.Descriptor instead.
func (*Token) Descriptor() ([]byte, []int) {
	return file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescGZIP(), []int{0}
}

func (x *Token) GetResult() Token_Code {
	if x != nil {
		return x.Result
	}
	return Token_Complete
}

func (x *Token) GetPassword() string {
	if x != nil {
		return x.Password
	}
	return ""
}

func (x *Token) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

type SetAutopilotModelRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Model *AutopilotModelDef `protobuf:"bytes,1,opt,name=Model,proto3" json:"Model,omitempty"`
}

func (x *SetAutopilotModelRequest) Reset() {
	*x = SetAutopilotModelRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SetAutopilotModelRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SetAutopilotModelRequest) ProtoMessage() {}

func (x *SetAutopilotModelRequest) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SetAutopilotModelRequest.ProtoReflect.Descriptor instead.
func (*SetAutopilotModelRequest) Descriptor() ([]byte, []int) {
	return file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescGZIP(), []int{1}
}

func (x *SetAutopilotModelRequest) GetModel() *AutopilotModelDef {
	if x != nil {
		return x.Model
	}
	return nil
}

type TrjxMavlink struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MavlinkMessage []byte `protobuf:"bytes,5,opt,name=MavlinkMessage,proto3" json:"MavlinkMessage,omitempty"` // mavlink binary
}

func (x *TrjxMavlink) Reset() {
	*x = TrjxMavlink{}
	if protoimpl.UnsafeEnabled {
		mi := &file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TrjxMavlink) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TrjxMavlink) ProtoMessage() {}

func (x *TrjxMavlink) ProtoReflect() protoreflect.Message {
	mi := &file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TrjxMavlink.ProtoReflect.Descriptor instead.
func (*TrjxMavlink) Descriptor() ([]byte, []int) {
	return file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescGZIP(), []int{2}
}

func (x *TrjxMavlink) GetMavlinkMessage() []byte {
	if x != nil {
		return x.MavlinkMessage
	}
	return nil
}

var File_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto protoreflect.FileDescriptor

var file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDesc = []byte{
	0x0a, 0x53, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61,
	0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x6a, 0x70, 0x2f, 0x74, 0x72, 0x6a, 0x78, 0x2d, 0x76,
	0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x6d, 0x61, 0x76, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65, 0x72, 0x5f,
	0x6f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69,
	0x6e, 0x6b, 0x1a, 0x56, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74,
	0x72, 0x61, 0x6a, 0x65, 0x63, 0x74, 0x6f, 0x72, 0x79, 0x6a, 0x70, 0x2f, 0x74, 0x72, 0x6a, 0x78,
	0x2d, 0x76, 0x65, 0x68, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x2f, 0x6d, 0x61, 0x76, 0x5f, 0x63, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x6c, 0x65,
	0x72, 0x5f, 0x6f, 0x75, 0x74, 0x73, 0x69, 0x64, 0x65, 0x2f, 0x74, 0x72, 0x6a, 0x78, 0x63, 0x6f,
	0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xb8, 0x01, 0x0a, 0x05, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x2f, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x17, 0x2e, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69,
	0x6e, 0x6b, 0x2e, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x2e, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x06, 0x52,
	0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x50, 0x61, 0x73, 0x73, 0x77, 0x6f, 0x72,
	0x64, 0x12, 0x14, 0x0a, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x54, 0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x4c, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12,
	0x0c, 0x0a, 0x08, 0x43, 0x6f, 0x6d, 0x70, 0x6c, 0x65, 0x74, 0x65, 0x10, 0x00, 0x12, 0x0c, 0x0a,
	0x08, 0x41, 0x63, 0x63, 0x65, 0x70, 0x74, 0x65, 0x64, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x44,
	0x65, 0x6e, 0x79, 0x10, 0x02, 0x12, 0x0b, 0x0a, 0x07, 0x49, 0x6e, 0x76, 0x61, 0x6c, 0x69, 0x64,
	0x10, 0x03, 0x12, 0x11, 0x0a, 0x0d, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x45, 0x72,
	0x72, 0x6f, 0x72, 0x10, 0x0a, 0x22, 0x50, 0x0a, 0x18, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74, 0x6f,
	0x70, 0x69, 0x6c, 0x6f, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x34, 0x0a, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1e, 0x2e, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x41,
	0x75, 0x74, 0x6f, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x44, 0x65, 0x66,
	0x52, 0x05, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x22, 0x35, 0x0a, 0x0b, 0x54, 0x72, 0x6a, 0x78, 0x4d,
	0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x12, 0x26, 0x0a, 0x0e, 0x4d, 0x61, 0x76, 0x6c, 0x69, 0x6e,
	0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e,
	0x4d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0xba,
	0x02, 0x0a, 0x12, 0x54, 0x72, 0x6a, 0x78, 0x4d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x53, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4c, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x53, 0x65, 0x72, 0x69,
	0x76, 0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74, 0x65, 0x12, 0x12, 0x2e,
	0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x6d, 0x70, 0x74,
	0x79, 0x1a, 0x1e, 0x2e, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x53, 0x65, 0x72, 0x69, 0x76, 0x69, 0x63, 0x65, 0x41, 0x74, 0x74, 0x72, 0x69, 0x62, 0x75, 0x74,
	0x65, 0x22, 0x00, 0x12, 0x31, 0x0a, 0x05, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x12, 0x12, 0x2e, 0x74,
	0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x45, 0x6d, 0x70, 0x74, 0x79,
	0x1a, 0x12, 0x2e, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x54,
	0x6f, 0x6b, 0x65, 0x6e, 0x22, 0x00, 0x12, 0x51, 0x0a, 0x11, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x6f, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x12, 0x25, 0x2e, 0x74, 0x72,
	0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x53, 0x65, 0x74, 0x41, 0x75, 0x74,
	0x6f, 0x70, 0x69, 0x6c, 0x6f, 0x74, 0x4d, 0x6f, 0x64, 0x65, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b,
	0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x22, 0x00, 0x12, 0x50, 0x0a, 0x14, 0x43, 0x6f, 0x6d,
	0x6d, 0x75, 0x6e, 0x69, 0x63, 0x61, 0x74, 0x65, 0x4f, 0x6e, 0x4d, 0x61, 0x76, 0x6c, 0x69, 0x6e,
	0x6b, 0x12, 0x18, 0x2e, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e,
	0x54, 0x72, 0x6a, 0x78, 0x4d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x1a, 0x18, 0x2e, 0x74, 0x72,
	0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x2e, 0x54, 0x72, 0x6a, 0x78, 0x4d, 0x61,
	0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x22, 0x00, 0x28, 0x01, 0x30, 0x01, 0x42, 0x3f, 0x5a, 0x3d, 0x67,
	0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x74, 0x72, 0x61, 0x6a, 0x65, 0x63,
	0x74, 0x6f, 0x72, 0x79, 0x6a, 0x70, 0x2f, 0x74, 0x72, 0x6a, 0x78, 0x2d, 0x76, 0x65, 0x68, 0x69,
	0x63, 0x6c, 0x65, 0x2d, 0x61, 0x70, 0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x5f, 0x67, 0x6f,
	0x2f, 0x74, 0x72, 0x6a, 0x78, 0x6d, 0x61, 0x76, 0x6c, 0x69, 0x6e, 0x6b, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescOnce sync.Once
	file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescData = file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDesc
)

func file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescGZIP() []byte {
	file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescOnce.Do(func() {
		file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescData = protoimpl.X.CompressGZIP(file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescData)
	})
	return file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDescData
}

var file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_goTypes = []interface{}{
	(Token_Code)(0),                  // 0: trjxmavlink.Token.Code
	(*Token)(nil),                    // 1: trjxmavlink.Token
	(*SetAutopilotModelRequest)(nil), // 2: trjxmavlink.SetAutopilotModelRequest
	(*TrjxMavlink)(nil),              // 3: trjxmavlink.TrjxMavlink
	(*AutopilotModelDef)(nil),        // 4: trjxmavlink.AutopilotModelDef
	(*Empty)(nil),                    // 5: trjxmavlink.Empty
	(*SeriviceAttribute)(nil),        // 6: trjxmavlink.SeriviceAttribute
	(*Result)(nil),                   // 7: trjxmavlink.Result
}
var file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_depIdxs = []int32{
	0, // 0: trjxmavlink.Token.Result:type_name -> trjxmavlink.Token.Code
	4, // 1: trjxmavlink.SetAutopilotModelRequest.Model:type_name -> trjxmavlink.AutopilotModelDef
	5, // 2: trjxmavlink.TrjxMavlinkService.GetSeriviceAttribute:input_type -> trjxmavlink.Empty
	5, // 3: trjxmavlink.TrjxMavlinkService.Login:input_type -> trjxmavlink.Empty
	2, // 4: trjxmavlink.TrjxMavlinkService.SetAutopilotModel:input_type -> trjxmavlink.SetAutopilotModelRequest
	3, // 5: trjxmavlink.TrjxMavlinkService.CommunicateOnMavlink:input_type -> trjxmavlink.TrjxMavlink
	6, // 6: trjxmavlink.TrjxMavlinkService.GetSeriviceAttribute:output_type -> trjxmavlink.SeriviceAttribute
	1, // 7: trjxmavlink.TrjxMavlinkService.Login:output_type -> trjxmavlink.Token
	7, // 8: trjxmavlink.TrjxMavlinkService.SetAutopilotModel:output_type -> trjxmavlink.Result
	3, // 9: trjxmavlink.TrjxMavlinkService.CommunicateOnMavlink:output_type -> trjxmavlink.TrjxMavlink
	6, // [6:10] is the sub-list for method output_type
	2, // [2:6] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() {
	file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_init()
}
func file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_init() {
	if File_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto != nil {
		return
	}
	file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxcommon_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Token); i {
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
		file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SetAutopilotModelRequest); i {
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
		file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TrjxMavlink); i {
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
			RawDescriptor: file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_goTypes,
		DependencyIndexes: file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_depIdxs,
		EnumInfos:         file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_enumTypes,
		MessageInfos:      file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_msgTypes,
	}.Build()
	File_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto = out.File
	file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_rawDesc = nil
	file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_goTypes = nil
	file_github_com_trajectoryjp_trjx_vehicle_api_proto_mav_controller_outside_trjxmav_proto_depIdxs = nil
}
