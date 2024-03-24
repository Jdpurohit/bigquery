// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: log.proto

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

type LogsTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Datetime string `protobuf:"bytes,1,opt,name=datetime,proto3" json:"datetime,omitempty"`
	Uid      int64  `protobuf:"varint,2,opt,name=uid,proto3" json:"uid,omitempty"`
	Hash     string `protobuf:"bytes,3,opt,name=hash,proto3" json:"hash,omitempty"`
	Url      string `protobuf:"bytes,4,opt,name=url,proto3" json:"url,omitempty"`
	Event    string `protobuf:"bytes,5,opt,name=event,proto3" json:"event,omitempty"`
	Unit     string `protobuf:"bytes,6,opt,name=unit,proto3" json:"unit,omitempty"`
	UnitName string `protobuf:"bytes,7,opt,name=unit_name,json=unitName,proto3" json:"unit_name,omitempty"`
	Details  string `protobuf:"bytes,8,opt,name=details,proto3" json:"details,omitempty"`
	ConfigId int64  `protobuf:"varint,9,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
}

func (x *LogsTable) Reset() {
	*x = LogsTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsTable) ProtoMessage() {}

func (x *LogsTable) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsTable.ProtoReflect.Descriptor instead.
func (*LogsTable) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{0}
}

func (x *LogsTable) GetDatetime() string {
	if x != nil {
		return x.Datetime
	}
	return ""
}

func (x *LogsTable) GetUid() int64 {
	if x != nil {
		return x.Uid
	}
	return 0
}

func (x *LogsTable) GetHash() string {
	if x != nil {
		return x.Hash
	}
	return ""
}

func (x *LogsTable) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *LogsTable) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *LogsTable) GetUnit() string {
	if x != nil {
		return x.Unit
	}
	return ""
}

func (x *LogsTable) GetUnitName() string {
	if x != nil {
		return x.UnitName
	}
	return ""
}

func (x *LogsTable) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *LogsTable) GetConfigId() int64 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

type LogsHBTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date          string  `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	Event         string  `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	ConfigId      int64   `protobuf:"varint,3,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	Device        string  `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`
	Geo           string  `protobuf:"bytes,5,opt,name=geo,proto3" json:"geo,omitempty"`
	CreativeSize  string  `protobuf:"bytes,6,opt,name=creative_size,json=creativeSize,proto3" json:"creative_size,omitempty"`
	Partner       string  `protobuf:"bytes,7,opt,name=partner,proto3" json:"partner,omitempty"`
	Revenue       float64 `protobuf:"fixed64,8,opt,name=revenue,proto3" json:"revenue,omitempty"`
	Currency      string  `protobuf:"bytes,9,opt,name=currency,proto3" json:"currency,omitempty"`
	IntegrationId int64   `protobuf:"varint,10,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	S2S           bool    `protobuf:"varint,11,opt,name=s2s,proto3" json:"s2s,omitempty"`
}

func (x *LogsHBTable) Reset() {
	*x = LogsHBTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsHBTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsHBTable) ProtoMessage() {}

func (x *LogsHBTable) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsHBTable.ProtoReflect.Descriptor instead.
func (*LogsHBTable) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{1}
}

func (x *LogsHBTable) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *LogsHBTable) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *LogsHBTable) GetConfigId() int64 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *LogsHBTable) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *LogsHBTable) GetGeo() string {
	if x != nil {
		return x.Geo
	}
	return ""
}

func (x *LogsHBTable) GetCreativeSize() string {
	if x != nil {
		return x.CreativeSize
	}
	return ""
}

func (x *LogsHBTable) GetPartner() string {
	if x != nil {
		return x.Partner
	}
	return ""
}

func (x *LogsHBTable) GetRevenue() float64 {
	if x != nil {
		return x.Revenue
	}
	return 0
}

func (x *LogsHBTable) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *LogsHBTable) GetIntegrationId() int64 {
	if x != nil {
		return x.IntegrationId
	}
	return 0
}

func (x *LogsHBTable) GetS2S() bool {
	if x != nil {
		return x.S2S
	}
	return false
}

type LogsFCTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date         string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
	ConfigId     int64  `protobuf:"varint,2,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	Geo          string `protobuf:"bytes,3,opt,name=geo,proto3" json:"geo,omitempty"`
	Device       string `protobuf:"bytes,4,opt,name=device,proto3" json:"device,omitempty"`
	CreativeSize string `protobuf:"bytes,5,opt,name=creative_size,json=creativeSize,proto3" json:"creative_size,omitempty"`
	Event        string `protobuf:"bytes,6,opt,name=event,proto3" json:"event,omitempty"`
	CreativeId   int64  `protobuf:"varint,7,opt,name=creative_id,json=creativeId,proto3" json:"creative_id,omitempty"`
}

