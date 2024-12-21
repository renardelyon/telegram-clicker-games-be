package config

import (
	"log"

	"github.com/caarlos0/env/v11"
	"github.com/joho/godotenv"
)

func Setup() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Warning: Could not load .env file, continuing with system environment variables")
	}

	var cfg Config
	err = env.Parse(&cfg)
	if err != nil {
		return nil, err
	}

	return &cfg, nil
}
