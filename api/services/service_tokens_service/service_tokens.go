package service_tokens

import (
	"api/database"
	"api/models"
	"time"

	"github.com/google/uuid"
)

func GetAllServiceTokens() []models.ServiceToken {
	var tokens []models.ServiceToken
	database.DB.Find(&tokens)
	return tokens
}

func GetServiceTokenById(token_id int) (*models.ServiceToken, error) {
	var token models.ServiceToken
	err := database.DB.First(&token, token_id).Error
	if err != nil {
		return nil, err
	}
	return &token, nil
}

func CreateServiceToken(team_id int, application_id int, expires_at time.Time) (*models.ServiceToken, error) {
	new_token := models.ServiceToken{
		Token:         uuid.New().String(),
		TeamID:        uint(team_id),
		ApplicationID: uint(application_id),
		ExpiresAt:     time.Now().AddDate(1, 0, 0),
	}
	err := database.DB.Create(&new_token).Error
	if err != nil {
		return nil, err
	}
	return &new_token, nil
}

func UpdateServiceToken(token_id int, updates models.ServiceToken) (*models.ServiceToken, error) {
	var token_to_update models.ServiceToken
	err := database.DB.First(&token_to_update, token_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&token_to_update).Updates(updates)

	return &token_to_update, nil
}
