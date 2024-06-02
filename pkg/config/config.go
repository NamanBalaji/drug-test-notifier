package config

import (
	"os"
	"strconv"
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
