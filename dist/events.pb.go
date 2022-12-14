// protoc --go_opt=module=github.com/pete-robinson/setmaker-proto --go_out=. src/events.proto

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.21.9
// source: src/events.proto

package dist

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

type Event_EventType int32

const (
	Event_EVENT_ARTIST_CREATED Event_EventType = 0
	Event_EVENT_ARTIST_DELETED Event_EventType = 1
)

// Enum value maps for Event_EventType.
var (
	Event_EventType_name = map[int32]string{
		0: "EVENT_ARTIST_CREATED",
		1: "EVENT_ARTIST_DELETED",
	}
	Event_EventType_value = map[string]int32{
		"EVENT_ARTIST_CREATED": 0,
		"EVENT_ARTIST_DELETED": 1,
	}
)

func (x Event_EventType) Enum() *Event_EventType {
	p := new(Event_EventType)
	*p = x
	return p
}

func (x Event_EventType) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Event_EventType) Descriptor() protoreflect.EnumDescriptor {
	return file_src_events_proto_enumTypes[0].Descriptor()
}

func (Event_EventType) Type() protoreflect.EnumType {
	return &file_src_events_proto_enumTypes[0]
}

func (x Event_EventType) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Event_EventType.Descriptor instead.
func (Event_EventType) EnumDescriptor() ([]byte, []int) {
	return file_src_events_proto_rawDescGZIP(), []int{0, 0}
}

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	EventType Event_EventType `protobuf:"varint,1,opt,name=eventType,proto3,enum=api.Event_EventType" json:"eventType,omitempty"` // type of event
	// Types that are assignable to MessageBody:
	//	*Event_ArtistCreated
	//	*Event_ArtistDeleted
	MessageBody isEvent_MessageBody `protobuf_oneof:"messageBody"`
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_events_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_src_events_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event.ProtoReflect.Descriptor instead.
func (*Event) Descriptor() ([]byte, []int) {
	return file_src_events_proto_rawDescGZIP(), []int{0}
}

func (x *Event) GetEventType() Event_EventType {
	if x != nil {
		return x.EventType
	}
	return Event_EVENT_ARTIST_CREATED
}

func (m *Event) GetMessageBody() isEvent_MessageBody {
	if m != nil {
		return m.MessageBody
	}
	return nil
}

func (x *Event) GetArtistCreated() *MessageBody_ArtistCreated {
	if x, ok := x.GetMessageBody().(*Event_ArtistCreated); ok {
		return x.ArtistCreated
	}
	return nil
}

func (x *Event) GetArtistDeleted() *MessageBody_ArtistDeleted {
	if x, ok := x.GetMessageBody().(*Event_ArtistDeleted); ok {
		return x.ArtistDeleted
	}
	return nil
}

type isEvent_MessageBody interface {
	isEvent_MessageBody()
}

type Event_ArtistCreated struct {
	ArtistCreated *MessageBody_ArtistCreated `protobuf:"bytes,2,opt,name=artistCreated,proto3,oneof"`
}

type Event_ArtistDeleted struct {
	ArtistDeleted *MessageBody_ArtistDeleted `protobuf:"bytes,3,opt,name=artistDeleted,proto3,oneof"`
}

func (*Event_ArtistCreated) isEvent_MessageBody() {}

func (*Event_ArtistDeleted) isEvent_MessageBody() {}

type MessageBody_ArtistCreated struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *MessageBody_ArtistCreated) Reset() {
	*x = MessageBody_ArtistCreated{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_events_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageBody_ArtistCreated) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBody_ArtistCreated) ProtoMessage() {}

