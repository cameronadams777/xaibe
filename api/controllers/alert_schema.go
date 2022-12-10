package controllers

import (
	"api/assertions"
	"api/models"
	"api/services/alert_schemas_service"
	"api/services/applications_service"
	"api/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateNewSchemaInput struct {
	ApplicationID int    `json:"applicationId" binding:"required"`
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

	_, fetch_err := applications_service.GetApplicationById(input.ApplicationID)

	if fetch_err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found.", "data": nil})
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

	new_alert_schema := models.AlertSchema{
		ApplicationID: uint(input.ApplicationID),
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
