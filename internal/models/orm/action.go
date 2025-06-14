package orm

type Action struct {
	ActionID int `gorm:"primaryKey"`
	Name     string

	Rates []ActionRate
}

// type Action struct {
// 	ActionID int    `gorm:"primaryKey" json:"action_id"`
// 	Name     string `gorm:"size:45" json:"name"`

// 	ActionRates []ActionRate `gorm:"foreignKey:ActionID"`
// }
