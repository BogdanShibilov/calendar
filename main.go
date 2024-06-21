package main

import (
	"embed"
	"hwCalendar/migrator"
	"hwCalendar/server/grpc"
	"hwCalendar/storage/postgres"
	"log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	db := postgres.GetDb()
	err := db.Ping()
	if err != nil {
		log.Panicf("failed to connect to db: %v", err)
	}

	migrator.SetBaseFs(embedMigrations)
	migrator.SetDB(db.DB)
	if err := migrator.Up("migrations"); err != nil {
		log.Panicf("failed to get up migrations: %v", err)
	}
	grpc.InitGrpc()
}
