package postgres

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
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
		dsn := os.Getenv("DATABASE_URL")

		db, err := sqlx.Open("pgx", dsn)
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
