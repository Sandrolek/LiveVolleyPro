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
