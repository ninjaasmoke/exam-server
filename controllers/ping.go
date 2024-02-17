package controllers

import (
	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists {
		c.JSON(200, gin.H{"message": "pong"})
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
		"role":    role,
	})
}
