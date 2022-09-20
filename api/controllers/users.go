package controllers

import "github.com/gin-gonic/gin"

func GetAllUsers(c *gin.Context) {
	// Return all non-deleted users
}

func GetUserById(c *gin.Context) {
	// TODO: Determine what is returned from here
}

func UpdateUser(c *gin.Context) {
	// Check to see if current user is admin or is the user being updated
	// If not, throw error
	// Update user
}

func DeleteUser(c *gin.Context) {
	// Update user deleted_at record to now
}
