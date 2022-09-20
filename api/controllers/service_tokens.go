package controllers

import "github.com/gin-gonic/gin"

func GetServiceTokensByTeam(c *gin.Context) {
	// Get team id from params
	// Get team by id
	// Check to see if user is in team
	// If not, throw error
	// If so, check to see if user is manager of team
	// If not, throw error
	// If so, retrieve all non-expired/non-deleted tokens for given team and return them
}

func GetServiceTokenByApplication(c *gin.Context) {
	// Get application id from params
	// Get application by id
	// Get team by application
	// Check to see if user is in team
	// If not, throw error
	// If so, check to see if user is manager of team
	// If not, throw error
	// If so, return non-expired/non-deleted service token
}

func CreateNewToken(c *gin.Context) {
	// Get team id and application id from request body
	// Get team by id
	// Ensure that team owns application
	// Ensure that user is manager of team
	// Ensure that application does not already have a non-expired token
	// If token exists and is not expired, throw error
	// If token exists but is expired, updated deleted_at time and continue
	// Create token
	// Return token
}

func DeleteToken(c *gin.Context) {
	// Get token id from params
	// Get application by token id
	// Get team by application id
	// Check to see if user is in team
	// If not, throw error
	// Check to see if user is manager of team
	// If not, throw error
	// Update deleted_at field in token record
}
