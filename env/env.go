package env

import (
	"os"
	"some-rest-api/logger"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Message("No .env file found", "fatal")
	}
}

// GetEnvVar := give the key of the environment variable you'd like and it will return the value
func GetEnvVar(key string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	logger.Message(key+" : that env var was not found", "fatal")
	return ""
}
