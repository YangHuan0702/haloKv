// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.21.9
// source: src/pb/server.proto

package pb

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

// RaftRpcClient is the client API for RaftRpc service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RaftRpcClient interface {
	VoteRequest(ctx context.Context, in *RequestVote, opts ...grpc.CallOption) (*ResponseVote, error)
	AppendEntries(ctx context.Context, in *RequestAppendEntries, opts ...grpc.CallOption) (*ResponseAppendEntries, error)
}

type raftRpcClient struct {
	cc grpc.ClientConnInterface
}

func NewRaftRpcClient(cc grpc.ClientConnInterface) RaftRpcClient {
	return &raftRpcClient{cc}
}

func (c *raftRpcClient) VoteRequest(ctx context.Context, in *RequestVote, opts ...grpc.CallOption) (*ResponseVote, error) {
	out := new(ResponseVote)
	err := c.cc.Invoke(ctx, "/RaftRpc/VoteRequest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *raftRpcClient) AppendEntries(ctx context.Context, in *RequestAppendEntries, opts ...grpc.CallOption) (*ResponseAppendEntries, error) {
	out := new(ResponseAppendEntries)
	err := c.cc.Invoke(ctx, "/RaftRpc/AppendEntries", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RaftRpcServer is the server API for RaftRpc service.
// All implementations must embed UnimplementedRaftRpcServer
// for forward compatibility
type RaftRpcServer interface {
	VoteRequest(context.Context, *RequestVote) (*ResponseVote, error)
	AppendEntries(context.Context, *RequestAppendEntries) (*ResponseAppendEntries, error)
	mustEmbedUnimplementedRaftRpcServer()
}

// UnimplementedRaftRpcServer must be embedded to have forward compatible implementations.
type UnimplementedRaftRpcServer struct {
}

func (UnimplementedRaftRpcServer) VoteRequest(context.Context, *RequestVote) (*ResponseVote, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VoteRequest not implemented")
}
func (UnimplementedRaftRpcServer) AppendEntries(context.Context, *RequestAppendEntries) (*ResponseAppendEntries, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AppendEntries not implemented")
}
func (UnimplementedRaftRpcServer) mustEmbedUnimplementedRaftRpcServer() {}

// UnsafeRaftRpcServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RaftRpcServer will
// result in compilation errors.
type UnsafeRaftRpcServer interface {
	mustEmbedUnimplementedRaftRpcServer()
}

func RegisterRaftRpcServer(s grpc.ServiceRegistrar, srv RaftRpcServer) {
	s.RegisterService(&RaftRpc_ServiceDesc, srv)
}

func _RaftRpc_VoteRequest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestVote)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftRpcServer).VoteRequest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RaftRpc/VoteRequest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftRpcServer).VoteRequest(ctx, req.(*RequestVote))
	}
	return interceptor(ctx, in, info, handler)
}

func _RaftRpc_AppendEntries_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestAppendEntries)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RaftRpcServer).AppendEntries(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/RaftRpc/AppendEntries",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RaftRpcServer).AppendEntries(ctx, req.(*RequestAppendEntries))
	}
	return interceptor(ctx, in, info, handler)
}

// RaftRpc_ServiceDesc is the grpc.ServiceDesc for RaftRpc service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RaftRpc_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "RaftRpc",
	HandlerType: (*RaftRpcServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VoteRequest",
			Handler:    _RaftRpc_VoteRequest_Handler,
		},
		{
			MethodName: "AppendEntries",
			Handler:    _RaftRpc_AppendEntries_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "src/pb/server.proto",
}
