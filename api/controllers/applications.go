package controllers

import (
	"api/assertions"
	"api/initializers/cache"
	"api/models"
	"api/services/applications_service"
	"api/services/service_tokens_service"
	"api/structs"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type AlertSchemaInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	Link        string `json:"link" binding:"required"`
}

type CreateNewApplicationInput struct {
	TeamId      *uint             `json:"team_id" binding:"required_without=UserId"`
	UserId      *uint             `json:"user_id" binding:"required_without=TeamId"`
	Name        string            `json:"application_name" binding:"required"`
	AlertSchema *AlertSchemaInput `json:"alert_schema" binding:"-"`
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

	created_application, creation_err := applications_service.CreateApplication(new_application)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the requested application.", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Application created.", "data": created_application})
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

func GetApplicationAlerts(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := strconv.Atoi(application_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting application by id.", "data": nil})
		return
	}

	application, err := applications_service.GetApplicationById(application_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Application not found when requesting alerts.", "data": nil})
		return
	}

	var owner_id string

	if application.TeamID != nil {
		team_id := strconv.Itoa(int(*application.TeamID))
		owner_id = "team_" + team_id
	} else if application.UserID != nil {
		user_id := strconv.Itoa(int(*application.UserID))
		owner_id = "user_" + user_id
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred posting alert to application.", "data": nil})
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

	var alerts_as_json []map[string]interface{}

	scan_key := owner_id + ":" + "application_" + application_input_param + ":" + application.UniqueId + ":*"

	iter := cache.RedisClient.Scan(0, scan_key, 0).Iterator()
	for iter.Next() {
		var alert map[string]interface{}
		alert_trx := cache.RedisClient.Get(iter.Val())
		alert_string := alert_trx.Val()
		if alert_string != redis.Nil.Error() {
			json.Unmarshal([]byte(alert_trx.Val()), &alert)
			alerts_as_json = append(alerts_as_json, alert)
		}
	}
	if iter_err := iter.Err(); iter_err != nil {
		panic(iter_err)
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Alerts retrieved.", "data": alerts_as_json})
}
