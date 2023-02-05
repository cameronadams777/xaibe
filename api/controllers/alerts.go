package controllers

import (
	"api/assertions"
	"api/initializers/cache"
	"api/models"
	"api/services/applications_service"
	"api/services/teams_service"
	"api/services/users_service"
	"api/structs"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/google/uuid"
)

type QueryObject struct {
	owner_type  string
	owner_id    string
	application models.Application
}

type ApplicationAlertResponse struct {
	AlertSchema models.AlertSchema
	Alerts      []map[string]interface{}
}

func GetAllAlerts(c *gin.Context) {
	// TODO: Refactor for improved performance due to query size
	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	applications := []QueryObject{}

	user, _ := users_service.GetUserById(authScope.UserID)

	for _, a := range user.Applications {
		application, err := applications_service.GetApplicationById(a.ID)

		if err != nil {
			continue
		}

		applications = append(applications, QueryObject{
			owner_type:  "user",
			owner_id:    user.ID.String(),
			application: *application,
		})
	}

	for _, t := range user.Teams {
		team, fetch_team_err := teams_service.GetTeamById(t.ID)
		if fetch_team_err != nil {
			continue
		}
		for _, a := range team.Applications {
			application, err := applications_service.GetApplicationById(a.ID)

			if err != nil {
				continue
			}

			applications = append(applications, QueryObject{
				owner_type:  "team",
				owner_id:    user.ID.String(),
				application: *application,
			})
		}
	}

	alerts_as_json := make(map[string]ApplicationAlertResponse)

	for _, queryObject := range applications {
		application_id := queryObject.application.ID.String()
		scan_key := queryObject.owner_type + "_" + queryObject.owner_id + ":" + "application_" + application_id + ":" + queryObject.application.UniqueId + ":*"

		var alerts []map[string]interface{}

		iter := cache.RedisClient.Scan(0, scan_key, 0).Iterator()
		for iter.Next() {
			var alert map[string]interface{}
			alert_trx := cache.RedisClient.Get(iter.Val())
			alert_string := alert_trx.Val()
			if alert_string != redis.Nil.Error() {
				json.Unmarshal([]byte(alert_trx.Val()), &alert)
			}
			alerts = append(alerts, alert)
		}
		if iter_err := iter.Err(); iter_err != nil {
			panic(iter_err)
		}
		alerts_as_json[application_id] = ApplicationAlertResponse{
			AlertSchema: queryObject.application.AlertSchema,
			Alerts:      alerts,
		}
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Alerts retrieved.", "data": alerts_as_json})
}

func GetApplicationAlerts(c *gin.Context) {
	application_input_param := c.Param("application_id")
	application_id, conv_err := uuid.Parse(application_input_param)

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
		team_id := application.TeamID.String()
		owner_id = "team_" + team_id
	} else if application.UserID != nil {
		user_id := application.UserID.String()
		owner_id = "user_" + user_id
	} else {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred retrieving alerts for application.", "data": nil})
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
