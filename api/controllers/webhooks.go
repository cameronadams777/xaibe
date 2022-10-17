package controllers

import (
	"api/cache"
	"api/services/applications_service"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
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

	body_as_string := string(body_as_byte_array)

	var body_as_json map[string]interface{}
	json.Unmarshal([]byte(body_as_string), &body_as_json)

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

	alerts := cache.RedisClient.Get(cache_key)

	var alerts_to_post []map[string]interface{}

	if alerts.Val() == redis.Nil.Error() {
		alerts_to_post = append(alerts_to_post, body_as_json)
	} else {
		var alerts_as_json []map[string]interface{}
		json.Unmarshal([]byte(alerts.Val()), &alerts_as_json)

		alerts_to_post = append(alerts_as_json, body_as_json)
	}

	// TODO: Publish alert to socket connection collection in pool identified by cache_key

	// Take body data and push to redis under cache key
	one_month_expiration := time.Hour * 24 * 30
	alerts_to_post_bytes, marshal_err := json.Marshal(&alerts_to_post)

	if marshal_err != nil {
		log.Println(marshal_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred persisting alert.", "data": nil})
		return
	}

	redis_err := cache.RedisClient.Set(cache_key, string(alerts_to_post_bytes), time.Duration(one_month_expiration)).Err()

	if redis_err != nil {
		log.Println(redis_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred persisting alert.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Alert posted.", "data": nil})
}
