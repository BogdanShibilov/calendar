package config

import (
	"github.com/joho/godotenv"
	"os"
	"sync"
)

var (
	once   sync.Once
	config *Config
)

type Config struct {
	Env         string
	DatabaseUrl string
}

func Get() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

		config = &Config{
			Env:         os.Getenv("ENV"),
			DatabaseUrl: os.Getenv("DATABASE_URL"),
		}
	})

	return *config
}
