package organizations_service

import (
	"api/database"
	"api/models"
)

func GetAllOrganizations() []models.Organization {
	var organizations []models.Organization
	database.DB.Preload("Teams").Find(&organizations)
	return organizations
}

func GetOrganizationById(organization_id int) (*models.Organization, error) {
	var organization models.Organization
	err := database.DB.Find(&organization, organization_id).Error
	if err != nil {
		return nil, err
	}
	return &organization, nil
}

func CreateOrganization(name string) error {
	organization := models.Organization{
		Name: name,
	}
	err := database.DB.Create(&organization).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateOrganization(organization_id int, updates models.Organization) (*models.Organization, error) {
	var organization_to_update models.Organization
	err := database.DB.First(&organization_to_update, organization_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&organization_to_update).Updates(updates)

	return &organization_to_update, nil
}
