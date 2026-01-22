package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel       string
	ServerAddress  string
	DBDSN          string
	AuthServiceUrl string
}

func GetConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.LogLevel = getEnvStringValue("LOG_LEVEL")
	cfg.ServerAddress = getEnvStringValue("SERVER_ADDRESS")
	cfg.DBDSN = getEnvStringValue("DBDSN")
	cfg.AuthServiceUrl = getEnvStringValue("AUTH_SERVICE_URL")

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
