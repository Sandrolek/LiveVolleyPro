package dto

import "time"

type CreateRoundDTO struct {
	SerialNumber   uint8      `json:"serial_number" binding:"required"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	ChampionshipID int        `json:"championship_id" binding:"required"`
}

type UpdateRoundDTO struct {
	SerialNumber   *uint8     `json:"serial_number,omitempty"`
	StartDate      *time.Time `json:"start_date,omitempty"`
	EndDate        *time.Time `json:"end_date,omitempty"`
	ChampionshipID *int       `json:"championship_id,omitempty"`
}
