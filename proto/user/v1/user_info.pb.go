// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: proto/user/v1/user_info.proto

package userv1

import (
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2/options"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's id
type SEXUAL int32

const (
	SEXUAL_SEXUAL_UNSPECIFIED SEXUAL = 0
	SEXUAL_SEXUAL_MALE        SEXUAL = 1
	SEXUAL_SEXUAL_FEMALE      SEXUAL = 2
)

// Enum value maps for SEXUAL.
var (
	SEXUAL_name = map[int32]string{
		0: "SEXUAL_UNSPECIFIED",
		1: "SEXUAL_MALE",
		2: "SEXUAL_FEMALE",
	}
	SEXUAL_value = map[string]int32{
		"SEXUAL_UNSPECIFIED": 0,
		"SEXUAL_MALE":        1,
		"SEXUAL_FEMALE":      2,
	}
)

func (x SEXUAL) Enum() *SEXUAL {
	p := new(SEXUAL)
	*p = x
	return p
}

func (x SEXUAL) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (SEXUAL) Descriptor() protoreflect.EnumDescriptor {
	return file_proto_user_v1_user_info_proto_enumTypes[0].Descriptor()
}

func (SEXUAL) Type() protoreflect.EnumType {
	return &file_proto_user_v1_user_info_proto_enumTypes[0]
}

func (x SEXUAL) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use SEXUAL.Descriptor instead.
func (SEXUAL) EnumDescriptor() ([]byte, []int) {
	return file_proto_user_v1_user_info_proto_rawDescGZIP(), []int{0}
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string            `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name   string            `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Age    int32             `protobuf:"varint,3,opt,name=age,proto3" json:"age,omitempty"`
	Email  string            `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Sexual SEXUAL            `protobuf:"varint,5,opt,name=sexual,proto3,enum=user.v1.SEXUAL" json:"sexual,omitempty"`
	Tasks  map[string]string `protobuf:"bytes,6,rep,name=tasks,proto3" json:"tasks,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_user_v1_user_info_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_proto_user_v1_user_info_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_proto_user_v1_user_info_proto_rawDescGZIP(), []int{0}
}

func (x *UserInfo) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *UserInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *UserInfo) GetAge() int32 {
	if x != nil {
		return x.Age
	}
	return 0
}

func (x *UserInfo) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

func (x *UserInfo) GetSexual() SEXUAL {
	if x != nil {
		return x.Sexual
	}
	return SEXUAL_SEXUAL_UNSPECIFIED
}

func (x *UserInfo) GetTasks() map[string]string {
	if x != nil {
		return x.Tasks
	}
	return nil
}

var File_proto_user_v1_user_info_proto protoreflect.FileDescriptor

