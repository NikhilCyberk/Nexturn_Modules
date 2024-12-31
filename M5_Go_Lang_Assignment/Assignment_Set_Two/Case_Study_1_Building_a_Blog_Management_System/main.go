package main

import (
	"log"
	"net/http"

	db "Case_Study_1_Building_a_Blog_Management_System/config"
	"Case_Study_1_Building_a_Blog_Management_System/controller"
	"Case_Study_1_Building_a_Blog_Management_System/middleware"
	"Case_Study_1_Building_a_Blog_Management_System/repository"
	"Case_Study_1_Building_a_Blog_Management_System/service"
)
func main() {
    // Initialize database
    if err := db.InitializeDatabase(); err != nil {
        log.Fatal(err)
    }
    defer db.DB.Close()

    // Initialize dependencies
    blogRepo := repository.NewBlogRepository(db.GetDB())
    blogService := service.NewBlogService(blogRepo)
    blogController := controller.NewBlogController(blogService)

    // Set up routing
    mux := http.NewServeMux()
    
    // Apply middleware chain to all routes
    withMiddleware := middleware.Chain(
        middleware.LoggingMiddleware(),
        middleware.AuthMiddleware(db.DB),
        middleware.JSONValidator(),
    )

    // Register routes with middleware
    mux.Handle("/blogs", withMiddleware(http.HandlerFunc(blogController.HandleBlogs)))
    mux.Handle("/blogs/", withMiddleware(http.HandlerFunc(blogController.HandleBlogByID)))

    // Start server
    log.Println("Server starting on :8080")
    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}