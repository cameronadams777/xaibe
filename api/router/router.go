package router

import (
	"api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(app *gin.Engine) {
	// Group all endpoints under /api path
	api := app.Group("/api")

	// Authentication
	api.POST("/login", controllers.Login)
	api.POST("/register", controllers.Register)
}
