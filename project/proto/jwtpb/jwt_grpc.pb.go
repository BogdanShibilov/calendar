// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v5.27.1
// source: jwt.proto

package jwtpb

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

// JwtServiceClient is the client API for JwtService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type JwtServiceClient interface {
	GenerateTokens(ctx context.Context, in *GenerateTokensRequest, opts ...grpc.CallOption) (*GenerateTokensResponse, error)
	RefreshTokens(ctx context.Context, in *RefreshTokensRequest, opts ...grpc.CallOption) (*RefreshTokensResponse, error)
	RemoveAllTokensForUser(ctx context.Context, in *RemoveAllTokensForUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error)
	ParseAccessToken(ctx context.Context, in *ParseAccessTokenRequest, opts ...grpc.CallOption) (*ParseAccessTokenResponse, error)
}

type jwtServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewJwtServiceClient(cc grpc.ClientConnInterface) JwtServiceClient {
	return &jwtServiceClient{cc}
}

func (c *jwtServiceClient) GenerateTokens(ctx context.Context, in *GenerateTokensRequest, opts ...grpc.CallOption) (*GenerateTokensResponse, error) {
	out := new(GenerateTokensResponse)
	err := c.cc.Invoke(ctx, "/jwt.JwtService/GenerateTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtServiceClient) RefreshTokens(ctx context.Context, in *RefreshTokensRequest, opts ...grpc.CallOption) (*RefreshTokensResponse, error) {
	out := new(RefreshTokensResponse)
	err := c.cc.Invoke(ctx, "/jwt.JwtService/RefreshTokens", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtServiceClient) RemoveAllTokensForUser(ctx context.Context, in *RemoveAllTokensForUserRequest, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, "/jwt.JwtService/RemoveAllTokensForUser", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *jwtServiceClient) ParseAccessToken(ctx context.Context, in *ParseAccessTokenRequest, opts ...grpc.CallOption) (*ParseAccessTokenResponse, error) {
	out := new(ParseAccessTokenResponse)
	err := c.cc.Invoke(ctx, "/jwt.JwtService/ParseAccessToken", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// JwtServiceServer is the server API for JwtService service.
// All implementations must embed UnimplementedJwtServiceServer
// for forward compatibility
type JwtServiceServer interface {
	GenerateTokens(context.Context, *GenerateTokensRequest) (*GenerateTokensResponse, error)
	RefreshTokens(context.Context, *RefreshTokensRequest) (*RefreshTokensResponse, error)
	RemoveAllTokensForUser(context.Context, *RemoveAllTokensForUserRequest) (*emptypb.Empty, error)
	ParseAccessToken(context.Context, *ParseAccessTokenRequest) (*ParseAccessTokenResponse, error)
	mustEmbedUnimplementedJwtServiceServer()
}

// UnimplementedJwtServiceServer must be embedded to have forward compatible implementations.
type UnimplementedJwtServiceServer struct {
}

func (UnimplementedJwtServiceServer) GenerateTokens(context.Context, *GenerateTokensRequest) (*GenerateTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GenerateTokens not implemented")
}
func (UnimplementedJwtServiceServer) RefreshTokens(context.Context, *RefreshTokensRequest) (*RefreshTokensResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshTokens not implemented")
}
func (UnimplementedJwtServiceServer) RemoveAllTokensForUser(context.Context, *RemoveAllTokensForUserRequest) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RemoveAllTokensForUser not implemented")
}
func (UnimplementedJwtServiceServer) ParseAccessToken(context.Context, *ParseAccessTokenRequest) (*ParseAccessTokenResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseAccessToken not implemented")
}
func (UnimplementedJwtServiceServer) mustEmbedUnimplementedJwtServiceServer() {}

// UnsafeJwtServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to JwtServiceServer will
// result in compilation errors.
type UnsafeJwtServiceServer interface {
	mustEmbedUnimplementedJwtServiceServer()
}

func RegisterJwtServiceServer(s grpc.ServiceRegistrar, srv JwtServiceServer) {
	s.RegisterService(&JwtService_ServiceDesc, srv)
}

func _JwtService_GenerateTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GenerateTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).GenerateTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwt.JwtService/GenerateTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).GenerateTokens(ctx, req.(*GenerateTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JwtService_RefreshTokens_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshTokensRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).RefreshTokens(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwt.JwtService/RefreshTokens",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).RefreshTokens(ctx, req.(*RefreshTokensRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JwtService_RemoveAllTokensForUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RemoveAllTokensForUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).RemoveAllTokensForUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwt.JwtService/RemoveAllTokensForUser",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).RemoveAllTokensForUser(ctx, req.(*RemoveAllTokensForUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _JwtService_ParseAccessToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ParseAccessTokenRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(JwtServiceServer).ParseAccessToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/jwt.JwtService/ParseAccessToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(JwtServiceServer).ParseAccessToken(ctx, req.(*ParseAccessTokenRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// JwtService_ServiceDesc is the grpc.ServiceDesc for JwtService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var JwtService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "jwt.JwtService",
	HandlerType: (*JwtServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GenerateTokens",
			Handler:    _JwtService_GenerateTokens_Handler,
		},
		{
			MethodName: "RefreshTokens",
			Handler:    _JwtService_RefreshTokens_Handler,
		},
		{
			MethodName: "RemoveAllTokensForUser",
			Handler:    _JwtService_RemoveAllTokensForUser_Handler,
		},
		{
			MethodName: "ParseAccessToken",
			Handler:    _JwtService_ParseAccessToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "jwt.proto",
}
