package lib

import (
	"fmt"

	"github.com/joho/godotenv"
)

func LoadEnv() error {
	err := godotenv.Load("../.env")
	if err != nil {
		return fmt.Errorf("error occured when loading .env file.\n%s", err)
	}
	return nil
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
