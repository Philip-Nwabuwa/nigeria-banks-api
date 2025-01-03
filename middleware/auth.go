package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nigeria-banks-api/config"
)

func AuthMiddleware(cfg *config.Config) gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")

		if cfg.APIKey == "" {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "API key not configured on server"})
			c.Abort()
			return
		}

		if apiKey == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "API key is required"})
			c.Abort()
			return
		}

		if apiKey != cfg.APIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid API key"})
			c.Abort()
			return
		}

		c.Next()
	}
}