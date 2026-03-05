package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
)

type Db struct {
	Url string
}

type Config struct {
	Db Db
}

func NewConfig(mode string) (*Config, error) {
	if mode == "localhost" {
		if _, err := os.Stat(".env/.env.localhost"); err == nil {
			if loadErr := godotenv.Load(".env/.env.localhost"); loadErr != nil {
				return nil, loadErr
			}
			log.Info().Msg("Loaded .env/.env.localhost")
		} else {
			log.Warn().Msg(".env/.env.localhost not found, using environment variables directly")
		}
	}

	return &Config{
		Db: Db{
			Url: os.Getenv("DATABASE_URL"),
		},
	}, nil
}
