package main

import (
	"embed"
	"hwCalendar/migrator"
	"hwCalendar/server/grpc"
	"log"
)

//go:embed migrations/*.sql
var embedMigrations embed.FS

func main() {
	migrator.SetBaseFs(embedMigrations)
	migrator.SetDB(db.DB)
	if err := migrator.Up("migrations"); err != nil {
		log.Panicf("failed to get up migrations: %v", err)
	}
	grpc.InitGrpc()
}
