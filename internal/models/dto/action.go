package dto

type CreateActionDTO struct {
	Name string `json:"name" binding:"required"`
}

type UpdateActionDTO struct {
	Name *string `json:"name,omitempty"`
}
