package orm

type ActionRate struct {
	ActionRateID int `gorm:"primaryKey"`
	HelpText     string
	Signature    string

	ActionID int
	Action   Action

	_ struct{} `gorm:"uniqueIndex:idx_action_signature,priority:1"`
}
