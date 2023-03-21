package main

import (
	"api/controllers"
	"api/initializers/cache"
	"api/initializers/database"
	"api/initializers/sparkpost"
	"api/initializers/stripe_client"
	"api/router"
	"api/websockets"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	websockets.CreateNewPool()
	go websockets.Pool.Run()

	// Create SparkPost Client
	sparkpost.CreateSparkPostClient()

  // Create Stripe Client
  stripe_client.CreateStripeClient();

	// Connect to postgres database
	database.ConnectDB()

	// Connect to Redis cache
	cache.ConnectRedis()

	app := gin.Default()

	app.Use(cors.Default())

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "I am health! 💪",
		})
	})

  // Auto-Update Xaibe Desktop
  app.GET("/auto_update/:platform/:current_version", controllers.CheckLatestAppVersion)
 

	router.SetupRouter(app)

	app.Run(":5000")
}
