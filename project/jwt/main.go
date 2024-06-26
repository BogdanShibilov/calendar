package main

import (
	"context"
	"hwCalendar/jwt/server/grpc"
	"hwCalendar/jwt/storage/redis"
)

func main() {
	err := redis.GetDb(redis.AccessTokenDb).Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	grpc.InitServer()
}
