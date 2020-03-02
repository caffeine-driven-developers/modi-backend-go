package lib

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

// get api key from .env file
func GoDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}