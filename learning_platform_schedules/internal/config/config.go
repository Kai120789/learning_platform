package config

import "github.com/joho/godotenv"

type Config struct {
}

func GetConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{}

	return cfg
}
