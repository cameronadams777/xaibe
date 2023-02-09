package controllers

import (
	"api/config"
	"api/services/sparkpost_service"
	"api/services/users_service"
  "api/structs"
  "fmt"
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
	PasswordConfirmation string `json:"passwordConfirmation" binding:"required"`
}

type ResetPasswordTemplateElements struct {
	Email string
	Link  string
}

func SendResetPasswordEmail(c *gin.Context) {
	// Retreive email from request body
	var input SendResetPasswordEmailInput

	if err := c.BindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}
	// See if user exists based on email
	user, err := users_service.GetUserByEmail(input.Email)

	if err != nil {
		log.Panicln(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred sending password reset email."})
		return
	}

	// Update user with random uuid token, and timestamp that is 15 minutes from now
	updates := map[string]interface{}{
		"ResetPasswordCode":   uuid.NewString(),
		"ResetPasswordExpiry": time.Now().Add(time.Minute * 15),
	}

	updated_user, update_err := users_service.UpdateUserNullish(user.ID, updates)

	if update_err != nil {
		log.Panicln(update_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred sending password reset email."})
		return
	}

	template_elements := ResetPasswordTemplateElements{
		Email: user.Email,
		Link:  config.Get("APP_HOST_NAME") + "/reset-password/" + updated_user.ResetPasswordCode,
	}

	// Send email with link to users email
	send_err := sparkpost_service.SendEmail(user.Email, "reset_password", template_elements)

	if send_err != nil {
		log.Panicln(send_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred sending password reset email."})
		return
	}

	c.JSON(http.StatusOK, nil)
}

func ValidateResetPasswordCode(c *gin.Context) {
	// Receive code from body
	var input ValidateResetPasswordCodeInput

	if err := c.BindJSON(&input); err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	// Find user by reset password token and reset password expiry
	_, err := users_service.GetUserByPasswordCode(input.Code, time.Now())

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, structs.ErrorMessage{Message: "Password reset token is invalid or has expired."})
		return
	}

	c.JSON(http.StatusOK, true)
}

func ResetUserPassword(c *gin.Context) {
	// Receive password and code from body
	var input ResetUserPasswordInput

	if err := c.BindJSON(&input); err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Invalid request body."})
		return
	}

	if input.Password != input.PasswordConfirmation {
		c.AbortWithStatusJSON(http.StatusBadRequest, structs.ErrorMessage{Message: "Password and Confirmation do not match."})
		return
	}

	// Fetch user by reset password token and reset password expiry
	user, err := users_service.GetUserByPasswordCode(input.Code, time.Now())

	if err != nil {
    fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusUnauthorized, structs.ErrorMessage{Message: "Password reset token is invalid or has expired."})
		return
	}

	// Salt the updated password
	password, _ := bcrypt.GenerateFromPassword([]byte(input.Password), 14)

	// Update user password as well as clear the reset psasword token and password expiry
	_, update_err := users_service.UpdateUserNullish(user.ID, map[string]interface{}{"Password": string(password), "ResetPasswordCode": nil, "ResetPasswordExpiry": nil})

	if update_err != nil {
    fmt.Println(update_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, structs.ErrorMessage{Message: "An error occurred during password reset process."})
		return
	}

	// Return success
	c.JSON(http.StatusOK, true)
}
