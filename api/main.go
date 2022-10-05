package main

import (
	"api/cache"
	"api/config"
	"api/database"
	"api/router"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	// Connect to postgres database
	database.ConnectDB()

	// Connect to Redis cache
	cache.ConnectRedis()

	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{config.Get("APP_HOST_NAME")},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
	}))

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "I am health! 💪",
		})
	})

	router.SetupRouter(app)

	app.Run(":5000")
}
