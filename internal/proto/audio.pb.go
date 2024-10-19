// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.0
// source: internal/proto/audio.proto

package proto

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

type AudioRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId string `protobuf:"bytes,1,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
}

func (x *AudioRequest) Reset() {
	*x = AudioRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_audio_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioRequest) ProtoMessage() {}

func (x *AudioRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_audio_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioRequest.ProtoReflect.Descriptor instead.
func (*AudioRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_audio_proto_rawDescGZIP(), []int{0}
}

func (x *AudioRequest) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

type AudioResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AudioData []byte `protobuf:"bytes,1,opt,name=audio_data,json=audioData,proto3" json:"audio_data,omitempty"`
	FileName  string `protobuf:"bytes,2,opt,name=file_name,json=fileName,proto3" json:"file_name,omitempty"`
}

func (x *AudioResponse) Reset() {
	*x = AudioResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_audio_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AudioResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AudioResponse) ProtoMessage() {}

func (x *AudioResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_audio_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AudioResponse.ProtoReflect.Descriptor instead.
func (*AudioResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_audio_proto_rawDescGZIP(), []int{1}
}

func (x *AudioResponse) GetAudioData() []byte {
	if x != nil {
		return x.AudioData
	}
	return nil
}

func (x *AudioResponse) GetFileName() string {
	if x != nil {
		return x.FileName
	}
	return ""
}

type NewsContentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NewsId  string `protobuf:"bytes,1,opt,name=news_id,json=newsId,proto3" json:"news_id,omitempty"`
	Content string `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
}

func (x *NewsContentRequest) Reset() {
	*x = NewsContentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_audio_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsContentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsContentRequest) ProtoMessage() {}

func (x *NewsContentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_audio_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsContentRequest.ProtoReflect.Descriptor instead.
func (*NewsContentRequest) Descriptor() ([]byte, []int) {
	return file_internal_proto_audio_proto_rawDescGZIP(), []int{2}
}

func (x *NewsContentRequest) GetNewsId() string {
	if x != nil {
		return x.NewsId
	}
	return ""
}

func (x *NewsContentRequest) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

type NewsContentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Success bool   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *NewsContentResponse) Reset() {
	*x = NewsContentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_proto_audio_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *NewsContentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NewsContentResponse) ProtoMessage() {}

func (x *NewsContentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_proto_audio_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NewsContentResponse.ProtoReflect.Descriptor instead.
func (*NewsContentResponse) Descriptor() ([]byte, []int) {
	return file_internal_proto_audio_proto_rawDescGZIP(), []int{3}
}

func (x *NewsContentResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *NewsContentResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

var File_internal_proto_audio_proto protoreflect.FileDescriptor

var file_internal_proto_audio_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x6e, 0x65,
	0x77, 0x73, 0x22, 0x27, 0x0a, 0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x22, 0x4b, 0x0a, 0x0d, 0x41,
	0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1d, 0x0a, 0x0a,
	0x61, 0x75, 0x64, 0x69, 0x6f, 0x5f, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x09, 0x61, 0x75, 0x64, 0x69, 0x6f, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1b, 0x0a, 0x09, 0x66,
	0x69, 0x6c, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08,
	0x66, 0x69, 0x6c, 0x65, 0x4e, 0x61, 0x6d, 0x65, 0x22, 0x47, 0x0a, 0x12, 0x4e, 0x65, 0x77, 0x73,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17,
	0x0a, 0x07, 0x6e, 0x65, 0x77, 0x73, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x6e, 0x65, 0x77, 0x73, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x22, 0x49, 0x0a, 0x13, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x73, 0x75, 0x63, 0x63,
	0x65, 0x73, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x73, 0x75, 0x63, 0x63, 0x65,
	0x73, 0x73, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x32, 0x96, 0x01, 0x0a,
	0x0c, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x39, 0x0a,
	0x0c, 0x47, 0x65, 0x74, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x46, 0x69, 0x6c, 0x65, 0x12, 0x12, 0x2e,
	0x6e, 0x65, 0x77, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x13, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x41, 0x75, 0x64, 0x69, 0x6f, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4b, 0x0a, 0x12, 0x52, 0x65, 0x63, 0x65,
	0x69, 0x76, 0x65, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12, 0x18,
	0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e, 0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x19, 0x2e, 0x6e, 0x65, 0x77, 0x73, 0x2e,
	0x4e, 0x65, 0x77, 0x73, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x30, 0x5a, 0x2e, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e,
	0x63, 0x6f, 0x6d, 0x2f, 0x79, 0x6f, 0x75, 0x72, 0x75, 0x73, 0x65, 0x72, 0x6e, 0x61, 0x6d, 0x65,
	0x2f, 0x67, 0x61, 0x74, 0x65, 0x77, 0x61, 0x79, 0x2f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_proto_audio_proto_rawDescOnce sync.Once
	file_internal_proto_audio_proto_rawDescData = file_internal_proto_audio_proto_rawDesc
)

func file_internal_proto_audio_proto_rawDescGZIP() []byte {
	file_internal_proto_audio_proto_rawDescOnce.Do(func() {
		file_internal_proto_audio_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_proto_audio_proto_rawDescData)
	})
	return file_internal_proto_audio_proto_rawDescData
}

var file_internal_proto_audio_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_proto_audio_proto_goTypes = []any{
	(*AudioRequest)(nil),        // 0: news.AudioRequest
	(*AudioResponse)(nil),       // 1: news.AudioResponse
	(*NewsContentRequest)(nil),  // 2: news.NewsContentRequest
	(*NewsContentResponse)(nil), // 3: news.NewsContentResponse
}
var file_internal_proto_audio_proto_depIdxs = []int32{
	0, // 0: news.AudioService.GetAudioFile:input_type -> news.AudioRequest
	2, // 1: news.AudioService.ReceiveNewsContent:input_type -> news.NewsContentRequest
	1, // 2: news.AudioService.GetAudioFile:output_type -> news.AudioResponse
	3, // 3: news.AudioService.ReceiveNewsContent:output_type -> news.NewsContentResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_proto_audio_proto_init() }
func file_internal_proto_audio_proto_init() {
	if File_internal_proto_audio_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_proto_audio_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*AudioRequest); i {
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
		file_internal_proto_audio_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*AudioResponse); i {
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
		file_internal_proto_audio_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*NewsContentRequest); i {
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
		file_internal_proto_audio_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*NewsContentResponse); i {
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
			RawDescriptor: file_internal_proto_audio_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_proto_audio_proto_goTypes,
		DependencyIndexes: file_internal_proto_audio_proto_depIdxs,
		MessageInfos:      file_internal_proto_audio_proto_msgTypes,
	}.Build()
	File_internal_proto_audio_proto = out.File
	file_internal_proto_audio_proto_rawDesc = nil
	file_internal_proto_audio_proto_goTypes = nil
	file_internal_proto_audio_proto_depIdxs = nil
}
