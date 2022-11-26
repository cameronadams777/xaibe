package websockets

import (
	"sync"
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
)

var Pool *ConnectionPool

type Connection struct {
	mu     sync.Mutex
	Socket *websocket.Conn
	Send   chan []byte
}

type ConnectionPool struct {
	Rooms      map[string]map[*Connection]bool
	Broadcast  chan Message
	Register   chan Subscription
	Unregister chan Subscription
}

type Subscription struct {
	Conn *Connection
	Room string
}

type Message struct {
	Data []byte
	Room string
}

func (s *Subscription) WritePump() {
	c := s.Conn
	c.mu.Lock()
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		c.mu.Unlock()
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
		case s := <-p.Register:
			connections := p.Rooms[s.Room]
			if connections == nil {
				connections = make(map[*Connection]bool)
				p.Rooms[s.Room] = connections
			}
			p.Rooms[s.Room][s.Conn] = true
		case s := <-p.Unregister:
			connections := p.Rooms[s.Room]
			if connections != nil {
				if _, ok := connections[s.Conn]; ok {
					delete(connections, s.Conn)
					close(s.Conn.Send)
					if len(connections) == 0 {
						delete(p.Rooms, s.Room)
					}
				}
			}
		case m := <-p.Broadcast:
			connections := p.Rooms[m.Room]
			for c := range connections {
				select {
				case c.Send <- m.Data:
				default:
					close(c.Send)
					delete(connections, c)
					if len(connections) == 0 {
						delete(p.Rooms, m.Room)
					}
				}
			}
		}
	}
}

func CreateNewPool() {
	var connection_pool = ConnectionPool{
		Broadcast:  make(chan Message),
		Register:   make(chan Subscription),
		Unregister: make(chan Subscription),
		Rooms:      make(map[string]map[*Connection]bool),
	}
	Pool = &connection_pool
}
