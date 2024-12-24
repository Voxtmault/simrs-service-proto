// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: encounter.proto

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

type GetEncounterDetailsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PageSize        int64  `protobuf:"varint,1,opt,name=pageSize,proto3" json:"pageSize,omitempty"`
	PageNumber      int64  `protobuf:"varint,2,opt,name=pageNumber,proto3" json:"pageNumber,omitempty"`
	IdEncounter     int64  `protobuf:"varint,3,opt,name=idEncounter,proto3" json:"idEncounter,omitempty"`
	MrNumber        string `protobuf:"bytes,4,opt,name=mrNumber,proto3" json:"mrNumber,omitempty"`
	ServiceType     int32  `protobuf:"varint,5,opt,name=serviceType,proto3" json:"serviceType,omitempty"`
	EncounterStatus int32  `protobuf:"varint,6,opt,name=encounterStatus,proto3" json:"encounterStatus,omitempty"`
	RangeDateStart  string `protobuf:"bytes,7,opt,name=rangeDateStart,proto3" json:"rangeDateStart,omitempty"`
	RangeDateEnd    string `protobuf:"bytes,8,opt,name=rangeDateEnd,proto3" json:"rangeDateEnd,omitempty"`
	CustomStr       string `protobuf:"bytes,9,opt,name=customStr,proto3" json:"customStr,omitempty"`
}

func (x *GetEncounterDetailsRequest) Reset() {
	*x = GetEncounterDetailsRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEncounterDetailsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEncounterDetailsRequest) ProtoMessage() {}

func (x *GetEncounterDetailsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEncounterDetailsRequest.ProtoReflect.Descriptor instead.
func (*GetEncounterDetailsRequest) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{0}
}

func (x *GetEncounterDetailsRequest) GetPageSize() int64 {
	if x != nil {
		return x.PageSize
	}
	return 0
}

func (x *GetEncounterDetailsRequest) GetPageNumber() int64 {
	if x != nil {
		return x.PageNumber
	}
	return 0
}

func (x *GetEncounterDetailsRequest) GetIdEncounter() int64 {
	if x != nil {
		return x.IdEncounter
	}
	return 0
}

func (x *GetEncounterDetailsRequest) GetMrNumber() string {
	if x != nil {
		return x.MrNumber
	}
	return ""
}

func (x *GetEncounterDetailsRequest) GetServiceType() int32 {
	if x != nil {
		return x.ServiceType
	}
	return 0
}

func (x *GetEncounterDetailsRequest) GetEncounterStatus() int32 {
	if x != nil {
		return x.EncounterStatus
	}
	return 0
}

func (x *GetEncounterDetailsRequest) GetRangeDateStart() string {
	if x != nil {
		return x.RangeDateStart
	}
	return ""
}

func (x *GetEncounterDetailsRequest) GetRangeDateEnd() string {
	if x != nil {
		return x.RangeDateEnd
	}
	return ""
}

func (x *GetEncounterDetailsRequest) GetCustomStr() string {
	if x != nil {
		return x.CustomStr
	}
	return ""
}

type GetEncounterDetailsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	GenericResponse *GenericEncounterServiceResponse `protobuf:"bytes,1,opt,name=genericResponse,proto3" json:"genericResponse,omitempty"`
	Encounters      []*Encounter                     `protobuf:"bytes,2,rep,name=encounters,proto3" json:"encounters,omitempty"`
}

func (x *GetEncounterDetailsResponse) Reset() {
	*x = GetEncounterDetailsResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetEncounterDetailsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetEncounterDetailsResponse) ProtoMessage() {}

func (x *GetEncounterDetailsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetEncounterDetailsResponse.ProtoReflect.Descriptor instead.
func (*GetEncounterDetailsResponse) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{1}
}

func (x *GetEncounterDetailsResponse) GetGenericResponse() *GenericEncounterServiceResponse {
	if x != nil {
		return x.GenericResponse
	}
	return nil
}

func (x *GetEncounterDetailsResponse) GetEncounters() []*Encounter {
	if x != nil {
		return x.Encounters
	}
	return nil
}

type AddEncounterRoomHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdEncounter  int64  `protobuf:"varint,1,opt,name=idEncounter,proto3" json:"idEncounter,omitempty"`
	IdRoom       int64  `protobuf:"varint,2,opt,name=idRoom,proto3" json:"idRoom,omitempty"`
	IdActor      int64  `protobuf:"varint,3,opt,name=idActor,proto3" json:"idActor,omitempty"`
	OptionalNote string `protobuf:"bytes,4,opt,name=optionalNote,proto3" json:"optionalNote,omitempty"`
}

func (x *AddEncounterRoomHistoryRequest) Reset() {
	*x = AddEncounterRoomHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddEncounterRoomHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddEncounterRoomHistoryRequest) ProtoMessage() {}

func (x *AddEncounterRoomHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddEncounterRoomHistoryRequest.ProtoReflect.Descriptor instead.
func (*AddEncounterRoomHistoryRequest) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{2}
}

func (x *AddEncounterRoomHistoryRequest) GetIdEncounter() int64 {
	if x != nil {
		return x.IdEncounter
	}
	return 0
}

func (x *AddEncounterRoomHistoryRequest) GetIdRoom() int64 {
	if x != nil {
		return x.IdRoom
	}
	return 0
}

func (x *AddEncounterRoomHistoryRequest) GetIdActor() int64 {
	if x != nil {
		return x.IdActor
	}
	return 0
}

func (x *AddEncounterRoomHistoryRequest) GetOptionalNote() string {
	if x != nil {
		return x.OptionalNote
	}
	return ""
}

type DeleteEncounterRoomHistoryRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdEncounter int64 `protobuf:"varint,1,opt,name=idEncounter,proto3" json:"idEncounter,omitempty"`
	IdRoom      int64 `protobuf:"varint,2,opt,name=idRoom,proto3" json:"idRoom,omitempty"`
}

func (x *DeleteEncounterRoomHistoryRequest) Reset() {
	*x = DeleteEncounterRoomHistoryRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteEncounterRoomHistoryRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteEncounterRoomHistoryRequest) ProtoMessage() {}

func (x *DeleteEncounterRoomHistoryRequest) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteEncounterRoomHistoryRequest.ProtoReflect.Descriptor instead.
func (*DeleteEncounterRoomHistoryRequest) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{3}
}

func (x *DeleteEncounterRoomHistoryRequest) GetIdEncounter() int64 {
	if x != nil {
		return x.IdEncounter
	}
	return 0
}

func (x *DeleteEncounterRoomHistoryRequest) GetIdRoom() int64 {
	if x != nil {
		return x.IdRoom
	}
	return 0
}

// Model definitions for Encounter
type Encounter struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	IdEncounter     int64                       `protobuf:"varint,1,opt,name=idEncounter,proto3" json:"idEncounter,omitempty"`
	MrNumber        string                      `protobuf:"bytes,2,opt,name=mrNumber,proto3" json:"mrNumber,omitempty"`
	AdmittionDate   string                      `protobuf:"bytes,3,opt,name=admittionDate,proto3" json:"admittionDate,omitempty"`
	DischargedDate  string                      `protobuf:"bytes,4,opt,name=dischargedDate,proto3" json:"dischargedDate,omitempty"`
	PaymentMethod   *HelperVar                  `protobuf:"bytes,5,opt,name=paymentMethod,proto3" json:"paymentMethod,omitempty"`
	EntryMethod     *HelperVar                  `protobuf:"bytes,6,opt,name=entryMethod,proto3" json:"entryMethod,omitempty"`
	EncounterStatus *HelperVar                  `protobuf:"bytes,7,opt,name=encounterStatus,proto3" json:"encounterStatus,omitempty"`
	EncounterType   *HelperVar                  `protobuf:"bytes,8,opt,name=encounterType,proto3" json:"encounterType,omitempty"`
	RoomHistory     []*EncounterRoomHistory     `protobuf:"bytes,9,rep,name=roomHistory,proto3" json:"roomHistory,omitempty"`
	StatusLog       []*EncounterUpdateStatusLog `protobuf:"bytes,10,rep,name=statusLog,proto3" json:"statusLog,omitempty"`
	Pic             *EncounterPersonInCharge    `protobuf:"bytes,11,opt,name=pic,proto3" json:"pic,omitempty"`
	Note            string                      `protobuf:"bytes,12,opt,name=note,proto3" json:"note,omitempty"`
	Accounting      *Accounting                 `protobuf:"bytes,13,opt,name=accounting,proto3" json:"accounting,omitempty"`
}

