package application_rooms

import "api/structs"

var ApplicationRooms map[string]*structs.ApplicationWSRoom

func Create() {
	application_rooms := make(map[string]*structs.ApplicationWSRoom)
	ApplicationRooms = application_rooms
}
