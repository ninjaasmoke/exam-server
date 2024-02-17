package middlewares

import (
	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("role")

		if !exists {
			c.JSON(500, gin.H{"error": "Role information not found", "data": "role information not found"})
			c.Abort()
			return
		}

		if roleInt, ok := role.(uint); !(ok && roleInt >= 2) {
			c.JSON(401, gin.H{"error": "unauthorized", "data": "user is not an admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
