// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.12
// source: simrs.proto

package __

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PatientServiceClient is the client API for PatientService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PatientServiceClient interface {
	GetPatients(ctx context.Context, in *GetPatientsRequest, opts ...grpc.CallOption) (*GetPatientsResponse, error)
}

type patientServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPatientServiceClient(cc grpc.ClientConnInterface) PatientServiceClient {
	return &patientServiceClient{cc}
}

func (c *patientServiceClient) GetPatients(ctx context.Context, in *GetPatientsRequest, opts ...grpc.CallOption) (*GetPatientsResponse, error) {
	out := new(GetPatientsResponse)
	err := c.cc.Invoke(ctx, "/simrs.PatientService/GetPatients", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PatientServiceServer is the server API for PatientService service.
// All implementations must embed UnimplementedPatientServiceServer
// for forward compatibility
type PatientServiceServer interface {
	GetPatients(context.Context, *GetPatientsRequest) (*GetPatientsResponse, error)
	mustEmbedUnimplementedPatientServiceServer()
}

// UnimplementedPatientServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPatientServiceServer struct {
}

func (UnimplementedPatientServiceServer) GetPatients(context.Context, *GetPatientsRequest) (*GetPatientsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPatients not implemented")
}
func (UnimplementedPatientServiceServer) mustEmbedUnimplementedPatientServiceServer() {}

// UnsafePatientServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PatientServiceServer will
// result in compilation errors.
type UnsafePatientServiceServer interface {
	mustEmbedUnimplementedPatientServiceServer()
}

func RegisterPatientServiceServer(s grpc.ServiceRegistrar, srv PatientServiceServer) {
	s.RegisterService(&PatientService_ServiceDesc, srv)
}

func _PatientService_GetPatients_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPatientsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PatientServiceServer).GetPatients(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.PatientService/GetPatients",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PatientServiceServer).GetPatients(ctx, req.(*GetPatientsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PatientService_ServiceDesc is the grpc.ServiceDesc for PatientService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PatientService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simrs.PatientService",
	HandlerType: (*PatientServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPatients",
			Handler:    _PatientService_GetPatients_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simrs.proto",
}

// EncounterServiceClient is the client API for EncounterService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EncounterServiceClient interface {
	GetEncounterDetails(ctx context.Context, in *GetEncounterDetailsRequest, opts ...grpc.CallOption) (*GetEncounterDetailsResponse, error)
	AddEncounterRoomHistory(ctx context.Context, in *AddEncounterRoomHistoryRequest, opts ...grpc.CallOption) (*GenericEncounterServiceResponse, error)
	DeleteEncounterRoomHistory(ctx context.Context, in *DeleteEncounterRoomHistoryRequest, opts ...grpc.CallOption) (*GenericEncounterServiceResponse, error)
}

type encounterServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEncounterServiceClient(cc grpc.ClientConnInterface) EncounterServiceClient {
	return &encounterServiceClient{cc}
}

