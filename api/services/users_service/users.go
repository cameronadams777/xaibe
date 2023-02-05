package users_service

import (
	"api/initializers/database"
	"api/models"
	"time"

  "github.com/google/uuid"
)

// TODO: Find way to specify preloads when querying, to prevent excess queries

func GetAllUsers() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}

func GetUserById(user_id uuid.UUID) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("Teams").Preload("Applications").First(&user, user_id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
  err := database.DB.Preload("Teams").Preload("Applications").Where(models.User{ Email: email }).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByPasswordCode(reset_password_code string, reset_password_expiry time.Time) (*models.User, error) {
	var user models.User
	err := database.DB.Where("reset_password_code = ? AND reset_password_expiry >= ?", reset_password_code, reset_password_expiry).First(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user_id uuid.UUID, updates models.User) (*models.User, error) {
	// Get user that we want to update
	var user_to_update models.User
	err := database.DB.Preload("Teams").Preload("Applications").First(&user_to_update, user_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&user_to_update).Updates(updates)

	return &user_to_update, nil
}

func UpdateUserNullish(user_id uuid.UUID, updates map[string]interface{}) (*models.User, error) {
	// Get user that we want to update
	var user_to_update models.User
	err := database.DB.Preload("Teams").Preload("Applications").First(&user_to_update, user_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&user_to_update).Updates(updates)

	return &user_to_update, nil
}
