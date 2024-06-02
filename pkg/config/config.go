package config

import (
	"os"
)

type Config struct {
	ProgramId   string
	Username    string
	Password    string
	AppPassword string
	SenderEmail string
}

func LoadConfig() Config {

	// Get the value of the environment variables
	return Config{
		ProgramId:   os.Getenv("PROGRAM_ID"),
		Username:    os.Getenv("EMAIL"),
		Password:    os.Getenv("PASSWORD"),
		AppPassword: os.Getenv("APP_PASSWORD"),
		SenderEmail: os.Getenv("SENDER_EMAIL"),
	}
}
