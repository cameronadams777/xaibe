package controllers

import (
	"api/assertions"
	"api/config"
	"api/models"
	"api/services/stripe_service"
	"api/services/teams_service"
	"api/services/users_service"
	"api/structs"
	"api/structs/invite_status"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CreateNewTeamInput struct {
	Name            string `json:"teamName" binding:"required"`
  NumberOfSeats   uint   `json:"numberOfSeats" binding:"required"`
}

type InviteExistingUserToTeamInput struct {
	UserId          string `json:"user_id" binding:"required"`
	TeamId          string `json:"team_id" binding:"required"`
}

type UpdateTeamInviteInput struct {
	InviteId        int `json:"invite_id" binding:"required"`
	Status          int `json:"invite_status" binding:"required"`
}

type CreateNewTeamResponse struct {
  Team            *models.Team  `json:"team"`
  IntentID        string        `json:"intentId"`
  ClientSecret    string        `json:"clientSecret"`
}

func GetAllTeams(c *gin.Context) {
	// Fetch paginated teams list
	teams := teams_service.GetAllTeams()
	c.JSON(http.StatusOK, teams)
}

func GetTeamById(c *gin.Context) {
	// Get ID from params
	team_input_param := c.Param("team_id")
	team_id, conv_err := uuid.Parse(team_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting team by id."})
		return
	}

	// Fetch team by id
	team, err := teams_service.GetTeamById(team_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Team not found."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	membership_err := assertions.UserIsMemberOfTeam(team.ID, authScope.UserID)

	if membership_err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "User not a member of the specified team."})
		return
	}

	// Return team info
	c.JSON(http.StatusOK, team)
}

