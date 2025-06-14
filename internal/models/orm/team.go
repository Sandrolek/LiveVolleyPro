package orm

type Team struct {
	TeamID  int `gorm:"primaryKey"`
	Name    string
	UserID  int
	User    User
	Players []Player
}

// type Team struct {
// 	gorm.Model
// 	TeamID int    `gorm:"primaryKey;autoIncrement" json:"team_id"`
// 	Name   string `gorm:"size:45" json:"name"`

// 	UserID int
// 	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

// 	// UserID int  `json:"user_id"`
// 	// User   User `gorm:"foreignKey:UserID;references:UserID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`

// 	//User    User     `gorm:"foreignKey:TeamID"`
// 	// Players []Player `gorm:"foreignKey:TeamID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
// 	//Games   []Game   `gorm:"foreignKey:TeamID"`
// }

type CreateTeam struct {
	UserID int    `json:"user_id" binding:"required"`
	Name   string `json:"name" binding:"required"`
}

type UpdateTeam struct {
	Name string `json:"name" binding:"required"`
}
