package db

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadEnv loads the .env file
func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}
}
