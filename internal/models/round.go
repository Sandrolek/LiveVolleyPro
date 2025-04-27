package models

import "time"

type CreateRound struct {
	SerialNumber   uint8  `json:"serial_number" binding:"required"`
	ChampionshipID int    `json:"championship_id" binding:"required"`
	StartDate      string `json:"start_date"` // "YYYY-MM-DD"
	EndDate        string `json:"end_date"`   // "YYYY-MM-DD"
}

type UpdateRound struct {
	SerialNumber *uint8 `json:"serial_number"`
	StartDate    string `json:"start_date"`
	EndDate      string `json:"end_date"`
}

type Round struct {
	RoundID        int        `gorm:"primaryKey" json:"round_id"`
	SerialNumber   uint8      `json:"serial_number"`
	ChampionshipID int        `json:"championship_id"`
	StartDate      *time.Time `json:"start_date"`
	EndDate        *time.Time `json:"end_date"`

	//Championship Championship `gorm:"foreignKey:ChampionshipID"`
	Games []Game `gorm:"foreignKey:RoundID"`
}
