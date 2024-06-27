package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/jwt/server/grpc/handler/jwt/generate"
	"hwCalendar/jwt/server/grpc/handler/jwt/parseAccess"
	"hwCalendar/jwt/server/grpc/handler/jwt/refresh"
	"hwCalendar/jwt/server/grpc/handler/jwt/removeAllTokens"
	"hwCalendar/proto/jwtpb"
	"log"
	"net"
)

type Server struct {
	*jwtpb.UnimplementedJwtServiceServer
}

func InitServer() {
	ln, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(onlyLocalhost))
	reflection.Register(grpcServer)
	jwtpb.RegisterJwtServiceServer(grpcServer, &Server{})
	_ = grpcServer.Serve(ln)
}

func (s Server) GenerateTokens(ctx context.Context, request *jwtpb.GenerateTokensRequest) (*jwtpb.GenerateTokensResponse, error) {
	return generate.Handle(ctx, request)
}

func (s Server) RefreshTokens(ctx context.Context, request *jwtpb.RefreshTokensRequest) (*jwtpb.RefreshTokensResponse, error) {
	return refresh.Handle(ctx, request)
}

func (s Server) RemoveAllTokensForUser(ctx context.Context, request *jwtpb.RemoveAllTokensForUserRequest) (*emptypb.Empty, error) {
	return removeAllTokens.Handle(ctx, request)
}

func (s Server) ParseAccessToken(ctx context.Context, request *jwtpb.ParseAccessTokenRequest) (*jwtpb.ParseAccessTokenResponse, error) {
	return parseAccess.Handle(ctx, request)
}
