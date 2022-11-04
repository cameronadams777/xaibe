package assertions

import (
	"api/services/applications_service"
	"api/services/teams_service"
	"errors"
)

func UserOwnsApplication(application_id uint, user_id uint) error {
	application, app_find_err := applications_service.GetApplicationById(int(application_id))
	if app_find_err != nil {
		return errors.New("UserOwnsApplication: Application does not exist")
	}

	if application.UserID == nil {
		return errors.New("UserOwnsApplication: Application does not belong to a specific user")
	}

	if *application.UserID == user_id {
		return nil
	}

	return errors.New("UserOwnsApplication: User does not own this application")
}

func UserIsMemberOfTeamApplication(application_id uint, user_id uint) error {
	application, app_find_err := applications_service.GetApplicationById(int(application_id))
	if app_find_err != nil {
		return errors.New("UserIsManagerOfTeamApplication: Application does not exist")
	}

	if application.TeamID == nil {
		return errors.New("UserIsManagerOfTeamApplication: Application does not belong to a team")
	}

	team, _ := teams_service.GetTeamById(int(*application.TeamID))

	for _, user := range team.Users {
		if user.ID == user_id {
			return nil
		}
	}

	return errors.New("UserIsManagerOfTeamApplication: User is not a member of the team that owns this application")
}

func UserIsManagerOfTeamApplication(application_id uint, user_id uint) error {
	application, app_find_err := applications_service.GetApplicationById(int(application_id))
	if app_find_err != nil {
		return errors.New("UserIsManagerOfTeamApplication: Application does not exist")
	}

	if application.TeamID == nil {
		return errors.New("UserIsManagerOfTeamApplication: Application does not belong to a team")
	}

	team, _ := teams_service.GetTeamById(int(*application.TeamID))

	for _, manager := range team.Managers {
		if manager.ID == user_id {
			return nil
		}
	}

	return errors.New("UserIsManagerOfTeamApplication: User is not a manager of the team that owns this application")
}

func UserIsManagerOfTeam(team_id uint, user_id uint) error {
	team, team_find_err := teams_service.GetTeamById(int(team_id))
	if team_find_err != nil {
		return errors.New("UserIsManagerOfTeam: Team does not exist")
	}

	for _, manager := range team.Managers {
		if manager.ID == user_id {
			return nil
		}
	}

	return errors.New("UserIsManagerOfTeam: User not a manager of the specified team")
}

func UserIsMemberOfTeam(team_id uint, user_id uint) error {
	team, team_find_err := teams_service.GetTeamById(int(team_id))
	if team_find_err != nil {
		return errors.New("UserIsMemberOfTeam: Team does not exist")
	}

	for _, user := range team.Users {
		if user.ID == user_id {
			return nil
		}
	}

	return errors.New("UserIsMemberOfTeam: User not a member of the specified team")
}
