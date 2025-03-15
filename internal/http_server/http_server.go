package http_server

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/enteresanlikk/go-modular-monolith/internal/users/application"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/infrastructure"
	"github.com/enteresanlikk/go-modular-monolith/internal/users/presentation"
	"github.com/gin-gonic/gin"
)

func Start() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	router := gin.Default()

	// Initialize user module
	userRepo := infrastructure.NewInMemoryUserRepository()
	authService := application.NewAuthService(userRepo)
	presentation.RegisterRoutes(router, authService)

	server := &http.Server{
		Addr:    host + ":" + port,
		Handler: router,
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
