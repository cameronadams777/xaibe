package controllers

import (
	"api/assertions"
	"api/models"
	"api/services/service_tokens_service"
	"api/structs"
	"fmt"
	"net/http"
	"time"

  "github.com/google/uuid"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateNewTokenInput struct {
	ApplicationID string `json:"applicationId" binding:"required"`
}

type CreateNewTokenResponse struct {
  Token string `json:"token"`
  ExpiresAt time.Time `json:"expires_at"`
}

func CreateNewToken(c *gin.Context) {
	// Get application id from request body
	var input CreateNewTokenInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

  parsed_application_id, uuid_err := uuid.Parse(input.ApplicationID)

  if uuid_err != nil {
    fmt.Println(uuid_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Invalid Application ID"})
  }

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(parsed_application_id, authScope.UserID)
	team_manager_error := assertions.UserIsManagerOfTeamApplication(parsed_application_id, authScope.UserID)

	if user_ownership_error != nil && team_manager_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You do not have permission to perform that action."})
		return
	}

	// Create token
	new_token, creation_err := service_tokens_service.CreateServiceToken(parsed_application_id)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred while creating the requested team."})
		return
	}

	// Return token
	c.JSON(http.StatusCreated, CreateNewTokenResponse{Token: new_token.Token, ExpiresAt: new_token.ExpiresAt})
}

func DeleteToken(c *gin.Context) {
	// Get token id from params
	token_input_param := c.Param("token_id")
	token_id, conv_err := uuid.Parse(token_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting token by id."})
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
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Application not found."})
		return
	}

	if token_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Service token has already been deleted."})
		return
	}

	deleted_token, _ := service_tokens_service.UpdateServiceToken(token_id, models.ServiceToken{Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now()}}})

	c.JSON(http.StatusOK, deleted_token)
}
