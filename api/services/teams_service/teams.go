package teams_service

import (
	"api/database"
	"api/models"
)

func GetAllTeams() []models.Team {
	var teams []models.Team
	database.DB.Find(&teams)
	return teams
}

func GetTeamById(team_id int) (*models.Team, error) {
	var team models.Team
	err := database.DB.First(&team, team_id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func CreateTeam(name string) error {
	team := models.Team{
		Name: name,
	}
	err := database.DB.Create(&team).Error
	if err != nil {
		return err
	}
	return nil
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
