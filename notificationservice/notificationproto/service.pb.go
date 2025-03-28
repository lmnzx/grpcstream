// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: notificationservice/notificationproto/service.proto

package notificationproto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type NotificationRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *NotificationRequest) Reset() {
	*x = NotificationRequest{}
	mi := &file_notificationservice_notificationproto_service_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *NotificationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*NotificationRequest) ProtoMessage() {}

func (x *NotificationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_notificationproto_service_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use NotificationRequest.ProtoReflect.Descriptor instead.
func (*NotificationRequest) Descriptor() ([]byte, []int) {
	return file_notificationservice_notificationproto_service_proto_rawDescGZIP(), []int{0}
}

func (x *NotificationRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type Notification struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	UserId        string                 `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Content       string                 `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	CreatedAt     int64                  `protobuf:"varint,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Notification) Reset() {
	*x = Notification{}
	mi := &file_notificationservice_notificationproto_service_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Notification) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Notification) ProtoMessage() {}

func (x *Notification) ProtoReflect() protoreflect.Message {
	mi := &file_notificationservice_notificationproto_service_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Notification.ProtoReflect.Descriptor instead.
func (*Notification) Descriptor() ([]byte, []int) {
	return file_notificationservice_notificationproto_service_proto_rawDescGZIP(), []int{1}
}

func (x *Notification) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Notification) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Notification) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

var File_notificationservice_notificationproto_service_proto protoreflect.FileDescriptor

const file_notificationservice_notificationproto_service_proto_rawDesc = "" +
	"\n" +
	"3notificationservice/notificationproto/service.proto\x12\x11notificationproto\".\n" +
	"\x13NotificationRequest\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\"`\n" +
	"\fNotification\x12\x17\n" +
	"\auser_id\x18\x01 \x01(\tR\x06userId\x12\x18\n" +
	"\acontent\x18\x02 \x01(\tR\acontent\x12\x1d\n" +
	"\n" +
	"created_at\x18\x03 \x01(\x03R\tcreatedAt2u\n" +
	"\x13NotificationService\x12^\n" +
	"\x0fGetNotification\x12&.notificationproto.NotificationRequest\x1a\x1f.notificationproto.Notification\"\x000\x01B!Z\x1fgrpcstreaming/notificationprotob\x06proto3"

var (
	file_notificationservice_notificationproto_service_proto_rawDescOnce sync.Once
	file_notificationservice_notificationproto_service_proto_rawDescData []byte
)

func file_notificationservice_notificationproto_service_proto_rawDescGZIP() []byte {
	file_notificationservice_notificationproto_service_proto_rawDescOnce.Do(func() {
		file_notificationservice_notificationproto_service_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_notificationservice_notificationproto_service_proto_rawDesc), len(file_notificationservice_notificationproto_service_proto_rawDesc)))
	})
	return file_notificationservice_notificationproto_service_proto_rawDescData
}

var file_notificationservice_notificationproto_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_notificationservice_notificationproto_service_proto_goTypes = []any{
	(*NotificationRequest)(nil), // 0: notificationproto.NotificationRequest
	(*Notification)(nil),        // 1: notificationproto.Notification
}
var file_notificationservice_notificationproto_service_proto_depIdxs = []int32{
	0, // 0: notificationproto.NotificationService.GetNotification:input_type -> notificationproto.NotificationRequest
	1, // 1: notificationproto.NotificationService.GetNotification:output_type -> notificationproto.Notification
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_notificationservice_notificationproto_service_proto_init() }
func file_notificationservice_notificationproto_service_proto_init() {
	if File_notificationservice_notificationproto_service_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_notificationservice_notificationproto_service_proto_rawDesc), len(file_notificationservice_notificationproto_service_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_notificationservice_notificationproto_service_proto_goTypes,
		DependencyIndexes: file_notificationservice_notificationproto_service_proto_depIdxs,
		MessageInfos:      file_notificationservice_notificationproto_service_proto_msgTypes,
	}.Build()
	File_notificationservice_notificationproto_service_proto = out.File
	file_notificationservice_notificationproto_service_proto_goTypes = nil
	file_notificationservice_notificationproto_service_proto_depIdxs = nil
}
