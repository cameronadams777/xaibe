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

	organizations := api.Group("/organizations", middleware.Protected(), middleware.IsSystemUser())
	organizations.GET("/", controllers.GetAllOrganizations)
	organizations.GET("/:organization_id", controllers.GetOrganizationById)
	organizations.DELETE("/:organization_id", controllers.DeleteOrganization)

	users := api.Group("/users", middleware.Protected())
	users.GET("/", middleware.IsAdmin(), controllers.GetAllUsers)
	users.PATCH("/", controllers.UpdateUser)
	users.GET("/:user_id", controllers.GetUserById)
	users.DELETE("/:user_id", middleware.IsAdmin(), controllers.DeleteUser)

	teams := api.Group("/teams", middleware.Protected())
	teams.GET("/", middleware.IsAdmin(), controllers.GetAllTeams)
	teams.GET("/:team_id", controllers.GetTeamById)
	teams.POST("/", middleware.IsAdmin(), controllers.CreateNewTeam)
	teams.DELETE("/:team_id", controllers.DeleteTeam)

	tokens := api.Group("/service_tokens", middleware.Protected())
	tokens.GET("/:application_id", controllers.GetServiceTokenByApplication)
	tokens.POST("/:application_id", controllers.CreateNewToken)
	tokens.DELETE("/", controllers.DeleteToken)

}
