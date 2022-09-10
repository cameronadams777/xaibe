package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Alert struct {
	Message string `json:"message"`
}

type AlertStore struct {
	alerts []Alert
}

var store *AlertStore

func setup_memory_store() {
	store = &AlertStore{
		alerts: make([]Alert, 0),
	}
}

func main() {
	app := gin.Default()

	setup_memory_store()

	app.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Am Health! 🚀",
		})
	})

	app.GET("/messages", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "success", "data": store.alerts})
	})

	app.POST("/webhook", func(c *gin.Context) {
		var alert Alert

		c.Bind(&alert)

		store.alerts = append(store.alerts, alert)

		c.Done()
	})

	app.Run(":5000")
}
