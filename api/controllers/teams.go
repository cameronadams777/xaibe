package controllers

import (
	"api/assertions"
	"api/models"
	"api/services/teams_service"
	"api/services/users_service"
	"api/structs"
	"api/structs/invite_status"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CreateNewTeamInput struct {
	Name string `json:"teamName" binding:"required"`
}

type InviteExistingUserToTeamInput struct {
	UserId int `json:"user_id" binding:"required"`
	TeamId int `json:"team_id" binding:"required"`
}

type UpdateTeamInviteInput struct {
	InviteId int `json:"invite_id" binding:"required"`
	Status   int `json:"invite_status" binding:"required"`
}

func GetAllTeams(c *gin.Context) {
	// Fetch paginated teams list
	teams := teams_service.GetAllTeams()
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Teams found.", "data": teams})
}

func GetTeamById(c *gin.Context) {
	// Get ID from params
	team_input_param := c.Param("team_id")
	team_id, conv_err := strconv.Atoi(team_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting team by id.", "data": nil})
		return
	}

	// Fetch team by id
	team, err := teams_service.GetTeamById(team_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Team not found.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	membership_err := assertions.UserIsMemberOfTeam(team.ID, uint(authScope.UserID))

	if membership_err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "User not a member of the specified team.", "data": nil})
		return
	}

	// Return team info
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Team found.", "data": team})
}

func CreateNewTeam(c *gin.Context) {
	// Generate a new team with the name
	// of the team gathered from the request body
	var input CreateNewTeamInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	current_user, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred.", "data": nil})
		return
	}

	created_team, creation_err := teams_service.CreateTeam(input.Name, *current_user)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the requested team.", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Team created.", "data": created_team})
}

func DeleteTeam(c *gin.Context) {
	team_input_param := c.Param("team_id")
	team_id, conv_err := strconv.Atoi(team_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting team by id.", "data": nil})
		return
	}

	team_to_delete, err := teams_service.GetTeamById(team_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Team not found.", "data": nil})
		return
	}

	if team_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Team has already been deleted.", "data": err})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	team_manager_err := assertions.UserIsManagerOfTeam(team_to_delete.ID, uint(authScope.UserID))

	if team_manager_err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	deleted_team, _ := teams_service.DeleteTeam(team_id)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Team successfully deleted.", "data": deleted_team})
}

func RemoveUserFromTeam(c *gin.Context) {
	team_input_param := c.Param("team_id")
	team_id, _ := strconv.Atoi(team_input_param)

	user_input_param := c.Param("user_id")
	user_id, _ := strconv.Atoi(user_input_param)

	team, err := teams_service.GetTeamById(team_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Team not found.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	if user_id == authScope.UserID {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You cannot remove yourself from a team.", "data": nil})
		return
	}

	team_manager_err := assertions.UserIsManagerOfTeam(team.ID, uint(authScope.UserID))

	if team_manager_err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	user_is_member_of_team := false

	for _, user := range team.Users {
		if int(user.ID) == user_id {
			user_is_member_of_team = true
			break
		}
	}

	if !user_is_member_of_team {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "User does not belong to the specified.", "data": nil})
		return
	}

	updated_team, update_err := teams_service.RemoveUserFromTeam(team_id, user_id)

	if update_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred removing the user.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User successfully removed.", "data": updated_team})
}

func GetTeamInvites(c *gin.Context) {
	// Retrieve all invites for a given user
	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user, err := users_service.GetUserById(authScope.UserID)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred retrieving current user.", "data": nil})
		return
	}

	invites, invite_err := teams_service.GetPendingTeamInvites(user.Email)

	if invite_err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred retrieving the user's invites.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Invites retrieved.", "data": &invites})
}

func InviteExistingUserToTeam(c *gin.Context) {
	// Invite an existing user to a team passed in request body
	var input InviteExistingUserToTeamInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	user, find_user_err := users_service.GetUserById(input.UserId)

	if find_user_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Could not find requested user.", "data": nil})
		return
	}

  data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

  if user.ID == uint(authScope.UserID) {
 		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You cannot invite yourself to a team.", "data": nil})
		return 
  }

	team, find_team_err := teams_service.GetTeamById(input.TeamId)

	if find_team_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting team by id.", "data": nil})
		return
	}
	
	team_manager_err := assertions.UserIsManagerOfTeam(team.ID, uint(authScope.UserID))

	if team_manager_err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	user_is_member_of_team := false

	for _, user := range team.Users {
		if int(user.ID) == input.UserId {
			user_is_member_of_team = true
			break
		}
	}

	if user_is_member_of_team {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "User is already a member of this team.", "data": nil})
		return
	}

	_, invite_create_err := teams_service.CreateInvite(uint(input.TeamId), uint(authScope.UserID), user.Email)

	if invite_create_err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "An error occurred inviting the requested user to your team.", "data": nil})
		return
	}

	// TODO: ush new event to websocket notifying user that there is a pending invitation
}

func UpdateTeamInviteStatus(c *gin.Context) {
	// Update the status of an invite
	var input UpdateTeamInviteInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	updated_invite, err := teams_service.UpdateInvite(uint(input.InviteId), *&models.TeamInvite{Status: invite_status.InviteStatus(input.Status)})

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while updating your invite status.", "data": nil})
		return
	}

	if invite_status.InviteStatus(input.Status) == invite_status.ACCEPTED {
		// Associate user team id on invite
		_, update_err := teams_service.AddUserToTeam(int(updated_invite.TeamID), authScope.UserID)

		if update_err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred removing the user.", "data": nil})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Status updated.", "data": updated_invite})
}
