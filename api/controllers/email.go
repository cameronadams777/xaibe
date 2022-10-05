package controllers

import (
	"api/config"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hanzoai/gochimp3"
)

type SubscribeNewUserInput struct {
	Email string `json:"email" binding:"required"`
}

func SubscribeNewUser(c *gin.Context) {
	var input SubscribeNewUserInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error on login request.", "data": err})
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
}
