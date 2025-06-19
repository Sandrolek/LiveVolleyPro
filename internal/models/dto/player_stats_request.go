package dto

type PlayerStatsRequestDTO struct {
	PlayerID int  `json:"player_id" binding:"required"`
	SetID    *int `json:"set_id,omitempty"`  // optional
	GameID   *int `json:"game_id,omitempty"` // optional
}
