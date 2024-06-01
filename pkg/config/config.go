package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	ProgramId   string
	Username    string
	Password    string
	AppPassword string
	SenderEmail string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the value of the environment variables
	return Config{
		ProgramId:   os.Getenv("PROGRAM_ID"),
		Username:    os.Getenv("EMAIL"),
		Password:    os.Getenv("PASSWORD"),
		AppPassword: os.Getenv("APP_PASSWORD"),
		SenderEmail: os.Getenv("SENDER_EMAIL"),
	}
}
