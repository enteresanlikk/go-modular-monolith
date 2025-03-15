package main

import (
	"fmt"
	"log"
	"os"

	"github.com/enteresanlikk/go-modular-monolith/internal/http_server"
	"github.com/joho/godotenv"
)

var requiredEnvVars = []string{"PORT", "HOST"}

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

	http_server.Start()
}
