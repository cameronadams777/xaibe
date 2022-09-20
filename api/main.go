package main

import (
	"api/config"
	"api/database"
	"api/router"
	"log"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

var RedisClient *redis.Client

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     config.Get("REDIS_HOST"),
		Password: config.Get("REDIS_PASSWORD"),
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	RedisClient = client

	database.ConnectDB()

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
