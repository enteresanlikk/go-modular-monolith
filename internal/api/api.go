package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	common_infrastructure "github.com/enteresanlikk/go-modular-monolith/internal/common/infrastructure"
	todos_module "github.com/enteresanlikk/go-modular-monolith/internal/todos"
	users_module "github.com/enteresanlikk/go-modular-monolith/internal/users"
	"github.com/gorilla/mux"
)

func Start() {
	host := os.Getenv("HOST")
	port := os.Getenv("PORT")

	mux := mux.NewRouter()

	db := common_infrastructure.NewPostgresDB()

	users_module.Register(mux, db)
	todos_module.Register(mux, db)

	server := &http.Server{
		Addr:         host + ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
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
