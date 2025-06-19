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
