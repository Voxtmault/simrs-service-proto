// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: front_desk.proto

package __

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

type UpdateQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *Queue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *UpdateQueueRequest) Reset() {
	*x = UpdateQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQueueRequest) ProtoMessage() {}

func (x *UpdateQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQueueRequest.ProtoReflect.Descriptor instead.
func (*UpdateQueueRequest) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{0}
}

func (x *UpdateQueueRequest) GetQueue() *Queue {
	if x != nil {
		return x.Queue
	}
	return nil
}

type UpdateQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *UpdateQueueResponse) Reset() {
	*x = UpdateQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateQueueResponse) ProtoMessage() {}

func (x *UpdateQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateQueueResponse.ProtoReflect.Descriptor instead.
func (*UpdateQueueResponse) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{1}
}

func (x *UpdateQueueResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *UpdateQueueResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *UpdateQueueResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type CompensateUpdateQueueEncounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId int64 `protobuf:"varint,1,opt,name=queueId,proto3" json:"queueId,omitempty"`
	ActorId int64 `protobuf:"varint,2,opt,name=actorId,proto3" json:"actorId,omitempty"`
}

func (x *CompensateUpdateQueueEncounterRequest) Reset() {
	*x = CompensateUpdateQueueEncounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompensateUpdateQueueEncounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompensateUpdateQueueEncounterRequest) ProtoMessage() {}

func (x *CompensateUpdateQueueEncounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompensateUpdateQueueEncounterRequest.ProtoReflect.Descriptor instead.
func (*CompensateUpdateQueueEncounterRequest) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{2}
}

func (x *CompensateUpdateQueueEncounterRequest) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *CompensateUpdateQueueEncounterRequest) GetActorId() int64 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

type CompensateUpdateQueueEncounterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *CompensateUpdateQueueEncounterResponse) Reset() {
	*x = CompensateUpdateQueueEncounterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompensateUpdateQueueEncounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompensateUpdateQueueEncounterResponse) ProtoMessage() {}

func (x *CompensateUpdateQueueEncounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompensateUpdateQueueEncounterResponse.ProtoReflect.Descriptor instead.
func (*CompensateUpdateQueueEncounterResponse) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{3}
}

func (x *CompensateUpdateQueueEncounterResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CompensateUpdateQueueEncounterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CompensateUpdateQueueEncounterResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type GetFrontDeskQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdQueue int64 `protobuf:"varint,1,opt,name=idQueue,proto3" json:"idQueue,omitempty"`
}

func (x *GetFrontDeskQueueRequest) Reset() {
	*x = GetFrontDeskQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFrontDeskQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFrontDeskQueueRequest) ProtoMessage() {}

func (x *GetFrontDeskQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFrontDeskQueueRequest.ProtoReflect.Descriptor instead.
func (*GetFrontDeskQueueRequest) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{4}
}

func (x *GetFrontDeskQueueRequest) GetIdQueue() int64 {
	if x != nil {
		return x.IdQueue
	}
	return 0
}

type GetFrontDeskQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string   `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage string   `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
	Data         []*Queue `protobuf:"bytes,4,rep,name=data,proto3" json:"data,omitempty"`
}

func (x *GetFrontDeskQueueResponse) Reset() {
	*x = GetFrontDeskQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetFrontDeskQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetFrontDeskQueueResponse) ProtoMessage() {}

func (x *GetFrontDeskQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetFrontDeskQueueResponse.ProtoReflect.Descriptor instead.
func (*GetFrontDeskQueueResponse) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{5}
}