func CreateNewTeam(c *gin.Context) {
	// Generate a new team with the name
	// of the team gathered from the request body
	var input CreateNewTeamInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	current_user, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil || len(current_user.StripeId) == 0 {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred."})
		return
	}

	created_team, creation_err := teams_service.CreateTeam(input.Name, input.NumberOfSeats, *current_user)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred while creating the requested team."})
		return
	}

  subscription, subscription_create_err := stripe_service.CreateSubscription(stripe_service.SubscriptionData{
    CustomerId: current_user.StripeId,
    PriceId: config.Get("STRIPE_TEAM_PRODUCT_ID"),
    Quantity: input.NumberOfSeats,
    Metadata: map[string]string{
      "TeamId": created_team.ID.String(),
    },
  })

  if subscription_create_err != nil {
    fmt.Println(subscription_create_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred while creating the requested team."})
    teams_service.PermDeleteTeam(created_team.ID)
    return 
  }

  updated_team, update_err := teams_service.UpdateTeam(created_team.ID, models.Team{
    SubscriptionId: &subscription.ID,
  })

  if update_err != nil {
    fmt.Println(subscription_create_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred while creating the requested team."})
    teams_service.PermDeleteTeam(created_team.ID)
    return 
  }

  c.JSON(http.StatusCreated, CreateNewTeamResponse{Team: updated_team, IntentID: subscription.LatestInvoice.PaymentIntent.ID, ClientSecret: subscription.LatestInvoice.PaymentIntent.ClientSecret})
}

func DeleteTeam(c *gin.Context) {
	team_input_param := c.Param("team_id")
	team_id, conv_err := uuid.Parse(team_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting team by id."})
		return
	}

	team_to_delete, err := teams_service.GetTeamById(team_id)

  subscription_cancellation_err := stripe_service.CancelSubscription(*team_to_delete.SubscriptionId)

  if subscription_cancellation_err != nil {
    fmt.Println(subscription_cancellation_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred attempting to delete the requested team."})
    return 
  }

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Team not found."})
		return
	}

	if team_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Team has already been deleted."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	team_manager_err := assertions.UserIsManagerOfTeam(team_to_delete.ID, authScope.UserID)

	if team_manager_err != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	deleted_team, _ := teams_service.DeleteTeam(team_id)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Team successfully deleted.", "data": deleted_team})
}

func RemoveUserFromTeam(c *gin.Context) {
	team_input_param := c.Param("team_id")
	team_id, _ := uuid.Parse(team_input_param)

	user_input_param := c.Param("user_id")
	user_id, _ := uuid.Parse(user_input_param)

	team, err := teams_service.GetTeamById(team_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Team not found."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	if user_id == authScope.UserID {
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You cannot remove yourself from a team."})
		return
	}

	team_manager_err := assertions.UserIsManagerOfTeam(team.ID, authScope.UserID)

	if team_manager_err != nil {
    fmt.Println(team_manager_err)
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You do not have permission to perform that action."})
		return
	}

	user_is_member_of_team := false

	for _, user := range team.Users {
		if user.ID == user_id {
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
    fmt.Println(update_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred removing the user."})
		return
	}

	c.JSON(http.StatusOK, updated_team)
}

func GetTeamInvites(c *gin.Context) {
	// Retrieve all invites for a given user
	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user, err := users_service.GetUserById(authScope.UserID)

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred retrieving current user."})
		return
	}

	invites, invite_err := teams_service.GetPendingTeamInvites(user.Email)

	if invite_err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred retrieving the user's invites."})
		return
	}

	c.JSON(http.StatusOK, &invites)
}

func InviteExistingUserToTeam(c *gin.Context) {
	// Invite an existing user to a team passed in request body
	var input InviteExistingUserToTeamInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

  parsed_user_id, user_uuid_err := uuid.Parse(input.UserId)

  if user_uuid_err != nil {
    fmt.Println(user_uuid_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Invalid User ID"})
    return
  }

	user, find_user_err := users_service.GetUserById(parsed_user_id)

	if find_user_err != nil {
    fmt.Println(find_user_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Could not find requested user."})
		return
	}

  data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

  if user.ID == authScope.UserID {
 		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You cannot invite yourself to a team."})
		return 
  }

  parsed_team_id, team_uuid_err := uuid.Parse(input.TeamId)

  if team_uuid_err != nil {
    fmt.Println(team_uuid_err)
    c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You cannot invite yourself to a team."})
    return 
  }

	team, find_team_err := teams_service.GetTeamById(parsed_team_id)

	if find_team_err != nil {
    fmt.Println(find_team_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting team by id."})
		return
	}
	
	team_manager_err := assertions.UserIsManagerOfTeam(team.ID, authScope.UserID)

	if team_manager_err != nil {
    fmt.Println(team_manager_err)
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You do not have permission to perform that action."})
		return
	}

	user_is_member_of_team := false

	for _, user := range team.Users {
		if user.ID == parsed_user_id {
			user_is_member_of_team = true
			break
		}
	}

	if user_is_member_of_team {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "User is already a member of this team.", "data": nil})
		return
	}

	_, invite_create_err := teams_service.CreateInvite(parsed_team_id, authScope.UserID, user.Email)

	if invite_create_err != nil {
    fmt.Println(invite_create_err)
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "An error occurred inviting the requested user to your team."})
		return
	}

	// TODO: Push new event to websocket notifying user that there is a pending invitation
}

func UpdateTeamInviteStatus(c *gin.Context) {
	// Update the status of an invite
	var input UpdateTeamInviteInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	updated_invite, err := teams_service.UpdateInvite(uint(input.InviteId), *&models.TeamInvite{Status: invite_status.InviteStatus(input.Status)})

	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred while updating your invite status."})
		return
	}

	if invite_status.InviteStatus(input.Status) == invite_status.ACCEPTED {
		// Associate user team id on invite
		_, update_err := teams_service.AddUserToTeam(updated_invite.TeamID, authScope.UserID)

		if update_err != nil {
      fmt.Println(update_err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred removing the user."})
			return
		}
	}

	c.JSON(http.StatusOK, updated_invite)
}
