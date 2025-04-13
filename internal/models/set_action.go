package models

type CreateSetAction struct {
	SetID        int `json:"set_id" binding:"required"`
	PlayerID     int `json:"player_id" binding:"required"`
	ActionRateID int `json:"action_rate_id" binding:"required"`
}

type UpdateSetAction struct {
	SetID        *int `json:"set_id"`
	PlayerID     *int `json:"player_id"`
	ActionRateID *int `json:"action_rate_id"`
}

type SetAction struct {
	SetActionID  int `gorm:"primaryKey" json:"set_action_id"`
	SetID        int `json:"set_id"`
	PlayerID     int `json:"player_id"`
	ActionRateID int `json:"action_rate_id"`

	Set        Set        `gorm:"foreignKey:SetID"`
	Player     Player     `gorm:"foreignKey:PlayerID"`
	ActionRate ActionRate `gorm:"foreignKey:ActionRateID"`
}
