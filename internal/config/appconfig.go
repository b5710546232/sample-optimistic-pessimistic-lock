package config

import (
	"context"
	"log"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
)

type AppConfig struct {
	DbConnectionStr string `env:"DATABASE_URL,required"`
}

func NewAppConfig() AppConfig {
	config := createConfig()
	return config
}

func createConfig() AppConfig {
	godotenv.Load(".env")

	ctx := context.Background()
	var appConfig AppConfig

	if err := envconfig.Process(ctx, &appConfig); err != nil {
		log.Fatal(err)
	}

	return appConfig
}
