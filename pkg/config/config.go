package config

import (
	"os"
)

type Config struct {
	APIBaseURL string
}

func LoadConfig() *Config {
	return &Config{
		APIBaseURL: getEnv("API_BASE_URL", "https://mock-node-wgqbnxruha-as.a.run.app"),
	}
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
