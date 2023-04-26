package applications_service

import (
	"api/initializers/cache"
	"api/initializers/database"
	"api/models"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func GetAllApplications(filters models.Application) []models.Application {
	var applications []models.Application
	database.DB.Preload("AlertSchema").Find(&applications, &filters)
	return applications
}

func GetApplicationById(application_id uuid.UUID) (*models.Application, error) {
	var application models.Application

  application_tx := cache.RedisClient.Get("applications:" + application_id.String()) 
  conv_err := json.Unmarshal([]byte(application_tx.Val()), &application) 

  if conv_err == nil {
    return &application, nil
  }

	err := database.DB.Preload("AlertSchema").Find(&application, application_id).Error
	
  if err != nil {
		return nil, err
	}
  
  application_as_string, marshal_err := json.Marshal(&application)

  if marshal_err == nil {
    cache.RedisClient.Set("applications:" + application_id.String(), application_as_string, time.Hour)
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

func UpdateApplication(application_id uuid.UUID, updates models.Application) (*models.Application, error) {
	var application_to_update models.Application
	err := database.DB.Preload("AlertSchema").First(&application_to_update, application_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&application_to_update).Updates(updates)

  application_as_string, marshal_err := json.Marshal(&application_to_update)

  if marshal_err == nil {
    cache.RedisClient.Set("applications:" + application_id.String(), application_as_string, time.Hour)
  }

	return &application_to_update, nil
}

func AddSchemaToApplication(application_id uuid.UUID, alert_schema models.AlertSchema) (*models.Application, error) {
	var application models.Application
	application_err := database.DB.First(&application, application_id).Error

	if application_err != nil {
		return nil, application_err
	}

	database.DB.Model(&application).Association("AlertSchema").Append(alert_schema)

  application_as_string, marshal_err := json.Marshal(&application)

  if marshal_err == nil {
    cache.RedisClient.Set("applications:" + application_id.String(), application_as_string, time.Hour)
  }

	return &application, nil
}

func DeleteApplication(application_id uuid.UUID) (*models.Application, error) {
	var application_to_delete models.Application
	err := database.DB.First(&application_to_delete, application_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Delete(&application_to_delete)

  application_as_string, marshal_err := json.Marshal(&application_to_delete)

  if marshal_err == nil {
    cache.RedisClient.Set("applications:" + application_id.String(), application_as_string, time.Hour)
  }

	return &application_to_delete, nil
}
