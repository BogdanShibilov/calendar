package grpc

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	"hwCalendar/calendar/server/grpc/handler/event/add"
	"hwCalendar/calendar/server/grpc/handler/event/all"
	"hwCalendar/calendar/server/grpc/handler/event/byid"
	"hwCalendar/calendar/server/grpc/handler/event/deleteEvent"
	"hwCalendar/calendar/server/grpc/handler/event/update"
	"hwCalendar/proto/eventpb"
	"log"
	"net"
)

type Server struct {
	*eventpb.UnimplementedEventServiceServer
}

func InitServer() {
	ln, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer(grpc.UnaryInterceptor(onlyLocalhost))
	reflection.Register(grpcServer)
	eventpb.RegisterEventServiceServer(grpcServer, &Server{})
	_ = grpcServer.Serve(ln)
}

func (s *Server) AddEvent(ctx context.Context, req *eventpb.AddEventRequest) (*eventpb.AddEventResponse, error) {
	return add.Handle(ctx, req)
}

func (s *Server) UpdateEvent(ctx context.Context, req *eventpb.UpdateEventRequest) (*emptypb.Empty, error) {
	return update.Handle(ctx, req)
}

func (s *Server) DeleteEvent(ctx context.Context, req *eventpb.DeleteEventRequest) (*emptypb.Empty, error) {
	return deleteEvent.Handle(ctx, req)
}

func (s *Server) EventById(ctx context.Context, req *eventpb.EventByIdRequest) (*eventpb.EventByIdResponse, error) {
	return byid.Handle(ctx, req)
}

func (s *Server) AllEvents(ctx context.Context, req *eventpb.AllEventsRequest) (*eventpb.AllEventsResponse, error) {
	return all.Handle(ctx, req)
}
