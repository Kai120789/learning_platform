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

	cfg.LogLevel = getEnvStringValue("LOG_LEVEL")
	cfg.GRPCServerAddress = getEnvStringValue("GRPC_SERVER_ADDRESS")
	cfg.DBDSN = getEnvStringValue("DBDSN")

	return cfg
}

func getEnvStringValue(envName string) string {
	var cfgValue string

	if envValue := os.Getenv(envName); envValue != "" {
		cfgValue = envValue
	} else {
		cfgValue = ""
	}

	return cfgValue
}
