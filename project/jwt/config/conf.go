package config

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
)

var (
	config *Config
	once   sync.Once
)

type Config struct {
	RedisAddr  string
	Secret     string
	AccessTTL  string
	RefreshTTL string
}

func Get() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

		config = &Config{
			RedisAddr:  os.Getenv("REDIS_ADDR"),
			Secret:     os.Getenv("SECRET"),
			AccessTTL:  os.Getenv("ACCESS_TTL"),
			RefreshTTL: os.Getenv("REFRESH_TTL"),
		}
	})

	return *config
}
