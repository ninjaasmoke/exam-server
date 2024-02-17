package middlewares

import (
	"github.com/gin-gonic/gin"
)

// IsStudent is a middleware to check if the user is a student

func IsStudent() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		if !exists {
			c.JSON(500, gin.H{"error": "Role information not found", "data": "role information not found"})
			c.Abort()
			return
		}

		if roleInt, ok := role.(int); ok && roleInt <= 1 {
			c.JSON(401, gin.H{"error": "unauthorized", "data": "user is not a student"})
			c.Abort()
			return
		}

		c.Next()
	}
}
