package main

import (
	"api/initializers/application_rooms"
	"api/initializers/cache"
	"api/initializers/database"
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

	// Create rooms
	application_rooms.Create()

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
