package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Db struct {
	Url string
}

type Config struct {
	Db Db
}

func NewConfig(mode string) (*Config, error) {
	if mode == "localhost" {
		err := godotenv.Load(".env.localhost")
		if err != nil {
			return nil, err
		}
	}

	return &Config{
		Db: Db{
			Url: os.Getenv("DATABASE_URL"),
		},
	}, nil
}
