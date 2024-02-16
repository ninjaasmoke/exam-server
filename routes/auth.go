package routes

import (
	"exam-server/controllers"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(route *gin.Engine) {
	route.POST("/register", controllers.Register)
	route.POST("/login", controllers.Login)
	// route.GET("/users", controllers.Users)
	// route.GET("/user/:id", controllers.UserByID)
	route.GET("/logout", controllers.Logout)
	route.GET("/home", controllers.Home)
}
