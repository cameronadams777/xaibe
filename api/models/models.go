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
  FirstName           string                                  `json:"first_name"`
	LastName            string                                  `json:"last_name"`
  Email               string                                  `json:"email"`
  Password            string                                  `json:"password"`
  StripeId            string                                  `json:"stripe_id"`
  IsAdmin             bool                                    `json:"is_admin"`
  IsVerified          bool                                    `json:"is_verified"`
  ResetPasswordCode   string                                  `json:"reset_password_code"`
  ResetPasswordExpiry time.Time                               `json:"reset_password_expiry"`
  Applications        []Application                           `json:"applications"`
  Teams               []*Team `gorm:"many2many:team_users;"`  `json:"teams"`
}

type Team struct {
  UUIDBaseModel
  Name         string                                   `json:"name"`
  Users        []*User `gorm:"many2many:team_users;"`   `json:"users"`
  Managers     []*User `gorm:"many2many:team_managers"` `json:"managers"`
  Applications []Application                            `json:"applications"`
}

type Application struct {
  UUIDBaseModel
  Name          string                                                               `json:"name"`
  TeamID        *uuid.UUID                                                           `json:"team_id"`
  Team          *Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`        `json:"team"`
  UserID        *uuid.UUID                                                           `json:"user_id"` 
  User          *User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`        `json:"user"`
  UniqueId      string                                                               `json:"unique_id"`
  AlertSchemaID *uuid.UUID                                                           `json:"alert_schema_id"`
  AlertSchema   AlertSchema `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`  `json:"alert_schema"`
  ServiceTokens []ServiceToken                                                       `json:"service_tokens"`
}

type ServiceToken struct {
	gorm.Model
  Token         string                                                               `json:"token"`
  ApplicationID uuid.UUID                                                            `json:"application_id"`
  Application   Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`  `json:"application"`
  ExpiresAt     time.Time                                                            `json:"expires_at"`
}

type AlertSchema struct {
  UUIDBaseModel
  ApplicationID uuid.UUID  `json:"application_id"`
  Title         string     `json:"title"`
  Description   string     `json:"description"`
  Link          string     `json:"link"`
}

type TeamInvite struct {
	gorm.Model
  SenderID uuid.UUID                                                      `json:"sender_id"`
  Sender   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`   `json:"sender"`
  TeamID   uuid.UUID                                                      `json:"team_id"`
  Team     Team `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`   `json:"team"`
  Email    string                                                         `json:"email"`
  Status   invite_status.InviteStatus                                     `json:"status"`
}
