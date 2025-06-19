package orm

type Amplua struct {
	AmpluaID int `gorm:"primaryKey"`
	Name     string
	Players  []Player
}
