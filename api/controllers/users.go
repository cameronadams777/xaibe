package controllers

import (
	"api/models"
	"api/services/users_service"
	"api/structs"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// TODO: Write helper function that returns a JSON response but filters out sensitive data

type UpdateUserInput struct {
	UserId  int         `json:"user_id" binding:"required"`
	Updates models.User `json:"user" binding:"-"`
}

func GetAllUsers(c *gin.Context) {
	users := users_service.GetAllUsers()
	c.JSON(http.StatusOK, gin.H{"status": "error", "message": "", "data": gin.H{"users": users}})
}

func GetUserById(c *gin.Context) {
	user_input_param := c.Param("user_id")
	user_id, conv_err := strconv.Atoi(user_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting user by id.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	_, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred.", "data": nil})
		return
	}

	user, err := users_service.GetUserById(user_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User found.", "data": gin.H{"user": user}})
}

func UpdateUser(c *gin.Context) {
	var input UpdateUserInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	data, _ := c.Get("authScope")
	authScope := data.(structs.AuthScope)

	current_user, current_user_err := users_service.GetUserById(authScope.UserID)

	if current_user_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred.", "data": nil})
		return
	}

	if int(current_user.ID) != input.UserId && !current_user.IsAdmin {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"status": "error", "message": "You do not have permission to update this user's account", "data": nil})
		return
	}

	updated_user, err := users_service.UpdateUser(input.UserId, input.Updates)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User successfully updated.", "data": gin.H{"user": updated_user}})
}

func DeleteUser(c *gin.Context) {
	// Find user by id
	// Check to see if the user is not already deleted
	// Update user's deleted_at field
	// Return user with updated deleted_at field
	user_input_param := c.Param("user_id")
	user_id, conv_err := strconv.Atoi(user_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting user by id.", "data": nil})
		return
	}

	user_to_delete, err := users_service.GetUserById(user_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "User not found.", "data": nil})
		return
	}

	if user_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "User has already been deleted.", "data": err})
		return
	}

	deleted_user, _ := users_service.UpdateUser(user_id, models.User{Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now()}}})

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User successfully deleted.", "data": gin.H{"user": deleted_user}})
}