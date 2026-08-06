package config

import (
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	GRPCServerAddress string
	LogLevel          string
	MinioUrl          string
	MinioRootUser     string
	MinioRootPassword string
	MinioBucket       string
}

func GetConfig() *Config {
	_ = godotenv.Load()

	var cfg Config

	cfg.GRPCServerAddress = getEnvStringValue("GRPC_SERVER_ADDRESS")
	cfg.LogLevel = getEnvStringValue("LOG_LEVEL")
	cfg.MinioUrl = getEnvStringValue("MINIO_URL")
	cfg.MinioBucket = getEnvStringValue("MINIO_BUCKET")
	cfg.MinioRootUser = getEnvStringValue("MINIO_ROOT_USER")
	cfg.MinioRootPassword = getEnvStringValue("MINIO_ROOT_PASSWORD")

	return &cfg
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
