package routes

import (
	"exam-server/controllers"

	"github.com/gin-gonic/gin"
)

func PingRoutes(route *gin.Engine) {
	route.GET("/ping", controllers.Ping)
}
