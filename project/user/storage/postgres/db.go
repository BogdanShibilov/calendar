package postgres

import (
	"github.com/jmoiron/sqlx"
	"hwCalendar/user/config"
	"log"
	"sync"

	_ "github.com/jackc/pgx/v5/stdlib"
)

var (
	once   sync.Once
	single *DB
)

type DB struct {
	*sqlx.DB
}

func GetDb() *DB {
	once.Do(func() {
		dsn := config.Get().DatabaseUrl

		db, err := sqlx.Connect("pgx", dsn)
		if err != nil {
			log.Panicf("failed to connect to database: %v\n", err)
		}

		single = &DB{db}
	})

	return single
}

func (db *DB) Close() {
	_ = db.DB.Close()
}
