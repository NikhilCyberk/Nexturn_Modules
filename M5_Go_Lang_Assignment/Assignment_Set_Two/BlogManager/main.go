package main

import (
	db "blogmanager/config"
	"blogmanager/controller"
	"blogmanager/middleware"
	"blogmanager/repository"
	"blogmanager/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitializeDatabase()

	// Create repository, service, and controller for products
	blogRepo := repository.NewBlogRepository(db.GetDB())
	blogService := service.NewBlogService(blogRepo)
	blogController := controller.NewBlogController(blogService)

	// Initialize Gin router
	r := gin.Default()

	// Apply logging middleware globally
	r.Use(middleware.LoggingMiddleware())

	// Group routes and apply authentication middleware
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware(db.GetDB()))

	// Routes for users
	api.POST("/blog", blogController.CreateBlog)
	api.GET("/blog/:id", blogController.GetBlog)
	api.GET("/blog", blogController.GetAllBlogs)
	api.PUT("/blog/:id", blogController.UpdateBlog)
	api.DELETE("/blog/:id", blogController.DeleteBlog)

	// Start server on port 8080
	r.Run(":8080")
}
