package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	commonInfrastructure "github.com/enteresanlikk/go-modular-monolith/internal/common/infrastructure"
	tenantsModule "github.com/enteresanlikk/go-modular-monolith/internal/tenants"
	todosModule "github.com/enteresanlikk/go-modular-monolith/internal/todos"
	usersModule "github.com/enteresanlikk/go-modular-monolith/internal/users"
	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	postgresConfig := &commonInfrastructure.PostgresDBConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	app := fiber.New(fiber.Config{
		JSONEncoder:  json.Marshal,
		JSONDecoder:  json.Unmarshal,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	db := commonInfrastructure.NewPostgresDB(postgresConfig)

	usersModule.Register(app, db)
	todosModule.Register(app, db)
	tenantsModule.Register(app, db)

	go func() {
		log.Printf("Server starting on http://%s:%s", host, port)
		if err := app.Listen(host + ":" + port); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	_, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := app.Shutdown(); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}
