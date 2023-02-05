package models

import (
	"api/structs/invite_status"
	"time"

	"gorm.io/gorm"
  "github.com/google/uuid"
)

type UUIDBaseModel struct {
  ID        uuid.UUID      `gorm:"primary_key" json:"id"`
  CreatedAt time.Time      `json:"created_at"` 
  UpdatedAt time.Time      `json:"updated_at"`
  DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func (base *UUIDBaseModel) BeforeCreate(tx *gorm.DB) error {
  uuid := uuid.NewString()
  tx.Statement.SetColumn("ID", uuid)
  return nil
}

type User struct {
  UUIDBaseModel
	FirstName           string
	LastName            string
	Email               string
	Password            string
  StripeId            string
	IsAdmin             bool
	IsVerified          bool
	ResetPasswordCode   string
	ResetPasswordExpiry time.Time
	Applications        []Application
	Teams               []*Team `gorm:"many2many:team_users;"`
}

type Team struct {
  UUIDBaseModel
	Name         string
	Users        []*User `gorm:"many2many:team_users;"`
	Managers     []*User `gorm:"many2many:team_managers"`
	Applications []Application
}

type Application struct {
  UUIDBaseModel
	Name          string
	TeamID        *uuid.UUID
	Team          *Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID        *uuid.UUID
	User          *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UniqueId      string
	AlertSchemaID *uuid.UUID
	AlertSchema   AlertSchema `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ServiceTokens []ServiceToken
}

type ServiceToken struct {
	gorm.Model
	Token         string
	ApplicationID uuid.UUID
	Application   Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	ExpiresAt     time.Time
}

type AlertSchema struct {
  UUIDBaseModel
	ApplicationID uuid.UUID
	Title         string
	Description   string
	Link          string
}

type TeamInvite struct {
	gorm.Model
	SenderID uuid.UUID
	Sender   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	TeamID   uuid.UUID
	Team     Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Email    string
	Status   invite_status.InviteStatus
}
