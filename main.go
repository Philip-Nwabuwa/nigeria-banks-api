package main

import (
	"github.com/gin-gonic/gin"
	"github.com/nigeria-banks-api/controllers"
	"github.com/nigeria-banks-api/database"
	"github.com/nigeria-banks-api/middleware"
)

func main() {
	// Initialize database
	database.InitDB()

	// Create Gin router
	r := gin.Default()

	// Apply rate limiter middleware (5 requests per second with burst of 10)
	r.Use(middleware.RateLimiter(5, 10))

	// Routes
	r.GET("/api/banks", controllers.GetBanks)
	r.POST("/api/banks", controllers.AddBank)

	// Start server
	r.Run(":8080")
}
