package alert_schemas_service

import (
	"api/initializers/database"
	"api/models"

  "github.com/google/uuid"
)

func GetAlertSchemaByApplicationId(application_id uuid.UUID) (*models.AlertSchema, error) {
	var alert_schema models.AlertSchema
	err := database.DB.Where("application_id = ?", application_id).First(&alert_schema).Error
	if err != nil {
		return nil, err
	}
	return &alert_schema, nil
}

func CreateNewAlertSchema(new_alert_schema models.AlertSchema) (*models.AlertSchema, error) {
	err := database.DB.Create(&new_alert_schema).Error
	if err != nil {
		return nil, err
	}
	return &new_alert_schema, nil
}
