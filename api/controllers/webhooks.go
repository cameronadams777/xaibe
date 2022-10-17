package controllers

import (
	"api/cache"
	"api/services/applications_service"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func WebHook(c *gin.Context) {
	application_input_param := c.Query("application_id")
	application_id, conv_err := strconv.Atoi(application_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Application not specified.", "data": nil})
		return
	}

	body_as_byte_array, err := io.ReadAll(c.Request.Body)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "An error occurred posting alert to application.", "data": nil})
		return
	}

	body := string(body_as_byte_array)

	application, retrieve_app_error := applications_service.GetApplicationById(application_id)

	if retrieve_app_error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Could not find application to post alert to.", "data": nil})
		return
	}

	var owner_id string

	if application.TeamID != 0 {
		team_id := strconv.Itoa(int(application.TeamID))
		owner_id = "team_" + team_id
	} else if application.UserID != 0 {
		user_id := strconv.Itoa(int(application.UserID))
		owner_id = "user_" + user_id
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred posting alert to application.", "data": nil})
		return
	}

	cache_key := owner_id + ":" + "application_" + application_input_param + ":" + application.UniqueId

	// TODO: Publish alert to socket connection collection in pool identified by cache_key

	// Take body data and push to redis under cache key
	one_month_expiration := time.Hour * 24 * 30
	redis_err := cache.RedisClient.HSet(cache_key, body, time.Duration(one_month_expiration)).Err()

	if redis_err != nil {
		log.Println(redis_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred persisting alert.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Alert posted.", "data": nil})
}
