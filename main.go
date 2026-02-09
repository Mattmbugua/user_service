package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
	"user_service/internal/handler"
	"user_service/internal/repository"
	"user_service/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize in-memory repository
	userRepo := repository.NewInMemoryUserRepository()

	// Initialize service with repository
	userService := service.NewUserService(userRepo)

	// Initialize handler with service
	userHandler := handler.NewUserHandler(userService)

	// Setup Gin router
	router := gin.Default()
	userHandler.RegisterRoutes(router)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}
	// Start server with error handling for graceful shutdown

	go func() {
		log.Println("Server running on http://localhost:8080")
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down the server")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exiting")

}
