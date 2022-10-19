// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.5
// source: internal/pb/analytics.proto

package pb

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

type Event struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Event) Reset() {
	*x = Event{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_analytics_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event) ProtoMessage() {}

func (x *Event) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_analytics_proto_msgTypes[0]
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
	return file_internal_pb_analytics_proto_rawDescGZIP(), []int{0}
}

type SyncSummary struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Invocation_UUID string         `protobuf:"bytes,1,opt,name=Invocation_UUID,json=InvocationUUID,proto3" json:"Invocation_UUID,omitempty"` // unique ID for each CLI invocation (This could be the same for some rows if multiple sources or destinations are being synced)
	Message_UUID    string         `protobuf:"bytes,2,opt,name=Message_UUID,json=MessageUUID,proto3" json:"Message_UUID,omitempty"`          // server-side UUID to uniquely identify every message and do at-least-once write
	Timestamp       int64          `protobuf:"varint,3,opt,name=Timestamp,proto3" json:"Timestamp,omitempty"`                                // server-side timestamp for when message was received
	SourcePath      string         `protobuf:"bytes,4,opt,name=Source_path,json=SourcePath,proto3" json:"Source_path,omitempty"`             // path of source plugin; registry is assumed to be Github
	SourceVersion   string         `protobuf:"bytes,5,opt,name=Source_version,json=SourceVersion,proto3" json:"Source_version,omitempty"`    // version of source plugin used
	Destinations    []*Destination `protobuf:"bytes,6,rep,name=Destinations,proto3" json:"Destinations,omitempty"`
	Resources       int64          `protobuf:"varint,8,opt,name=Resources,proto3" json:"Resources,omitempty"`                              // number of resources fetched
	Errors          int64          `protobuf:"varint,9,opt,name=Errors,proto3" json:"Errors,omitempty"`                                    // number of errors that occurred
	Panics          int64          `protobuf:"varint,10,opt,name=Panics,proto3" json:"Panics,omitempty"`                                   // number of panics that occurred
	ClientVersion   string         `protobuf:"bytes,11,opt,name=Client_version,json=ClientVersion,proto3" json:"Client_version,omitempty"` // client (CLI) version used
}

func (x *SyncSummary) Reset() {
	*x = SyncSummary{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_analytics_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SyncSummary) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SyncSummary) ProtoMessage() {}

func (x *SyncSummary) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_analytics_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SyncSummary.ProtoReflect.Descriptor instead.
func (*SyncSummary) Descriptor() ([]byte, []int) {
	return file_internal_pb_analytics_proto_rawDescGZIP(), []int{1}
}

func (x *SyncSummary) GetInvocation_UUID() string {
	if x != nil {
		return x.Invocation_UUID
	}
	return ""
}

func (x *SyncSummary) GetMessage_UUID() string {
	if x != nil {
		return x.Message_UUID
	}
	return ""
}

func (x *SyncSummary) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *SyncSummary) GetSourcePath() string {
	if x != nil {
		return x.SourcePath
	}
	return ""
}

func (x *SyncSummary) GetSourceVersion() string {
	if x != nil {
		return x.SourceVersion
	}
	return ""
}

func (x *SyncSummary) GetDestinations() []*Destination {
	if x != nil {
		return x.Destinations
	}
	return nil
}

func (x *SyncSummary) GetResources() int64 {
	if x != nil {
		return x.Resources
	}
	return 0
}

func (x *SyncSummary) GetErrors() int64 {
	if x != nil {
		return x.Errors
	}
	return 0
}

func (x *SyncSummary) GetPanics() int64 {
	if x != nil {
		return x.Panics
	}
	return 0
}

func (x *SyncSummary) GetClientVersion() string {
	if x != nil {
		return x.ClientVersion
	}
	return ""
}

type Destination struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Path    string `protobuf:"bytes,1,opt,name=Path,proto3" json:"Path,omitempty"`       // path of destination plugin; registry is assumed to be Github
	Version string `protobuf:"bytes,2,opt,name=Version,proto3" json:"Version,omitempty"` // version of destination plugin used
}

func (x *Destination) Reset() {
	*x = Destination{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_analytics_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Destination) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Destination) ProtoMessage() {}

func (x *Destination) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_analytics_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Destination.ProtoReflect.Descriptor instead.
func (*Destination) Descriptor() ([]byte, []int) {
	return file_internal_pb_analytics_proto_rawDescGZIP(), []int{2}
}

func (x *Destination) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *Destination) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

type Event_Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	SyncSummary *SyncSummary `protobuf:"bytes,1,opt,name=sync_summary,json=syncSummary,proto3" json:"sync_summary,omitempty"`
}

func (x *Event_Request) Reset() {
	*x = Event_Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_analytics_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_Request) ProtoMessage() {}

func (x *Event_Request) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_analytics_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_Request.ProtoReflect.Descriptor instead.
func (*Event_Request) Descriptor() ([]byte, []int) {
	return file_internal_pb_analytics_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Event_Request) GetSyncSummary() *SyncSummary {
	if x != nil {
		return x.SyncSummary
	}
	return nil
}

type Event_Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Event_Response) Reset() {
	*x = Event_Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_pb_analytics_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Event_Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Event_Response) ProtoMessage() {}

