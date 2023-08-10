// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v4.23.4
// source: user/publish.proto

package publish

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

type CalcWorkCountRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId    int64 `protobuf:"varint,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	IsPublish bool  `protobuf:"varint,2,opt,name=isPublish,proto3" json:"isPublish,omitempty"`
}

func (x *CalcWorkCountRequest) Reset() {
	*x = CalcWorkCountRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_publish_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcWorkCountRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcWorkCountRequest) ProtoMessage() {}

func (x *CalcWorkCountRequest) ProtoReflect() protoreflect.Message {
	mi := &file_user_publish_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcWorkCountRequest.ProtoReflect.Descriptor instead.
func (*CalcWorkCountRequest) Descriptor() ([]byte, []int) {
	return file_user_publish_proto_rawDescGZIP(), []int{0}
}

func (x *CalcWorkCountRequest) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *CalcWorkCountRequest) GetIsPublish() bool {
	if x != nil {
		return x.IsPublish
	}
	return false
}

type CalcWorkCountResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `protobuf:"bytes,2,opt,name=status_msg,json=statusMsg,proto3" json:"status_msg,omitempty"`     // 返回状态描述
}

func (x *CalcWorkCountResponse) Reset() {
	*x = CalcWorkCountResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_user_publish_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CalcWorkCountResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CalcWorkCountResponse) ProtoMessage() {}

func (x *CalcWorkCountResponse) ProtoReflect() protoreflect.Message {
	mi := &file_user_publish_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CalcWorkCountResponse.ProtoReflect.Descriptor instead.
func (*CalcWorkCountResponse) Descriptor() ([]byte, []int) {
	return file_user_publish_proto_rawDescGZIP(), []int{1}
}

func (x *CalcWorkCountResponse) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *CalcWorkCountResponse) GetStatusMsg() string {
	if x != nil {
		return x.StatusMsg
	}
	return ""
}

var File_user_publish_proto protoreflect.FileDescriptor

var file_user_publish_proto_rawDesc = []byte{
	0x0a, 0x12, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x22, 0x4d, 0x0a,
	0x14, 0x43, 0x61, 0x6c, 0x63, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1c,
	0x0a, 0x09, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x08, 0x52, 0x09, 0x69, 0x73, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x22, 0x57, 0x0a, 0x15,
	0x43, 0x61, 0x6c, 0x63, 0x57, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x5f,
	0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x5f, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x4d, 0x73, 0x67, 0x42, 0x1f, 0x5a, 0x1d, 0x74, 0x69, 0x6e, 0x79, 0x74, 0x69, 0x6b,
	0x74, 0x6f, 0x6b, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70,
	0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_user_publish_proto_rawDescOnce sync.Once
	file_user_publish_proto_rawDescData = file_user_publish_proto_rawDesc
)

func file_user_publish_proto_rawDescGZIP() []byte {
	file_user_publish_proto_rawDescOnce.Do(func() {
		file_user_publish_proto_rawDescData = protoimpl.X.CompressGZIP(file_user_publish_proto_rawDescData)
	})
	return file_user_publish_proto_rawDescData
}

var file_user_publish_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_user_publish_proto_goTypes = []interface{}{
	(*CalcWorkCountRequest)(nil),  // 0: publish.CalcWorkCountRequest
	(*CalcWorkCountResponse)(nil), // 1: publish.CalcWorkCountResponse
}
var file_user_publish_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_user_publish_proto_init() }
func file_user_publish_proto_init() {
	if File_user_publish_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_user_publish_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcWorkCountRequest); i {
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
		file_user_publish_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CalcWorkCountResponse); i {
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
			RawDescriptor: file_user_publish_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_user_publish_proto_goTypes,
		DependencyIndexes: file_user_publish_proto_depIdxs,
		MessageInfos:      file_user_publish_proto_msgTypes,
	}.Build()
	File_user_publish_proto = out.File
	file_user_publish_proto_rawDesc = nil
	file_user_publish_proto_goTypes = nil
	file_user_publish_proto_depIdxs = nil
}