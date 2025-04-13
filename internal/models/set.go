package models

type CreateSet struct {
	GameID       int `json:"game_id" binding:"required"`
	SerialNumber int `json:"serial_number" binding:"required"`
	TeamScore    int `json:"team_score" binding:"required"`
	OppScore     int `json:"opp_score" binding:"required"`
}

type UpdateSet struct {
	SerialNumber *int `json:"serial_number"`
	TeamScore    *int `json:"team_score"`
	OppScore     *int `json:"opp_score"`
}

type Set struct {
	SetID        int `gorm:"primaryKey" json:"set_id"`
	GameID       int `json:"game_id"`
	SerialNumber int `json:"serial_number"`
	TeamScore    int `json:"team_score"`
	OppScore     int `json:"opp_score"`

	Game       Game        `gorm:"foreignKey:GameID"`
	SetActions []SetAction `gorm:"foreignKey:SetID"`
}
