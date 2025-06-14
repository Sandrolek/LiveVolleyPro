package dto

import "time"

type CreateGameDTO struct {
	Date      time.Time `json:"date" binding:"required"`
	Win       bool      `json:"win" binding:"required"`
	RoundID   int       `json:"round_id" binding:"required"`
	TeamID    int       `json:"team_id" binding:"required"`
	OppTeamID int       `json:"opp_team_id" binding:"required"`
}

type UpdateGameDTO struct {
	Date      *time.Time `json:"date,omitempty"`
	Win       *bool      `json:"win,omitempty"`
	RoundID   *int       `json:"round_id,omitempty"`
	TeamID    *int       `json:"team_id,omitempty"`
	OppTeamID *int       `json:"opp_team_id,omitempty"`
}
