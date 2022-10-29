package alert_schemas_service

import (
	"api/initializers/database"
	"api/models"
)

func CreateNewAlertSchema(new_alert_schema models.AlertSchema) (*models.AlertSchema, error) {
	err := database.DB.Create(&new_alert_schema).Error
	if err != nil {
		return nil, err
	}
	return &new_alert_schema, nil
}
