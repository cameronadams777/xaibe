package applications_service

import (
	"api/initializers/database"
	"api/models"
)

func GetAllApplications(filters models.Application) []models.Application {
	var applications []models.Application
	database.DB.Find(&applications, &filters)
	return applications
}

func GetApplicationById(application_id int) (*models.Application, error) {
	var application models.Application
	err := database.DB.Find(&application, application_id).Error
	if err != nil {
		return nil, err
	}
	return &application, err
}

func CreateApplication(new_application models.Application) error {
	err := database.DB.Create(&new_application).Error
	if err != nil {
		return err
	}
	return nil
}

func UpdateApplication(application_id int, updates models.Application) (*models.Application, error) {
	var application_to_update models.Application
	err := database.DB.First(&application_to_update, application_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&application_to_update).Updates(updates)

	return &application_to_update, nil
}
