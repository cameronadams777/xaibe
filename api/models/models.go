package models

import (
	"api/structs/invite_status"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName           string
	LastName            string
	Email               string
	Password            string
	IsAdmin             bool
	IsVerified          bool
	ResetPasswordCode   string
	ResetPasswordExpiry time.Time
	Applications        []Application
	Teams               []*Team `gorm:"many2many:team_users;"`
}

type Team struct {
	gorm.Model
	Name         string
	Users        []*User `gorm:"many2many:team_users;"`
	Managers     []*User `gorm:"many2many:team_managers"`
	Applications []Application
}

type Application struct {
	gorm.Model
	Name          string
	TeamID        *uint
	Team          *Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID        *uint
	User          *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UniqueId      string
	AlertSchemaID *uint
	AlertSchema   AlertSchema `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ServiceTokens []ServiceToken
}

type ServiceToken struct {
	gorm.Model
	Token         string
	ApplicationID uint
	Application   Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExpiresAt     time.Time
}

type AlertSchema struct {
	gorm.Model
	ApplicationID uint
	Title         string
	Description   string
	Link          string
}

type TeamInvite struct {
	gorm.Model
	SenderID uint
	Sender   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TeamID   uint
	Team     Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email    string
	Status   invite_status.InviteStatus
}
