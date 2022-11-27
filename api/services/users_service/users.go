package users_service

import (
	"api/initializers/database"
	"api/models"
)

// TODO: Find way to specify preloads when querying, to prevent excess queries

func GetAllUsers() []models.User {
	var users []models.User
	database.DB.Find(&users)
	return users
}

func GetUserById(user_id int) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("Teams").Preload("Applications").First(&user, user_id).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := database.DB.Preload("Teams").Preload("Applications").First(&user, email).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func UpdateUser(user_id int, updates models.User) (*models.User, error) {
	// Get user that we want to update
	var user_to_update models.User
	err := database.DB.Preload("Teams").Preload("Applications").First(&user_to_update, user_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&user_to_update).Updates(updates)

	return &user_to_update, nil
}
