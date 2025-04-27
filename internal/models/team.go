package models

type Team struct {
	TeamID int    `gorm:"primaryKey" json:"team_id"`
	UserID int    `json:"user_id"`
	Name   string `gorm:"size:45" json:"name"`

	//User    User     `gorm:"foreignKey:TeamID"`
	Players []Player `gorm:"foreignKey:TeamID"`
	//Games   []Game   `gorm:"foreignKey:TeamID"`
}

type CreateTeam struct {
	UserID int    `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type UpdateTeam struct {
	Name string `json:"name" binding:"required"`
}
