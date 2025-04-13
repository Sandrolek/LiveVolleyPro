package models

type ActionRate struct {
	ActionRateID int    `gorm:"primaryKey" json:"action_rate_id"`
	ActionID     int    `json:"action_id"`
	HelpText     string `gorm:"size:255" json:"help_text"`
	Signature    string `gorm:"size:45" json:"signature"`

	Action Action `gorm:"foreignKey:ActionID"`
}
