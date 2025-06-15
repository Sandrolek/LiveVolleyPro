package dto

import (
	"time"
)

type CreatePlayerDTO struct {
	FirstName string     `json:"first_name" binding:"required"`
	LastName  string     `json:"last_name" binding:"required"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	Gender    string     `json:"gender" binding:"required,oneof=male female"`
	Height    *float64   `json:"height,omitempty"`
	Number    *int       `json:"number,omitempty"`
	TeamID    int        `json:"team_id" binding:"required"`
	AmpluaID  int        `json:"amplua_id" binding:"required"`
}

type UpdatePlayerDTO struct {
	FirstName *string    `json:"first_name,omitempty"`
	LastName  *string    `json:"last_name,omitempty"`
	Birthdate *time.Time `json:"birthdate,omitempty"`
	Gender    *string    `json:"gender,omitempty"` // optionally validate in handler
	Height    *float64   `json:"height,omitempty"`
	Number    *int       `json:"number,omitempty"`
	TeamID    *int       `json:"team_id,omitempty"`
	AmpluaID  *int       `json:"amplua_id,omitempty"`
}
