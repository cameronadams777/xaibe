package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName    string
	LastName     string
	Email        string
	Password     string
	IsAdmin      bool
	IsVerified   bool
	Applications []Application
	Teams        []*Team `gorm:"many2many:team_users;"`
}

type Team struct {
	gorm.Model
	Name         string
	Users        []*User `gorm:"many2many:team_users;"`
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
	AlertSchema   *AlertSchema `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
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
}
