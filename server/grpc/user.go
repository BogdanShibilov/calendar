package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/protobuf/userpb"
	"hwCalendar/server/grpc/handler/user/add"
	"hwCalendar/server/grpc/handler/user/all"
	"hwCalendar/server/grpc/handler/user/byid"
	"hwCalendar/server/grpc/handler/user/deleteUser"
	"hwCalendar/server/grpc/handler/user/update"
	"log"
	"net"
)

type UserGprc struct {
	*userpb.UnimplementedUserServiceServer
}

func InitUserServer() {
	ln, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(onlyLocalhost))
	reflection.Register(grpcServer)
	userpb.RegisterUserServiceServer(grpcServer, &UserGprc{})
	_ = grpcServer.Serve(ln)
}

func (u UserGprc) AddUser(ctx context.Context, request *userpb.AddUserRequest) (*userpb.AddUserResponse, error) {
	return add.Handle(ctx, request)
}

func (u UserGprc) UpdateUser(ctx context.Context, request *userpb.UpdateUserRequest) (*emptypb.Empty, error) {
	return update.Handle(ctx, request)
}

func (u UserGprc) DeleteUser(ctx context.Context, request *userpb.DeleteUserRequest) (*emptypb.Empty, error) {
	return deleteUser.Handle(ctx, request)
}

func (u UserGprc) UserById(ctx context.Context, request *userpb.UserByIdRequest) (*userpb.UserByIdResponse, error) {
	return byid.Handle(ctx, request)
}

func (u UserGprc) AllUsers(ctx context.Context, empty *emptypb.Empty) (*userpb.AllUsersResponse, error) {
	return all.Handle(ctx, empty)
}
