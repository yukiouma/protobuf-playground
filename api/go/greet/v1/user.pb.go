// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.2
// source: pb/greet/v1/user.proto

package v1

import (
	v1 "github.com/yuki/api/go/user/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Symbols defined in public import of pb/user/v1/user.proto.

type Gender = v1.Gender

const Gender_MISS = v1.Gender_MISS
const Gender_MR = v1.Gender_MR

var Gender_name = v1.Gender_name
var Gender_value = v1.Gender_value

type User = v1.User

var File_pb_greet_v1_user_proto protoreflect.FileDescriptor

var file_pb_greet_v1_user_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x62, 0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73,
	0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x67, 0x72, 0x65, 0x65, 0x74, 0x1a,
	0x15, 0x70, 0x62, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x42, 0x21, 0x5a, 0x1f, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62,
	0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x75, 0x6b, 0x69, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x67, 0x6f,
	0x2f, 0x67, 0x72, 0x65, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x50, 0x00, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var file_pb_greet_v1_user_proto_goTypes = []interface{}{}
var file_pb_greet_v1_user_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_pb_greet_v1_user_proto_init() }
func file_pb_greet_v1_user_proto_init() {
	if File_pb_greet_v1_user_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_pb_greet_v1_user_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_pb_greet_v1_user_proto_goTypes,
		DependencyIndexes: file_pb_greet_v1_user_proto_depIdxs,
	}.Build()
	File_pb_greet_v1_user_proto = out.File
	file_pb_greet_v1_user_proto_rawDesc = nil
	file_pb_greet_v1_user_proto_goTypes = nil
	file_pb_greet_v1_user_proto_depIdxs = nil
}