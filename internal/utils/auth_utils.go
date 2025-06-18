package utils

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"math/rand"
	"time"
	"volley/internal/models/orm"
)

func HashPassword(password string) string {
	hashed, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashed)
}

func CheckPassword(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func GenerateToken(db *gorm.DB, userID uint) (string, error) {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, 32)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}

	token := string(b)

	authToken := orm.AuthToken{
		Token:     token,
		UserID:    userID,
		ExpiresAt: time.Now().Add(24 * time.Hour * 7), // 7 дней
	}

	// Правильное использование GORM
	if err := db.Create(&authToken).Error; err != nil {
		return "", err
	}

	return token, nil
}
