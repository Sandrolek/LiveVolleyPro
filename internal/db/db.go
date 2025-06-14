package db

import (
	"log"
	"volley/internal/models/orm"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	dsn := "host=localhost user=postgres password=secret dbname=volley port=5432 sslmode=disable"

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL:", err)
	}

	// Drop tables in reverse dependency order
	// err = DB.Migrator().DropTable(
	// 	&orm.GamePlayer{},
	// 	&orm.SetAction{},
	// 	&orm.ActionRate{},
	// 	&orm.Action{},
	// 	&orm.Set{},
	// 	&orm.Game{},
	// 	&orm.Round{},
	// 	&orm.Championship{},
	// 	&orm.Player{},
	// 	&orm.Team{},
	// 	&orm.Amplua{},
	// 	&orm.User{},
	// )
	// if err != nil {
	// 	log.Fatalf("Failed to drop tables: %v", err)
	// }

	// Create tables in dependency order
	err = DB.AutoMigrate(
		&orm.User{},
		&orm.Amplua{},
		&orm.Team{},
		&orm.Player{},
		&orm.Game{},
		&orm.Round{},
		&orm.Championship{},
		&orm.GamePlayer{},
		&orm.Set{},
		&orm.SetAction{},
		&orm.ActionRate{},
		&orm.Action{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
}
