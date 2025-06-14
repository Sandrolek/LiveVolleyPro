package dto

type CreateTeamDTO struct {
	Name   string `json:"name" binding:"required"`
	UserID int    `json:"user_id" binding:"required"`
}

type UpdateTeamDTO struct {
	Name   *string `json:"name,omitempty"`
	UserID *int    `json:"user_id,omitempty"`
}
