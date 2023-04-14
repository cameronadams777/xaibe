package controllers

import (
	"api/initializers/cache"
	"api/services/applications_service"
  "api/structs"
	"api/websockets"

	"encoding/json"
  "fmt"
	"io"
	"log"
	"net/http"
	"time"

  "github.com/google/uuid"
	"github.com/gin-gonic/gin"
)

// TODO: Look into cleaning up this function
func WebHook(c *gin.Context) {
	application_input_param := c.Query("application_id")
	application_id, conv_err := uuid.Parse(application_input_param)

	if conv_err != nil {
    fmt.Println(conv_err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Application not specified."})
		return
	}

	body_as_byte_array, err := io.ReadAll(c.Request.Body)
	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "An error occurred posting alert to application."})
		return
	}

	body_as_string := string(body_as_byte_array)
  
  fmt.Println("Test", body_as_string)

	var body_as_json map[string]interface{}
	json.Unmarshal([]byte(body_as_string), &body_as_json)

	application, retrieve_app_error := applications_service.GetApplicationById(application_id)

	if retrieve_app_error != nil {
    fmt.Println(retrieve_app_error)
		c.AbortWithStatusJSON(http.StatusNotFound, structs.ErrorMessage{Message: "Could not find application to post alert to."})
		return
	}

	body_as_json["application_id"] = application_id
	body_as_json["alert_schema"] = application.AlertSchema

	alert_as_byte_array, _ := json.Marshal(body_as_json)

	// Push alert data via websocket to any active clients
	alert := websockets.Message{
		Room: application_input_param + ":" + application.UniqueId,
		Data: alert_as_byte_array,
	}

	websockets.Pool.Broadcast <- alert

	// Store alert data for later in redis
	var owner_id string
	if application.TeamID != nil {
		team_id := application.TeamID.String()
		owner_id = "team_" + team_id
	} else if application.UserID != nil {
		user_id := application.UserID.String()
		owner_id = "user_" + user_id
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred posting alert to application."})
		return
	}

	cache_key := owner_id + ":" + "application_" + application_input_param + ":" + application.UniqueId + ":" + time.Now().Local().String()
	// Take body data and push to redis under cache key
	one_month_expiration := time.Hour * 24 * 30
	redis_err := cache.RedisClient.Set(cache_key, string(body_as_byte_array), time.Duration(one_month_expiration)).Err()
	if redis_err != nil {
		log.Println(redis_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred persisting alert."})
		return
	}

	c.JSON(http.StatusOK, nil)
}
