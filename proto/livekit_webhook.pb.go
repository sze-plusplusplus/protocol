// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.18.0
// source: livekit_webhook.proto

package livekit

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

type WebhookEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// one of room_started, room_finished, participant_joined, participant_left, recording_finished
	Event string `protobuf:"bytes,1,opt,name=event,proto3" json:"event,omitempty"`
	Room  *Room  `protobuf:"bytes,2,opt,name=room,proto3" json:"room,omitempty"`
	// set when event is participant_*
	Participant *ParticipantInfo `protobuf:"bytes,3,opt,name=participant,proto3" json:"participant,omitempty"`
	// set when event is recording_finished
	RecordingResult *RecordingResult `protobuf:"bytes,4,opt,name=recording_result,json=recordingResult,proto3" json:"recording_result,omitempty"`
}

func (x *WebhookEvent) Reset() {
	*x = WebhookEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_livekit_webhook_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *WebhookEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WebhookEvent) ProtoMessage() {}

func (x *WebhookEvent) ProtoReflect() protoreflect.Message {
	mi := &file_livekit_webhook_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WebhookEvent.ProtoReflect.Descriptor instead.
func (*WebhookEvent) Descriptor() ([]byte, []int) {
	return file_livekit_webhook_proto_rawDescGZIP(), []int{0}
}

func (x *WebhookEvent) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *WebhookEvent) GetRoom() *Room {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *WebhookEvent) GetParticipant() *ParticipantInfo {
	if x != nil {
		return x.Participant
	}
	return nil
}

func (x *WebhookEvent) GetRecordingResult() *RecordingResult {
	if x != nil {
		return x.RecordingResult
	}
	return nil
}

var File_livekit_webhook_proto protoreflect.FileDescriptor

var file_livekit_webhook_proto_rawDesc = []byte{
	0x0a, 0x15, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x77, 0x65, 0x62, 0x68, 0x6f, 0x6f,
	0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x1a, 0x14, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x73,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xc8, 0x01, 0x0a, 0x0c, 0x57, 0x65, 0x62, 0x68, 0x6f,
	0x6f, 0x6b, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x21, 0x0a,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6c, 0x69,
	0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x52, 0x6f, 0x6f, 0x6d, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x12, 0x3a, 0x0a, 0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e,
	0x50, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x49, 0x6e, 0x66, 0x6f, 0x52,
	0x0b, 0x70, 0x61, 0x72, 0x74, 0x69, 0x63, 0x69, 0x70, 0x61, 0x6e, 0x74, 0x12, 0x43, 0x0a, 0x10,
	0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x5f, 0x72, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74,
	0x2e, 0x52, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x52, 0x0f, 0x72, 0x65, 0x63, 0x6f, 0x72, 0x64, 0x69, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x42, 0x42, 0x5a, 0x29, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x6c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0xaa, 0x02,
	0x14, 0x4c, 0x69, 0x76, 0x65, 0x6b, 0x69, 0x74, 0x2e, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x2e,
	0x50, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_livekit_webhook_proto_rawDescOnce sync.Once
	file_livekit_webhook_proto_rawDescData = file_livekit_webhook_proto_rawDesc
)

func file_livekit_webhook_proto_rawDescGZIP() []byte {
	file_livekit_webhook_proto_rawDescOnce.Do(func() {
		file_livekit_webhook_proto_rawDescData = protoimpl.X.CompressGZIP(file_livekit_webhook_proto_rawDescData)
	})
	return file_livekit_webhook_proto_rawDescData
}

var file_livekit_webhook_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_livekit_webhook_proto_goTypes = []interface{}{
	(*WebhookEvent)(nil),    // 0: livekit.WebhookEvent
	(*Room)(nil),            // 1: livekit.Room
	(*ParticipantInfo)(nil), // 2: livekit.ParticipantInfo
	(*RecordingResult)(nil), // 3: livekit.RecordingResult
}
var file_livekit_webhook_proto_depIdxs = []int32{
	1, // 0: livekit.WebhookEvent.room:type_name -> livekit.Room
	2, // 1: livekit.WebhookEvent.participant:type_name -> livekit.ParticipantInfo
	3, // 2: livekit.WebhookEvent.recording_result:type_name -> livekit.RecordingResult
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_livekit_webhook_proto_init() }
func file_livekit_webhook_proto_init() {
	if File_livekit_webhook_proto != nil {
		return
	}
	file_livekit_models_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_livekit_webhook_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*WebhookEvent); i {
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
			RawDescriptor: file_livekit_webhook_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_livekit_webhook_proto_goTypes,
		DependencyIndexes: file_livekit_webhook_proto_depIdxs,
		MessageInfos:      file_livekit_webhook_proto_msgTypes,
	}.Build()
	File_livekit_webhook_proto = out.File
	file_livekit_webhook_proto_rawDesc = nil
	file_livekit_webhook_proto_goTypes = nil
	file_livekit_webhook_proto_depIdxs = nil
}
