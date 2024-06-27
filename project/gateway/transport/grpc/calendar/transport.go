package calendar

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hwCalendar/gateway/config"
	"hwCalendar/proto/eventpb"
	"sync"
)

var (
	once   sync.Once
	single eventpb.EventServiceClient
)

func GetClient() eventpb.EventServiceClient {
	once.Do(func() {
		conn, err := grpc.NewClient(config.Get().CalendarAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		single = eventpb.NewEventServiceClient(conn)
	})

	return single
}
