package controllers

import (
	"api/initializers/application_rooms"
	"api/services/users_service"
	"api/structs"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func HandleNewClientWS(c *gin.Context) {
	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	current_user, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred.", "data": nil})
		return
	}

	applications := current_user.Applications

	client_id := uuid.NewString()
	new_client := structs.NewClient(client_id, c.Request.RemoteAddr, current_user.ID, c.Writer, c.Request)

	for _, application := range applications {
		var owner_id string

		if application.TeamID != 0 {
			team_id := strconv.Itoa(int(application.TeamID))
			owner_id = "team_" + team_id
		} else if application.UserID != 0 {
			user_id := strconv.Itoa(int(application.UserID))
			owner_id = "user_" + user_id
		} else {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred posting alert to application.", "data": nil})
			return
		}
		application_id_as_string := strconv.Itoa(int(application.ID))
		room_id := owner_id + ":" + "application_" + application_id_as_string + ":" + application.UniqueId

		if _, ok := application_rooms.ApplicationRooms[room_id]; ok {
			if !ok {
				application_rooms.ApplicationRooms[room_id] = structs.NewRoom(application.ID)
			}
			application_rooms.ApplicationRooms[room_id].Register <- new_client
		}
	}
}
