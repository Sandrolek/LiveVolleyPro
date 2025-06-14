package orm

import "time"

type Game struct {
	GameID int `gorm:"primaryKey"`

	Date time.Time
	Win  bool

	RoundID int
	Round   Round

	TeamID int
	Team   Team

	OppTeamID int
	OppTeam   Team `gorm:"foreignKey:OppTeamID"`

	Sets        []Set
	GamePlayers []GamePlayer
}

type CreateGame struct {
	RoundID   int    `json:"round_id" binding:"required"`
	TeamID    int    `json:"team_id" binding:"required"`
	OppTeamID int    `json:"opp_team_id" binding:"required"`
	Date      string `json:"date" binding:"required"`
	Win       bool   `json:"win"`
}

type UpdateGame struct {
	Date string `json:"date"`
	Win  *bool  `json:"win"`
}

// type Game struct {
// 	GameID  int       `gorm:"primaryKey" json:"game_id"`
// 	RoundID int       `json:"round_id"`
// 	TeamID  int       `json:"team_id"`
// 	Date    time.Time `json:"date"`
// 	Win     bool      `json:"win"`

// 	//Rounds  []Round      `gorm:"foreignKey:GameID"`
// 	Players []GamePlayer `gorm:"foreignKey:GameID"`
// 	Team    Team         `gorm:"foreignKey:TeamID"`
// 	Sets    []Set        `gorm:"foreignKey:GameID"`
// }
