package controllers

import (
	"api/config"
	"api/models"
	"api/services/sparkpost_service"
	"api/services/teams_service"
	"api/services/users_service"
	"api/structs"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// TODO: Write helper function that returns a JSON response but filters out sensitive data

type UpdateUserInput struct {
	UserId  string      `json:"userId" binding:"required"`
	Updates models.User `json:"user" binding:"-"`
}

type InviteNewUserInput struct {
	Email  string `json:"email" binding:"required"`
	TeamId *string   `json:"teamId" binding:"-"`
}

func GetUserDetails(c *gin.Context) {
	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	_, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
    fmt.Println(current_user_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred."})
		return
	}

	user, err := users_service.GetUserById(authScope.UserID)

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "User not found."})
		return
	}

	c.JSON(http.StatusOK, user)
}

func GetAllUsers(c *gin.Context) {
	users := users_service.GetAllUsers()
	c.JSON(http.StatusOK, users)
}

func GetUserById(c *gin.Context) {
	user_input_param := c.Param("user_id")
	user_id, conv_err := uuid.Parse(user_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting user by id."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	_, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
    fmt.Println(current_user_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred."})
		return
	}

	user, err := users_service.GetUserById(user_id)

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "User not found."})
		return
	}

	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var input UpdateUserInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	current_user, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An unknown error occurred."})
		return
	}

  parsed_user_id, uuid_err := uuid.Parse(input.UserId)

  if uuid_err != nil {
    fmt.Println(uuid_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Invalid User ID"})
    return
  }

	if current_user.ID != parsed_user_id && !current_user.IsAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You do not have permission to update this user's account"})
		return
	}

	updated_user, err := users_service.UpdateUser(parsed_user_id, input.Updates)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "User not found."})
		return
	}

	c.JSON(http.StatusOK, updated_user)
}

func DeleteUser(c *gin.Context) {
	user_input_param := c.Param("user_id")
	user_id, conv_err := uuid.Parse(user_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting user by id."})
		return
	}

	user_to_delete, err := users_service.GetUserById(user_id)

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "User not found."})
		return
	}

	if user_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "User has already been deleted."})
		return
	}

  deleted_user, _ := users_service.UpdateUser(user_id, models.User{UUIDBaseModel: models.UUIDBaseModel{DeletedAt: gorm.DeletedAt{Time: time.Now()}}})

	c.JSON(http.StatusOK, deleted_user)
}

func InviteNewUser(c *gin.Context) {
	var input InviteNewUserInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	_, find_existing_user_err := users_service.GetUserByEmail(input.Email)

	if find_existing_user_err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "A user with that email already exists."})
		return
	}

	if input.TeamId != nil {
		// Add new record to invites table
    parsed_team_id, uuid_err := uuid.Parse(*input.TeamId)
    if uuid_err == nil {
      
      team, get_team_err := teams_service.GetTeamById(parsed_team_id)

      if get_team_err != nil {
        fmt.Println(get_team_err)
        c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Team not found."})
        return 
      }

      if len(team.Users) + 1 > int(team.ActiveNumberOfSeats) {
        c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Team's subscription does not support more users."})
        return
      }

		  teams_service.CreateInvite(parsed_team_id, authScope.UserID, input.Email)
    }
	}

	templateElements := ResetPasswordTemplateElements{
		Link: config.Get("APP_HOST_NAME"),
	}

	sparkpost_service.SendEmail(input.Email, "invite_new_user", templateElements)
}
