package main

import (
	"app_backend/controller"
	"app_backend/database"
	"app_backend/middleware"
	model "app_backend/models"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()
	loadDatabase()
	serveApplication()
}

func loadDatabase() {
	database.Connect()
	database.Database.AutoMigrate(&model.User{})
	database.Database.AutoMigrate(&model.Entry{})
	database.Database.AutoMigrate(&model.Customer{})
	database.Database.AutoMigrate(&model.Product{})
	database.Database.AutoMigrate(&model.Bank{})
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func serveApplication() {
	router := gin.Default()

	publicRoutes := router.Group("/auth")
	publicRoutes.POST("/register", controller.Register)
	publicRoutes.POST("/login", controller.Login)

	protectedRoutes := router.Group("/api")
	protectedRoutes.Use(middleware.JWTAuthMiddleware())
	protectedRoutes.POST("/entry", controller.AddEntry)
	protectedRoutes.GET("/entry", controller.GetAllEntries)
	protectedRoutes.POST("/customer", controller.AddCustomer)
	protectedRoutes.GET("/customer", controller.GetAllCustomer)

	router.Run(":8080")
	fmt.Println("Server running on port 8080")
}
