// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: video/server.proto

package server

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	video "tinytiktok/video/proto/video"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_video_server_proto protoreflect.FileDescriptor

var file_video_server_proto_rawDesc = []byte{
	0x0a, 0x12, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x75, 0x73, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72,
	0x1a, 0x11, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x32, 0x41, 0x0a, 0x0c, 0x56, 0x69, 0x64, 0x65, 0x6f, 0x53, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x12, 0x31, 0x0a, 0x04, 0x46, 0x65, 0x65, 0x64, 0x12, 0x12, 0x2e, 0x76, 0x69,
	0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a,
	0x13, 0x2e, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2e, 0x46, 0x65, 0x65, 0x64, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1f, 0x5a, 0x1d, 0x74, 0x69, 0x6e, 0x79, 0x74, 0x69,
	0x6b, 0x74, 0x6f, 0x6b, 0x2f, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x73, 0x65, 0x72, 0x76, 0x65, 0x72, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_video_server_proto_goTypes = []interface{}{
	(*video.FeedRequest)(nil),  // 0: video.FeedRequest
	(*video.FeedResponse)(nil), // 1: video.FeedResponse
}
var file_video_server_proto_depIdxs = []int32{
	0, // 0: userServer.VideoService.Feed:input_type -> video.FeedRequest
	1, // 1: userServer.VideoService.Feed:output_type -> video.FeedResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_server_proto_init() }
func file_video_server_proto_init() {
	if File_video_server_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_server_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_video_server_proto_goTypes,
		DependencyIndexes: file_video_server_proto_depIdxs,
	}.Build()
	File_video_server_proto = out.File
	file_video_server_proto_rawDesc = nil
	file_video_server_proto_goTypes = nil
	file_video_server_proto_depIdxs = nil
}