func (x *Encounter) Reset() {
	*x = Encounter{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Encounter) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Encounter) ProtoMessage() {}

func (x *Encounter) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Encounter.ProtoReflect.Descriptor instead.
func (*Encounter) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{4}
}

func (x *Encounter) GetIdEncounter() int64 {
	if x != nil {
		return x.IdEncounter
	}
	return 0
}

func (x *Encounter) GetMrNumber() string {
	if x != nil {
		return x.MrNumber
	}
	return ""
}

func (x *Encounter) GetAdmittionDate() string {
	if x != nil {
		return x.AdmittionDate
	}
	return ""
}

func (x *Encounter) GetDischargedDate() string {
	if x != nil {
		return x.DischargedDate
	}
	return ""
}

func (x *Encounter) GetPaymentMethod() *HelperVar {
	if x != nil {
		return x.PaymentMethod
	}
	return nil
}

func (x *Encounter) GetEntryMethod() *HelperVar {
	if x != nil {
		return x.EntryMethod
	}
	return nil
}

func (x *Encounter) GetEncounterStatus() *HelperVar {
	if x != nil {
		return x.EncounterStatus
	}
	return nil
}

func (x *Encounter) GetEncounterType() *HelperVar {
	if x != nil {
		return x.EncounterType
	}
	return nil
}

func (x *Encounter) GetRoomHistory() []*EncounterRoomHistory {
	if x != nil {
		return x.RoomHistory
	}
	return nil
}

func (x *Encounter) GetStatusLog() []*EncounterUpdateStatusLog {
	if x != nil {
		return x.StatusLog
	}
	return nil
}

func (x *Encounter) GetPic() *EncounterPersonInCharge {
	if x != nil {
		return x.Pic
	}
	return nil
}

func (x *Encounter) GetNote() string {
	if x != nil {
		return x.Note
	}
	return ""
}

func (x *Encounter) GetAccounting() *Accounting {
	if x != nil {
		return x.Accounting
	}
	return nil
}

type EncounterRoomHistory struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdEncounter   int64       `protobuf:"varint,2,opt,name=idEncounter,proto3" json:"idEncounter,omitempty"`
	Room          *HelperVar  `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
	AdmittionDate string      `protobuf:"bytes,4,opt,name=admittionDate,proto3" json:"admittionDate,omitempty"`
	DischargeDate string      `protobuf:"bytes,5,opt,name=DischargeDate,proto3" json:"DischargeDate,omitempty"`
	Sep           string      `protobuf:"bytes,6,opt,name=sep,proto3" json:"sep,omitempty"`
	Accounting    *Accounting `protobuf:"bytes,7,opt,name=accounting,proto3" json:"accounting,omitempty"`
}

func (x *EncounterRoomHistory) Reset() {
	*x = EncounterRoomHistory{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncounterRoomHistory) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncounterRoomHistory) ProtoMessage() {}

func (x *EncounterRoomHistory) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncounterRoomHistory.ProtoReflect.Descriptor instead.
func (*EncounterRoomHistory) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{5}
}

func (x *EncounterRoomHistory) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EncounterRoomHistory) GetIdEncounter() int64 {
	if x != nil {
		return x.IdEncounter
	}
	return 0
}

func (x *EncounterRoomHistory) GetRoom() *HelperVar {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *EncounterRoomHistory) GetAdmittionDate() string {
	if x != nil {
		return x.AdmittionDate
	}
	return ""
}

func (x *EncounterRoomHistory) GetDischargeDate() string {
	if x != nil {
		return x.DischargeDate
	}
	return ""
}

func (x *EncounterRoomHistory) GetSep() string {
	if x != nil {
		return x.Sep
	}
	return ""
}

func (x *EncounterRoomHistory) GetAccounting() *Accounting {
	if x != nil {
		return x.Accounting
	}
	return nil
}

type EncounterUpdateStatusLog struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64       `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	IdEncounter     int64       `protobuf:"varint,2,opt,name=idEncounter,proto3" json:"idEncounter,omitempty"`
	Room            *HelperVar  `protobuf:"bytes,3,opt,name=room,proto3" json:"room,omitempty"`
	EncounterStatus *HelperVar  `protobuf:"bytes,4,opt,name=encounterStatus,proto3" json:"encounterStatus,omitempty"`
	Accounting      *Accounting `protobuf:"bytes,5,opt,name=accounting,proto3" json:"accounting,omitempty"`
}

