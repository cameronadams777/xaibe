package controllers

import (
	"api/models"
	"api/services/alert_schemas_service"
	"api/services/applications_service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNewSchemaInput struct {
	ApplicationID int    `json:"application_id" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description" binding:"required"`
}

func CreateApplicationAlertSchema(c *gin.Context) {
	var input CreateNewSchemaInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	_, fetch_err := applications_service.GetApplicationById(input.ApplicationID)

	if fetch_err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found.", "data": nil})
		return
	}

	// TODO: Add below logic
	// Ensure that current user is either the owner of the application
	// or a manager of said team
	// If not, throw error
	// If so, store schema and return 200 response

	new_alert_schema := models.AlertSchema{
		ApplicationID: uint(input.ApplicationID),
		Title:         input.Title,
		Description:   input.Description,
	}

	created_alert_schema, creation_err := alert_schemas_service.CreateNewAlertSchema(new_alert_schema)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the requested application.", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Schema created.", "data": created_alert_schema})
}