func (x *MessageBody_ArtistCreated) ProtoReflect() protoreflect.Message {
	mi := &file_src_events_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBody_ArtistCreated.ProtoReflect.Descriptor instead.
func (*MessageBody_ArtistCreated) Descriptor() ([]byte, []int) {
	return file_src_events_proto_rawDescGZIP(), []int{1}
}

func (x *MessageBody_ArtistCreated) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *MessageBody_ArtistCreated) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type MessageBody_ArtistDeleted struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *MessageBody_ArtistDeleted) Reset() {
	*x = MessageBody_ArtistDeleted{}
	if protoimpl.UnsafeEnabled {
		mi := &file_src_events_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MessageBody_ArtistDeleted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MessageBody_ArtistDeleted) ProtoMessage() {}

func (x *MessageBody_ArtistDeleted) ProtoReflect() protoreflect.Message {
	mi := &file_src_events_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MessageBody_ArtistDeleted.ProtoReflect.Descriptor instead.
func (*MessageBody_ArtistDeleted) Descriptor() ([]byte, []int) {
	return file_src_events_proto_rawDescGZIP(), []int{2}
}

func (x *MessageBody_ArtistDeleted) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

var File_src_events_proto protoreflect.FileDescriptor

var file_src_events_proto_rawDesc = []byte{
	0x0a, 0x10, 0x73, 0x72, 0x63, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x03, 0x61, 0x70, 0x69, 0x22, 0x9b, 0x02, 0x0a, 0x05, 0x45, 0x76, 0x65, 0x6e,
	0x74, 0x12, 0x32, 0x0a, 0x09, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0e, 0x32, 0x14, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74,
	0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79, 0x70, 0x65, 0x52, 0x09, 0x65, 0x76, 0x65, 0x6e,
	0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x46, 0x0a, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61,
	0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x5f, 0x41,
	0x72, 0x74, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0d,
	0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x12, 0x46, 0x0a,
	0x0d, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x61, 0x70, 0x69, 0x2e, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x6f, 0x64, 0x79, 0x5f, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c,
	0x65, 0x74, 0x65, 0x64, 0x48, 0x00, 0x52, 0x0d, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x44, 0x65,
	0x6c, 0x65, 0x74, 0x65, 0x64, 0x22, 0x3f, 0x0a, 0x09, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x54, 0x79,
	0x70, 0x65, 0x12, 0x18, 0x0a, 0x14, 0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x52, 0x54, 0x49,
	0x53, 0x54, 0x5f, 0x43, 0x52, 0x45, 0x41, 0x54, 0x45, 0x44, 0x10, 0x00, 0x12, 0x18, 0x0a, 0x14,
	0x45, 0x56, 0x45, 0x4e, 0x54, 0x5f, 0x41, 0x52, 0x54, 0x49, 0x53, 0x54, 0x5f, 0x44, 0x45, 0x4c,
	0x45, 0x54, 0x45, 0x44, 0x10, 0x01, 0x42, 0x0d, 0x0a, 0x0b, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x6f, 0x64, 0x79, 0x22, 0x3f, 0x0a, 0x19, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x42, 0x6f, 0x64, 0x79, 0x5f, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2b, 0x0a, 0x19, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x42, 0x6f, 0x64, 0x79, 0x5f, 0x41, 0x72, 0x74, 0x69, 0x73, 0x74, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x69, 0x64, 0x42, 0x2e, 0x5a, 0x2c, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f,
	0x6d, 0x2f, 0x70, 0x65, 0x74, 0x65, 0x2d, 0x72, 0x6f, 0x62, 0x69, 0x6e, 0x73, 0x6f, 0x6e, 0x2f,
	0x73, 0x65, 0x74, 0x6d, 0x61, 0x6b, 0x65, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x64,
	0x69, 0x73, 0x74, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_src_events_proto_rawDescOnce sync.Once
	file_src_events_proto_rawDescData = file_src_events_proto_rawDesc
)

func file_src_events_proto_rawDescGZIP() []byte {
	file_src_events_proto_rawDescOnce.Do(func() {
		file_src_events_proto_rawDescData = protoimpl.X.CompressGZIP(file_src_events_proto_rawDescData)
	})
	return file_src_events_proto_rawDescData
}

var file_src_events_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_src_events_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_src_events_proto_goTypes = []interface{}{
	(Event_EventType)(0),              // 0: api.Event.EventType
	(*Event)(nil),                     // 1: api.Event
	(*MessageBody_ArtistCreated)(nil), // 2: api.MessageBody_ArtistCreated
	(*MessageBody_ArtistDeleted)(nil), // 3: api.MessageBody_ArtistDeleted
}
var file_src_events_proto_depIdxs = []int32{
	0, // 0: api.Event.eventType:type_name -> api.Event.EventType
	2, // 1: api.Event.artistCreated:type_name -> api.MessageBody_ArtistCreated
	3, // 2: api.Event.artistDeleted:type_name -> api.MessageBody_ArtistDeleted
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_src_events_proto_init() }
func file_src_events_proto_init() {
	if File_src_events_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_src_events_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event); i {
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
		file_src_events_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageBody_ArtistCreated); i {
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
		file_src_events_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MessageBody_ArtistDeleted); i {
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
	file_src_events_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*Event_ArtistCreated)(nil),
		(*Event_ArtistDeleted)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_src_events_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_src_events_proto_goTypes,
		DependencyIndexes: file_src_events_proto_depIdxs,
		EnumInfos:         file_src_events_proto_enumTypes,
		MessageInfos:      file_src_events_proto_msgTypes,
	}.Build()
	File_src_events_proto = out.File
	file_src_events_proto_rawDesc = nil
	file_src_events_proto_goTypes = nil
	file_src_events_proto_depIdxs = nil
}
