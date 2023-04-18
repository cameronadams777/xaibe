package controllers

import (
	"api/assertions"
	"api/models"
	"api/services/alert_schemas_service"
	"api/services/applications_service"
	"api/services/service_tokens_service"
	"api/structs"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type AlertSchemaInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

type CreateNewApplicationInput struct {
	TeamId      *string            `json:"teamId" binding:"required_without=UserId"`
	UserId      *string            `json:"userId" binding:"required_without=TeamId"`
	Name        string           `json:"applicationName" binding:"required"`
	AlertSchema AlertSchemaInput `json:"alertSchema" binding:"-"`
}

func GetApplicationById(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := uuid.Parse(application_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting application by id."})
		return
	}

	application, fetch_err := applications_service.GetApplicationById(application_id)

	if fetch_err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Application not found."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(application.ID, authScope.UserID)
	team_membership_error := assertions.UserIsMemberOfTeamApplication(application.ID, authScope.UserID)

	if user_ownership_error != nil && team_membership_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to perform that action.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, application)
}

func CreateNewApplication(c *gin.Context) {
	var input CreateNewApplicationInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	new_application := models.Application{
		UniqueId: uuid.NewString(),
		Name:     input.Name,
	}

  if input.UserId != nil {
    parsed_user_id, user_uuid_err := uuid.Parse(*input.UserId)

    if user_uuid_err != nil {
      fmt.Println(user_uuid_err)
      c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Invalid User ID"})
      return
    }
    new_application.UserID = &parsed_user_id
  } else if input.TeamId != nil {
    parsed_team_id, team_uuid_err := uuid.Parse(*input.TeamId)

    if team_uuid_err != nil {
      fmt.Println(team_uuid_err)
      c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Invalid Team ID"})
      return
    }
    new_application.TeamID = &parsed_team_id
  } else {
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "A valid User ID or Team ID must be provided"})
    return
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
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred while creating the requested application."})
		return
	}

	alert_schema, find_schema_err := alert_schemas_service.GetAlertSchemaByApplicationId(created_application.ID)

	if find_schema_err == nil {
		applications_service.UpdateApplication(created_application.ID, models.Application{AlertSchemaID: &alert_schema.ID})
	}

	application_with_schema, _ := applications_service.GetApplicationById(created_application.ID)

	c.JSON(http.StatusCreated, application_with_schema)
}

func DeleteApplication(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := uuid.Parse(application_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting application by id."})
		return
	}

	application_to_delete, err := applications_service.GetApplicationById(application_id)

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Application not found."})
		return
	}

	if application_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "Application has already been deleted."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(application_to_delete.ID, authScope.UserID)
	team_manager_error := assertions.UserIsManagerOfTeamApplication(application_to_delete.ID, authScope.UserID)

	if user_ownership_error != nil && team_manager_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You do not have permission to perform that action."})
		return
	}

	deleted_application, _ := applications_service.DeleteApplication(application_id)

	c.JSON(http.StatusOK, deleted_application)
}

func AddSchemaToApplication(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := uuid.Parse(application_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting application by id."})
		return
	}

	var input AlertSchemaInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	application_to_update, err := applications_service.GetApplicationById(application_id)

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Application not found."})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	user_ownership_error := assertions.UserOwnsApplication(application_to_update.ID, authScope.UserID)
	team_manager_error := assertions.UserIsManagerOfTeamApplication(application_to_update.ID, authScope.UserID)

	if user_ownership_error != nil && team_manager_error != nil {
		c.AbortWithStatusJSON(http.StatusForbidden, structs.ErrorMessage{Message: "You do not have permission to perform that action."})
		return
	}

	alert_schema_to_create := models.AlertSchema{
		ApplicationID: application_to_update.ID,
		Title:         input.Title,
		Description:   input.Description,
		Link:          input.Link,
	}

	created_schema, schema_create_err := alert_schemas_service.CreateNewAlertSchema(alert_schema_to_create)

	if schema_create_err != nil {
		fmt.Println(schema_create_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred while creating the alert schema for the specified application."})
		return
	}

  updated_application, update_err := applications_service.UpdateApplication(application_to_update.ID, models.Application{ 
    AlertSchemaID: &created_schema.ID, 
  })

  if update_err != nil {
    fmt.Println(update_err)
    c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred adding alert schema to application."})
    return
  }

	c.JSON(http.StatusOK, updated_application)
}

func GetApplicationServiceTokens(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := uuid.Parse(application_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Error requesting application by id."})
		return
	}

	_, err := applications_service.GetApplicationById(application_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Application not found when requesting service tokens."})
		return
	}

	service_tokens := service_tokens_service.GetAllServiceTokensByApplicationId(application_id)
	c.JSON(http.StatusOK, service_tokens)
}
