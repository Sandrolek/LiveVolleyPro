package orm

type GamePlayer struct {
	GamePlayerID int `gorm:"primaryKey"`
	PlayerID     int
	GameID       int

	Player Player
	Game   Game
}

type CreateGamePlayer struct {
	PlayerID int `json:"player_id" binding:"required"`
	GameID   int `json:"game_id" binding:"required"`
}

type UpdateGamePlayer struct {
	PlayerID *int `json:"player_id"`
	GameID   *int `json:"game_id"`
}

// type GamePlayer struct {
// 	GamePlayerID int `gorm:"primaryKey" json:"game_player_id"`
// 	PlayerID     int `json:"player_id"`
// 	GameID       int `json:"game_id"`

// 	//Player Player `gorm:"foreignKey:PlayerID"`
// 	//Game   Game   `gorm:"foreignKey:GameID;"`
// }
