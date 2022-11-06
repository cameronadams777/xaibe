package teams_service

import (
	"api/initializers/database"
	"api/models"
)

// TODO: Find way to specify preloads when querying, to prevent excess queries

func GetAllTeams() []models.Team {
	var teams []models.Team
	database.DB.Preload("Users").Preload("Managers").Preload("Applications").Find(&teams)
	return teams
}

func GetTeamByName(name string) (*models.Team, error) {
	var team models.Team
	err := database.DB.First(&team, name).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func GetTeamById(team_id int) (*models.Team, error) {
	var team models.Team
	err := database.DB.Preload("Users").Preload("Managers").Preload("Applications").First(&team, team_id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func CreateTeam(name string, creating_user models.User) (*models.Team, error) {
	team := models.Team{
		Name:     name,
		Users:    []*models.User{&creating_user},
		Managers: []*models.User{&creating_user},
	}
	err := database.DB.Create(&team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func UpdateTeam(team_id int, updates models.Team) (*models.Team, error) {
	var team_to_update models.Team
	err := database.DB.First(&team_to_update, team_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&team_to_update).Updates(updates)

	return &team_to_update, nil
}

func DeleteTeam(team_id int) (*models.Team, error) {
	var team_to_delete models.Team
	err := database.DB.First(&team_to_delete, team_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Delete(&team_to_delete)

	return &team_to_delete, nil
}
