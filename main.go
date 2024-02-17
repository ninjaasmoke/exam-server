// PATH: exam-server/main.go

package main

import (
	"exam-server/middlewares"
	"exam-server/models"
	"exam-server/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Create a new gin instance
	r := gin.Default()

	// Load .env file and Create a new connection to the database
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	config := models.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   os.Getenv("DB_NAME"),
		SSLMode:  os.Getenv("DB_SSLMODE"),
	}

	// Initialize DB
	models.InitDB(config)

	// Load the routes
	routes.PingRoutes(r)
	routes.AuthRoutes(r)

	// Authorized routes
	authorizedGroup := r.Group("/protected")
	authorizedGroup.Use(middlewares.IsAuthorized())
	authorizedGroup.GET("/ping", func(c *gin.Context) {
		role, exists := c.Get("role")
		if !exists {
			c.JSON(500, gin.H{"error": "Role information not found"})
			return
		}

		c.JSON(200, gin.H{"message": "Access granted", "role": role})
	})

	// Student routes
	studentGroup := r.Group("/student")
	studentGroup.Use(middlewares.IsAuthorized(), middlewares.IsStudent())
	routes.StudentRoutes(studentGroup)

	// Admin routes
	adminGroup := r.Group("/admin")
	adminGroup.Use(middlewares.IsAuthorized(), middlewares.IsAdmin())
	routes.AdminRoutes(adminGroup)

	r.Run("127.0.0.1:8080")
}
