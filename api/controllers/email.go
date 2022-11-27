package controllers

import (
	"api/config"
	"api/services/users_service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanzoai/gochimp3"
)

type SubscribeNewUserInput struct {
	Email string `json:"email" binding:"required"`
}

type SendTransactionalEmailInput struct {
	Email      string `json:"email" binding:"required"`
	TemplateID string `json:"template_id" binding:"required"`
}

func GetTransactionalEmail(key string) string {
	var transaction_email_template_ids = make(map[string]string)
	transaction_email_template_ids["reset_password"] = "reset_password"

	template := transaction_email_template_ids[key]

	return template
}

func SubscribeNewUser(c *gin.Context) {
	var input SubscribeNewUserInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request.", "data": err})
		return
	}

	client := gochimp3.New(config.Get("MAILCHIMP_API_KEY"))

	listID := config.Get("MAILCHIMP_AUDIENCE_ID")

	// Fetch list
	list, err := client.GetList(listID, nil)
	if err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Could not find mailing list.", "data": err})
		return
	}

	// Add subscriber
	req := &gochimp3.MemberRequest{
		EmailAddress: input.Email,
		Status:       "subscribed",
	}

	if _, err := list.CreateMember(req); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Failed to subscribe user.", "data": err})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "User subscribed.", "data": nil})
}

func SendTransactionalEmail(c *gin.Context) {
	var input SendTransactionalEmailInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on login request.", "data": err})
		return
	}

	_, err := users_service.GetUserByEmail(input.Email)

	// Do not proceed if we cannot find a user with the email provided but
	// do not respond with error so as to prevent malicious email querying
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"status": "success", "message": "", "data": nil})
		return
	}

	// template_id := GetTransactionalEmail(input.TemplateID)

	// Send email

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Email sent.", "data": nil})
}
