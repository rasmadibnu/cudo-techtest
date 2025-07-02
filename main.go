package main

import (
	"cudo-techtest/config"
	"cudo-techtest/controller"
	"cudo-techtest/repository"
	"cudo-techtest/service"
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		fmt.Printf("Unable to load .env file")
	}

	db := config.InitialDB()

	sqlDB, err := db.DB.DB()
	if err != nil {
		log.Fatalf("Failed to get database instance: %v", err)
	}
	defer func() {
		if err := sqlDB.Close(); err != nil {
			log.Fatalf("Failed to close the database connection: %v", err)
		}
	}()
	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"PUT", "PATCH", "POST", "GET", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	v1 := r.Group("/api/v1")

	transactionRepo := repository.NewTransactionRepository(db)
	transactionService := service.NewTransactionService(transactionRepo)
	controller.NewTransactionController(v1, transactionService)

	r.Run(":6565")
}
