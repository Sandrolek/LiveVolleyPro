package orm

type Action struct {
	ActionID int `gorm:"primaryKey"`
	Name     string

	Rates []ActionRate
}
