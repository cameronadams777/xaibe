package controllers

import "github.com/gin-gonic/gin"

func GetAllTeams(c *gin.Context) {
	// Fetch paginated teams list
}

func GetTeamById(c *gin.Context) {
	// Get ID from params
	// Fetch team by id
	// If current user is not member of team and is not admin,
	// throw an error
	// Return team info
}

func CreateNewTeam(c *gin.Context) {
	// Generate a new team with the name
	// of the team gathered from the request body
}

func DeleteTeam(c *gin.Context) {
	// If user is not within the team, is not of type manager on
	// the current team or is not an admin, throw an error
	// Update deleted at property on team
}