func (x *LogsFCTable) Reset() {
	*x = LogsFCTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsFCTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsFCTable) ProtoMessage() {}

func (x *LogsFCTable) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsFCTable.ProtoReflect.Descriptor instead.
func (*LogsFCTable) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{2}
}

func (x *LogsFCTable) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *LogsFCTable) GetConfigId() int64 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *LogsFCTable) GetGeo() string {
	if x != nil {
		return x.Geo
	}
	return ""
}

func (x *LogsFCTable) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *LogsFCTable) GetCreativeSize() string {
	if x != nil {
		return x.CreativeSize
	}
	return ""
}

func (x *LogsFCTable) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *LogsFCTable) GetCreativeId() int64 {
	if x != nil {
		return x.CreativeId
	}
	return 0
}

type LogsBDTable struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp     int64   `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	IntegrationId int64   `protobuf:"varint,2,opt,name=integration_id,json=integrationId,proto3" json:"integration_id,omitempty"`
	ConfigId      int64   `protobuf:"varint,3,opt,name=config_id,json=configId,proto3" json:"config_id,omitempty"`
	Event         string  `protobuf:"bytes,4,opt,name=event,proto3" json:"event,omitempty"`
	Partner       string  `protobuf:"bytes,5,opt,name=partner,proto3" json:"partner,omitempty"`
	Device        string  `protobuf:"bytes,6,opt,name=device,proto3" json:"device,omitempty"`
	Geo           string  `protobuf:"bytes,7,opt,name=geo,proto3" json:"geo,omitempty"`
	CreativeSize  string  `protobuf:"bytes,8,opt,name=creative_size,json=creativeSize,proto3" json:"creative_size,omitempty"`
	Revenue       float64 `protobuf:"fixed64,9,opt,name=revenue,proto3" json:"revenue,omitempty"`
	Currency      string  `protobuf:"bytes,10,opt,name=currency,proto3" json:"currency,omitempty"`
	S2S           bool    `protobuf:"varint,11,opt,name=s2s,proto3" json:"s2s,omitempty"`
}

func (x *LogsBDTable) Reset() {
	*x = LogsBDTable{}
	if protoimpl.UnsafeEnabled {
		mi := &file_log_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LogsBDTable) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LogsBDTable) ProtoMessage() {}

func (x *LogsBDTable) ProtoReflect() protoreflect.Message {
	mi := &file_log_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LogsBDTable.ProtoReflect.Descriptor instead.
func (*LogsBDTable) Descriptor() ([]byte, []int) {
	return file_log_proto_rawDescGZIP(), []int{3}
}

func (x *LogsBDTable) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *LogsBDTable) GetIntegrationId() int64 {
	if x != nil {
		return x.IntegrationId
	}
	return 0
}

func (x *LogsBDTable) GetConfigId() int64 {
	if x != nil {
		return x.ConfigId
	}
	return 0
}

func (x *LogsBDTable) GetEvent() string {
	if x != nil {
		return x.Event
	}
	return ""
}

func (x *LogsBDTable) GetPartner() string {
	if x != nil {
		return x.Partner
	}
	return ""
}

func (x *LogsBDTable) GetDevice() string {
	if x != nil {
		return x.Device
	}
	return ""
}

func (x *LogsBDTable) GetGeo() string {
	if x != nil {
		return x.Geo
	}
	return ""
}

func (x *LogsBDTable) GetCreativeSize() string {
	if x != nil {
		return x.CreativeSize
	}
	return ""
}

func (x *LogsBDTable) GetRevenue() float64 {
	if x != nil {
		return x.Revenue
	}
	return 0
}

func (x *LogsBDTable) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *LogsBDTable) GetS2S() bool {
	if x != nil {
		return x.S2S
	}
	return false
}

var File_log_proto protoreflect.FileDescriptor

var file_log_proto_rawDesc = []byte{
	0x0a, 0x09, 0x6c, 0x6f, 0x67, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08, 0x62, 0x69, 0x67,
	0x71, 0x75, 0x65, 0x72, 0x79, 0x22, 0xdd, 0x01, 0x0a, 0x09, 0x4c, 0x6f, 0x67, 0x73, 0x54, 0x61,
	0x62, 0x6c, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x64, 0x61, 0x74, 0x65, 0x74, 0x69, 0x6d, 0x65, 0x12,
	0x10, 0x0a, 0x03, 0x75, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x03, 0x75, 0x69,
	0x64, 0x12, 0x12, 0x0a, 0x04, 0x68, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x68, 0x61, 0x73, 0x68, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x75, 0x6e, 0x69, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x6e, 0x69,
	0x74, 0x12, 0x1b, 0x0a, 0x09, 0x75, 0x6e, 0x69, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x6e, 0x69, 0x74, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x64, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x49, 0x64, 0x22, 0xac, 0x02, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x73, 0x48, 0x42,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12,
	0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x16, 0x0a, 0x06,
	0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65,
	0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x65, 0x6f, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x67, 0x65, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69,
	0x76, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63,
	0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70,
	0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61,
	0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65,
	0x18, 0x08, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x12,
	0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x63, 0x79, 0x12, 0x25, 0x0a, 0x0e, 0x69,
	0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x0a, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x32, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x73, 0x32, 0x73, 0x22, 0xc4, 0x01, 0x0a, 0x0b, 0x4c, 0x6f, 0x67, 0x73, 0x46, 0x43, 0x54,
	0x61, 0x62, 0x6c, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6e,
	0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x65, 0x6f, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x67, 0x65, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63,
	0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65,
	0x53, 0x69, 0x7a, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x0a, 0x0b, 0x63, 0x72,
	0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x07, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x49, 0x64, 0x22, 0xb6, 0x02, 0x0a, 0x0b,
	0x4c, 0x6f, 0x67, 0x73, 0x42, 0x44, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74,
	0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x25, 0x0a, 0x0e, 0x69, 0x6e, 0x74,
	0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0d, 0x69, 0x6e, 0x74, 0x65, 0x67, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x64,
	0x12, 0x1b, 0x0a, 0x09, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x63, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x49, 0x64, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x76,
	0x65, 0x6e, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70, 0x61, 0x72, 0x74, 0x6e, 0x65, 0x72, 0x12, 0x16, 0x0a,
	0x06, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64,
	0x65, 0x76, 0x69, 0x63, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x67, 0x65, 0x6f, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x67, 0x65, 0x6f, 0x12, 0x23, 0x0a, 0x0d, 0x63, 0x72, 0x65, 0x61, 0x74,
	0x69, 0x76, 0x65, 0x5f, 0x73, 0x69, 0x7a, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x63, 0x72, 0x65, 0x61, 0x74, 0x69, 0x76, 0x65, 0x53, 0x69, 0x7a, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x72, 0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x18, 0x09, 0x20, 0x01, 0x28, 0x01, 0x52, 0x07, 0x72,
	0x65, 0x76, 0x65, 0x6e, 0x75, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e,
	0x63, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x32, 0x73, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x08, 0x52,
	0x03, 0x73, 0x32, 0x73, 0x42, 0x06, 0x5a, 0x04, 0x2e, 0x2f, 0x70, 0x62, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_log_proto_rawDescOnce sync.Once
	file_log_proto_rawDescData = file_log_proto_rawDesc
)

func file_log_proto_rawDescGZIP() []byte {
	file_log_proto_rawDescOnce.Do(func() {
		file_log_proto_rawDescData = protoimpl.X.CompressGZIP(file_log_proto_rawDescData)
	})
	return file_log_proto_rawDescData
}

var file_log_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_log_proto_goTypes = []interface{}{
	(*LogsTable)(nil),   // 0: bigquery.LogsTable
	(*LogsHBTable)(nil), // 1: bigquery.LogsHBTable
	(*LogsFCTable)(nil), // 2: bigquery.LogsFCTable
	(*LogsBDTable)(nil), // 3: bigquery.LogsBDTable
}
var file_log_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_log_proto_init() }
func file_log_proto_init() {
	if File_log_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_log_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsTable); i {
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
		file_log_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsHBTable); i {
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
		file_log_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsFCTable); i {
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
		file_log_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LogsBDTable); i {
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
			RawDescriptor: file_log_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_log_proto_goTypes,
		DependencyIndexes: file_log_proto_depIdxs,
		MessageInfos:      file_log_proto_msgTypes,
	}.Build()
	File_log_proto = out.File
	file_log_proto_rawDesc = nil
	file_log_proto_goTypes = nil
	file_log_proto_depIdxs = nil
}
