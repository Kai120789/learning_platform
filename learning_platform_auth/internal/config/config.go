package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

type Config struct {
	GRPCServerAddress    string
	LogLevel             string
	DBDSN                string
	RedisUrl             string
	UserServiceUrl       string
	AccessTokenLifeTime  int64
	RefreshTokenLifeTime int64
	Issuer               string
	SignedKey            string
	Salt                 int64
}

func GetConfig() *Config {
	_ = godotenv.Load()

	cfg := &Config{}

	cfg.LogLevel = getEnvStringValue("LOG_LEVEL")
	cfg.GRPCServerAddress = getEnvStringValue("GRPC_SERVER_ADDRESS")
	cfg.DBDSN = getEnvStringValue("DBDSN")
	cfg.RedisUrl = getEnvStringValue("REDIS_URL")
	cfg.UserServiceUrl = getEnvStringValue("USER_SERVICE_URL")
	cfg.Issuer = getEnvStringValue("ISSUER")
	cfg.SignedKey = getEnvStringValue("SIGNED_KEY")

	cfg.Salt = getEnvIntValue("SALT", 10)
	cfg.AccessTokenLifeTime = getEnvIntValue("ACCESS_TOKEN_LIVE_TIME", 5)
	cfg.RefreshTokenLifeTime = getEnvIntValue("REFRESH_TOKEN_LIVE_TIME", 7)

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
