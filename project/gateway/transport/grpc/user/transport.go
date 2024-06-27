package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hwCalendar/gateway/config"
	"hwCalendar/proto/userpb"
	"sync"
)

var (
	once   sync.Once
	single userpb.UserServiceClient
)

func GetClient() userpb.UserServiceClient {
	once.Do(func() {
		conn, err := grpc.NewClient(config.Get().UserAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		single = userpb.NewUserServiceClient(conn)
	})

	return single
}
