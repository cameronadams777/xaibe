package controllers

import (
	"api/config"
	"api/models"
	"api/services/sparkpost_service"
	"api/services/users_service"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type SendResetPasswordEmailInput struct {
	Email string `json:"email" binding:"required"`
}

type ValidateResetPasswordCodeInput struct {
	Code string `json:"code" binding:"required"`
}

type ResetUserPasswordInput struct {
	Code                 string `json:"code" binding:"required"`
	Password             string `json:"password" binding:"required"`
	PasswordConfirmation string `json:"password_confirmation" binding:"required"`
}

type ResetPasswordTemplateElements struct {
  Email string  
  Link  string
}

func SendResetPasswordEmail(c *gin.Context) {
	// Retreive email from request body
	var input SendResetPasswordEmailInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": err})
		return
	}
	// See if user exists based on email
	user, err := users_service.GetUserByEmail(input.Email)

	if err != nil {
    log.Panicln(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred sending password reset email.", "data": err})
		return
	}

	// Update user with random uuid token, and timestamp that is 15 minutes from now
	updates := models.User{
		ResetPasswordCode:   uuid.NewString(),
		ResetPasswordExpiry: time.Now().Add(time.Minute * 15),
	}
	_, update_err := users_service.UpdateUser(int(user.ID), updates)

	if update_err != nil {
    log.Panicln(update_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred sending password reset email.", "data": err})
		return
	}

  templateElements := ResetPasswordTemplateElements{
    Email: user.Email,
    Link: config.Get("APP_HOST_NAME") + "/reset-password?hash=" + user.ResetPasswordCode,
  }
  
	// Send email with link to users email
	send_err := sparkpost_service.SendEmail(user.Email, "reset_password", templateElements)

	if send_err != nil {
    log.Panicln(send_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred sending password reset email.", "data": err})
		return
	}

  c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email sent.", "data": nil})
}

func ValidateResetPasswordCode(c *gin.Context) {
	// Receive code from body
	var input ValidateResetPasswordCodeInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": err})
		return
	}

	// Find user by reset password token and reset password expiry
	_, err := users_service.GetUserByPasswordCode(input.Code, time.Now())

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Password reset token is invalid or has expired.", "data": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Validated.", "data": true})
}

func ResetUserPassword(c *gin.Context) {
	// Receive password and code from body
	var input ResetUserPasswordInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": err})
		return
	}

	if input.Password != input.PasswordConfirmation {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Password and Confirmation do not match.", "data": nil})
		return
	}

	// Fetch user by reset password token and reset password expiry
	user, err := users_service.GetUserByPasswordCode(input.Code, time.Now())

	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "error", "message": "Password reset token is invalid or has expired.", "data": false})
		return
	}

	// Salt the updated password
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	// Update user password as well as clear the reset psasword token and password expiry
	_, update_err := users_service.UpdateUser(int(user.ID), models.User{Password: string(password)})
	// Send password update success email

	if update_err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An error occurred during password reset process.", "data": false})
		return
	}

	// Return success
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Password updated.", "data": true})
}
