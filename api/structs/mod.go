package structs

import "github.com/google/uuid"

type AuthScope struct {
	UserID uuid.UUID
}

type ErrorMessage struct {
  Message string `json:"message" binding:"required"` 
}
