package models

import "time"

type Player struct {
	PlayerID  int        `gorm:"primaryKey" json:"player_id"`
	AmpluaID  int        `json:"amplua_id"`
	FirstName string     `gorm:"size:45" json:"first_name"`
	LastName  string     `gorm:"size:45" json:"last_name"`
	Birthdate *time.Time `json:"birthdate"`
	Gender    string     `gorm:"type:varchar(10)" json:"gender"`
	Height    *float64   `json:"height"`
	Number    *uint8     `json:"number"`
	TeamID    int        `json:"team_id"`

	//Amplua Amplua `gorm:"foreignKey:PlayerID"`
	//Team   Team   `gorm:"foreignKey:TeamID;"`
	SetActions []SetAction `gorm:"foreignKey:PlayerID"`
}

type CreatePlayer struct {
	AmpluaID  int      `json:"amplua_id" binding:"required"`
	FirstName string   `json:"first_name" binding:"required"`
	LastName  string   `json:"last_name" binding:"required"`
	Birthdate string   `json:"birthdate"`
	Gender    string   `gorm:"type:varchar(10)" json:"gender"`
	Height    *float64 `json:"height"`
	Number    *uint8   `json:"number"`
	TeamID    int      `json:"team_id" binding:"required"`
}

type UpdatePlayer struct {
	FirstName string     `json:"first_name"`
	LastName  string     `json:"last_name"`
	Birthdate *time.Time `json:"birthdate"`
	Gender    string     `gorm:"type:varchar(10)" json:"gender"`
	Height    *float64   `json:"height"`
	Number    *uint8     `json:"number"`
}
