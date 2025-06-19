package dto

type CreateTeamDTO struct {
	Name string `json:"name" binding:"required"`
}

type UpdateTeamDTO struct {
	Name *string `json:"name,omitempty"`
}
