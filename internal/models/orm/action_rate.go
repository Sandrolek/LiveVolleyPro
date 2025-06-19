package orm

type ActionRate struct {
	ActionRateID int `gorm:"primaryKey"`
	HelpText     string
	Signature    string

	ActionID int
	Action   Action

	_ struct{} `gorm:"uniqueIndex:idx_action_signature,priority:1"`
}

// type ActionRate struct {
// 	ActionRateID int    `gorm:"primaryKey" json:"action_rate_id"`
// 	ActionID     int    `json:"action_id"`
// 	HelpText     string `gorm:"size:255" json:"help_text"`
// 	Signature    string `gorm:"size:45" json:"signature"`

// 	//Action Action `gorm:"foreignKey:ActionID"`
// 	SetActions []SetAction `gorm:"foreignKey:ActionRateID"`
// }
