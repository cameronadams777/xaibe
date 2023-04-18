package models

import (
	"api/structs/invite_status"
	"time"

	"gorm.io/gorm"
  "github.com/google/uuid"
)

type UUIDBaseModel struct {
  ID        uuid.UUID      `gorm:"primaryKey" json:"id"`
  CreatedAt time.Time      `json:"createdAt"` 
  UpdatedAt time.Time      `json:"updatedAt"`
  DeletedAt gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (base *UUIDBaseModel) BeforeCreate(tx *gorm.DB) error {
  uuid := uuid.NewString()
  tx.Statement.SetColumn("ID", uuid)
  return nil
}

type User struct {
  UUIDBaseModel
  FirstName           string        `json:"firstName"`
	LastName            string        `json:"lastName"`
  Email               string        `json:"email"`
  Password            string        `json:"password"`
  StripeId            string        `json:"stripeId"`
  IsAdmin             bool          `json:"isAdmin"`
  IsVerified          bool          `json:"isVerified"`
  ResetPasswordCode   string        `json:"resetPasswordCode"`
  ResetPasswordExpiry time.Time     `json:"resetPasswordExpiry"`
  Applications        []Application `json:"applications"`
  Teams               []Team       `gorm:"many2many:team_users;" json:"teams"`
}

type Team struct {
  UUIDBaseModel
  Name                string          `json:"name"`
  SubscriptionId      *string         `json:"subscriptionId"`
  ActiveNumberOfSeats uint            `json:"activeNumberOfSeats"` 
  Users               []User          `gorm:"many2many:team_users;" json:"users"`
  Managers            []User          `gorm:"many2many:team_managers" json:"managers"`
  Applications        []Application   `json:"applications"`
}

type Application struct {
  UUIDBaseModel
  Name          string          `json:"name"`
  TeamID        *uuid.UUID      `json:"teamId"`
  Team          *Team           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team"`
  UserID        *uuid.UUID      `json:"userId"` 
  User          *User           `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json"user"`
  UniqueId      string          `json:"uniqueId"`
  AlertSchemaID *uuid.UUID      `json:"alertSchemaId"`
  AlertSchema   AlertSchema     `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"alertSchema"`
  ServiceTokens []ServiceToken  `json:"serviceTokens"`
}

type ServiceToken struct {
  UUIDBaseModel
  Token         string      `json:"token"`
  ApplicationID uuid.UUID   `json:"applicationId"`
  Application   Application `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"application"`
  ExpiresAt     time.Time   `json:"expiresAt"`
}

type AlertSchema struct {
  UUIDBaseModel
  ApplicationID uuid.UUID   `json:"applicationId"`
  Title         string      `json:"title"`
  Description   string      `json:"description"`
  Link          string      `json:"link"`
}

type TeamInvite struct {
  UUIDBaseModel
  SenderID uuid.UUID                    `json:"senderId"`
  Sender   User                         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"sender"`
  TeamID   uuid.UUID                    `json:"teamId"`
  Team     Team                         `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;" json:"team"`
  Email    string                       `json:"email"`
  Status   invite_status.InviteStatus   `json:"status"`
}
