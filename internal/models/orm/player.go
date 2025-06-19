package orm

import "time"

type Player struct {
	PlayerID int `gorm:"primaryKey"`

	FirstName string
	LastName  string
	Birthdate *time.Time
	Gender    string `gorm:"type:text;check:gender IN ('male','female')"`
	Height    *float64
	Number    *int

	TeamID int
	Team   Team

	AmpluaID int
	Amplua   Amplua
}