var file_proto_user_v1_user_info_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f,
	0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65,
	0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x2d,
	0x67, 0x65, 0x6e, 0x2d, 0x6f, 0x70, 0x65, 0x6e, 0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x6f, 0x70,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xed, 0x01, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72,
	0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x03, 0x61, 0x67, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c,
	0x12, 0x27, 0x0a, 0x06, 0x73, 0x65, 0x78, 0x75, 0x61, 0x6c, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x0f, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x45, 0x58, 0x55, 0x41,
	0x4c, 0x52, 0x06, 0x73, 0x65, 0x78, 0x75, 0x61, 0x6c, 0x12, 0x32, 0x0a, 0x05, 0x74, 0x61, 0x73,
	0x6b, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e,
	0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x2e, 0x54, 0x61, 0x73, 0x6b,
	0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x05, 0x74, 0x61, 0x73, 0x6b, 0x73, 0x1a, 0x38, 0x0a,
	0x0a, 0x54, 0x61, 0x73, 0x6b, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b,
	0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61,
	0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x2a, 0x44, 0x0a, 0x06, 0x53, 0x45, 0x58, 0x55, 0x41,
	0x4c, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x45, 0x58, 0x55, 0x41, 0x4c, 0x5f, 0x55, 0x4e, 0x53, 0x50,
	0x45, 0x43, 0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0f, 0x0a, 0x0b, 0x53, 0x45, 0x58,
	0x55, 0x41, 0x4c, 0x5f, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x45,
	0x58, 0x55, 0x41, 0x4c, 0x5f, 0x46, 0x45, 0x4d, 0x41, 0x4c, 0x45, 0x10, 0x02, 0x42, 0xb7, 0x01,
	0x0a, 0x0b, 0x63, 0x6f, 0x6d, 0x2e, 0x75, 0x73, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x42, 0x0d, 0x55,
	0x73, 0x65, 0x72, 0x49, 0x6e, 0x66, 0x6f, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x36,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x61, 0x64, 0x6f, 0x6e, 0x39,
	0x38, 0x38, 0x2f, 0x67, 0x6f, 0x2d, 0x61, 0x72, 0x74, 0x69, 0x63, 0x6c, 0x65, 0x2d, 0x61, 0x70,
	0x69, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x3b,
	0x75, 0x73, 0x65, 0x72, 0x76, 0x31, 0xa2, 0x02, 0x03, 0x55, 0x58, 0x58, 0xaa, 0x02, 0x07, 0x55,
	0x73, 0x65, 0x72, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x07, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31,
	0xe2, 0x02, 0x13, 0x55, 0x73, 0x65, 0x72, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50, 0x42, 0x4d, 0x65,
	0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x08, 0x55, 0x73, 0x65, 0x72, 0x3a, 0x3a, 0x56,
	0x31, 0x92, 0x41, 0x23, 0x12, 0x05, 0x32, 0x03, 0x31, 0x2e, 0x30, 0x2a, 0x01, 0x01, 0x72, 0x17,
	0x12, 0x15, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x68, 0x6f,
	0x73, 0x74, 0x3a, 0x38, 0x30, 0x38, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_user_v1_user_info_proto_rawDescOnce sync.Once
	file_proto_user_v1_user_info_proto_rawDescData = file_proto_user_v1_user_info_proto_rawDesc
)

func file_proto_user_v1_user_info_proto_rawDescGZIP() []byte {
	file_proto_user_v1_user_info_proto_rawDescOnce.Do(func() {
		file_proto_user_v1_user_info_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_user_v1_user_info_proto_rawDescData)
	})
	return file_proto_user_v1_user_info_proto_rawDescData
}

var file_proto_user_v1_user_info_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_proto_user_v1_user_info_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_user_v1_user_info_proto_goTypes = []interface{}{
	(SEXUAL)(0),      // 0: user.v1.SEXUAL
	(*UserInfo)(nil), // 1: user.v1.UserInfo
	nil,              // 2: user.v1.UserInfo.TasksEntry
}
var file_proto_user_v1_user_info_proto_depIdxs = []int32{
	0, // 0: user.v1.UserInfo.sexual:type_name -> user.v1.SEXUAL
	2, // 1: user.v1.UserInfo.tasks:type_name -> user.v1.UserInfo.TasksEntry
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_user_v1_user_info_proto_init() }
func file_proto_user_v1_user_info_proto_init() {
	if File_proto_user_v1_user_info_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_user_v1_user_info_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
			RawDescriptor: file_proto_user_v1_user_info_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_proto_user_v1_user_info_proto_goTypes,
		DependencyIndexes: file_proto_user_v1_user_info_proto_depIdxs,
		EnumInfos:         file_proto_user_v1_user_info_proto_enumTypes,
		MessageInfos:      file_proto_user_v1_user_info_proto_msgTypes,
	}.Build()
	File_proto_user_v1_user_info_proto = out.File
	file_proto_user_v1_user_info_proto_rawDesc = nil
	file_proto_user_v1_user_info_proto_goTypes = nil
	file_proto_user_v1_user_info_proto_depIdxs = nil
}
