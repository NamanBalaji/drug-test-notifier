package config

import (
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Config struct {
	ProgramId string
	Username  string
	Password  string
}

func LoadConfig() Config {
	envPath, err := filepath.Abs("../.env")
	if err != nil {
		log.Fatalf("Error getting absolute path to .env file: %v", err)
	}
	err = godotenv.Load(envPath)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// Get the value of the environment variables
	return Config{
		ProgramId: os.Getenv("PROGRAM_ID"),
		Username:  os.Getenv("EMAIL"),
		Password:  os.Getenv("PASSWORD"),
	}
}
