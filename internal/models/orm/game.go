package orm

import "time"

type Game struct {
	GameID int `gorm:"primaryKey"`

	Date time.Time
	Win  *bool

	RoundID int
	Round   Round

	TeamID    int
	Team      Team
	OppTeamID int
	OppTeam   Team

	TeamScore    *int
	OppTeamScore *int

	Sets        []Set
	GamePlayers []GamePlayer
}
