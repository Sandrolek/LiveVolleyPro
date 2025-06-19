package dto

import "time"

type CreateChampionshipDTO struct {
	Title     string    `json:"title" binding:"required"`
	StartDate time.Time `json:"start_date" binding:"required"`
	EndDate   time.Time `json:"end_date" binding:"required,gtfield=StartDate"`
}

type UpdateChampionshipDTO struct {
	Title     *string    `json:"title,omitempty"`
	StartDate *time.Time `json:"start_date,omitempty"`
	EndDate   *time.Time `json:"end_date,omitempty"`
}
