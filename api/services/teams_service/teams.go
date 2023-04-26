package teams_service

import (
	"api/initializers/cache"
	"api/initializers/database"
	"api/models"
	"api/structs/invite_status"
	"encoding/json"
	"time"

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

  team_as_string, marshal_err := json.Marshal(&team)

  if marshal_err == nil {
    cache.RedisClient.Set("teams:" + team.ID.String(), team_as_string, time.Hour)
  }

	return &team, nil
}

func GetTeamById(team_id uuid.UUID) (*models.Team, error) {
	var team models.Team

  team_tx := cache.RedisClient.Get("teams:" + team_id.String()) 
  conv_err := json.Unmarshal([]byte(team_tx.Val()), &team) 

  if conv_err == nil {
    return &team, nil 
  }

	err := database.DB.Preload("Users").Preload("Managers").Preload("Applications").First(&team, team_id).Error

	if err != nil {
		return nil, err
	}

  team_as_string, marshal_err := json.Marshal(&team)

  if marshal_err == nil {
    cache.RedisClient.Set("teams:" + team.ID.String(), team_as_string, time.Hour)
  }

	return &team, nil
}

func CreateTeam(name string, numberOfSeats uint,  creating_user models.User) (*models.Team, error) {
	team := models.Team{
		Name:     name,
    ActiveNumberOfSeats: numberOfSeats,
		Users:    []models.User{creating_user},
		Managers: []models.User{creating_user},
	}
  err := database.DB.Omit("Users.*", "Managers.*").Create(&team).Error
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

  team_as_string, marshal_err := json.Marshal(&team_to_update)

  if marshal_err == nil {
    cache.RedisClient.Set("teams:" + team_id.String(), team_as_string, time.Hour)
  }

	return &team_to_update, nil
}

func DeleteTeam(team_id uuid.UUID) (*models.Team, error) {
	var team_to_delete models.Team
	err := database.DB.First(&team_to_delete, team_id).Error

	if err != nil {
		return nil, err
	}

	database.DB.Delete(&team_to_delete)

  team_as_string, marshal_err := json.Marshal(&team_to_delete)

  if marshal_err == nil {
    cache.RedisClient.Set("teams:" + team_id.String(), team_as_string, time.Hour)
  }

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

  cache.RedisClient.Del("teams:" + team_id.String())

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

  team_as_string, marshal_err := json.Marshal(&team)

  if marshal_err == nil {
    cache.RedisClient.Set("teams:" + team_id.String(), team_as_string, time.Hour)
  }

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

  team_as_string, marshal_err := json.Marshal(&team)

  if marshal_err == nil {
    cache.RedisClient.Set("teams:" + team_id.String(), team_as_string, time.Hour)
  }

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
