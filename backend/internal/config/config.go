package config

import (
	"os"
	"strconv"
)

const (
	DefaultHTTPPort = "3080"

	EnvLocal = "local"
	Prod     = "prod"
)

type JWTConfig struct {
	JWTSecretKey []byte
}

type Config struct {
	JWT JWTConfig
}

func New() *Config {
	return &Config{
		JWT: JWTConfig{
			JWTSecretKey: getEnvAsByte("JWT_SECRET_KEY"),
		},
	}
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}

func getEnvAsByte(key string) []byte {
	valueStr := getEnv(key, "")
	return []byte(valueStr)
}

func getEnvAsInt(name string, defaultVal int) int {
	valueStr := getEnv(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}

	return defaultVal
}
