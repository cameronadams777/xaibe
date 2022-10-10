package controllers

import (
	"api/models"
	"api/services/service_tokens_service"
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

func GetServiceTokenById(c *gin.Context) {
	// Get application id from params
	// Get application by id
	// Get ID from params
	token_input_param := c.Param("token_id")
	token_id, conv_err := strconv.Atoi(token_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting application by id.", "data": nil})
		return
	}

	service_token, err := service_tokens_service.GetServiceTokenById(token_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Token not found.", "data": nil})
		return
	}

	// TODO: ADD LOGIC BELOW
	// Get application from service_token
	// Get team by application
	// Check to see if user is in team
	// If not, throw error
	// If so, check to see if user is manager of team
	// If not, throw error

	// If so, return non-expired/non-deleted service token
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Token found.", "data": gin.H{"token": service_token.Token, "expires_at": service_token.ExpiresAt}})
}

func CreateNewToken(c *gin.Context) {
	// Get application id from request body
	var input CreateNewTokenInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	// TODO: ADD LOGIC BELOW
	// Get team from application
	// Ensure that team owns application
	// Ensure that user is manager of team

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
