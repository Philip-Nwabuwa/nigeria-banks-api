package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/nigeria-banks-api/config"
	"github.com/nigeria-banks-api/controllers"
	"github.com/nigeria-banks-api/database"
	"github.com/nigeria-banks-api/middleware"
)

func main() {
	cfg := config.LoadConfig()

	database.InitDB()
	defer database.CloseDB()

	r := gin.Default()
	r.Use(middleware.RateLimiter(5, 10))

	r.GET("/api/banks", controllers.GetBanks)
	r.POST("/api/banks", middleware.AuthMiddleware(cfg), controllers.AddBank)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		log.Println("Shutting down server...")
		database.CloseDB()
		os.Exit(0)
	}()

	log.Printf("Server starting on port %s", cfg.Port)
	r.Run(fmt.Sprintf(":%s", cfg.Port))
}