func (c *encounterServiceClient) GetEncounterDetails(ctx context.Context, in *GetEncounterDetailsRequest, opts ...grpc.CallOption) (*GetEncounterDetailsResponse, error) {
	out := new(GetEncounterDetailsResponse)
	err := c.cc.Invoke(ctx, "/simrs.EncounterService/GetEncounterDetails", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterServiceClient) AddEncounterRoomHistory(ctx context.Context, in *AddEncounterRoomHistoryRequest, opts ...grpc.CallOption) (*GenericEncounterServiceResponse, error) {
	out := new(GenericEncounterServiceResponse)
	err := c.cc.Invoke(ctx, "/simrs.EncounterService/AddEncounterRoomHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *encounterServiceClient) DeleteEncounterRoomHistory(ctx context.Context, in *DeleteEncounterRoomHistoryRequest, opts ...grpc.CallOption) (*GenericEncounterServiceResponse, error) {
	out := new(GenericEncounterServiceResponse)
	err := c.cc.Invoke(ctx, "/simrs.EncounterService/DeleteEncounterRoomHistory", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EncounterServiceServer is the server API for EncounterService service.
// All implementations must embed UnimplementedEncounterServiceServer
// for forward compatibility
type EncounterServiceServer interface {
	GetEncounterDetails(context.Context, *GetEncounterDetailsRequest) (*GetEncounterDetailsResponse, error)
	AddEncounterRoomHistory(context.Context, *AddEncounterRoomHistoryRequest) (*GenericEncounterServiceResponse, error)
	DeleteEncounterRoomHistory(context.Context, *DeleteEncounterRoomHistoryRequest) (*GenericEncounterServiceResponse, error)
	mustEmbedUnimplementedEncounterServiceServer()
}

// UnimplementedEncounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEncounterServiceServer struct {
}

func (UnimplementedEncounterServiceServer) GetEncounterDetails(context.Context, *GetEncounterDetailsRequest) (*GetEncounterDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncounterDetails not implemented")
}
func (UnimplementedEncounterServiceServer) AddEncounterRoomHistory(context.Context, *AddEncounterRoomHistoryRequest) (*GenericEncounterServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddEncounterRoomHistory not implemented")
}
func (UnimplementedEncounterServiceServer) DeleteEncounterRoomHistory(context.Context, *DeleteEncounterRoomHistoryRequest) (*GenericEncounterServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteEncounterRoomHistory not implemented")
}
func (UnimplementedEncounterServiceServer) mustEmbedUnimplementedEncounterServiceServer() {}

// UnsafeEncounterServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EncounterServiceServer will
// result in compilation errors.
type UnsafeEncounterServiceServer interface {
	mustEmbedUnimplementedEncounterServiceServer()
}

func RegisterEncounterServiceServer(s grpc.ServiceRegistrar, srv EncounterServiceServer) {
	s.RegisterService(&EncounterService_ServiceDesc, srv)
}

func _EncounterService_GetEncounterDetails_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetEncounterDetailsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).GetEncounterDetails(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.EncounterService/GetEncounterDetails",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).GetEncounterDetails(ctx, req.(*GetEncounterDetailsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterService_AddEncounterRoomHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddEncounterRoomHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).AddEncounterRoomHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.EncounterService/AddEncounterRoomHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).AddEncounterRoomHistory(ctx, req.(*AddEncounterRoomHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EncounterService_DeleteEncounterRoomHistory_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeleteEncounterRoomHistoryRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EncounterServiceServer).DeleteEncounterRoomHistory(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.EncounterService/DeleteEncounterRoomHistory",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EncounterServiceServer).DeleteEncounterRoomHistory(ctx, req.(*DeleteEncounterRoomHistoryRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EncounterService_ServiceDesc is the grpc.ServiceDesc for EncounterService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EncounterService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simrs.EncounterService",
	HandlerType: (*EncounterServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetEncounterDetails",
			Handler:    _EncounterService_GetEncounterDetails_Handler,
		},
		{
			MethodName: "AddEncounterRoomHistory",
			Handler:    _EncounterService_AddEncounterRoomHistory_Handler,
		},
		{
			MethodName: "DeleteEncounterRoomHistory",
			Handler:    _EncounterService_DeleteEncounterRoomHistory_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simrs.proto",
}

// QueueServiceClient is the client API for QueueService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type QueueServiceClient interface {
	UpdateQueue(ctx context.Context, in *UpdateQueueRequest, opts ...grpc.CallOption) (*UpdateQueueResponse, error)
	CompensateUpdateQueueEncounter(ctx context.Context, in *CompensateUpdateQueueEncounterRequest, opts ...grpc.CallOption) (*CompensateUpdateQueueEncounterResponse, error)
}

type queueServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewQueueServiceClient(cc grpc.ClientConnInterface) QueueServiceClient {
	return &queueServiceClient{cc}
}

func (c *queueServiceClient) UpdateQueue(ctx context.Context, in *UpdateQueueRequest, opts ...grpc.CallOption) (*UpdateQueueResponse, error) {
	out := new(UpdateQueueResponse)
	err := c.cc.Invoke(ctx, "/simrs.QueueService/UpdateQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queueServiceClient) CompensateUpdateQueueEncounter(ctx context.Context, in *CompensateUpdateQueueEncounterRequest, opts ...grpc.CallOption) (*CompensateUpdateQueueEncounterResponse, error) {
	out := new(CompensateUpdateQueueEncounterResponse)
	err := c.cc.Invoke(ctx, "/simrs.QueueService/CompensateUpdateQueueEncounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueueServiceServer is the server API for QueueService service.
// All implementations must embed UnimplementedQueueServiceServer
// for forward compatibility
type QueueServiceServer interface {
	UpdateQueue(context.Context, *UpdateQueueRequest) (*UpdateQueueResponse, error)
	CompensateUpdateQueueEncounter(context.Context, *CompensateUpdateQueueEncounterRequest) (*CompensateUpdateQueueEncounterResponse, error)
	mustEmbedUnimplementedQueueServiceServer()
}

// UnimplementedQueueServiceServer must be embedded to have forward compatible implementations.
type UnimplementedQueueServiceServer struct {
}

func (UnimplementedQueueServiceServer) UpdateQueue(context.Context, *UpdateQueueRequest) (*UpdateQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateQueue not implemented")
}
func (UnimplementedQueueServiceServer) CompensateUpdateQueueEncounter(context.Context, *CompensateUpdateQueueEncounterRequest) (*CompensateUpdateQueueEncounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompensateUpdateQueueEncounter not implemented")
}
func (UnimplementedQueueServiceServer) mustEmbedUnimplementedQueueServiceServer() {}

// UnsafeQueueServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to QueueServiceServer will
// result in compilation errors.
type UnsafeQueueServiceServer interface {
	mustEmbedUnimplementedQueueServiceServer()
}

func RegisterQueueServiceServer(s grpc.ServiceRegistrar, srv QueueServiceServer) {
	s.RegisterService(&QueueService_ServiceDesc, srv)
}

func _QueueService_UpdateQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdateQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).UpdateQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.QueueService/UpdateQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).UpdateQueue(ctx, req.(*UpdateQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _QueueService_CompensateUpdateQueueEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompensateUpdateQueueEncounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueueServiceServer).CompensateUpdateQueueEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.QueueService/CompensateUpdateQueueEncounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueueServiceServer).CompensateUpdateQueueEncounter(ctx, req.(*CompensateUpdateQueueEncounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// QueueService_ServiceDesc is the grpc.ServiceDesc for QueueService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var QueueService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simrs.QueueService",
	HandlerType: (*QueueServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "UpdateQueue",
			Handler:    _QueueService_UpdateQueue_Handler,
		},
		{
			MethodName: "CompensateUpdateQueueEncounter",
			Handler:    _QueueService_CompensateUpdateQueueEncounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simrs.proto",
}

// PolyclinicServiceClient is the client API for PolyclinicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PolyclinicServiceClient interface {
	CreatePolyclinicQueue(ctx context.Context, in *CreatePolyclinicQueueRequest, opts ...grpc.CallOption) (*CreatePolyclinicQueueResponse, error)
	CompensateCreatePolyQueueEncounter(ctx context.Context, in *CompensateCreatePolyQueueEncounterRequest, opts ...grpc.CallOption) (*CompensateCreatePolyQueueEncounterResponse, error)
}

type polyclinicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPolyclinicServiceClient(cc grpc.ClientConnInterface) PolyclinicServiceClient {
	return &polyclinicServiceClient{cc}
}

func (c *polyclinicServiceClient) CreatePolyclinicQueue(ctx context.Context, in *CreatePolyclinicQueueRequest, opts ...grpc.CallOption) (*CreatePolyclinicQueueResponse, error) {
	out := new(CreatePolyclinicQueueResponse)
	err := c.cc.Invoke(ctx, "/simrs.PolyclinicService/CreatePolyclinicQueue", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *polyclinicServiceClient) CompensateCreatePolyQueueEncounter(ctx context.Context, in *CompensateCreatePolyQueueEncounterRequest, opts ...grpc.CallOption) (*CompensateCreatePolyQueueEncounterResponse, error) {
	out := new(CompensateCreatePolyQueueEncounterResponse)
	err := c.cc.Invoke(ctx, "/simrs.PolyclinicService/CompensateCreatePolyQueueEncounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PolyclinicServiceServer is the server API for PolyclinicService service.
// All implementations must embed UnimplementedPolyclinicServiceServer
// for forward compatibility
type PolyclinicServiceServer interface {
	CreatePolyclinicQueue(context.Context, *CreatePolyclinicQueueRequest) (*CreatePolyclinicQueueResponse, error)
	CompensateCreatePolyQueueEncounter(context.Context, *CompensateCreatePolyQueueEncounterRequest) (*CompensateCreatePolyQueueEncounterResponse, error)
	mustEmbedUnimplementedPolyclinicServiceServer()
}

// UnimplementedPolyclinicServiceServer must be embedded to have forward compatible implementations.
type UnimplementedPolyclinicServiceServer struct {
}

func (UnimplementedPolyclinicServiceServer) CreatePolyclinicQueue(context.Context, *CreatePolyclinicQueueRequest) (*CreatePolyclinicQueueResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePolyclinicQueue not implemented")
}
func (UnimplementedPolyclinicServiceServer) CompensateCreatePolyQueueEncounter(context.Context, *CompensateCreatePolyQueueEncounterRequest) (*CompensateCreatePolyQueueEncounterResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompensateCreatePolyQueueEncounter not implemented")
}
func (UnimplementedPolyclinicServiceServer) mustEmbedUnimplementedPolyclinicServiceServer() {}

// UnsafePolyclinicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PolyclinicServiceServer will
// result in compilation errors.
type UnsafePolyclinicServiceServer interface {
	mustEmbedUnimplementedPolyclinicServiceServer()
}

func RegisterPolyclinicServiceServer(s grpc.ServiceRegistrar, srv PolyclinicServiceServer) {
	s.RegisterService(&PolyclinicService_ServiceDesc, srv)
}

func _PolyclinicService_CreatePolyclinicQueue_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePolyclinicQueueRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolyclinicServiceServer).CreatePolyclinicQueue(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.PolyclinicService/CreatePolyclinicQueue",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolyclinicServiceServer).CreatePolyclinicQueue(ctx, req.(*CreatePolyclinicQueueRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PolyclinicService_CompensateCreatePolyQueueEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompensateCreatePolyQueueEncounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PolyclinicServiceServer).CompensateCreatePolyQueueEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.PolyclinicService/CompensateCreatePolyQueueEncounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PolyclinicServiceServer).CompensateCreatePolyQueueEncounter(ctx, req.(*CompensateCreatePolyQueueEncounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// PolyclinicService_ServiceDesc is the grpc.ServiceDesc for PolyclinicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PolyclinicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simrs.PolyclinicService",
	HandlerType: (*PolyclinicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePolyclinicQueue",
			Handler:    _PolyclinicService_CreatePolyclinicQueue_Handler,
		},
		{
			MethodName: "CompensateCreatePolyQueueEncounter",
			Handler:    _PolyclinicService_CompensateCreatePolyQueueEncounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simrs.proto",
}

// EmergencyRoomServiceClient is the client API for EmergencyRoomService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type EmergencyRoomServiceClient interface {
	AddERPatient(ctx context.Context, in *AddERPatientRequest, opts ...grpc.CallOption) (*GenericERServiceResponse, error)
	CompensateAddERPatientEncounter(ctx context.Context, in *CompensateAddERPatientEncounterRequest, opts ...grpc.CallOption) (*GenericERServiceResponse, error)
}

type emergencyRoomServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewEmergencyRoomServiceClient(cc grpc.ClientConnInterface) EmergencyRoomServiceClient {
	return &emergencyRoomServiceClient{cc}
}

func (c *emergencyRoomServiceClient) AddERPatient(ctx context.Context, in *AddERPatientRequest, opts ...grpc.CallOption) (*GenericERServiceResponse, error) {
	out := new(GenericERServiceResponse)
	err := c.cc.Invoke(ctx, "/simrs.EmergencyRoomService/AddERPatient", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *emergencyRoomServiceClient) CompensateAddERPatientEncounter(ctx context.Context, in *CompensateAddERPatientEncounterRequest, opts ...grpc.CallOption) (*GenericERServiceResponse, error) {
	out := new(GenericERServiceResponse)
	err := c.cc.Invoke(ctx, "/simrs.EmergencyRoomService/CompensateAddERPatientEncounter", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// EmergencyRoomServiceServer is the server API for EmergencyRoomService service.
// All implementations must embed UnimplementedEmergencyRoomServiceServer
// for forward compatibility
type EmergencyRoomServiceServer interface {
	AddERPatient(context.Context, *AddERPatientRequest) (*GenericERServiceResponse, error)
	CompensateAddERPatientEncounter(context.Context, *CompensateAddERPatientEncounterRequest) (*GenericERServiceResponse, error)
	mustEmbedUnimplementedEmergencyRoomServiceServer()
}

// UnimplementedEmergencyRoomServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEmergencyRoomServiceServer struct {
}

func (UnimplementedEmergencyRoomServiceServer) AddERPatient(context.Context, *AddERPatientRequest) (*GenericERServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddERPatient not implemented")
}
func (UnimplementedEmergencyRoomServiceServer) CompensateAddERPatientEncounter(context.Context, *CompensateAddERPatientEncounterRequest) (*GenericERServiceResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CompensateAddERPatientEncounter not implemented")
}
func (UnimplementedEmergencyRoomServiceServer) mustEmbedUnimplementedEmergencyRoomServiceServer() {}

// UnsafeEmergencyRoomServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to EmergencyRoomServiceServer will
// result in compilation errors.
type UnsafeEmergencyRoomServiceServer interface {
	mustEmbedUnimplementedEmergencyRoomServiceServer()
}

func RegisterEmergencyRoomServiceServer(s grpc.ServiceRegistrar, srv EmergencyRoomServiceServer) {
	s.RegisterService(&EmergencyRoomService_ServiceDesc, srv)
}

func _EmergencyRoomService_AddERPatient_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddERPatientRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyRoomServiceServer).AddERPatient(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.EmergencyRoomService/AddERPatient",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyRoomServiceServer).AddERPatient(ctx, req.(*AddERPatientRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _EmergencyRoomService_CompensateAddERPatientEncounter_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CompensateAddERPatientEncounterRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EmergencyRoomServiceServer).CompensateAddERPatientEncounter(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/simrs.EmergencyRoomService/CompensateAddERPatientEncounter",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EmergencyRoomServiceServer).CompensateAddERPatientEncounter(ctx, req.(*CompensateAddERPatientEncounterRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// EmergencyRoomService_ServiceDesc is the grpc.ServiceDesc for EmergencyRoomService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var EmergencyRoomService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "simrs.EmergencyRoomService",
	HandlerType: (*EmergencyRoomServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "AddERPatient",
			Handler:    _EmergencyRoomService_AddERPatient_Handler,
		},
		{
			MethodName: "CompensateAddERPatientEncounter",
			Handler:    _EmergencyRoomService_CompensateAddERPatientEncounter_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "simrs.proto",
}
