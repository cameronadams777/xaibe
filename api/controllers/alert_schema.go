package controllers

import (
	"api/assertions"
	"api/models"
	"api/services/alert_schemas_service"
	"api/services/applications_service"
	"api/structs"
	"fmt"
	"net/http"

  "github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

type CreateNewSchemaInput struct {
	ApplicationID string    `json:"applicationId" binding:"required"`
	Title         string `json:"title" binding:"required"`
	Description   string `json:"description" binding:"required"`
	Link          string `json:"link" binding:"required"`
}

func CreateApplicationAlertSchema(c *gin.Context) {
	var input CreateNewSchemaInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

  parsed_application_id, uuid_err := uuid.Parse(input.ApplicationID)

  if uuid_err != nil {
    fmt.Println(uuid_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Invalid Application ID", "data": nil})
    return
  }

	_, fetch_err := applications_service.GetApplicationById(parsed_application_id)

	if fetch_err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(parsed_application_id, authScope.UserID)
	team_manager_error := assertions.UserIsManagerOfTeamApplication(parsed_application_id, authScope.UserID)

	if user_ownership_error != nil && team_manager_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	new_alert_schema := models.AlertSchema{
		ApplicationID: parsed_application_id,
		Title:         input.Title,
		Description:   input.Description,
		Link:          input.Link,
	}

	created_alert_schema, creation_err := alert_schemas_service.CreateNewAlertSchema(new_alert_schema)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the requested application.", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Schema created.", "data": created_alert_schema})
}
