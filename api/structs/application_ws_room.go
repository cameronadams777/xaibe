package structs

type ApplicationWSRoom struct {
	ApplicationID   uint
	ApplicationName string
	Client          map[*Client]bool
	Broadcast       chan *string
	Register        chan *Client
	Unregister      chan *Client
}

func NewRoom(application_id uint) *ApplicationWSRoom {
	room := &ApplicationWSRoom{
		ApplicationID: application_id,
		Broadcast:     make(chan *string),
		Register:      make(chan *Client),
		Unregister:    make(chan *Client),
		Client:        make(map[*Client]bool),
	}

	go room.run()
	return room
}

func (r *ApplicationWSRoom) run() {
	for {
		select {
		case client := <-r.Register:
			r.Client[client] = true
		case client := <-r.Unregister:
			if _, ok := r.Client[client]; ok {
				delete(r.Client, client)
				close(client.Send)
			}
		case message := <-r.Broadcast:
			for client := range r.Client {
				client.Send <- message
			}
		}
	}
}
