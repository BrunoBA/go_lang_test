package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

var DB = make(map[string]string)

func setupRouter() *gin.Engine {

	if os.Getenv("ENV_MODE") == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, os.Getenv("ENV_MODE"))
	})

	// Get user value
	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := DB[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	return r
}

func main() {
	godotenv.Load()
	r := setupRouter()
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	// Listen and Server in 0.0.0.0:8080
	r.Run(":" + port)
}
