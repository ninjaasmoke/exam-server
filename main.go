// PATH: exam-server/main.go

package main

import (
	"exam-server/middlewares"
	"exam-server/models"
	"exam-server/routes"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const maxRetries = 5
const retryInterval = 5 * time.Second // Adjust the retry interval as needed

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

	for i := 0; i < maxRetries; i++ {
		err := models.InitDB(config)
		if err != nil {
			log.Printf("Failed to connect to the database. Retrying in %s...", retryInterval)
			time.Sleep(retryInterval)
		} else {
			break
		}
	}

	if err != nil {
		log.Fatalf("Failed to initialize the database after %d attempts. Exiting.", maxRetries)
	}

	// Load the routes
	routes.PingRoutes(r)
	routes.AuthRoutes(r)

	// Authorized routes
	authorizedGroup := r.Group("/protected")
	authorizedGroup.Use(middlewares.IsAuthorized())
	authorizedGroup.GET("/ping", func(c *gin.Context) {
		claims, exists := c.Get("claims")
		if !exists {
			c.JSON(500, gin.H{"error": "Role information not found"})
			return
		}

		c.JSON(200, gin.H{"message": "Access granted", "role": claims.(*models.Claims).Role})
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
