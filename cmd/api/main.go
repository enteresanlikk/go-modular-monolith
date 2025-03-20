package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	commonInfrastructure "github.com/enteresanlikk/go-modular-monolith/internal/common/infrastructure"
	todosModule "github.com/enteresanlikk/go-modular-monolith/internal/todos"
	usersModule "github.com/enteresanlikk/go-modular-monolith/internal/users"
	"github.com/fasthttp/router"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
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

	r := router.New()

	db := commonInfrastructure.NewPostgresDB(postgresConfig)

	usersModule.Register(r, db)
	todosModule.Register(r, db)

	server := &fasthttp.Server{
		Handler:      r.Handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go func() {
		log.Printf("Server starting on http://%s:%s", host, port)
		if err := server.ListenAndServe(host + ":" + port); err != nil {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	_, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}
