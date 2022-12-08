package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


func LoadENV(envName string) string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading URI, %s\n", envName)
	}

	return os.Getenv(envName)
}