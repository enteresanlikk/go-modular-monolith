package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enteresanlikk/go-modular-monolith/internal/api"
	"github.com/joho/godotenv"
)

var requiredEnvVars = []string{
	"PORT",
	"HOST",
	"DB_HOST",
	"DB_PORT",
	"DB_USER",
	"DB_PASSWORD",
	"DB_NAME",
	"DB_SSLMODE",
	"BCRYPT_COST",
	"JWT_SECRET",
	"JWT_REFRESH_SECRET",
}

func checkRequiredEnvironmentVariables() error {
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			return fmt.Errorf("%s is not set", envVar)
		}
	}
	return nil
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = checkRequiredEnvironmentVariables()
	if err != nil {
		log.Fatal(err)
	}

	api.Start()
}
