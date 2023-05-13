package entity

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
)

type Config struct {
	ServerPort int    `env:"SERVER_PORT"`
	PGport     int    `env:"PG_PORT"`
	PGhost     string `env:"PG_HOST"`
	PGuser     string `env:"PG_USER"`
	PGpassword string `env:"PG_PASSWORD"`
	PGname     string `env:"PG_NAME"`
	RedisAddr  string `env:"REDIS_ADDR"`
}

func ParseConfig() (Config, error) {
	err := godotenv.Load()
	if err != nil {
		return Config{}, err
	}

	var cfg Config

	err = env.Parse(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
