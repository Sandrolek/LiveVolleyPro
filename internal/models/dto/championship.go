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

// type CreateChampionship struct {
// 	Title     string `json:"title" binding:"required"`
// 	StartDate string `json:"start_date" binding:"required"` // формат: YYYY-MM-DD
// 	EndDate   string `json:"end_date" binding:"required"`   // формат: YYYY-MM-DD
// }

// type UpdateChampionship struct {
// 	Title     *string `json:"title"`
// 	StartDate string  `json:"start_date"`
// 	EndDate   string  `json:"end_date"`
// }

// type Championship struct {
// 	ChampionshipID int       `gorm:"primaryKey" json:"championship_id"`
// 	Title          string    `gorm:"size:255" json:"title"`
// 	StartDate      time.Time `json:"start_date"`
// 	EndDate        time.Time `json:"end_date"`

// 	Rounds []Round `gorm:"foreignKey:ChampionshipID;references:ChampionshipID"`
// }
