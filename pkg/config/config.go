package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        int
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

	port := os.Getenv("PORT")
	intPort, err := strconv.Atoi(port)
	if err != nil {
		intPort = 8080
	}

	return Config{
		Port:        intPort,
		ProgramId:   os.Getenv("PROGRAM_ID"),
		Username:    os.Getenv("EMAIL"),
		Password:    os.Getenv("PASSWORD"),
		AppPassword: os.Getenv("APP_PASSWORD"),
		SenderEmail: os.Getenv("SENDER_EMAIL"),
	}
}
