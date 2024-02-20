// PATH: exam-server/controllers/auth.go

package controllers

import (
	"exam-server/models"
	"os"
	"time"

	"exam-server/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	if user.UserName == "" || user.Password == "" {
		c.JSON(400, gin.H{"error": "email and password are required"})
		return
	}

	if err := models.DB.Where("username = ?", user.UserName).First(&existingUser).Error; err != nil {
		c.JSON(400, gin.H{"error": "user does not exist", "data": err.Error()})
		return
	}

	if existingUser.ID == 0 {
		c.JSON(400, gin.H{"error": "user does not exist"})
		return
	}

	errHash := utils.CompareHashPassword(user.Password, existingUser.Password)

	if !errHash {
		c.JSON(400, gin.H{"error": "invalid password"})
		return
	}

	expirationTime := time.Now().Add(480 * time.Hour)

	claims := &models.Claims{
		Role:     existingUser.Role,
		UserID:   existingUser.ID,
		UserName: existingUser.UserName,
		StandardClaims: jwt.StandardClaims{
			Subject:   existingUser.Email,
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	var jwtKey = []byte(os.Getenv("JWT_KEY"))

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		c.JSON(500, gin.H{"error": "could not generate token"})
		return
	}

	c.SetCookie("token", tokenString, int(expirationTime.Unix()), "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged in", "token": tokenString, "role": existingUser.Role, "username": existingUser.UserName, "email": existingUser.Email, "id": existingUser.ID})
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var existingUser models.User

	models.DB.Where("email = ?", user.Email).First(&existingUser)

	if existingUser.ID != 0 {
		c.JSON(400, gin.H{"error": "user already exists"})
		return
	}

	var errHash error
	user.Password, errHash = utils.GenerateHashPassword(user.Password)

	if errHash != nil {
		c.JSON(500, gin.H{"error": "could not generate password hash"})
		return
	}

	user.Role = 0

	if err := models.DB.Create(&user).Error; err != nil {
		c.JSON(500, gin.H{"error": "could not create user", "data": err.Error()})
		return
	}

	c.JSON(200, gin.H{"success": "user created", "data": user})
}

func Home(c *gin.Context) {

	cookie, err := c.Cookie("token")

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	claims, err := utils.ParseToken(cookie)

	if err != nil {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	if claims.Role < 1 {
		c.JSON(401, gin.H{"error": "unauthorized"})
		return
	}

	c.JSON(200, gin.H{"success": "home page", "role": claims.Role})
}

func Logout(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "localhost", false, true)
	c.JSON(200, gin.H{"success": "user logged out"})
}
