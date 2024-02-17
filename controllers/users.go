package controllers

import (
	"exam-server/models"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {
	var users []models.User
	if err := models.DB.Find(&users).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(200, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(400, gin.H{"error": "Record not found!"})
		return
	}
	c.JSON(200, gin.H{"data": user})
}
