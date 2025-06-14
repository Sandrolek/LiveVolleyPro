package dto

type CreateAmpluaDTO struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAmpluaDTO struct {
	Name *string `json:"name,omitempty"`
}
