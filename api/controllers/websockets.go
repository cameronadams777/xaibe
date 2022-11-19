package controllers

import (
	"api/config"
	"api/services/users_service"
	"api/websockets"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // TODO: Need to make this work like CORS
	},
}

func ServeWS(c *gin.Context) {
	token_string := c.Query("token")
	token, _ := jwt.Parse(token_string, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.Get("AUTH_TOKEN_SECRET")), nil
	})
	claims := token.Claims.(jwt.MapClaims)
	user_id, _ := strconv.Atoi(claims["iss"].(string))
	current_user, _ := users_service.GetUserById(user_id)

	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Println(err.Error())
		return
	}

	conn := &websockets.Connection{Send: make(chan []byte, 256), Socket: ws}
	// TODO: Also take into account team applications
	for _, application := range current_user.Applications {
		room_id := strconv.Itoa(int(application.ID)) + ":" + application.UniqueId
		s := websockets.Subscription{Conn: conn, Room: room_id}
		websockets.Pool.Register <- s
		go s.WritePump()
		go s.ReadPump()
	}

}
