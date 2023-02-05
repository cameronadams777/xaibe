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
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred.", "data": nil})
		return
	}

	user, err := users_service.GetUserById(authScope.UserID)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User found.", "data": user})
}

func GetAllUsers(c *gin.Context) {
	users := users_service.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"status": "error", "message": "", "data": users})
}

func GetUserById(c *gin.Context) {
	user_input_param := c.Param("user_id")
	user_id, conv_err := uuid.Parse(user_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting user by id.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	_, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred.", "data": nil})
		return
	}

	user, err := users_service.GetUserById(user_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User found.", "data": gin.H{"user": user}})
}

func UpdateUser(c *gin.Context) {
	var input UpdateUserInput

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

  parsed_user_id, uuid_err := uuid.Parse(input.UserId)

  if uuid_err != nil {
    fmt.Println(uuid_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Invalid User ID", "data": nil})
    return
  }

	if current_user.ID != parsed_user_id && !current_user.IsAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to update this user's account", "data": nil})
		return
	}

	updated_user, err := users_service.UpdateUser(parsed_user_id, input.Updates)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User successfully updated.", "data": gin.H{"user": updated_user}})
}

func DeleteUser(c *gin.Context) {
	user_input_param := c.Param("user_id")
	user_id, conv_err := uuid.Parse(user_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting user by id.", "data": nil})
		return
	}

	user_to_delete, err := users_service.GetUserById(user_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	if user_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "User has already been deleted.", "data": err})
		return
	}

  deleted_user, _ := users_service.UpdateUser(user_id, models.User{UUIDBaseModel: models.UUIDBaseModel{DeletedAt: gorm.DeletedAt{Time: time.Now()}}})

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User successfully deleted.", "data": gin.H{"user": deleted_user}})
}

func InviteNewUser(c *gin.Context) {
	var input InviteNewUserInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	_, find_existing_user_err := users_service.GetUserByEmail(input.Email)

	if find_existing_user_err == nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "A user with that email already exists.", "data": nil})
		return
	}

	if input.TeamId != nil {
		// Add new record to invites table
    parsed_team_id, uuid_err := uuid.Parse(*input.TeamId)
    if uuid_err == nil {
		  teams_service.CreateInvite(parsed_team_id, authScope.UserID, input.Email)
    }
	}

	templateElements := ResetPasswordTemplateElements{
		Link: config.Get("APP_HOST_NAME"),
	}

	sparkpost_service.SendEmail(input.Email, "invite_new_user", templateElements)
}
