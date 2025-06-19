package orm

import (
	"time"
)

type User struct {
	UserID int `gorm:"primaryKey"`

	Name     string
	Password string
	Email    string `gorm:"unique;not null"`

	Teams []Team
}

type AuthToken struct {
	Token     string `gorm:"unique;not null"`
	UserID    uint   `gorm:"not null"`
	ExpiresAt time.Time
}

// type User struct {
// 	gorm.Model
// 	UserID   int    `gorm:"primaryKey" json:"user_id"`
// 	Name     string `gorm:"size:45" json:"name"`
// 	Password string `gorm:"size:255" json:"password"`
// 	Email    string `gorm:"size:255" json:"email"`

// 	// Teams []Team `gorm:"foreignKey:UserID"`
// 	// Teams []Team `gorm:"foreignKey:UserID" json:"teams"`
// }
