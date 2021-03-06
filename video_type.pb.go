// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.14.0
// source: video_type.proto

package live

import (
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

// 视频类型
type VideoType int32

const (
	VideoType_AUDIO    VideoType = 0
	VideoType_P360     VideoType = 1
	VideoType_P480     VideoType = 2
	VideoType_P720     VideoType = 3
	VideoType_P1080    VideoType = 4
	VideoType_ORIGINAL VideoType = 5
)

// Enum value maps for VideoType.
var (
	VideoType_name = map[int32]string{
		0: "AUDIO",
		1: "P360",
		2: "P480",
		3: "P720",
		4: "P1080",
		5: "ORIGINAL",
	}
	VideoType_value = map[string]int32{
		"AUDIO":    0,
		"P360":     1,
		"P480":     2,
		"P720":     3,
		"P1080":    4,
		"ORIGINAL": 5,
	}
)

func (x VideoType) Enum() *VideoType {
	p := new(VideoType)
	*p = x
	return p
}

func (x VideoType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (VideoType) Descriptor() protoreflect.EnumDescriptor {
	return file_video_type_proto_enumTypes[0].Descriptor()
}

func (VideoType) Type() protoreflect.EnumType {
	return &file_video_type_proto_enumTypes[0]
}

func (x VideoType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use VideoType.Descriptor instead.
func (VideoType) EnumDescriptor() ([]byte, []int) {
	return file_video_type_proto_rawDescGZIP(), []int{0}
}

var File_video_type_proto protoreflect.FileDescriptor

var file_video_type_proto_rawDesc = []byte{
	0x0a, 0x10, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x04, 0x6c, 0x69, 0x76, 0x65, 0x2a, 0x4d, 0x0a, 0x09, 0x56, 0x69, 0x64, 0x65,
	0x6f, 0x54, 0x79, 0x70, 0x65, 0x12, 0x09, 0x0a, 0x05, 0x41, 0x55, 0x44, 0x49, 0x4f, 0x10, 0x00,
	0x12, 0x08, 0x0a, 0x04, 0x50, 0x33, 0x36, 0x30, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x34,
	0x38, 0x30, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x37, 0x32, 0x30, 0x10, 0x03, 0x12, 0x09,
	0x0a, 0x05, 0x50, 0x31, 0x30, 0x38, 0x30, 0x10, 0x04, 0x12, 0x0c, 0x0a, 0x08, 0x4f, 0x52, 0x49,
	0x47, 0x49, 0x4e, 0x41, 0x4c, 0x10, 0x05, 0x42, 0x08, 0x5a, 0x06, 0x2e, 0x3b, 0x6c, 0x69, 0x76,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_video_type_proto_rawDescOnce sync.Once
	file_video_type_proto_rawDescData = file_video_type_proto_rawDesc
)

func file_video_type_proto_rawDescGZIP() []byte {
	file_video_type_proto_rawDescOnce.Do(func() {
		file_video_type_proto_rawDescData = protoimpl.X.CompressGZIP(file_video_type_proto_rawDescData)
	})
	return file_video_type_proto_rawDescData
}

var file_video_type_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_video_type_proto_goTypes = []interface{}{
	(VideoType)(0), // 0: live.VideoType
}
var file_video_type_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_video_type_proto_init() }
func file_video_type_proto_init() {
	if File_video_type_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_video_type_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_video_type_proto_goTypes,
		DependencyIndexes: file_video_type_proto_depIdxs,
		EnumInfos:         file_video_type_proto_enumTypes,
	}.Build()
	File_video_type_proto = out.File
	file_video_type_proto_rawDesc = nil
	file_video_type_proto_goTypes = nil
	file_video_type_proto_depIdxs = nil
}
