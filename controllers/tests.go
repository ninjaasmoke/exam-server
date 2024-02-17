package controllers

import (
	"exam-server/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": err})
		return
	}
	if err := models.DB.Create(&question).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": question, "message": "Question created successfully"})
}

func Createquestions(c *gin.Context) {
	var questions []models.Question
	if err := c.ShouldBindJSON(&questions); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": err})
		return
	}
	if err := models.DB.Create(&questions).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": questions, "message": "Questions created successfully"})

}

func GetQuestions(c *gin.Context) {
	var questions []models.Question
	if err := models.DB.Find(&questions).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": questions})
}

func GetQuestion(c *gin.Context) {
	var question models.Question
	if err := models.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": question})
}

func UpdateQuestion(c *gin.Context) {
	var question models.Question
	if err := models.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	if err := c.ShouldBindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": err})
		return
	}
	models.DB.Save(&question)
	c.JSON(http.StatusOK, gin.H{"data": question, "message": "Question updated successfully"})
}

func DeleteQuestion(c *gin.Context) {
	var question models.Question
	if err := models.DB.Where("id = ?", c.Param("id")).First(&question).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	// soft delete
	models.DB.Save(&models.Question{Model: gorm.Model{ID: question.ID}, Deleted: true})
	c.JSON(http.StatusOK, gin.H{"data": question, "message": "Question deleted successfully"})
}

func CreateTest(c *gin.Context) {
	var test models.Test
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": err})
		return
	}
	if err := models.DB.Create(&test).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": test, "message": "Test created successfully"})
}

func GetTests(c *gin.Context) {
	var tests []models.Test
	if err := models.DB.Find(&tests).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": tests})
}

func GetTest(c *gin.Context) {
	var test models.Test
	if err := models.DB.Where("id = ?", c.Param("id")).First(&test).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": test})
}

func UpdateTest(c *gin.Context) {
	var test models.Test
	if err := models.DB.Where("id = ?", c.Param("id")).First(&test).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	if err := c.ShouldBindJSON(&test); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": err})
		return
	}
	if err := models.DB.Save(&test).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": test, "message": "Test updated successfully"})
}

func DeleteTest(c *gin.Context) {
	var test models.Test
	if err := models.DB.Where("id = ?", c.Param("id")).First(&test).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	// soft delete
	if err := models.DB.Save(&models.Test{Model: gorm.Model{ID: test.ID}, Deleted: true}).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": test, "message": "Test deleted successfully"})
}

func CreateResponse(c *gin.Context) {
	var response models.Response
	if err := c.ShouldBindJSON(&response); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": err})
		return
	}
	if err := models.DB.Create(&response).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response, "message": "Response created successfully"})
}

func CreateResponses(c *gin.Context) {
	var responses []models.Response
	if err := c.ShouldBindJSON(&responses); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "data": err})
		return
	}

	if err := models.DB.Create(&responses).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error(), "data": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": responses, "message": "Responses created successfully"})
}

func GetAllResponses(c *gin.Context) {
	var responses []models.Response
	if err := models.DB.Find(&responses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": responses})
}

func GetResponse(c *gin.Context) {
	var response models.Response
	if err := models.DB.Where("id = ?", c.Param("id")).First(&response).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": response})
}

func GetResponsesPerTest(c *gin.Context) {
	var responses []models.Response
	if err := models.DB.Where("test_id = ?", c.Param("id")).Find(&responses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": responses})
}

func GetResponsePerUser(c *gin.Context) {
	var responses []models.Response
	if err := models.DB.Where("user_id = ?", c.Param("id")).Find(&responses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": responses})
}

func GetResponsePerUserPerTest(c *gin.Context) {
	var responses []models.Response
	if err := models.DB.Where("user_id = ? AND test_id = ?", c.Param("user_id"), c.Param("test_id")).Find(&responses).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": responses})
}

func GetResultsPerUser(c *gin.Context) {
	var results []models.Result
	if err := models.DB.Where("user_id = ?", c.Param("user_id")).Find(&results).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": results})
}

func GetResultsPerTest(c *gin.Context) {
	var results []models.Result
	if err := models.DB.Where("test_id = ?", c.Param("id")).Find(&results).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!", "data": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": results})
}
