package config

import (
	"log"

	"github.com/joho/godotenv"
)

// function to read config.env file
func ReadEnvVariables() {
	err := godotenv.Load("./config/config.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
