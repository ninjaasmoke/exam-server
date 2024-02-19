package middlewares

import (
	"exam-server/models"

	"github.com/gin-gonic/gin"
)

func IsAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		claims, exists := c.Get("claims")

		if !exists {
			c.JSON(500, gin.H{"error": "Role information not found", "data": "role information not found"})
			c.Abort()
			return
		}

		if claims.(*models.Claims).Role < 2 {
			c.JSON(401, gin.H{"error": "unauthorized", "data": "user is not an admin"})
			c.Abort()
			return
		}

		c.Next()
	}
}
