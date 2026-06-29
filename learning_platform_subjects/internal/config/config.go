package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	GRPCServerAddress string
	LogLevel          string
	DBDSN             string
}

func GetConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.GRPCServerAddress = getEnvStringValue("GRPC_SERVER_ADDRESS")
	cfg.DBDSN = getEnvStringValue("DBDSN")
	cfg.LogLevel = getEnvStringValue("LOG_LEVEL")

	return cfg
}

func getEnvStringValue(envName string) string {
	if val := os.Getenv(envName); val != "" {
		return val
	}

	return ""
}
