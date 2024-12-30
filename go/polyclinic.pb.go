// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: polyclinic.proto

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

type CreatePolyclinicQueueRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Queue *PolyclinicQueue `protobuf:"bytes,1,opt,name=queue,proto3" json:"queue,omitempty"`
}

func (x *CreatePolyclinicQueueRequest) Reset() {
	*x = CreatePolyclinicQueueRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyclinic_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolyclinicQueueRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolyclinicQueueRequest) ProtoMessage() {}

func (x *CreatePolyclinicQueueRequest) ProtoReflect() protoreflect.Message {
	mi := &file_polyclinic_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolyclinicQueueRequest.ProtoReflect.Descriptor instead.
func (*CreatePolyclinicQueueRequest) Descriptor() ([]byte, []int) {
	return file_polyclinic_proto_rawDescGZIP(), []int{0}
}

func (x *CreatePolyclinicQueueRequest) GetQueue() *PolyclinicQueue {
	if x != nil {
		return x.Queue
	}
	return nil
}

type CreatePolyclinicQueueResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *CreatePolyclinicQueueResponse) Reset() {
	*x = CreatePolyclinicQueueResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyclinic_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreatePolyclinicQueueResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreatePolyclinicQueueResponse) ProtoMessage() {}

func (x *CreatePolyclinicQueueResponse) ProtoReflect() protoreflect.Message {
	mi := &file_polyclinic_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreatePolyclinicQueueResponse.ProtoReflect.Descriptor instead.
func (*CreatePolyclinicQueueResponse) Descriptor() ([]byte, []int) {
	return file_polyclinic_proto_rawDescGZIP(), []int{1}
}

func (x *CreatePolyclinicQueueResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CreatePolyclinicQueueResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CreatePolyclinicQueueResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

type CompensateCreatePolyQueueEncounterRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	QueueId     int64 `protobuf:"varint,1,opt,name=queueId,proto3" json:"queueId,omitempty"`
	ActorId     int64 `protobuf:"varint,2,opt,name=actorId,proto3" json:"actorId,omitempty"`
	EncounterId int64 `protobuf:"varint,3,opt,name=encounterId,proto3" json:"encounterId,omitempty"`
}

func (x *CompensateCreatePolyQueueEncounterRequest) Reset() {
	*x = CompensateCreatePolyQueueEncounterRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyclinic_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompensateCreatePolyQueueEncounterRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompensateCreatePolyQueueEncounterRequest) ProtoMessage() {}

func (x *CompensateCreatePolyQueueEncounterRequest) ProtoReflect() protoreflect.Message {
	mi := &file_polyclinic_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompensateCreatePolyQueueEncounterRequest.ProtoReflect.Descriptor instead.
func (*CompensateCreatePolyQueueEncounterRequest) Descriptor() ([]byte, []int) {
	return file_polyclinic_proto_rawDescGZIP(), []int{2}
}

func (x *CompensateCreatePolyQueueEncounterRequest) GetQueueId() int64 {
	if x != nil {
		return x.QueueId
	}
	return 0
}

func (x *CompensateCreatePolyQueueEncounterRequest) GetActorId() int64 {
	if x != nil {
		return x.ActorId
	}
	return 0
}

func (x *CompensateCreatePolyQueueEncounterRequest) GetEncounterId() int64 {
	if x != nil {
		return x.EncounterId
	}
	return 0
}

type CompensateCreatePolyQueueEncounterResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *CompensateCreatePolyQueueEncounterResponse) Reset() {
	*x = CompensateCreatePolyQueueEncounterResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyclinic_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CompensateCreatePolyQueueEncounterResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CompensateCreatePolyQueueEncounterResponse) ProtoMessage() {}

func (x *CompensateCreatePolyQueueEncounterResponse) ProtoReflect() protoreflect.Message {
	mi := &file_polyclinic_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CompensateCreatePolyQueueEncounterResponse.ProtoReflect.Descriptor instead.
func (*CompensateCreatePolyQueueEncounterResponse) Descriptor() ([]byte, []int) {
	return file_polyclinic_proto_rawDescGZIP(), []int{3}
}

func (x *CompensateCreatePolyQueueEncounterResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *CompensateCreatePolyQueueEncounterResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *CompensateCreatePolyQueueEncounterResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

// Model definitions
type PolyclinicQueue struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdQueue             int64       `protobuf:"varint,1,opt,name=idQueue,proto3" json:"idQueue,omitempty"`
	QueueNumber         int64       `protobuf:"varint,2,opt,name=queueNumber,proto3" json:"queueNumber,omitempty"`
	IDEncounter         int64       `protobuf:"varint,3,opt,name=iDEncounter,proto3" json:"iDEncounter,omitempty"`
	MrNumber            string      `protobuf:"bytes,4,opt,name=mrNumber,proto3" json:"mrNumber,omitempty"`
	Room                *HelperVar  `protobuf:"bytes,5,opt,name=room,proto3" json:"room,omitempty"`
	QueueStatus         *HelperVar  `protobuf:"bytes,6,opt,name=queueStatus,proto3" json:"queueStatus,omitempty"`
	GeneralPractitioner *HelperVar  `protobuf:"bytes,7,opt,name=generalPractitioner,proto3" json:"generalPractitioner,omitempty"`
	Accounting          *Accounting `protobuf:"bytes,8,opt,name=accounting,proto3" json:"accounting,omitempty"`
}

func (x *PolyclinicQueue) Reset() {
	*x = PolyclinicQueue{}
	if protoimpl.UnsafeEnabled {
		mi := &file_polyclinic_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PolyclinicQueue) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PolyclinicQueue) ProtoMessage() {}

func (x *PolyclinicQueue) ProtoReflect() protoreflect.Message {
	mi := &file_polyclinic_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PolyclinicQueue.ProtoReflect.Descriptor instead.
func (*PolyclinicQueue) Descriptor() ([]byte, []int) {
	return file_polyclinic_proto_rawDescGZIP(), []int{4}
}

func (x *PolyclinicQueue) GetIdQueue() int64 {
	if x != nil {
		return x.IdQueue
	}
	return 0
}

func (x *PolyclinicQueue) GetQueueNumber() int64 {
	if x != nil {
		return x.QueueNumber
	}
	return 0
}

func (x *PolyclinicQueue) GetIDEncounter() int64 {
	if x != nil {
		return x.IDEncounter
	}
	return 0
}

func (x *PolyclinicQueue) GetMrNumber() string {
	if x != nil {
		return x.MrNumber
	}
	return ""
}

func (x *PolyclinicQueue) GetRoom() *HelperVar {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *PolyclinicQueue) GetQueueStatus() *HelperVar {
	if x != nil {
		return x.QueueStatus
	}
	return nil
}

func (x *PolyclinicQueue) GetGeneralPractitioner() *HelperVar {
	if x != nil {
		return x.GeneralPractitioner
	}
	return nil
}

func (x *PolyclinicQueue) GetAccounting() *Accounting {
	if x != nil {
		return x.Accounting
	}
	return nil
}

var File_polyclinic_proto protoreflect.FileDescriptor

var file_polyclinic_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x6f, 0x6c, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x05, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x1a, 0x0c, 0x68, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4c, 0x0a, 0x1c, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x50, 0x6f, 0x6c, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2c, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x50,
	0x6f, 0x6c, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x05,
	0x71, 0x75, 0x65, 0x75, 0x65, 0x22, 0x71, 0x0a, 0x1d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x6f, 0x6c, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x51, 0x75, 0x65, 0x75, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65,
	0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73,
	0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x81, 0x01, 0x0a, 0x29, 0x43, 0x6f, 0x6d,
	0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x6f, 0x6c,
	0x79, 0x51, 0x75, 0x65, 0x75, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x71, 0x75, 0x65, 0x75, 0x65, 0x49, 0x64,
	0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x61, 0x63, 0x74, 0x6f, 0x72, 0x49, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x65, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x49, 0x64, 0x22, 0x7e, 0x0a, 0x2a,
	0x43, 0x6f, 0x6d, 0x70, 0x65, 0x6e, 0x73, 0x61, 0x74, 0x65, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x50, 0x6f, 0x6c, 0x79, 0x51, 0x75, 0x65, 0x75, 0x65, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f,
	0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22, 0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f,
	0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0xdc, 0x02, 0x0a,
	0x0f, 0x50, 0x6f, 0x6c, 0x79, 0x63, 0x6c, 0x69, 0x6e, 0x69, 0x63, 0x51, 0x75, 0x65, 0x75, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x69, 0x64, 0x51, 0x75, 0x65, 0x75, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x71, 0x75,
	0x65, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x44, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x69, 0x44, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x1a,
	0x0a, 0x08, 0x6d, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x6d, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x6f,
	0x6f, 0x6d, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73,
	0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x56, 0x61, 0x72, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d,
	0x12, 0x32, 0x0a, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65,
	0x6c, 0x70, 0x65, 0x72, 0x56, 0x61, 0x72, 0x52, 0x0b, 0x71, 0x75, 0x65, 0x75, 0x65, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x12, 0x42, 0x0a, 0x13, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x50,
	0x72, 0x61, 0x63, 0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x56, 0x61, 0x72, 0x52, 0x13, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61, 0x6c, 0x50, 0x72, 0x61, 0x63,
	0x74, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x65, 0x72, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73,
	0x69, 0x6d, 0x72, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52,
	0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x42, 0x03, 0x5a, 0x01, 0x2e,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_polyclinic_proto_rawDescOnce sync.Once
	file_polyclinic_proto_rawDescData = file_polyclinic_proto_rawDesc
)

func file_polyclinic_proto_rawDescGZIP() []byte {
	file_polyclinic_proto_rawDescOnce.Do(func() {
		file_polyclinic_proto_rawDescData = protoimpl.X.CompressGZIP(file_polyclinic_proto_rawDescData)
	})
	return file_polyclinic_proto_rawDescData
}

var file_polyclinic_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_polyclinic_proto_goTypes = []interface{}{
	(*CreatePolyclinicQueueRequest)(nil),               // 0: simrs.CreatePolyclinicQueueRequest
	(*CreatePolyclinicQueueResponse)(nil),              // 1: simrs.CreatePolyclinicQueueResponse
	(*CompensateCreatePolyQueueEncounterRequest)(nil),  // 2: simrs.CompensateCreatePolyQueueEncounterRequest
	(*CompensateCreatePolyQueueEncounterResponse)(nil), // 3: simrs.CompensateCreatePolyQueueEncounterResponse
	(*PolyclinicQueue)(nil),                            // 4: simrs.PolyclinicQueue
	(*HelperVar)(nil),                                  // 5: simrs.HelperVar
	(*Accounting)(nil),                                 // 6: simrs.Accounting
}
var file_polyclinic_proto_depIdxs = []int32{
	4, // 0: simrs.CreatePolyclinicQueueRequest.queue:type_name -> simrs.PolyclinicQueue
	5, // 1: simrs.PolyclinicQueue.room:type_name -> simrs.HelperVar
	5, // 2: simrs.PolyclinicQueue.queueStatus:type_name -> simrs.HelperVar
	5, // 3: simrs.PolyclinicQueue.generalPractitioner:type_name -> simrs.HelperVar
	6, // 4: simrs.PolyclinicQueue.accounting:type_name -> simrs.Accounting
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_polyclinic_proto_init() }
func file_polyclinic_proto_init() {
	if File_polyclinic_proto != nil {
		return
	}
	file_helper_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_polyclinic_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolyclinicQueueRequest); i {
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
		file_polyclinic_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreatePolyclinicQueueResponse); i {
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
		file_polyclinic_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompensateCreatePolyQueueEncounterRequest); i {
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
		file_polyclinic_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CompensateCreatePolyQueueEncounterResponse); i {
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
		file_polyclinic_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PolyclinicQueue); i {
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
			RawDescriptor: file_polyclinic_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_polyclinic_proto_goTypes,
		DependencyIndexes: file_polyclinic_proto_depIdxs,
		MessageInfos:      file_polyclinic_proto_msgTypes,
	}.Build()
	File_polyclinic_proto = out.File
	file_polyclinic_proto_rawDesc = nil
	file_polyclinic_proto_goTypes = nil
	file_polyclinic_proto_depIdxs = nil
}
