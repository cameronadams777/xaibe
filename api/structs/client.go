package structs

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

const (
	// Time allowed to write a message to the peer
	write_wait = 10 * time.Second

	// Time allowed to read the next pong message from the peer
	pong_wait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	ping_period = (pong_wait * 9) / 10

	// Maximum message size allowed from peer
	max_message_size = 512
)

type Client struct {
	ID         string
	IP         string
	UserID     uint
	Connection *websocket.Conn
	Send       chan *string
}

func (c *Client) receiveMessages() {
	defer func() {
		// TODO: Also unregister client from all rooms that it's in
		c.Connection.Close()
	}()

	c.Connection.SetReadLimit(max_message_size)
	c.Connection.SetReadDeadline(time.Now().Add(pong_wait))

	c.Connection.SetPongHandler(func(app_data string) error {
		c.Connection.SetReadDeadline(time.Now().Add(pong_wait))
		return nil
	})
}

func (c *Client) sendMessages() {
	ticker := time.NewTicker(ping_period)

	defer func() {
		ticker.Stop()
		c.Connection.Close()
	}()

	for {
		select {
		case message, ok := <-c.Send:
			c.Connection.SetWriteDeadline(time.Now().Add(write_wait))
			if !ok {
				// The hub closed the channel.
				c.Connection.WriteMessage(websocket.CloseMessage, []byte("connection closed"))
				return
			}
			c.Connection.WriteJSON(message)
		case <-ticker.C:
			fmt.Println("ticker hit")
			c.Connection.SetWriteDeadline(time.Now().Add(write_wait))
			if err := c.Connection.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

func NewClient(client_id string, client_ip string, user_id uint, w http.ResponseWriter, r *http.Request) *Client {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println(err)
	}

	client := &Client{ID: client_id, IP: client_ip, UserID: user_id, Connection: conn, Send: make(chan *string, 256)}

	go client.sendMessages()
	go client.receiveMessages()

	return client
}
