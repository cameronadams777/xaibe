package controllers

import (
	"api/models"
	"api/services/organizations_service"
	"api/services/teams_service"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CreateNewOrganizationInput struct {
	Name string `json:"organization_name" binding:"required"`
}

func GetAllOrganizations(c *gin.Context) {
	organizations := organizations_service.GetAllOrganizations()
	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "", "data": organizations})
}

func GetOrganizationById(c *gin.Context) {
	// Get ID from params
	organization_input_param := c.Param("organization_id")
	organization_id, conv_err := strconv.Atoi(organization_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting organization by id.", "data": nil})
		return
	}

	organization, err := organizations_service.GetOrganizationById(organization_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Organization not found.", "data": nil})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Organization found.", "data": organization})
}

func CreateNewOrganization(c *gin.Context) {
	// Generate a new organization with the name
	// of the organization gathered from the request body
	var input CreateNewOrganizationInput

	if err := c.BindJSON(&input); err != nil {
		fmt.Println(err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Invalid request body.", "data": nil})
		return
	}

	creation_err := organizations_service.CreateOrganization(input.Name)

	if creation_err != nil {
		fmt.Println(creation_err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "An unknown error occurred while creating the requested organization.", "data": nil})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"status": "success", "message": "Organization created.", "data": nil})
}

func DeleteOrganization(c *gin.Context) {
	// Update deleted at property on team
	organization_input_param := c.Param("organization_id")
	organization_id, conv_err := strconv.Atoi(organization_input_param)

	if conv_err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"status": "error", "message": "Error requesting organization by id.", "data": nil})
		return
	}

	organization_to_delete, err := organizations_service.GetOrganizationById(organization_id)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"status": "error", "message": "Organization not found.", "data": nil})
		return
	}

	if organization_to_delete.DeletedAt.Valid {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"status": "error", "message": "Organization has already been deleted.", "data": err})
		return
	}

	deleted_organization, _ := teams_service.UpdateTeam(organization_id, models.Team{Model: gorm.Model{DeletedAt: gorm.DeletedAt{Time: time.Now()}}})

	c.JSON(http.StatusOK, gin.H{"status": "success", "message": "Organization successfully deleted.", "data": gin.H{"organization": deleted_organization}})
}
