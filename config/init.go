package config

import (
	"log"

	"github.com/joho/godotenv"
)

// ReadEnvVariables is a function that reads the environment variables from the .env file.
func ReadEnvVariables() {
	err := godotenv.Load("./config/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
