package routes

import (
	"exam-server/controllers"

	"github.com/gin-gonic/gin"
)

func StudentRoutes(route *gin.RouterGroup) {
	route.GET("/ping", controllers.Ping)
	route.GET("/test/:id", controllers.GetTest)
	route.POST("/response", controllers.CreateResponse)
	route.POST("/responses", controllers.CreateResponses)
	route.GET("/responses/:user_id/:test_id", controllers.GetResponsePerUserPerTest)
	route.GET("/results/:user_id", controllers.GetResultsPerUser)
}

func AdminRoutes(route *gin.RouterGroup) {
	// all the routes for admin
	route.GET("/ping", controllers.Ping)
	route.GET("/tests", controllers.GetTests)
	route.GET("/test/:id", controllers.GetTest)
	route.POST("/test", controllers.CreateTest)
	route.PUT("/test/:id", controllers.UpdateTest)
	route.DELETE("/test/:id", controllers.DeleteTest)
	route.GET("/responses/:test_id", controllers.GetResponsesPerTest)
	route.GET("/results/:test_id", controllers.GetResultsPerTest)
	route.GET("/users", controllers.GetUsers)
	route.GET("/user/:id", controllers.GetUser)
	route.GET("/response/:id", controllers.GetResponse)
	route.GET("/questions", controllers.GetQuestions)
	route.GET("/question/:id", controllers.GetQuestion)
	route.POST("/question", controllers.CreateQuestion)
	route.POST("/questions", controllers.CreateQuestions)
	route.PUT("/question/:id", controllers.UpdateQuestion)
	route.DELETE("/question/:id", controllers.DeleteQuestion)
}
