package config

import (
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	LogLevel            string
	ServerAddress       string
	AuthServiceUrl      string
	UserServiceUrl      string
	RedisUrl            string
	SignedKey           string
	AccessTokenLiveTime int64
}

func GetConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.LogLevel = getEnvStringValue("LOG_LEVEL")
	cfg.ServerAddress = getEnvStringValue("SERVER_ADDRESS")
	cfg.AuthServiceUrl = getEnvStringValue("AUTH_SERVICE_URL")
	cfg.SignedKey = getEnvStringValue("SIGNED_KEY")
	cfg.AccessTokenLiveTime = getEnvIntValue("ACCESS_TOKEN_LIVE_TIME", 5)
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

func getEnvIntValue(envName string, fallbackValue int64) int64 {
	cfgValue, err := strconv.Atoi(os.Getenv(envName))
	if err != nil {
		fmt.Errorf("string to int convert error: %w", err)
		return fallbackValue
	}

	return int64(cfgValue)
}