func (x *Event_Response) ProtoReflect() protoreflect.Message {
	mi := &file_internal_pb_analytics_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Event_Response.ProtoReflect.Descriptor instead.
func (*Event_Response) Descriptor() ([]byte, []int) {
	return file_internal_pb_analytics_proto_rawDescGZIP(), []int{0, 1}
}

var File_internal_pb_analytics_proto protoreflect.FileDescriptor

var file_internal_pb_analytics_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x70, 0x62, 0x2f, 0x61, 0x6e,
	0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1c, 0x63,
	0x6c, 0x6f, 0x75, 0x64, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e,
	0x64, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x22, 0x6c, 0x0a, 0x05, 0x45,
	0x76, 0x65, 0x6e, 0x74, 0x1a, 0x57, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x4c, 0x0a, 0x0c, 0x73, 0x79, 0x6e, 0x63, 0x5f, 0x73, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x71, 0x75, 0x65,
	0x72, 0x79, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79,
	0x74, 0x69, 0x63, 0x73, 0x2e, 0x53, 0x79, 0x6e, 0x63, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79,
	0x52, 0x0b, 0x73, 0x79, 0x6e, 0x63, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x1a, 0x0a, 0x0a,
	0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x83, 0x03, 0x0a, 0x0b, 0x53, 0x79,
	0x6e, 0x63, 0x53, 0x75, 0x6d, 0x6d, 0x61, 0x72, 0x79, 0x12, 0x27, 0x0a, 0x0f, 0x49, 0x6e, 0x76,
	0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x55, 0x55, 0x49, 0x44, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0e, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x55, 0x55,
	0x49, 0x44, 0x12, 0x21, 0x0a, 0x0c, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x5f, 0x55, 0x55,
	0x49, 0x44, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x55, 0x55, 0x49, 0x44, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74,
	0x61, 0x6d, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x70, 0x61,
	0x74, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x50, 0x61, 0x74, 0x68, 0x12, 0x25, 0x0a, 0x0e, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x53, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4d, 0x0a, 0x0c, 0x44,
	0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x06, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2e, 0x62,
	0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73,
	0x2e, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x44, 0x65,
	0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x1c, 0x0a, 0x09, 0x52, 0x65,
	0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x52,
	0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x45, 0x72, 0x72, 0x6f,
	0x72, 0x73, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x73,
	0x12, 0x16, 0x0a, 0x06, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x73, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x06, 0x50, 0x61, 0x6e, 0x69, 0x63, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x43, 0x6c, 0x69, 0x65,
	0x6e, 0x74, 0x5f, 0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x43, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22,
	0x3b, 0x0a, 0x0b, 0x44, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12,
	0x0a, 0x04, 0x50, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x50, 0x61,
	0x74, 0x68, 0x12, 0x18, 0x0a, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x07, 0x56, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x32, 0x73, 0x0a, 0x09,
	0x41, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69, 0x63, 0x73, 0x12, 0x66, 0x0a, 0x09, 0x53, 0x65, 0x6e,
	0x64, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x2b, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x71, 0x75,
	0x65, 0x72, 0x79, 0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x6e, 0x61, 0x6c,
	0x79, 0x74, 0x69, 0x63, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x1a, 0x2c, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x71, 0x75, 0x65, 0x72, 0x79,
	0x2e, 0x62, 0x61, 0x63, 0x6b, 0x65, 0x6e, 0x64, 0x2e, 0x61, 0x6e, 0x61, 0x6c, 0x79, 0x74, 0x69,
	0x63, 0x73, 0x2e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x42, 0x05, 0x5a, 0x03, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_internal_pb_analytics_proto_rawDescOnce sync.Once
	file_internal_pb_analytics_proto_rawDescData = file_internal_pb_analytics_proto_rawDesc
)

func file_internal_pb_analytics_proto_rawDescGZIP() []byte {
	file_internal_pb_analytics_proto_rawDescOnce.Do(func() {
		file_internal_pb_analytics_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_pb_analytics_proto_rawDescData)
	})
	return file_internal_pb_analytics_proto_rawDescData
}

var file_internal_pb_analytics_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_internal_pb_analytics_proto_goTypes = []interface{}{
	(*Event)(nil),          // 0: cloudquery.backend.analytics.Event
	(*SyncSummary)(nil),    // 1: cloudquery.backend.analytics.SyncSummary
	(*Destination)(nil),    // 2: cloudquery.backend.analytics.Destination
	(*Event_Request)(nil),  // 3: cloudquery.backend.analytics.Event.Request
	(*Event_Response)(nil), // 4: cloudquery.backend.analytics.Event.Response
}
var file_internal_pb_analytics_proto_depIdxs = []int32{
	2, // 0: cloudquery.backend.analytics.SyncSummary.Destinations:type_name -> cloudquery.backend.analytics.Destination
	1, // 1: cloudquery.backend.analytics.Event.Request.sync_summary:type_name -> cloudquery.backend.analytics.SyncSummary
	3, // 2: cloudquery.backend.analytics.Analytics.SendEvent:input_type -> cloudquery.backend.analytics.Event.Request
	4, // 3: cloudquery.backend.analytics.Analytics.SendEvent:output_type -> cloudquery.backend.analytics.Event.Response
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_internal_pb_analytics_proto_init() }
func file_internal_pb_analytics_proto_init() {
	if File_internal_pb_analytics_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_pb_analytics_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_internal_pb_analytics_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SyncSummary); i {
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
		file_internal_pb_analytics_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Destination); i {
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
		file_internal_pb_analytics_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_Request); i {
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
		file_internal_pb_analytics_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Event_Response); i {
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
			RawDescriptor: file_internal_pb_analytics_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_pb_analytics_proto_goTypes,
		DependencyIndexes: file_internal_pb_analytics_proto_depIdxs,
		MessageInfos:      file_internal_pb_analytics_proto_msgTypes,
	}.Build()
	File_internal_pb_analytics_proto = out.File
	file_internal_pb_analytics_proto_rawDesc = nil
	file_internal_pb_analytics_proto_goTypes = nil
	file_internal_pb_analytics_proto_depIdxs = nil
}
