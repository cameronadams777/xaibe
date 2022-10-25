package controllers

import (
	"api/models"
	"api/services/teams_service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateNewTeamInput struct {
	Name string `json:"team_name" binding:"required"`
}

func GetAllTeams(c *gin.Context) {
	// Fetch paginated teams list
	teams := teams_service.GetAllTeams()
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Teams found.", "data": teams})
}

func GetTeamById(c *gin.Context) {
	// Get ID from params
	team_input_param := c.Param("team_id")
	team_id, conv_err := strconv.Atoi(team_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting team by id.", "data": nil})
		return
	}

	// Fetch team by id
	team, err := teams_service.GetTeamById(team_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Team not found.", "data": nil})
		return
	}

	// TODO: ADD BELOW LOGIC
	// If current user is not member of team and is not admin,
	// throw an error

	// Return team info
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Team found.", "data": team})
}

func CreateNewTeam(c *gin.Context) {
	// Generate a new team with the name
	// of the team gathered from the request body
	var input CreateNewTeamInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	created_team, creation_err := teams_service.CreateTeam(input.Name)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred while creating the requested team.", "data": nil})
		return
	}

	fmt.Println(created_team)

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Team created.", "data": created_team})
}

func DeleteTeam(c *gin.Context) {
	team_input_param := c.Param("team_id")
	team_id, conv_err := strconv.Atoi(team_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting team by id.", "data": nil})
		return
	}

	team_to_delete, err := teams_service.GetTeamById(team_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Team not found.", "data": nil})
		return
	}

	if team_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Team has already been deleted.", "data": err})
		return
	}

	// TODO: ADD BELOW LOGIC
	// If user is not within the team, is not of type manager on
	// the current team or is not an admin, throw an error

	deleted_team, _ := teams_service.UpdateTeam(team_id, models.Team{Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now()}}})

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Team successfully deleted.", "data": gin.H{"team": deleted_team}})
}

// func GetAllTeamApplications(c *gin.Context) {
// 	team_input_param := c.Param("team_id")
// 	team_id, conv_err := strconv.Atoi(team_input_param)

// 	if conv_err != nil {
// 		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting team by id.", "data": nil})
// 		return
// 	}

// 	// TODO: ADD BELOW LOGIC
// 	// If user is not within the team, is not of type manager on
// 	// the current team or is not an admin, throw an error

// 	applications := applications_service.GetAllApplications(models.Application{TeamID: uint(team_id)})

// 	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Successfully retrieved applications for team.", "data": gin.H{"applications": applications}})
// }
