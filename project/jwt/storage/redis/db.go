package redis

import (
	"github.com/redis/go-redis/v9"
	"os"
	"sync"
)

type db int

const (
	AccessTokenDb db = iota
	RefreshTokenDb
)

var (
	accessOnce     sync.Once
	accessTokenDb  *redis.Client
	refreshOnce    sync.Once
	refreshTokenDb *redis.Client
)

func createRedisClient(db db) *redis.Client {
	addr := os.Getenv("REDIS_ADDR")

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: "",
		DB:       int(db),
	})
}

func GetDb(db db) *redis.Client {
	switch db {
	case AccessTokenDb:
		accessOnce.Do(func() {
			accessTokenDb = createRedisClient(db)
		})
		return accessTokenDb
	case RefreshTokenDb:
		refreshOnce.Do(func() {
			refreshTokenDb = createRedisClient(db)
		})
		return refreshTokenDb
	default:
		panic("non existing db")
	}
}
