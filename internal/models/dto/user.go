package dto

type CreateUserDTO struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}

type UpdateUserDTO struct {
	Name     *string `json:"name,omitempty"`
	Password *string `json:"password,omitempty"`
	Email    *string `json:"email,omitempty"`
}

type UserInput struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}
