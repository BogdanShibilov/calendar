package jwt

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"hwCalendar/gateway/config"
	"hwCalendar/proto/jwtpb"
	"sync"
)

var (
	once   sync.Once
	single jwtpb.JwtServiceClient
)

func GetClient() jwtpb.JwtServiceClient {
	once.Do(func() {
		conn, err := grpc.NewClient(config.Get().JwtAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		single = jwtpb.NewJwtServiceClient(conn)
	})

	return single
}
