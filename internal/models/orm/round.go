package orm

import "time"

type Round struct {
	RoundID      int `gorm:"primaryKey"`
	SerialNumber uint8
	StartDate    *time.Time
	EndDate      *time.Time

	ChampionshipID int
	Championship   Championship

	Games []Game
}

type CreateRound struct {
	SerialNumber   uint8  `json:"serial_number" binding:"required"`
	ChampionshipID int    `json:"championship_id" binding:"required"`
	StartDate      string `json:"start_date"` // "YYYY-MM-DD"
	EndDate        string `json:"end_date"`   // "YYYY-MM-DD"
}

// type Round struct {
// 	RoundID        int        `gorm:"primaryKey" json:"round_id"`
// 	SerialNumber   uint8      `json:"serial_number"`
// 	ChampionshipID int        `json:"championship_id"`
// 	StartDate      *time.Time `json:"start_date"`
// 	EndDate        *time.Time `json:"end_date"`

// 	//Championship Championship `gorm:"foreignKey:ChampionshipID"`
// 	Games []Game `gorm:"foreignKey:RoundID"`
// }
