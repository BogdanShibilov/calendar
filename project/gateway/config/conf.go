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
	CalendarAddr string
	JwtAddr      string
	UserAddr     string
}

func Get() Config {
	once.Do(func() {
		err := godotenv.Load()
		if err != nil {
			panic(err)
		}

		config = &Config{
			CalendarAddr: os.Getenv("CALENDAR_ADDR"),
			JwtAddr:      os.Getenv("JWT_ADDR"),
			UserAddr:     os.Getenv("USER_ADDR"),
		}
	})

	return *config
}
