package dto

type RecordActionDTO struct {
	SetID  int    `json:"set_id" binding:"required"`
	Record string `json:"record" binding:"required"` // e.g. "12A++"
}
