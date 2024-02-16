package controllers

import (
	"exam-server/models"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Ping(c *gin.Context) {
	var users = models.DB.Find(&models.User{})

	fmt.Println(users)

	c.JSON(200, gin.H{
		"message": "pong",
	})
}
