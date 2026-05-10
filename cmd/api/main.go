package main

import (
	"log"
	"study/config"
	"study/internal/handler"
	"study/internal/repository"
	"study/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize Database
	db, err := config.InitDB()
	if err != nil {
		log.Fatalf("Could not connect to database: %v", err)
	}

	// Dependency Injection
	repo := repository.NewRentalRepository(db)
	uc := usecase.NewRentalUsecase(repo)
	h := handler.NewRentalHandler(uc)

	// Router Setup
	r := gin.Default()

	// Endpoints
	r.POST("/vehicles", h.AddVehicle)
	r.POST("/customers", h.AddCustomer)
	r.POST("/rentals", h.CreateRental)

	log.Println("Rental API server running on :8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
