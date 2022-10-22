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

	email := api.Group("/email")
	email.POST("/subscribe", controllers.SubscribeNewUser)

	users := api.Group("/users", middleware.Protected())
	users.GET("/me", controllers.GetUserDetails)
	users.GET("/", middleware.IsAdmin(), controllers.GetAllUsers)
	users.PATCH("/", controllers.UpdateUser)
	users.GET("/:user_id", controllers.GetUserById)
	users.DELETE("/:user_id", middleware.IsAdmin(), controllers.DeleteUser)

	teams := api.Group("/teams", middleware.Protected())
	teams.GET("/", middleware.IsAdmin(), controllers.GetAllTeams)
	teams.GET("/:team_id", controllers.GetTeamById)
	teams.GET("/:team_id/applications", controllers.GetAllTeamApplications)
	teams.POST("/", middleware.IsAdmin(), controllers.CreateNewTeam)
	teams.DELETE("/:team_id", controllers.DeleteTeam)

	applications := api.Group("/applications", middleware.Protected())
	applications.GET("/:application_id", controllers.GetApplicationById)
	applications.GET("/:application_id/service_tokens", controllers.GetApplicationServiceTokens)
	applications.GET("/:application_id/alerts", controllers.GetApplicationAlerts)
	applications.POST("/", controllers.CreateNewApplication)
	applications.DELETE("/:application_id", controllers.DeleteApplication)

	tokens := api.Group("/service_tokens", middleware.Protected())
	tokens.POST("/:application_id", controllers.CreateNewToken)
	tokens.DELETE("/:token_id", controllers.DeleteToken)

}
