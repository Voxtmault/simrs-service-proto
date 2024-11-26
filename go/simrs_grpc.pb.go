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

// EncounterServiceServer is the server API for EncounterService service.
// All implementations must embed UnimplementedEncounterServiceServer
// for forward compatibility
type EncounterServiceServer interface {
	GetEncounterDetails(context.Context, *GetEncounterDetailsRequest) (*GetEncounterDetailsResponse, error)
	mustEmbedUnimplementedEncounterServiceServer()
}

// UnimplementedEncounterServiceServer must be embedded to have forward compatible implementations.
type UnimplementedEncounterServiceServer struct {
}

func (UnimplementedEncounterServiceServer) GetEncounterDetails(context.Context, *GetEncounterDetailsRequest) (*GetEncounterDetailsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetEncounterDetails not implemented")
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