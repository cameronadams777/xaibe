package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	FirstName string
	LastName  string
	Email     string
	Password  string
}

type Team struct {
	gorm.Model
}

type TeamUser struct {
	TeamID  uint
	Team    Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID  uint
	User    User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	IsAdmin bool
}

type Application struct {
	gorm.Model
	TeamID uint
	Team   Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}
