package main

import (
	"hwCalendar/server/grpc"
	"hwCalendar/storage/postgres"
)

func main() {
	defer postgres.GetDb().Close()
	grpc.InitGrpc()
}
