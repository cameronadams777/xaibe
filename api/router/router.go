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
	api.POST("/refresh_token", controllers.RefreshToken)

	api.POST("/webhook", controllers.WebHook)
	api.GET("/ws", controllers.ServeWS)

	reset_password := api.Group("/reset_password")
	reset_password.POST("/send_code", controllers.SendResetPasswordEmail)
	reset_password.POST("/validate", controllers.ValidateResetPasswordCode)
	reset_password.POST("/reset", controllers.ResetUserPassword)

	users := api.Group("/users", middleware.Protected())
	users.GET("/me", controllers.GetUserDetails)
	users.GET("/", controllers.GetAllUsers)
	users.GET("/:user_id", controllers.GetUserById)
	users.POST("/", controllers.InviteNewUser)
	users.PATCH("/", controllers.UpdateUser)
	users.DELETE("/:user_id", middleware.IsAdmin(), controllers.DeleteUser)

	teams := api.Group("/teams", middleware.Protected())
	teams.GET("/", controllers.GetAllTeams)
	teams.GET("/:team_id", controllers.GetTeamById)
	teams.GET("/invites", controllers.GetTeamInvites)
	teams.POST("/", controllers.CreateNewTeam)
	teams.POST("/invites", controllers.InviteExistingUserToTeam)
	teams.PATCH("/invites", controllers.UpdateTeamInviteStatus)
	teams.DELETE("/:team_id", controllers.DeleteTeam)
	teams.DELETE("/:team_id/user/:user_id", controllers.RemoveUserFromTeam)

	applications := api.Group("/applications", middleware.Protected())
	applications.GET("/:application_id", controllers.GetApplicationById)
	applications.GET("/:application_id/service_tokens", controllers.GetApplicationServiceTokens)
	applications.POST("/", controllers.CreateNewApplication)
	applications.PATCH("/:application_id/alert_schema", controllers.AddSchemaToApplication)
	applications.DELETE("/:application_id", controllers.DeleteApplication)

	alerts := api.Group("/alerts", middleware.Protected())
	alerts.GET("/", controllers.GetAllAlerts)
	alerts.GET("/applications/:application_id", controllers.GetApplicationAlerts)

	schemas := api.Group("/schemas", middleware.Protected())
	schemas.POST("/", controllers.CreateApplicationAlertSchema)

	tokens := api.Group("/service_tokens", middleware.Protected())
	tokens.POST("/:application_id", controllers.CreateNewToken)
	tokens.DELETE("/:token_id", controllers.DeleteToken)

  payments := api.Group("/payments", middleware.Protected())
  payments.POST("/create-customer", controllers.CreateNewStripeCustomer)
  payments.POST("/create-new-team-subscription", controllers.CreateNewTeamSubscription)
}
