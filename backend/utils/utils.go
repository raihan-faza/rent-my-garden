package utils

import (
	"github.com/joho/godotenv"
	"github.com/pkg/errors"
)

func LoadEnv() error {
	if err := godotenv.Load(); err != nil {
		return errors.Errorf("failed to load .env file: %s", err)
	}
	return nil
}
