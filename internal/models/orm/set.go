package orm

type Set struct {
	SetID int `gorm:"primaryKey"`

	SerialNumber int
	TeamScore    *int
	OppScore     *int

	GameID int
	Game   Game

	SetActions []SetAction
}
