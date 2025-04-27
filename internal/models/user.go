package models

type CreateUser struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Email    string `json:"email" binding:"required"`
}

type User struct {
	UserID   int    `gorm:"primaryKey" json:"user_id"`
	Name     string `gorm:"size:45" json:"name"`
	Password string `gorm:"size:255" json:"password"`
	Email    string `gorm:"size:255" json:"email"`

	Teams []Team `gorm:"foreignKey:UserID" json:"teams"`
}
