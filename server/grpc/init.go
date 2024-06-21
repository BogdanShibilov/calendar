package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/protobuf/eventpb"
	"hwCalendar/server/grpc/handler/event/add"
	"hwCalendar/server/grpc/handler/event/all"
	"hwCalendar/server/grpc/handler/event/byid"
	"hwCalendar/server/grpc/handler/event/deleteEvent"
	"hwCalendar/server/grpc/handler/event/update"
	"log"
	"net"
	"strings"
	"time"
)

type Server struct {
	*eventpb.UnimplementedEventServiceServer
}

func InitGrpc() {
	ln, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(onlyLocalhostInterceptor))
	reflection.Register(grpcServer)
	eventpb.RegisterEventServiceServer(grpcServer, &Server{})
	_ = grpcServer.Serve(ln)
}

func (s *Server) AddEvent(ctx context.Context, req *eventpb.AddEventRequest) (*eventpb.AddEventResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	return add.Handle(ctx, req)
}

func (s *Server) UpdateEvent(ctx context.Context, req *eventpb.UpdateEventRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	return update.Handle(ctx, req)
}

func (s *Server) DeleteEvent(ctx context.Context, req *eventpb.DeleteEventRequest) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	return deleteEvent.Handle(ctx, req)
}

func (s *Server) EventById(ctx context.Context, req *eventpb.EventByIdRequest) (*eventpb.EventByIdResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	return byid.Handle(ctx, req)
}

func (s *Server) AllEvents(ctx context.Context, req *emptypb.Empty) (*eventpb.AllEventsResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	return all.Handle(ctx, req)
}

const localhost = "127.0.0.1"

func onlyLocalhostInterceptor(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.PermissionDenied, "invalid peer")
	}
	peerIpAddr, _, _ := strings.Cut(p.Addr.String(), ":")
	if peerIpAddr != localhost {
		return nil, status.Errorf(codes.PermissionDenied, "invalid peer")
	}

	return handler(ctx, req)
}
