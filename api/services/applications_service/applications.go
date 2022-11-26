package applications_service

import (
	"api/initializers/database"
	"api/models"
)

func GetAllApplications(filters models.Application) []models.Application {
	var applications []models.Application
	database.DB.Preload("AlertSchema").Find(&applications, &filters)
	return applications
}

func GetApplicationById(application_id int) (*models.Application, error) {
	var application models.Application
	err := database.DB.Preload("AlertSchema").Find(&application, application_id).Error
	if err != nil {
		return nil, err
	}
	return &application, err
}

func CreateApplication(new_application models.Application) (*models.Application, error) {
	err := database.DB.Create(&new_application).Error
	if err != nil {
		return nil, err
	}
	return &new_application, nil
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

func AddSchemaToApplication(application_id int, alert_schema models.AlertSchema) (*models.Application, error) {
	var application models.Application
	application_err := database.DB.First(&application, application_id).Error

	if application_err != nil {
		return nil, application_err
	}

	database.DB.Model(&application).Association("AlertSchema").Append(alert_schema)

	return &application, nil
}

func DeleteApplication(application_id int) (*models.Application, error) {
	var application_to_delete models.Application
	err := database.DB.First(&application_to_delete, application_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Delete(&application_to_delete)

	return &application_to_delete, nil
}
