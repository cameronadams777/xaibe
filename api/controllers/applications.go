package controllers

import (
	"api/cache"
	"api/models"
	"api/services/applications_service"
	"api/services/service_tokens_service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type CreateNewApplicationInput struct {
	TeamId int    `json:"team_id" binding:"required_without=UserId"`
	UserId int    `json:"user_id" binding:"required_without=TeamId"`
	Name   string `json:"application_name" binding:"required"`
}

func CreateNewApplication(c *gin.Context) {
	var input CreateNewApplicationInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	new_application := models.Application{
		TeamID:   uint(input.TeamId),
		UserID:   uint(input.UserId),
		UniqueId: uuid.NewString(),
		Name:     input.Name,
	}

	creation_err := applications_service.CreateApplication(new_application)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the requested application.", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Application created.", "data": nil})
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

	// TODO: ADD BELOW LOGIC
	// If user is not within the application and is not of type manager on
	// the current application or is not an admin, throw an error

	deleted_application, _ := applications_service.UpdateApplication(application_id, models.Application{Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now()}}})

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Application successfully deleted.", "data": gin.H{"application": deleted_application}})
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

	// TODO: ADD BELOW LOGIC
	// Check to see if user either owns or is a member of the application
	// or if the user is an admin
	// If not, throw error

	// TODO: Append owner type plus id to the front of the cache key
	// to ensure security and no data overlap
	cache_key := application_input_param + "_" + "application" + ":" + application.UniqueId

	alerts := cache.RedisClient.Get(cache_key)

	if alerts.String() == redis.Nil.Error() {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Alerts not found.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Alerts retrieved.", "data": gin.H{"alerts": alerts}})
}
