package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/proto/userpb"
	"hwCalendar/user/server/grpc/handler/user/add"
	"hwCalendar/user/server/grpc/handler/user/all"
	"hwCalendar/user/server/grpc/handler/user/byid"
	"hwCalendar/user/server/grpc/handler/user/byusername"
	"hwCalendar/user/server/grpc/handler/user/checkCreds"
	"hwCalendar/user/server/grpc/handler/user/deleteUser"
	"hwCalendar/user/server/grpc/handler/user/update"
	"log"
	"net"
)

type UserGprc struct {
	*userpb.UnimplementedUserServiceServer
}

func InitServer() {
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

func (u UserGprc) UserByUsername(ctx context.Context, request *userpb.UserByUsernameRequest) (*userpb.UserByUsernameResponse, error) {
	return byusername.Handle(ctx, request)
}

func (u UserGprc) AllUsers(ctx context.Context, empty *emptypb.Empty) (*userpb.AllUsersResponse, error) {
	return all.Handle(ctx, empty)
}

func (u UserGprc) CheckCredentials(ctx context.Context, request *userpb.CheckCredentialsRequest) (*userpb.CheckCredentialsResponse, error) {
	return checkCreds.Handle(ctx, request)
}
