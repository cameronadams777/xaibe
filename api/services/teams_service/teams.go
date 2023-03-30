package teams_service

import (
	"api/initializers/database"
	"api/models"
	"api/structs/invite_status"
	"github.com/google/uuid"
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

func GetTeamById(team_id uuid.UUID) (*models.Team, error) {
	var team models.Team
	err := database.DB.Preload("Users").Preload("Managers").Preload("Applications").First(&team, team_id).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func CreateTeam(name string, numberOfSeats uint,  creating_user models.User) (*models.Team, error) {
	team := models.Team{
		Name:     name,
		Users:    []*models.User{&creating_user},
		Managers: []*models.User{&creating_user},
	}
	err := database.DB.Create(team).Error
	if err != nil {
		return nil, err
	}
	return &team, nil
}

func UpdateTeam(team_id uuid.UUID, updates models.Team) (*models.Team, error) {
	var team_to_update models.Team
	err := database.DB.First(&team_to_update, team_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&team_to_update).Updates(updates)

	return &team_to_update, nil
}

func DeleteTeam(team_id uuid.UUID) (*models.Team, error) {
	var team_to_delete models.Team
	err := database.DB.First(&team_to_delete, team_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Delete(&team_to_delete)

	return &team_to_delete, nil
}

func PermDeleteTeam(team_id uuid.UUID) error {
	var team_to_delete models.Team
	err := database.DB.First(&team_to_delete, team_id).Error

	if err != nil {
		return err
	}

  del_err := database.DB.Unscoped().Delete(team_to_delete).Error

  if del_err != nil {
    return del_err
  }

  return nil
}

func AddUserToTeam(team_id uuid.UUID, user_id uuid.UUID) (*models.Team, error) {
	var team models.Team
	team_err := database.DB.First(&team, team_id).Error

	if team_err != nil {
		return nil, team_err
	}

	var user models.User
	user_err := database.DB.First(&user, user_id).Error

	if user_err != nil {
		return nil, user_err
	}

	database.DB.Model(&team).Association("Users").Append(&user)

	return &team, nil
}

func RemoveUserFromTeam(team_id uuid.UUID, user_id uuid.UUID) (*models.Team, error) {
	var team models.Team
	team_err := database.DB.First(&team, team_id).Error

	if team_err != nil {
		return nil, team_err
	}

	var user models.User
	user_err := database.DB.First(&user, user_id).Error

	if user_err != nil {
		return nil, user_err
	}

	database.DB.Model(&team).Association("Users").Delete(&user)
	database.DB.Model(&team).Association("Managers").Delete(&user)

	return &team, nil
}

func GetPendingTeamInvites(email string) (*[]models.TeamInvite, error) {
	var invites []models.TeamInvite
	err := database.DB.Preload("Team").Where("email = ? AND status = ?", email, invite_status.PENDING).Find(&invites).Error
	if err != nil {
		return nil, err
	}
	return &invites, nil
}

func CreateInvite(team_id uuid.UUID, sender_id uuid.UUID, email string) (*models.TeamInvite, error) {
	invite := models.TeamInvite{
		TeamID:   team_id,
		SenderID: sender_id,
		Email:    email,
		Status:   invite_status.PENDING,
	}
	err := database.DB.Create(&invite).Error
	if err != nil {
		return nil, err
	}
	return &invite, nil
}

func UpdateInvite(invite_id uint, updates models.TeamInvite) (*models.TeamInvite, error) {
	var invite_to_update models.TeamInvite
	err := database.DB.Preload("Team").First(&invite_to_update, invite_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Model(&invite_to_update).Updates(updates)

	return &invite_to_update, nil
}
