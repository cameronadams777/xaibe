package main

import (
	"api/initializers/cache"
	"api/initializers/database"
	"api/router"
	"api/websockets"
	"fmt"
	"log"

	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	socketio "github.com/googollee/go-socket.io"
)

func main() {

	go websockets.Pool.Run()

	server := socketio.NewServer(nil)

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Connected:", s.ID())
		return nil
	})

	server.OnEvent("/", "msg", func(s socketio.Conn, msg string) string {
		fmt.Println("Receive Message : " + msg)
		s.Emit("reply", "OK")
		return "recv " + msg
	})

	server.OnError("/", func(s socketio.Conn, e error) {
		log.Println("meet error:", e)
	})

	server.OnDisconnect("/", func(s socketio.Conn, msg string) {
		fmt.Println("Somebody just close the connection ")
	})

	go server.Serve()
	defer server.Close()

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

	app.GET("/socket.io/", gin.WrapH(server))

	router.SetupRouter(app)

	app.Run(":5000")
}
