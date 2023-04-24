package main

import (
	"api/controllers"
	"api/initializers/cache"
	"api/initializers/database"
	"api/initializers/sparkpost"
	"api/initializers/stripe_client"
	"api/router"
	"api/websockets"

  "log"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
  "github.com/getsentry/sentry-go"
)

func main() {
  err := sentry.Init(sentry.ClientOptions{
    EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
  })
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

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
