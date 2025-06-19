package orm

type GamePlayer struct {
	GamePlayerID int `gorm:"primaryKey"`
	PlayerID     int
	GameID       int

	Player Player
	Game   Game
}
