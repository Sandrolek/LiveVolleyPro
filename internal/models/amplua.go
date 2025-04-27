package models

type CreateAmplua struct {
	Name string `json:"name" binding:"required"`
}

type UpdateAmplua struct {
	Name string `json:"name" binding:"required"`
}

type Amplua struct {
	AmpluaID int    `gorm:"primaryKey" json:"amplua_id"`
	Name     string `gorm:"size:45" json:"name"`

	Players []Player `gorm:"foreignKey:AmpluaID"`
}
