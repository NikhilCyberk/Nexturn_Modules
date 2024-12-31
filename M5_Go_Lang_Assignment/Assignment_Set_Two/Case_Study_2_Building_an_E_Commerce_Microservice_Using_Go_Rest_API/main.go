package main

import (
	"log"

	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/auth"
	db "Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/config"
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/controller"
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/middleware"
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/repository"
	"Case_Study_2_Building_an_E_Commerce_Microservice_Using_Go_Rest_API/service"

	"github.com/gin-gonic/gin"
)

func main() {
    // Initialize database
    if err := db.InitializeDatabase(); err != nil {
        log.Fatal(err)
    }
    defer db.DB.Close()

    // Initialize dependencies
    productRepo := repository.NewProductRepository(db.GetDB())
    productService := service.NewProductService(productRepo)
    productController := controller.NewProductController(productService)

    // Initialize Gin router
    router := gin.Default()

    // Configure JWT middleware
    secretKey := []byte("nikhil") // Replace with your actual secret key
    jwtConfig := middleware.JWTConfig{
        SecretKey: secretKey,
    }

    // Apply custom middleware
    router.Use(middleware.Logger())
    router.Use(middleware.RateLimiter())

    // Public routes (no authentication required)
    router.POST("/login", auth.LoginHandler(secretKey))

    // Protected routes
    api := router.Group("/api")
    api.Use(middleware.JWTAuth(jwtConfig))
    {
        products := api.Group("/products")
        {
            products.POST("/", productController.CreateProduct)
            products.GET("/", productController.GetProducts)
            products.GET("/:id", productController.GetProduct)
            products.PUT("/:id", productController.UpdateProduct)
            products.DELETE("/:id", productController.DeleteProduct)
            products.PUT("/:id/stock", productController.UpdateStock)
        }
    }

    // Start server
    router.Run(":8080")
}