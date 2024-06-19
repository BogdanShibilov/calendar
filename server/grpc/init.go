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
)

type Server struct {
	*eventpb.UnimplementedEventServiceServer
}

func InitGrpc() {
	ln, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	eventpb.RegisterEventServiceServer(grpcServer, &Server{})
	_ = grpcServer.Serve(ln)
}

func (s *Server) AddEvent(ctx context.Context, req *eventpb.AddEventRequest) (*eventpb.AddEventResponse, error) {
	err := checkIfLocalhost(ctx)
	if err != nil {
		return nil, err
	}

	return add.Handle(ctx, req)
}

func (s *Server) UpdateEvent(ctx context.Context, req *eventpb.UpdateEventRequest) (*emptypb.Empty, error) {
	err := checkIfLocalhost(ctx)
	if err != nil {
		return nil, err
	}

	return update.Handle(ctx, req)
}

func (s *Server) DeleteEvent(ctx context.Context, req *eventpb.DeleteEventRequest) (*emptypb.Empty, error) {
	err := checkIfLocalhost(ctx)
	if err != nil {
		return nil, err
	}

	return deleteEvent.Handle(ctx, req)
}

func (s *Server) EventById(ctx context.Context, req *eventpb.EventByIdRequest) (*eventpb.EventByIdResponse, error) {
	err := checkIfLocalhost(ctx)
	if err != nil {
		return nil, err
	}

	return byid.Handle(ctx, req)
}

func (s *Server) AllEvents(ctx context.Context, req *emptypb.Empty) (*eventpb.AllEventsResponse, error) {
	err := checkIfLocalhost(ctx)
	if err != nil {
		return nil, err
	}

	return all.Handle(ctx, req)
}

func checkIfLocalhost(ctx context.Context) error {
	p, ok := peer.FromContext(ctx)
	if !ok {
		return status.Errorf(codes.PermissionDenied, "invalid peer")
	}
	peerIpAddr, _, _ := strings.Cut(p.Addr.String(), ":")
	if peerIpAddr != "127.0.0.1" {
		return status.Errorf(codes.PermissionDenied, "invalid peer")
	}

	return nil
}
