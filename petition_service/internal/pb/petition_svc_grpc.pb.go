// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.24.2
// source: petition_svc.proto

package pb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	PetitionService_GetPetitionById_FullMethodName        = "/proto.PetitionService/GetPetitionById"
	PetitionService_CreatePetition_FullMethodName         = "/proto.PetitionService/CreatePetition"
	PetitionService_GetPetitions_FullMethodName           = "/proto.PetitionService/GetPetitions"
	PetitionService_UpdatePetitionStatus_FullMethodName   = "/proto.PetitionService/UpdatePetitionStatus"
	PetitionService_DeletePetition_FullMethodName         = "/proto.PetitionService/DeletePetition"
	PetitionService_ValidatePetitionId_FullMethodName     = "/proto.PetitionService/ValidatePetitionId"
	PetitionService_CreateVote_FullMethodName             = "/proto.PetitionService/CreateVote"
	PetitionService_GetUserPetitions_FullMethodName       = "/proto.PetitionService/GetUserPetitions"
	PetitionService_GetUserVotedPetitions_FullMethodName  = "/proto.PetitionService/GetUserVotedPetitions"
	PetitionService_CheckIfPetitionExpired_FullMethodName = "/proto.PetitionService/CheckIfPetitionExpired"
)

// PetitionServiceClient is the client API for PetitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetitionServiceClient interface {
	GetPetitionById(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*Petition, error)
	CreatePetition(ctx context.Context, in *CreatePetitionRequest, opts ...grpc.CallOption) (*PetitionId, error)
	GetPetitions(ctx context.Context, in *GetPetitionsRequest, opts ...grpc.CallOption) (*GetPetitionsResponse, error)
	UpdatePetitionStatus(ctx context.Context, in *UpdatePetitionStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	DeletePetition(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	// returns empty if is valid, error if not valid
	ValidatePetitionId(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*emptypb.Empty, error)
	CreateVote(ctx context.Context, in *CreateVoteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetUserPetitions(ctx context.Context, in *GetUserPetitionsRequest, opts ...grpc.CallOption) (*GetUserPetitionsResponse, error)
	GetUserVotedPetitions(ctx context.Context, in *GetUserVotedPetitionsRequest, opts ...grpc.CallOption) (*GetUserVotedPetitionsResponse, error)
	CheckIfPetitionExpired(ctx context.Context, in *Petition, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type petitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetitionServiceClient(cc grpc.ClientConnInterface) PetitionServiceClient {
	return &petitionServiceClient{cc}
}

func (c *petitionServiceClient) GetPetitionById(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*Petition, error) {
	out := new(Petition)
	err := c.cc.Invoke(ctx, PetitionService_GetPetitionById_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) CreatePetition(ctx context.Context, in *CreatePetitionRequest, opts ...grpc.CallOption) (*PetitionId, error) {
	out := new(PetitionId)
	err := c.cc.Invoke(ctx, PetitionService_CreatePetition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) GetPetitions(ctx context.Context, in *GetPetitionsRequest, opts ...grpc.CallOption) (*GetPetitionsResponse, error) {
	out := new(GetPetitionsResponse)
	err := c.cc.Invoke(ctx, PetitionService_GetPetitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) UpdatePetitionStatus(ctx context.Context, in *UpdatePetitionStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PetitionService_UpdatePetitionStatus_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) DeletePetition(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PetitionService_DeletePetition_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) ValidatePetitionId(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PetitionService_ValidatePetitionId_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) CreateVote(ctx context.Context, in *CreateVoteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PetitionService_CreateVote_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) GetUserPetitions(ctx context.Context, in *GetUserPetitionsRequest, opts ...grpc.CallOption) (*GetUserPetitionsResponse, error) {
	out := new(GetUserPetitionsResponse)
	err := c.cc.Invoke(ctx, PetitionService_GetUserPetitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) GetUserVotedPetitions(ctx context.Context, in *GetUserVotedPetitionsRequest, opts ...grpc.CallOption) (*GetUserVotedPetitionsResponse, error) {
	out := new(GetUserVotedPetitionsResponse)
	err := c.cc.Invoke(ctx, PetitionService_GetUserVotedPetitions_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) CheckIfPetitionExpired(ctx context.Context, in *Petition, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, PetitionService_CheckIfPetitionExpired_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetitionServiceServer is the server API for PetitionService service.
// All implementations should embed UnimplementedPetitionServiceServer
// for forward compatibility
type PetitionServiceServer interface {
	GetPetitionById(context.Context, *PetitionId) (*Petition, error)
	CreatePetition(context.Context, *CreatePetitionRequest) (*PetitionId, error)
	GetPetitions(context.Context, *GetPetitionsRequest) (*GetPetitionsResponse, error)
	UpdatePetitionStatus(context.Context, *UpdatePetitionStatusRequest) (*emptypb.Empty, error)
	DeletePetition(context.Context, *PetitionId) (*emptypb.Empty, error)
	// returns empty if is valid, error if not valid
	ValidatePetitionId(context.Context, *PetitionId) (*emptypb.Empty, error)
	CreateVote(context.Context, *CreateVoteRequest) (*emptypb.Empty, error)
	GetUserPetitions(context.Context, *GetUserPetitionsRequest) (*GetUserPetitionsResponse, error)
	GetUserVotedPetitions(context.Context, *GetUserVotedPetitionsRequest) (*GetUserVotedPetitionsResponse, error)
	CheckIfPetitionExpired(context.Context, *Petition) (*emptypb.Empty, error)
}

// UnimplementedPetitionServiceServer should be embedded to have forward compatible implementations.
type UnimplementedPetitionServiceServer struct {
}

func (UnimplementedPetitionServiceServer) GetPetitionById(context.Context, *PetitionId) (*Petition, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPetitionById not implemented")
}
func (UnimplementedPetitionServiceServer) CreatePetition(context.Context, *CreatePetitionRequest) (*PetitionId, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreatePetition not implemented")
}
func (UnimplementedPetitionServiceServer) GetPetitions(context.Context, *GetPetitionsRequest) (*GetPetitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPetitions not implemented")
}
func (UnimplementedPetitionServiceServer) UpdatePetitionStatus(context.Context, *UpdatePetitionStatusRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePetitionStatus not implemented")
}
func (UnimplementedPetitionServiceServer) DeletePetition(context.Context, *PetitionId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePetition not implemented")
}
func (UnimplementedPetitionServiceServer) ValidatePetitionId(context.Context, *PetitionId) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePetitionId not implemented")
}
func (UnimplementedPetitionServiceServer) CreateVote(context.Context, *CreateVoteRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVote not implemented")
}
func (UnimplementedPetitionServiceServer) GetUserPetitions(context.Context, *GetUserPetitionsRequest) (*GetUserPetitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPetitions not implemented")
}
func (UnimplementedPetitionServiceServer) GetUserVotedPetitions(context.Context, *GetUserVotedPetitionsRequest) (*GetUserVotedPetitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserVotedPetitions not implemented")
}
func (UnimplementedPetitionServiceServer) CheckIfPetitionExpired(context.Context, *Petition) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfPetitionExpired not implemented")
}

// UnsafePetitionServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PetitionServiceServer will
// result in compilation errors.
type UnsafePetitionServiceServer interface {
	mustEmbedUnimplementedPetitionServiceServer()
}

func RegisterPetitionServiceServer(s grpc.ServiceRegistrar, srv PetitionServiceServer) {
	s.RegisterService(&PetitionService_ServiceDesc, srv)
}

func _PetitionService_GetPetitionById_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetitionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).GetPetitionById(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_GetPetitionById_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).GetPetitionById(ctx, req.(*PetitionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_CreatePetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreatePetitionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).CreatePetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_CreatePetition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).CreatePetition(ctx, req.(*CreatePetitionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_GetPetitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetPetitionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).GetPetitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_GetPetitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).GetPetitions(ctx, req.(*GetPetitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_UpdatePetitionStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(UpdatePetitionStatusRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).UpdatePetitionStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_UpdatePetitionStatus_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).UpdatePetitionStatus(ctx, req.(*UpdatePetitionStatusRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_DeletePetition_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetitionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).DeletePetition(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_DeletePetition_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).DeletePetition(ctx, req.(*PetitionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_ValidatePetitionId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetitionId)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).ValidatePetitionId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_ValidatePetitionId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).ValidatePetitionId(ctx, req.(*PetitionId))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_CreateVote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateVoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).CreateVote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_CreateVote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).CreateVote(ctx, req.(*CreateVoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_GetUserPetitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserPetitionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).GetUserPetitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_GetUserPetitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).GetUserPetitions(ctx, req.(*GetUserPetitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_GetUserVotedPetitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetUserVotedPetitionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).GetUserVotedPetitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_GetUserVotedPetitions_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).GetUserVotedPetitions(ctx, req.(*GetUserVotedPetitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_CheckIfPetitionExpired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Petition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).CheckIfPetitionExpired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: PetitionService_CheckIfPetitionExpired_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).CheckIfPetitionExpired(ctx, req.(*Petition))
	}
	return interceptor(ctx, in, info, handler)
}

// PetitionService_ServiceDesc is the grpc.ServiceDesc for PetitionService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var PetitionService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "proto.PetitionService",
	HandlerType: (*PetitionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetPetitionById",
			Handler:    _PetitionService_GetPetitionById_Handler,
		},
		{
			MethodName: "CreatePetition",
			Handler:    _PetitionService_CreatePetition_Handler,
		},
		{
			MethodName: "GetPetitions",
			Handler:    _PetitionService_GetPetitions_Handler,
		},
		{
			MethodName: "UpdatePetitionStatus",
			Handler:    _PetitionService_UpdatePetitionStatus_Handler,
		},
		{
			MethodName: "DeletePetition",
			Handler:    _PetitionService_DeletePetition_Handler,
		},
		{
			MethodName: "ValidatePetitionId",
			Handler:    _PetitionService_ValidatePetitionId_Handler,
		},
		{
			MethodName: "CreateVote",
			Handler:    _PetitionService_CreateVote_Handler,
		},
		{
			MethodName: "GetUserPetitions",
			Handler:    _PetitionService_GetUserPetitions_Handler,
		},
		{
			MethodName: "GetUserVotedPetitions",
			Handler:    _PetitionService_GetUserVotedPetitions_Handler,
		},
		{
			MethodName: "CheckIfPetitionExpired",
			Handler:    _PetitionService_CheckIfPetitionExpired_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petition_svc.proto",
}
