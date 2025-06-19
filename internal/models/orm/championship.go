package orm

import "time"

type Championship struct {
	ChampionshipID int `gorm:"primaryKey"`
	Title          string
	StartDate      time.Time
	EndDate        time.Time

	Rounds []Round
}
