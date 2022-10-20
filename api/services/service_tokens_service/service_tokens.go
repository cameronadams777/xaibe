package service_tokens_service

import (
	"api/initializers/database"
	"api/models"
	"time"

	"github.com/google/uuid"
)

func GetAllServiceTokens() []models.ServiceToken {
	var tokens []models.ServiceToken
	database.DB.Preload("Application").Find(&tokens)
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

// TODO: Refactor this so that we can just use `GetAllServiceTokens` rather than
// having a completely separate function
func GetAllServiceTokensByApplicationId(application_id int) []models.ServiceToken {
	var tokens []models.ServiceToken
	database.DB.Find(&tokens, models.ServiceToken{ApplicationID: uint(application_id)})
	return tokens
}

func CreateServiceToken(application_id int) (*models.ServiceToken, error) {
	new_token := models.ServiceToken{
		Token:         uuid.NewString(),
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
