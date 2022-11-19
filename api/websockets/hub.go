package websockets

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Connection struct {
	Socket *websocket.Conn
	Send   chan []byte
}

type ConnectionPool struct {
	Rooms      map[string]map[*Connection]bool
	Broadcast  chan Message
	Register   chan Subscription
	Unregister chan Subscription
}

var Pool = ConnectionPool{
	Broadcast:  make(chan Message),
	Register:   make(chan Subscription),
	Unregister: make(chan Subscription),
	Rooms:      make(map[string]map[*Connection]bool),
}

type Subscription struct {
	Conn *Connection
	Room string
}

type Message struct {
	Data []byte
	Room string
}

func (s Subscription) ReadPump() {
	c := s.Conn
	defer func() {
		Pool.Unregister <- s
		c.Socket.Close()
	}()
	c.Socket.SetReadLimit(maxMessageSize)
	c.Socket.SetReadDeadline(time.Now().Add(pongWait))
	c.Socket.SetPongHandler(func(string) error { c.Socket.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, msg, err := c.Socket.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		m := Message{msg, s.Room}
		Pool.Broadcast <- m
	}
}

func (s *Subscription) WritePump() {
	c := s.Conn
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Socket.Close()
	}()
	for {
		select {
		case message, ok := <-c.Send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}
			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Connection) write(mt int, payload []byte) error {
	c.Socket.SetWriteDeadline(time.Now().Add(writeWait))
	return c.Socket.WriteMessage(mt, payload)
}

func (p *ConnectionPool) Run() {
	for {
		select {
		case s := <-Pool.Register:
			connections := Pool.Rooms[s.Room]
			if connections == nil {
				connections = make(map[*Connection]bool)
				Pool.Rooms[s.Room] = connections
			}
			Pool.Rooms[s.Room][s.Conn] = true
		case s := <-Pool.Unregister:
			connections := Pool.Rooms[s.Room]
			if connections != nil {
				if _, ok := connections[s.Conn]; ok {
					delete(connections, s.Conn)
					close(s.Conn.Send)
					if len(connections) == 0 {
						delete(Pool.Rooms, s.Room)
					}
				}
			}
		case m := <-Pool.Broadcast:
			connections := Pool.Rooms[m.Room]
			for c := range connections {
				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(Pool.Rooms, m.Room)
					}
				}
			}
		}
	}
}
