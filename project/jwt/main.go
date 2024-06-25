package main

import (
	"context"
	"github.com/joho/godotenv"
	"hwCalendar/jwt/server/grpc"
	"hwCalendar/jwt/storage/redis"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	err = redis.GetDb(redis.AccessTokenDb).Ping(context.Background()).Err()
	if err != nil {
		panic(err)
	}
	grpc.InitServer()
}
