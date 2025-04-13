package models

type Action struct {
	ActionID int    `gorm:"primaryKey" json:"action_id"`
	Name     string `gorm:"size:45" json:"name"`
}
