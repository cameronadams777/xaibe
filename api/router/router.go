package router

import (
	"api/controllers"
	"api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	// Group all endpoints under /api path
	api := app.Group("/api")

	// Authentication
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)

	api.POST("/webhook", controllers.WebHook)

	tokens := api.Group("/service_tokens", middleware.Protected())
	tokens.GET("/:team_id", controllers.GetServiceTokensByTeam)
	tokens.GET("/:application_id", controllers.GetServiceTokenByApplication)
	tokens.POST("/:application_id", controllers.CreateNewToken)
	tokens.DELETE("/", controllers.DeleteToken)

}
