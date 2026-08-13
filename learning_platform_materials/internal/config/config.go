package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	LogLevel          string
	DBDSN             string
	GRPCServerAddress string
}

func GetConfig() *Config {
	_ = godotenv.Load()

	var cfg Config
	cfg.GRPCServerAddress = getEnvStringValue("GRPC_SERVER_ADDRESS")
	cfg.LogLevel = getEnvStringValue("LOG_LEVEL")
	cfg.DBDSN = getEnvStringValue("DBDSN")

	return &cfg
}

func getEnvStringValue(envName string) string {
	if val := os.Getenv(envName); val != "" {
		return val
	}

	return ""
}
