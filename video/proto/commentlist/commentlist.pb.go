// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: video/commentlist.proto

package commentlist

import (
	comment "tinytiktok/video/proto/comment"
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

type CommentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                     // 用户鉴权token
	VideoId int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` // 视频id
}

func (x *CommentListRequest) Reset() {
	*x = CommentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_commentlist_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListRequest) ProtoMessage() {}

func (x *CommentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_video_commentlist_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListRequest.ProtoReflect.Descriptor instead.
func (*CommentListRequest) Descriptor() ([]byte, []int) {
	return file_video_commentlist_proto_rawDescGZIP(), []int{0}
}

func (x *CommentListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommentListRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type CommentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode  int32              `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`   // 状态码，0-成功，其他值-失败
	StatusMsg   string             `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`       // 返回状态描述
	CommentList []*comment.Comment `protobuf:"bytes,3,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"` // 视频评论列表
}

func (x *CommentListResponse) Reset() {
	*x = CommentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_video_commentlist_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResponse) ProtoMessage() {}

func (x *CommentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_video_commentlist_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResponse.ProtoReflect.Descriptor instead.
func (*CommentListResponse) Descriptor() ([]byte, []int) {
	return file_video_commentlist_proto_rawDescGZIP(), []int{1}
}

func (x *CommentListResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CommentListResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

func (x *CommentListResponse) GetCommentList() []*comment.Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

var File_video_commentlist_proto protoreflect.FileDescriptor

var file_video_commentlist_proto_rawDesc = []byte{
	0x0a, 0x17, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x6c,
	0x69, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x6c, 0x69, 0x73, 0x74, 0x1a, 0x13, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x2f, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x45, 0x0a, 0x12, 0x43,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f,
	0x49, 0x64, 0x22, 0x8a, 0x01, 0x0a, 0x13, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69,
	0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x73, 0x67, 0x12, 0x33, 0x0a, 0x0c, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x42,
	0x32, 0x5a, 0x30, 0x2f, 0x67, 0x6f, 0x5f, 0x77, 0x6f, 0x72, 0x6b, 0x73, 0x70, 0x61, 0x63, 0x65,
	0x2f, 0x74, 0x69, 0x6e, 0x79, 0x74, 0x69, 0x6b, 0x74, 0x6f, 0x6b, 0x2f, 0x76, 0x69, 0x64, 0x65,
	0x6f, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x6c,
	0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_commentlist_proto_rawDescOnce sync.Once
	file_video_commentlist_proto_rawDescData = file_video_commentlist_proto_rawDesc
)

func file_video_commentlist_proto_rawDescGZIP() []byte {
	file_video_commentlist_proto_rawDescOnce.Do(func() {
		file_video_commentlist_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_commentlist_proto_rawDescData)
	})
	return file_video_commentlist_proto_rawDescData
}

var file_video_commentlist_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_video_commentlist_proto_goTypes = []interface{}{
	(*CommentListRequest)(nil),  // 0: commentlist.CommentListRequest
	(*CommentListResponse)(nil), // 1: commentlist.CommentListResponse
	(*comment.Comment)(nil),     // 2: comment.Comment
}
var file_video_commentlist_proto_depIdxs = []int32{
	2, // 0: commentlist.CommentListResponse.comment_list:type_name -> comment.Comment
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_video_commentlist_proto_init() }
func file_video_commentlist_proto_init() {
	if File_video_commentlist_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_video_commentlist_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListRequest); i {
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
		file_video_commentlist_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResponse); i {
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
			RawDescriptor: file_video_commentlist_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_video_commentlist_proto_goTypes,
		DependencyIndexes: file_video_commentlist_proto_depIdxs,
		MessageInfos:      file_video_commentlist_proto_msgTypes,
	}.Build()
	File_video_commentlist_proto = out.File
	file_video_commentlist_proto_rawDesc = nil
	file_video_commentlist_proto_goTypes = nil
	file_video_commentlist_proto_depIdxs = nil
}
