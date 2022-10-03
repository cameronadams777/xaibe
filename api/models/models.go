package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
	IsAdmin   bool
}

type Team struct {
	gorm.Model
	Name string
}

type TeamUser struct {
	TeamID    uint
	Team      Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID    uint
	User      User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsManager bool
}

type Application struct {
	gorm.Model
	TeamID uint
	Team   Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

type ServiceToken struct {
	gorm.Model
	Token         string
	TeamID        uint
	Team          Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ApplicationID uint
	Application   Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExpiresAt     time.Time
}
