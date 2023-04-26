package alert_schemas_service

import (
	"api/initializers/cache"
	"api/initializers/database"
	"api/models"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

func GetAlertSchemaByApplicationId(application_id uuid.UUID) (*models.AlertSchema, error) {
	var alert_schema models.AlertSchema
	err := database.DB.Where("application_id = ?", application_id).First(&alert_schema).Error

	if err != nil {
		return nil, err
	}

  schema_as_string, marshal_err := json.Marshal(&alert_schema)

  if marshal_err == nil {
    cache.RedisClient.Set("schemas:" + alert_schema.ID.String(), schema_as_string, time.Hour)
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
