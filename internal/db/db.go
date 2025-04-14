package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"volley/internal/models"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=secret dbname=volley port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	err = DB.AutoMigrate(
		&models.User{},
		&models.Team{},
		&models.Player{},
		&models.Amplua{},
		&models.Game{},
		&models.Round{},
		&models.Championship{},
		&models.GamePlayer{},
		&models.Set{},
		&models.SetAction{},
		&models.ActionRate{},
		&models.Action{},
	)

	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
