package controllers

import (
	"api/assertions"
	"api/helpers"
	"api/models"
	"api/services/alert_schemas_service"
	"api/services/applications_service"
	"api/services/service_tokens_service"
	"api/structs"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AlertSchemaInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

type CreateNewApplicationInput struct {
	TeamId      *uint            `json:"teamId" binding:"required_without=UserId"`
	UserId      *uint            `json:"userId" binding:"required_without=TeamId"`
	Name        string           `json:"applicationName" binding:"required"`
	AlertSchema AlertSchemaInput `json:"alertSchema" binding:"-"`
}

func GetApplicationById(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := strconv.Atoi(application_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting application by id.", "data": nil})
		return
	}

	application, fetch_err := applications_service.GetApplicationById(application_id)

	if fetch_err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(application.ID, uint(authScope.UserID))
	team_membership_error := assertions.UserIsMemberOfTeamApplication(application.ID, uint(authScope.UserID))

	if user_ownership_error != nil && team_membership_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Application found.", "data": application})
}

func CreateNewApplication(c *gin.Context) {
	var input CreateNewApplicationInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	new_application := models.Application{
		UserID:   input.UserId,
		TeamID:   input.TeamId,
		UniqueId: uuid.NewString(),
		Name:     input.Name,
	}

	if input.AlertSchema != (AlertSchemaInput{}) {
		new_application.AlertSchema = models.AlertSchema{
			Title:       input.AlertSchema.Title,
			Description: input.AlertSchema.Description,
			Link:        input.AlertSchema.Link,
		}
	}

	created_application, creation_err := applications_service.CreateApplication(new_application)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the requested application.", "data": nil})
		return
	}

	alert_schema, find_schema_err := alert_schemas_service.GetAlertSchemaByApplicationId(int(created_application.ID))

	if find_schema_err == nil {
		applications_service.UpdateApplication(int(created_application.ID), models.Application{AlertSchemaID: &alert_schema.ID})
	}

	application_with_schema, _ := applications_service.GetApplicationById(int(created_application.ID))

	fmt.Println(helpers.PrettyPrint(application_with_schema))

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Application created.", "data": application_with_schema})
}

func DeleteApplication(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := strconv.Atoi(application_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting application by id.", "data": nil})
		return
	}

	application_to_delete, err := applications_service.GetApplicationById(application_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found.", "data": nil})
		return
	}

	if application_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Application has already been deleted.", "data": err})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(application_to_delete.ID, uint(authScope.UserID))
	team_manager_error := assertions.UserIsManagerOfTeamApplication(application_to_delete.ID, uint(authScope.UserID))

	if user_ownership_error != nil && team_manager_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	deleted_application, _ := applications_service.DeleteApplication(application_id)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Application successfully deleted.", "data": deleted_application})
}

func AddSchemaToApplication(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := strconv.Atoi(application_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting application by id.", "data": nil})
		return
	}

	var input AlertSchemaInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	application_to_update, err := applications_service.GetApplicationById(application_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(application_to_update.ID, uint(authScope.UserID))
	team_manager_error := assertions.UserIsManagerOfTeamApplication(application_to_update.ID, uint(authScope.UserID))

	if user_ownership_error != nil && team_manager_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	alert_schema_to_create := models.AlertSchema{
		ApplicationID: application_to_update.ID,
		Title:         input.Title,
		Description:   input.Description,
		Link:          input.Link,
	}

	_, schema_create_err := alert_schemas_service.CreateNewAlertSchema(alert_schema_to_create)

	if schema_create_err != nil {
		fmt.Println(schema_create_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the alert schema for the specified application.", "data": nil})
		return
	}

	updated_application, _ := applications_service.GetApplicationById(application_id)

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Schema added to application.", "data": updated_application})
}

func GetApplicationServiceTokens(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := strconv.Atoi(application_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting application by id.", "data": nil})
		return
	}

	_, err := applications_service.GetApplicationById(application_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found when requesting service tokens.", "data": nil})
		return
	}

	service_tokens := service_tokens_service.GetAllServiceTokensByApplicationId(application_id)
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Tokens found.", "data": gin.H{"tokens": service_tokens}})
}
