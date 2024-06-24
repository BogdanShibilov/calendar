package main

import (
	"embed"
	"hwCalendar/calendar/migrator"
	"hwCalendar/calendar/server/grpc"
	"hwCalendar/calendar/storage/postgres"
	"log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	migrator.SetBaseFs(embedMigrations)
	migrator.SetDB(postgres.GetDb().DB)
	if err := migrator.Up("migrations"); err != nil {
		log.Panicf("failed to get up migrations: %v", err)
	}
	grpc.InitServer()
}
