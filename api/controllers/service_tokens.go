package controllers

import (
	"api/assertions"
	"api/models"
	"api/services/service_tokens_service"
	"api/structs"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateNewTokenInput struct {
	ApplicationID int `json:"application_id" binding:"required"`
}

func CreateNewToken(c *gin.Context) {
	// Get application id from request body
	var input CreateNewTokenInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(uint(input.ApplicationID), uint(authScope.UserID))
	team_manager_error := assertions.UserIsManagerOfTeamApplication(uint(input.ApplicationID), uint(authScope.UserID))

	if user_ownership_error != nil && team_manager_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	// Create token
	new_token, creation_err := service_tokens_service.CreateServiceToken(input.ApplicationID)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the requested team.", "data": nil})
		return
	}

	// Return token
	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Service token created.", "data": gin.H{"token": new_token.Token, "expires_at": new_token.ExpiresAt}})
}

func DeleteToken(c *gin.Context) {
	// Get token id from params
	token_input_param := c.Param("token_id")
	token_id, conv_err := strconv.Atoi(token_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting token by id.", "data": nil})
		return
	}

	// TODO: ADD LOGIC BELOW
	// Get application by token id
	// Get team by application id
	// Check to see if user is in team
	// If not, throw error
	// Check to see if user is manager of team
	// If not, throw error

	// Update deleted_at field in token record
	token_to_delete, err := service_tokens_service.GetServiceTokenById(token_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found.", "data": nil})
		return
	}

	if token_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Service token has already been deleted.", "data": err})
		return
	}

	deleted_token, _ := service_tokens_service.UpdateServiceToken(token_id, models.ServiceToken{Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now()}}})

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Application successfully deleted.", "data": gin.H{"application": deleted_token}})
}