func (x *EncounterUpdateStatusLog) Reset() {
	*x = EncounterUpdateStatusLog{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncounterUpdateStatusLog) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncounterUpdateStatusLog) ProtoMessage() {}

func (x *EncounterUpdateStatusLog) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncounterUpdateStatusLog.ProtoReflect.Descriptor instead.
func (*EncounterUpdateStatusLog) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{6}
}

func (x *EncounterUpdateStatusLog) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *EncounterUpdateStatusLog) GetIdEncounter() int64 {
	if x != nil {
		return x.IdEncounter
	}
	return 0
}

func (x *EncounterUpdateStatusLog) GetRoom() *HelperVar {
	if x != nil {
		return x.Room
	}
	return nil
}

func (x *EncounterUpdateStatusLog) GetEncounterStatus() *HelperVar {
	if x != nil {
		return x.EncounterStatus
	}
	return nil
}

func (x *EncounterUpdateStatusLog) GetAccounting() *Accounting {
	if x != nil {
		return x.Accounting
	}
	return nil
}

type EncounterPersonInCharge struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name     string     `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Phone    string     `protobuf:"bytes,2,opt,name=phone,proto3" json:"phone,omitempty"`
	Relation *HelperVar `protobuf:"bytes,3,opt,name=relation,proto3" json:"relation,omitempty"`
}

func (x *EncounterPersonInCharge) Reset() {
	*x = EncounterPersonInCharge{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EncounterPersonInCharge) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EncounterPersonInCharge) ProtoMessage() {}

func (x *EncounterPersonInCharge) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EncounterPersonInCharge.ProtoReflect.Descriptor instead.
func (*EncounterPersonInCharge) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{7}
}

func (x *EncounterPersonInCharge) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *EncounterPersonInCharge) GetPhone() string {
	if x != nil {
		return x.Phone
	}
	return ""
}

func (x *EncounterPersonInCharge) GetRelation() *HelperVar {
	if x != nil {
		return x.Relation
	}
	return nil
}

// Generic Encounter Service Response
type GenericEncounterServiceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code         int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Message      string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	ErrorMessage string `protobuf:"bytes,3,opt,name=errorMessage,proto3" json:"errorMessage,omitempty"`
}

func (x *GenericEncounterServiceResponse) Reset() {
	*x = GenericEncounterServiceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_encounter_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GenericEncounterServiceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GenericEncounterServiceResponse) ProtoMessage() {}

func (x *GenericEncounterServiceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_encounter_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GenericEncounterServiceResponse.ProtoReflect.Descriptor instead.
func (*GenericEncounterServiceResponse) Descriptor() ([]byte, []int) {
	return file_encounter_proto_rawDescGZIP(), []int{8}
}

func (x *GenericEncounterServiceResponse) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *GenericEncounterServiceResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *GenericEncounterServiceResponse) GetErrorMessage() string {
	if x != nil {
		return x.ErrorMessage
	}
	return ""
}

var File_encounter_proto protoreflect.FileDescriptor

var file_encounter_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x1a, 0x0c, 0x68, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcc, 0x02, 0x0a, 0x1a, 0x47, 0x65, 0x74, 0x45, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x70, 0x61, 0x67, 0x65, 0x53, 0x69, 0x7a,
	0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0a, 0x70, 0x61, 0x67, 0x65, 0x4e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x20, 0x0a, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54, 0x79, 0x70,
	0x65, 0x12, 0x28, 0x0a, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74,
	0x61, 0x74, 0x75, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x26, 0x0a, 0x0e, 0x72,
	0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x72, 0x74, 0x18, 0x07, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0e, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65, 0x53, 0x74,
	0x61, 0x72, 0x74, 0x12, 0x22, 0x0a, 0x0c, 0x72, 0x61, 0x6e, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x45, 0x6e, 0x64, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x72, 0x61, 0x6e, 0x67, 0x65,
	0x44, 0x61, 0x74, 0x65, 0x45, 0x6e, 0x64, 0x12, 0x1c, 0x0a, 0x09, 0x63, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x53, 0x74, 0x72, 0x18, 0x09, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x63, 0x75, 0x73, 0x74,
	0x6f, 0x6d, 0x53, 0x74, 0x72, 0x22, 0xa1, 0x01, 0x0a, 0x1b, 0x47, 0x65, 0x74, 0x45, 0x6e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x44, 0x65, 0x74, 0x61, 0x69, 0x6c, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x50, 0x0a, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63,
	0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26,
	0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x52, 0x0f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x30, 0x0a, 0x0a, 0x65, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69,
	0x6d, 0x72, 0x73, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x0a, 0x65,
	0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x73, 0x22, 0x98, 0x01, 0x0a, 0x1e, 0x41, 0x64,
	0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x48, 0x69,
	0x73, 0x74, 0x6f, 0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b,
	0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16,
	0x0a, 0x06, 0x69, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06,
	0x69, 0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x12, 0x18, 0x0a, 0x07, 0x69, 0x64, 0x41, 0x63, 0x74, 0x6f,
	0x72, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x69, 0x64, 0x41, 0x63, 0x74, 0x6f, 0x72,
	0x12, 0x22, 0x0a, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c, 0x4e, 0x6f, 0x74, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x61, 0x6c,
	0x4e, 0x6f, 0x74, 0x65, 0x22, 0x5d, 0x0a, 0x21, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x45, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f,
	0x72, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x45,
	0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x16, 0x0a, 0x06, 0x69,
	0x64, 0x52, 0x6f, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x69, 0x64, 0x52,
	0x6f, 0x6f, 0x6d, 0x22, 0xee, 0x04, 0x0a, 0x09, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x72, 0x4e, 0x75, 0x6d, 0x62, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x74, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x74, 0x74, 0x69, 0x6f,
	0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x26, 0x0a, 0x0e, 0x64, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72,
	0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0e, 0x64,
	0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x64, 0x44, 0x61, 0x74, 0x65, 0x12, 0x36, 0x0a,
	0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c,
	0x70, 0x65, 0x72, 0x56, 0x61, 0x72, 0x52, 0x0d, 0x70, 0x61, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x4d,
	0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x32, 0x0a, 0x0b, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x4d, 0x65,
	0x74, 0x68, 0x6f, 0x64, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d,
	0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x56, 0x61, 0x72, 0x52, 0x0b, 0x65, 0x6e,
	0x74, 0x72, 0x79, 0x4d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x12, 0x3a, 0x0a, 0x0f, 0x65, 0x6e, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65,
	0x72, 0x56, 0x61, 0x72, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x36, 0x0a, 0x0d, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73,
	0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x56, 0x61, 0x72, 0x52, 0x0d,
	0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x54, 0x79, 0x70, 0x65, 0x12, 0x3d, 0x0a,
	0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x09, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x1b, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x52,
	0x0b, 0x72, 0x6f, 0x6f, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x3d, 0x0a, 0x09,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x6f, 0x67, 0x18, 0x0a, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x1f, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x6f, 0x67,
	0x52, 0x09, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4c, 0x6f, 0x67, 0x12, 0x30, 0x0a, 0x03, 0x70,
	0x69, 0x63, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73,
	0x2e, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e,
	0x49, 0x6e, 0x43, 0x68, 0x61, 0x72, 0x67, 0x65, 0x52, 0x03, 0x70, 0x69, 0x63, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x6f, 0x74, 0x65, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x6f, 0x74,
	0x65, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18,
	0x0d, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x41, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x69, 0x6e, 0x67, 0x22, 0xff, 0x01, 0x0a, 0x14, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x52, 0x6f, 0x6f, 0x6d, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x20, 0x0a,
	0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x12,
	0x24, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e,
	0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72, 0x56, 0x61, 0x72, 0x52,
	0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x61, 0x64, 0x6d, 0x69, 0x74, 0x74, 0x69,
	0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0d, 0x61, 0x64,
	0x6d, 0x69, 0x74, 0x74, 0x69, 0x6f, 0x6e, 0x44, 0x61, 0x74, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x44,
	0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x44, 0x61, 0x74, 0x65, 0x18, 0x05, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x44, 0x69, 0x73, 0x63, 0x68, 0x61, 0x72, 0x67, 0x65, 0x44, 0x61, 0x74,
	0x65, 0x12, 0x10, 0x0a, 0x03, 0x73, 0x65, 0x70, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03,
	0x73, 0x65, 0x70, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e,
	0x67, 0x18, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e,
	0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a, 0x61, 0x63, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x22, 0xe1, 0x01, 0x0a, 0x18, 0x45, 0x6e, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x65, 0x72, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73,
	0x4c, 0x6f, 0x67, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x02, 0x69, 0x64, 0x12, 0x20, 0x0a, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x69, 0x64, 0x45, 0x6e, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70,
	0x65, 0x72, 0x56, 0x61, 0x72, 0x52, 0x04, 0x72, 0x6f, 0x6f, 0x6d, 0x12, 0x3a, 0x0a, 0x0f, 0x65,
	0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c,
	0x70, 0x65, 0x72, 0x56, 0x61, 0x72, 0x52, 0x0f, 0x65, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x31, 0x0a, 0x0a, 0x61, 0x63, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x69, 0x6e, 0x67, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x73, 0x69,
	0x6d, 0x72, 0x73, 0x2e, 0x41, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x52, 0x0a,
	0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x69, 0x6e, 0x67, 0x22, 0x71, 0x0a, 0x17, 0x45, 0x6e,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65, 0x72, 0x50, 0x65, 0x72, 0x73, 0x6f, 0x6e, 0x49, 0x6e, 0x43,
	0x68, 0x61, 0x72, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x68, 0x6f,
	0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x12,
	0x2c, 0x0a, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x10, 0x2e, 0x73, 0x69, 0x6d, 0x72, 0x73, 0x2e, 0x48, 0x65, 0x6c, 0x70, 0x65, 0x72,
	0x56, 0x61, 0x72, 0x52, 0x08, 0x72, 0x65, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x73, 0x0a,
	0x1f, 0x47, 0x65, 0x6e, 0x65, 0x72, 0x69, 0x63, 0x45, 0x6e, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x65,
	0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x12, 0x22,
	0x0a, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0c, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x4d, 0x65, 0x73, 0x73, 0x61,
	0x67, 0x65, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_encounter_proto_rawDescOnce sync.Once
	file_encounter_proto_rawDescData = file_encounter_proto_rawDesc
)

func file_encounter_proto_rawDescGZIP() []byte {
	file_encounter_proto_rawDescOnce.Do(func() {
		file_encounter_proto_rawDescData = protoimpl.X.CompressGZIP(file_encounter_proto_rawDescData)
	})
	return file_encounter_proto_rawDescData
}

var file_encounter_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_encounter_proto_goTypes = []interface{}{
	(*GetEncounterDetailsRequest)(nil),        // 0: simrs.GetEncounterDetailsRequest
	(*GetEncounterDetailsResponse)(nil),       // 1: simrs.GetEncounterDetailsResponse
	(*AddEncounterRoomHistoryRequest)(nil),    // 2: simrs.AddEncounterRoomHistoryRequest
	(*DeleteEncounterRoomHistoryRequest)(nil), // 3: simrs.DeleteEncounterRoomHistoryRequest
	(*Encounter)(nil),                         // 4: simrs.Encounter
	(*EncounterRoomHistory)(nil),              // 5: simrs.EncounterRoomHistory
	(*EncounterUpdateStatusLog)(nil),          // 6: simrs.EncounterUpdateStatusLog
	(*EncounterPersonInCharge)(nil),           // 7: simrs.EncounterPersonInCharge
	(*GenericEncounterServiceResponse)(nil),   // 8: simrs.GenericEncounterServiceResponse
	(*HelperVar)(nil),                         // 9: simrs.HelperVar
	(*Accounting)(nil),                        // 10: simrs.Accounting
}
var file_encounter_proto_depIdxs = []int32{
	8,  // 0: simrs.GetEncounterDetailsResponse.genericResponse:type_name -> simrs.GenericEncounterServiceResponse
	4,  // 1: simrs.GetEncounterDetailsResponse.encounters:type_name -> simrs.Encounter
	9,  // 2: simrs.Encounter.paymentMethod:type_name -> simrs.HelperVar
	9,  // 3: simrs.Encounter.entryMethod:type_name -> simrs.HelperVar
	9,  // 4: simrs.Encounter.encounterStatus:type_name -> simrs.HelperVar
	9,  // 5: simrs.Encounter.encounterType:type_name -> simrs.HelperVar
	5,  // 6: simrs.Encounter.roomHistory:type_name -> simrs.EncounterRoomHistory
	6,  // 7: simrs.Encounter.statusLog:type_name -> simrs.EncounterUpdateStatusLog
	7,  // 8: simrs.Encounter.pic:type_name -> simrs.EncounterPersonInCharge
	10, // 9: simrs.Encounter.accounting:type_name -> simrs.Accounting
	9,  // 10: simrs.EncounterRoomHistory.room:type_name -> simrs.HelperVar
	10, // 11: simrs.EncounterRoomHistory.accounting:type_name -> simrs.Accounting
	9,  // 12: simrs.EncounterUpdateStatusLog.room:type_name -> simrs.HelperVar
	9,  // 13: simrs.EncounterUpdateStatusLog.encounterStatus:type_name -> simrs.HelperVar
	10, // 14: simrs.EncounterUpdateStatusLog.accounting:type_name -> simrs.Accounting
	9,  // 15: simrs.EncounterPersonInCharge.relation:type_name -> simrs.HelperVar
	16, // [16:16] is the sub-list for method output_type
	16, // [16:16] is the sub-list for method input_type
	16, // [16:16] is the sub-list for extension type_name
	16, // [16:16] is the sub-list for extension extendee
	0,  // [0:16] is the sub-list for field type_name
}

func init() { file_encounter_proto_init() }
func file_encounter_proto_init() {
	if File_encounter_proto != nil {
		return
	}
	file_helper_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_encounter_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEncounterDetailsRequest); i {
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
		file_encounter_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetEncounterDetailsResponse); i {
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
		file_encounter_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddEncounterRoomHistoryRequest); i {
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
		file_encounter_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteEncounterRoomHistoryRequest); i {
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
		file_encounter_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Encounter); i {
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
		file_encounter_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncounterRoomHistory); i {
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
		file_encounter_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncounterUpdateStatusLog); i {
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
		file_encounter_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EncounterPersonInCharge); i {
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
		file_encounter_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GenericEncounterServiceResponse); i {
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
			RawDescriptor: file_encounter_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_encounter_proto_goTypes,
		DependencyIndexes: file_encounter_proto_depIdxs,
		MessageInfos:      file_encounter_proto_msgTypes,
	}.Build()
	File_encounter_proto = out.File
	file_encounter_proto_rawDesc = nil
	file_encounter_proto_goTypes = nil
	file_encounter_proto_depIdxs = nil
}
