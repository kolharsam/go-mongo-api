package env

import (
	"os"
	log "some-rest-api/logger"

	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Message("No .env file found", "fatal")
	}
}

// GetEnvVar := give the key of the environment variable you'd like and it will return the value
func GetEnvVar(key string) string {
	value, exists := os.LookupEnv(key)
	if exists {
		return value
	}
	log.Message(key+" : that env var was not found", "fatal")
	return ""
}
