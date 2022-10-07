package main

import (
	"api/cache"
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

	app.Use(cors.Default())

	app.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"health": "I am health! 💪",
		})
	})

	router.SetupRouter(app)

	app.Run(":5000")
}