func (x *GetFrontDeskQueueResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GetFrontDeskQueueResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GetFrontDeskQueueResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

func (x *GetFrontDeskQueueResponse) GetData() []*Queue {
	if x != nil {
		return x.Data
	}
	return nil
}

type ManualQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdEncounter int64       `protobuf:"varint,1,opt,name=idEncounter,proto3" json:"idEncounter,omitempty"`
	MrNumber    string      `protobuf:"bytes,2,opt,name=mrNumber,proto3" json:"mrNumber,omitempty"`
	Room        *HelperVar  `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
	Accounting  *Accounting `protobuf:"bytes,4,opt,name=accounting,proto3" json:"accounting,omitempty"`
}

func (x *ManualQueueRequest) Reset() {
	*x = ManualQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManualQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManualQueueRequest) ProtoMessage() {}

func (x *ManualQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManualQueueRequest.ProtoReflect.Descriptor instead.
func (*ManualQueueRequest) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{6}
}

func (x *ManualQueueRequest) GetIdEncounter() int64 {
	if x != nil {
		return x.IdEncounter
	}
	return 0
}

func (x *ManualQueueRequest) GetMrNumber() string {
	if x != nil {
		return x.MrNumber
	}
	return ""
}

func (x *ManualQueueRequest) GetRoom() *HelperVar {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *ManualQueueRequest) GetAccounting() *Accounting {
	if x != nil {
		return x.Accounting
	}
	return nil
}

type ManualQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *ManualQueueResponse) Reset() {
	*x = ManualQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ManualQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ManualQueueResponse) ProtoMessage() {}

func (x *ManualQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ManualQueueResponse.ProtoReflect.Descriptor instead.
func (*ManualQueueResponse) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{7}
}

func (x *ManualQueueResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *ManualQueueResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *ManualQueueResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type CompensateManualQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdQueue    int64       `protobuf:"varint,1,opt,name=idQueue,proto3" json:"idQueue,omitempty"`
	Accounting *Accounting `protobuf:"bytes,2,opt,name=accounting,proto3" json:"accounting,omitempty"`
}

func (x *CompensateManualQueueRequest) Reset() {
	*x = CompensateManualQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompensateManualQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompensateManualQueueRequest) ProtoMessage() {}

func (x *CompensateManualQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompensateManualQueueRequest.ProtoReflect.Descriptor instead.
func (*CompensateManualQueueRequest) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{8}
}

func (x *CompensateManualQueueRequest) GetIdQueue() int64 {
	if x != nil {
		return x.IdQueue
	}
	return 0
}

func (x *CompensateManualQueueRequest) GetAccounting() *Accounting {
	if x != nil {
		return x.Accounting
	}
	return nil
}

type CompensateManualQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *CompensateManualQueueResponse) Reset() {
	*x = CompensateManualQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompensateManualQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompensateManualQueueResponse) ProtoMessage() {}

func (x *CompensateManualQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompensateManualQueueResponse.ProtoReflect.Descriptor instead.
func (*CompensateManualQueueResponse) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{9}
}

func (x *CompensateManualQueueResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CompensateManualQueueResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CompensateManualQueueResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Model definitions
type Queue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdQueue     int64       `protobuf:"varint,1,opt,name=idQueue,proto3" json:"idQueue,omitempty"`
	QueueNumber string      `protobuf:"bytes,2,opt,name=queueNumber,proto3" json:"queueNumber,omitempty"`
	IDEncounter int64       `protobuf:"varint,3,opt,name=IDEncounter,proto3" json:"IDEncounter,omitempty"`
	MrNumber    string      `protobuf:"bytes,4,opt,name=mrNumber,proto3" json:"mrNumber,omitempty"`
	Room        *HelperVar  `protobuf:"bytes,5,opt,name=room,proto3" json:"room,omitempty"`
	QueueStatus *HelperVar  `protobuf:"bytes,6,opt,name=QueueStatus,proto3" json:"QueueStatus,omitempty"`
	Accounting  *Accounting `protobuf:"bytes,7,opt,name=accounting,proto3" json:"accounting,omitempty"`
}

func (x *Queue) Reset() {
	*x = Queue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_front_desk_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Queue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Queue) ProtoMessage() {}

func (x *Queue) ProtoReflect() protoreflect.Message {
	mi := &file_front_desk_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Queue.ProtoReflect.Descriptor instead.
func (*Queue) Descriptor() ([]byte, []int) {
	return file_front_desk_proto_rawDescGZIP(), []int{10}
}

func (x *Queue) GetIdQueue() int64 {
	if x != nil {
		return x.IdQueue
	}
	return 0
}

func (x *Queue) GetQueueNumber() string {
	if x != nil {
		return x.QueueNumber
	}
	return ""
}

func (x *Queue) GetIDEncounter() int64 {
	if x != nil {
		return x.IDEncounter
	}
	return 0
}

func (x *Queue) GetMrNumber() string {
	if x != nil {
		return x.MrNumber
	}
	return ""
}

func (x *Queue) GetRoom() *HelperVar {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *Queue) GetQueueStatus() *HelperVar {
	if x != nil {
		return x.QueueStatus
	}
	return nil
}

func (x *Queue) GetAccounting() *Accounting {
	if x != nil {
		return x.Accounting
	}
	return nil
}

var File_front_desk_proto protoreflect.FileDescriptor

var file_front_desk_proto_rawDesc = []byte{
	0x0a, 0x10, 0x66, 0x72, 0x6f, 0x6e, 0x74, 0x5f, 0x64, 0x65, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x1a, 0x0c, 0x68, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x12, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x22, 0x0a,
	0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73,
	0x69, 0x6d, 0x72, 0x73, 0x2e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x22, 0x67, 0x0a, 0x13, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72,
	0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x5b, 0x0a, 0x25, 0x43, 0x6f,
	0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75,
	0x65, 0x75, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64, 0x12, 0x18, 0x0a,
	0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07,
	0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x22, 0x7a, 0x0a, 0x26, 0x43, 0x6f, 0x6d, 0x70, 0x65,
	0x6e, 0x73, 0x61, 0x74, 0x65, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12,
	0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73,
	0x61, 0x67, 0x65, 0x22, 0x34, 0x0a, 0x18, 0x47, 0x65, 0x74, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x44,
	0x65, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x18, 0x0a, 0x07, 0x69, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x07, 0x69, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x22, 0x8f, 0x01, 0x0a, 0x19, 0x47, 0x65,
	0x74, 0x46, 0x72, 0x6f, 0x6e, 0x74, 0x44, 0x65, 0x73, 0x6b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x20, 0x0a, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x22, 0xab, 0x01, 0x0a, 0x12,
	0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72,
	0x12, 0x24, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x56, 0x61, 0x72,
	0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x6d,
	0x72, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61,
	0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x67, 0x0a, 0x13, 0x4d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x6b, 0x0a, 0x1c, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65,
	0x4d, 0x61, 0x6e, 0x75, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x31, 0x0a, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x11, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x22,
	0x71, 0x0a, 0x1d, 0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x4d, 0x61, 0x6e,
	0x75, 0x61, 0x6c, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x22, 0x8e, 0x02, 0x0a, 0x05, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x69, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69,
	0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e,
	0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0b, 0x71, 0x75, 0x65,
	0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b, 0x49, 0x44, 0x45, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x49,
	0x44, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x72,
	0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c,
	0x70, 0x65, 0x72, 0x56, 0x61, 0x72, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x32, 0x0a, 0x0b,
	0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x56, 0x61, 0x72, 0x52, 0x0b, 0x51, 0x75, 0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x41, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x69, 0x6e, 0x67, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_front_desk_proto_rawDescOnce sync.Once
	file_front_desk_proto_rawDescData = file_front_desk_proto_rawDesc
)

func file_front_desk_proto_rawDescGZIP() []byte {
	file_front_desk_proto_rawDescOnce.Do(func() {
		file_front_desk_proto_rawDescData = protoimpl.X.CompressGZIP(file_front_desk_proto_rawDescData)
	})
	return file_front_desk_proto_rawDescData
}

var file_front_desk_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_front_desk_proto_goTypes = []interface{}{
	(*UpdateQueueRequest)(nil),                     // 0: simrs.UpdateQueueRequest
	(*UpdateQueueResponse)(nil),                    // 1: simrs.UpdateQueueResponse
	(*CompensateUpdateQueueEncounterRequest)(nil),  // 2: simrs.CompensateUpdateQueueEncounterRequest
	(*CompensateUpdateQueueEncounterResponse)(nil), // 3: simrs.CompensateUpdateQueueEncounterResponse
	(*GetFrontDeskQueueRequest)(nil),               // 4: simrs.GetFrontDeskQueueRequest
	(*GetFrontDeskQueueResponse)(nil),              // 5: simrs.GetFrontDeskQueueResponse
	(*ManualQueueRequest)(nil),                     // 6: simrs.ManualQueueRequest
	(*ManualQueueResponse)(nil),                    // 7: simrs.ManualQueueResponse
	(*CompensateManualQueueRequest)(nil),           // 8: simrs.CompensateManualQueueRequest
	(*CompensateManualQueueResponse)(nil),          // 9: simrs.CompensateManualQueueResponse
	(*Queue)(nil),                                  // 10: simrs.Queue
	(*HelperVar)(nil),                              // 11: simrs.HelperVar
	(*Accounting)(nil),                             // 12: simrs.Accounting
}
var file_front_desk_proto_depIdxs = []int32{
	10, // 0: simrs.UpdateQueueRequest.queue:type_name -> simrs.Queue
	10, // 1: simrs.GetFrontDeskQueueResponse.data:type_name -> simrs.Queue
	11, // 2: simrs.ManualQueueRequest.room:type_name -> simrs.HelperVar
	12, // 3: simrs.ManualQueueRequest.accounting:type_name -> simrs.Accounting
	12, // 4: simrs.CompensateManualQueueRequest.accounting:type_name -> simrs.Accounting
	11, // 5: simrs.Queue.room:type_name -> simrs.HelperVar
	11, // 6: simrs.Queue.QueueStatus:type_name -> simrs.HelperVar
	12, // 7: simrs.Queue.accounting:type_name -> simrs.Accounting
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_front_desk_proto_init() }
func file_front_desk_proto_init() {
	if File_front_desk_proto != nil {
		return
	}
	file_helper_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_front_desk_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQueueRequest); i {
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
		file_front_desk_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateQueueResponse); i {
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
		file_front_desk_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompensateUpdateQueueEncounterRequest); i {
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
		file_front_desk_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompensateUpdateQueueEncounterResponse); i {
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
		file_front_desk_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFrontDeskQueueRequest); i {
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
		file_front_desk_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetFrontDeskQueueResponse); i {
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
		file_front_desk_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManualQueueRequest); i {
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
		file_front_desk_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ManualQueueResponse); i {
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
		file_front_desk_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompensateManualQueueRequest); i {
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
		file_front_desk_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompensateManualQueueResponse); i {
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
		file_front_desk_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Queue); i {
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
			RawDescriptor: file_front_desk_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_front_desk_proto_goTypes,
		DependencyIndexes: file_front_desk_proto_depIdxs,
		MessageInfos:      file_front_desk_proto_msgTypes,
	}.Build()
	File_front_desk_proto = out.File
	file_front_desk_proto_rawDesc = nil
	file_front_desk_proto_goTypes = nil
	file_front_desk_proto_depIdxs = nil
}
