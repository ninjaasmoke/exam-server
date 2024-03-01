package controllers

import (
	"exam-server/models"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	claims, exists := c.Get("claims")
	if !exists {
		c.JSON(200, gin.H{"message": "pong"})
		return
	}

	c.JSON(200, gin.H{
		"message": "pong",
		"role":    claims.(*models.Claims).Role,
	})
}
