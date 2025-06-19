package orm

type SetAction struct {
	SetActionID  int `gorm:"primaryKey"`
	SetID        int `gorm:"index"`
	PlayerID     int `gorm:"index"`
	ActionRateID int `gorm:"index"`

	Set        Set
	Player     Player
	ActionRate ActionRate
}
