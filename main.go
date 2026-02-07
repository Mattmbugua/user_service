package main

import (
	"log"
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

	// Start server
	log.Println("Server running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
