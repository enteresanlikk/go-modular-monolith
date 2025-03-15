package http_server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/enteresanlikk/go-modular-monolith/internal/config"
	users_presentation "github.com/enteresanlikk/go-modular-monolith/internal/users/presentation"
	"github.com/gorilla/mux"
)

func Start() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	mux := mux.NewRouter()

	// Initialize database
	db := config.NewPostgresDB()

	// Initialize user module with PostgreSQL repository
	users_presentation.RegisterRoutes(mux, db)

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: mux,
	}

	go func() {
		log.Printf("Server starting on http://%s:%s", host, port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("HTTP server error: %v", err)
		}
		log.Println("Stopped serving new connections.")
	}()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	shutdownCtx, shutdownRelease := context.WithTimeout(context.Background(), 10*time.Second)
	defer shutdownRelease()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("HTTP shutdown error: %v", err)
	}
	log.Println("Graceful shutdown complete.")
}
