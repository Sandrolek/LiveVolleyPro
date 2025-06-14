package dto

type CreateSetDTO struct {
	SerialNumber int `json:"serial_number" binding:"required"`
	TeamScore    int `json:"team_score" binding:"required"`
	OppScore     int `json:"opp_score" binding:"required"`
	GameID       int `json:"game_id" binding:"required"`
}

type UpdateSetDTO struct {
	SerialNumber *int `json:"serial_number,omitempty"`
	TeamScore    *int `json:"team_score,omitempty"`
	OppScore     *int `json:"opp_score,omitempty"`
	GameID       *int `json:"game_id,omitempty"`
}
