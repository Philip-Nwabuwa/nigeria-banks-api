package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/nigeria-banks-api/controllers"
	"github.com/nigeria-banks-api/database"
	"github.com/nigeria-banks-api/middleware"
)

func main() {
	database.InitDB()
	defer database.CloseDB()

	r := gin.Default()
	r.Use(middleware.RateLimiter(5, 10))

	r.GET("/api/banks", controllers.GetBanks)
	r.POST("/api/banks", controllers.AddBank)

	// Handle graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		database.CloseDB()
		os.Exit(0)
	}()

	r.Run(":8080")
}
