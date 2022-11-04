package controllers

import (
	"api/initializers/cache"
	"api/services/applications_service"
	"api/websockets"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// TODO: Look into cleaning up this function
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

	body_as_string := string(body_as_byte_array)

	var body_as_json map[string]interface{}
	json.Unmarshal([]byte(body_as_string), &body_as_json)

	application, retrieve_app_error := applications_service.GetApplicationById(application_id)

	if retrieve_app_error != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Could not find application to post alert to.", "data": nil})
		return
	}

	// Push alert data via websocket to any active clients
	alert := websockets.Message{
		Room: application_input_param + ":" + application.UniqueId,
		Data: body_as_byte_array,
	}

	websockets.Pool.Broadcast <- alert

	// Store alert data for later in redis
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

	cache_key := owner_id + ":" + "application_" + application_input_param + ":" + application.UniqueId + ":" + time.Now().Local().String()
	// Take body data and push to redis under cache key
	one_month_expiration := time.Hour * 24 * 30
	redis_err := cache.RedisClient.Set(cache_key, string(body_as_byte_array), time.Duration(one_month_expiration)).Err()
	if redis_err != nil {
		log.Println(redis_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred persisting alert.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Alert posted.", "data": nil})
}
