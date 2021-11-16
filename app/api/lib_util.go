package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/gommon/log"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		log.Fatal(err)
		os.Exit(1)
	}
}
