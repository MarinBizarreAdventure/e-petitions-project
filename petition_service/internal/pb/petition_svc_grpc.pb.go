// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.24.2
// source: petition_svc.proto

package pb

import (
	context "context"
	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// PetitionServiceClient is the client API for PetitionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type PetitionServiceClient interface {
	GetPetitionById(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*Petition, error)
	CreatePetition(ctx context.Context, in *CreatePetitionRequest, opts ...grpc.CallOption) (*PetitionId, error)
	GetPetitions(ctx context.Context, in *GetPetitionsRequest, opts ...grpc.CallOption) (*GetPetitionsResponse, error)
	UpdatePetitionStatus(ctx context.Context, in *UpdatePetitionStatusRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	DeletePetition(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*empty.Empty, error)
	// returns empty if is valid, error if not valid
	ValidatePetitionId(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*empty.Empty, error)
	CreateVote(ctx context.Context, in *CreateVoteRequest, opts ...grpc.CallOption) (*empty.Empty, error)
	GetUserPetitions(ctx context.Context, in *GetUserPetitionsRequest, opts ...grpc.CallOption) (*GetUserPetitionsResponse, error)
	GetUserVotedPetitions(ctx context.Context, in *GetUserVotedPetitionsRequest, opts ...grpc.CallOption) (*GetUserVotedPetitionsResponse, error)
	CheckIfPetitionsExpired(ctx context.Context, in *Petition, opts ...grpc.CallOption) (*emptypb.Empty, error)
	GetAllSimilarPetitions(ctx context.Context, in *PetitionSuggestionRequest, opts ...grpc.CallOption) (*PetitionSuggestionResponse, error)
	SearchPetitionsByTitle(ctx context.Context, in *SearchPetitionsByTitRequest, opts ...grpc.CallOption) (*PetitionSuggestionResponse, error)
}

type petitionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewPetitionServiceClient(cc grpc.ClientConnInterface) PetitionServiceClient {
	return &petitionServiceClient{cc}
}

func (c *petitionServiceClient) GetPetitionById(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*Petition, error) {
	out := new(Petition)
	err := c.cc.Invoke(ctx, "/proto.PetitionService/GetPetitionById", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) CreatePetition(ctx context.Context, in *CreatePetitionRequest, opts ...grpc.CallOption) (*PetitionId, error) {
	out := new(PetitionId)
	err := c.cc.Invoke(ctx, "/proto.PetitionService/CreatePetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) GetPetitions(ctx context.Context, in *GetPetitionsRequest, opts ...grpc.CallOption) (*GetPetitionsResponse, error) {
	out := new(GetPetitionsResponse)
	err := c.cc.Invoke(ctx, "/proto.PetitionService/GetPetitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}


func (c *petitionServiceClient) UpdatePetitionStatus(ctx context.Context, in *UpdatePetitionStatusRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)

	err := c.cc.Invoke(ctx, "/proto.PetitionService/UpdatePetitionStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) DeletePetition(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)

	err := c.cc.Invoke(ctx, "/proto.PetitionService/DeletePetition", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) ValidatePetitionId(ctx context.Context, in *PetitionId, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)

	err := c.cc.Invoke(ctx, "/proto.PetitionService/ValidatePetitionId", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) CreateVote(ctx context.Context, in *CreateVoteRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)

	err := c.cc.Invoke(ctx, "/proto.PetitionService/CreateVote", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) GetUserPetitions(ctx context.Context, in *GetUserPetitionsRequest, opts ...grpc.CallOption) (*GetUserPetitionsResponse, error) {
	out := new(GetUserPetitionsResponse)
	err := c.cc.Invoke(ctx, "/proto.PetitionService/GetUserPetitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) GetUserVotedPetitions(ctx context.Context, in *GetUserVotedPetitionsRequest, opts ...grpc.CallOption) (*GetUserVotedPetitionsResponse, error) {
	out := new(GetUserVotedPetitionsResponse)
	err := c.cc.Invoke(ctx, "/proto.PetitionService/GetUserVotedPetitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) CheckIfPetitionsExpired(ctx context.Context, in *Petition, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/proto.PetitionService/CheckIfPetitionsExpired", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) GetAllSimilarPetitions(ctx context.Context, in *PetitionSuggestionRequest, opts ...grpc.CallOption) (*PetitionSuggestionResponse, error) {
	out := new(PetitionSuggestionResponse)
	err := c.cc.Invoke(ctx, "/proto.PetitionService/GetAllSimilarPetitions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *petitionServiceClient) SearchPetitionsByTitle(ctx context.Context, in *SearchPetitionsByTitRequest, opts ...grpc.CallOption) (*PetitionSuggestionResponse, error) {
	out := new(PetitionSuggestionResponse)
	err := c.cc.Invoke(ctx, "/proto.PetitionService/SearchPetitionsByTitle", in, out, opts...)

	if err != nil {
		return nil, err
	}
	return out, nil
}

// PetitionServiceServer is the server API for PetitionService service.
// All implementations must embed UnimplementedPetitionServiceServer
// for forward compatibility
type PetitionServiceServer interface {
	GetPetitionById(context.Context, *PetitionId) (*Petition, error)
	CreatePetition(context.Context, *CreatePetitionRequest) (*PetitionId, error)
	GetPetitions(context.Context, *GetPetitionsRequest) (*GetPetitionsResponse, error)
	UpdatePetitionStatus(context.Context, *UpdatePetitionStatusRequest) (*empty.Empty, error)
	DeletePetition(context.Context, *PetitionId) (*empty.Empty, error)
	// returns empty if is valid, error if not valid
	ValidatePetitionId(context.Context, *PetitionId) (*empty.Empty, error)
	CreateVote(context.Context, *CreateVoteRequest) (*empty.Empty, error)
	GetUserPetitions(context.Context, *GetUserPetitionsRequest) (*GetUserPetitionsResponse, error)
	GetUserVotedPetitions(context.Context, *GetUserVotedPetitionsRequest) (*GetUserVotedPetitionsResponse, error)
	CheckIfPetitionsExpired(context.Context, *Petition) (*emptypb.Empty, error)
	GetAllSimilarPetitions(context.Context, *PetitionSuggestionRequest) (*PetitionSuggestionResponse, error)
	SearchPetitionsByTitle(context.Context, *SearchPetitionsByTitRequest) (*PetitionSuggestionResponse, error)
	mustEmbedUnimplementedPetitionServiceServer()

}

// UnimplementedPetitionServiceServer must be embedded to have forward compatible implementations.
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
func (UnimplementedPetitionServiceServer) UpdatePetitionStatus(context.Context, *UpdatePetitionStatusRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePetitionStatus not implemented")
}
func (UnimplementedPetitionServiceServer) DeletePetition(context.Context, *PetitionId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePetition not implemented")
}
func (UnimplementedPetitionServiceServer) ValidatePetitionId(context.Context, *PetitionId) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ValidatePetitionId not implemented")
}
func (UnimplementedPetitionServiceServer) CreateVote(context.Context, *CreateVoteRequest) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVote not implemented")
}
func (UnimplementedPetitionServiceServer) GetUserPetitions(context.Context, *GetUserPetitionsRequest) (*GetUserPetitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserPetitions not implemented")
}
func (UnimplementedPetitionServiceServer) GetUserVotedPetitions(context.Context, *GetUserVotedPetitionsRequest) (*GetUserVotedPetitionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetUserVotedPetitions not implemented")
}
func (UnimplementedPetitionServiceServer) CheckIfPetitionsExpired(context.Context, *Petition) (*empty.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CheckIfPetitionsExpired not implemented")
}
func (UnimplementedPetitionServiceServer) GetAllSimilarPetitions(context.Context, *PetitionSuggestionRequest) (*PetitionSuggestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllSimilarPetitions not implemented")
}
func (UnimplementedPetitionServiceServer) SearchPetitionsByTitle(context.Context, *SearchPetitionsByTitRequest) (*PetitionSuggestionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SearchPetitionsByTitle not implemented")
}
func (UnimplementedPetitionServiceServer) mustEmbedUnimplementedPetitionServiceServer() {}

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
		FullMethod: "/proto.PetitionService/GetPetitionById",
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
		FullMethod: "/proto.PetitionService/CreatePetition",
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
		FullMethod: "/proto.PetitionService/GetPetitions",
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
		FullMethod: "/proto.PetitionService/UpdatePetitionStatus",
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
		FullMethod: "/proto.PetitionService/DeletePetition",
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
		FullMethod: "/proto.PetitionService/ValidatePetitionId",
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
		FullMethod: "/proto.PetitionService/CreateVote",
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
		FullMethod: "/proto.PetitionService/GetUserPetitions",
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
		FullMethod: "/proto.PetitionService/GetUserVotedPetitions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).GetUserVotedPetitions(ctx, req.(*GetUserVotedPetitionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_CheckIfPetitionsExpired_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Petition)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).CheckIfPetitionsExpired(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PetitionService/CheckIfPetitionsExpired",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).CheckIfPetitionsExpired(ctx, req.(*Petition))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_GetAllSimilarPetitions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PetitionSuggestionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).GetAllSimilarPetitions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PetitionService/GetAllSimilarPetitions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).GetAllSimilarPetitions(ctx, req.(*PetitionSuggestionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _PetitionService_SearchPetitionsByTitle_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchPetitionsByTitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PetitionServiceServer).SearchPetitionsByTitle(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.PetitionService/SearchPetitionsByTitle",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PetitionServiceServer).SearchPetitionsByTitle(ctx, req.(*SearchPetitionsByTitRequest))
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
			MethodName: "CheckIfPetitionsExpired",
			Handler:    _PetitionService_CheckIfPetitionsExpired_Handler,
		},
		{
			MethodName: "GetAllSimilarPetitions",
			Handler:    _PetitionService_GetAllSimilarPetitions_Handler,
		},
		{
			MethodName: "SearchPetitionsByTitle",
			Handler:    _PetitionService_SearchPetitionsByTitle_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "petition_svc.proto",
}
