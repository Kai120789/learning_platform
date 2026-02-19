package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel       string
	ServerAddress  string
	AuthServiceUrl string
	UserServiceUrl string
	RedisUrl       string
}

func GetConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.LogLevel = getEnvStringValue("LOG_LEVEL")
	cfg.ServerAddress = getEnvStringValue("SERVER_ADDRESS")
	cfg.AuthServiceUrl = getEnvStringValue("AUTH_SERVICE_URL")
	cfg.UserServiceUrl = getEnvStringValue("USER_SERVICE_URL")
	cfg.RedisUrl = getEnvStringValue("REDIS_URL")

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
