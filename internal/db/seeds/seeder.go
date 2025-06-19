package seeds

import (
	"log"

	"gorm.io/gorm"
)

func SeedAll(db *gorm.DB) {
	log.Println("🔄 Seeding core data...")

	SeedActions(db)
	SeedActionRates(db)
	SeedAmpluas(db)

	log.Println("✅ Seeding completed")
}

func SeedOnce[T any](db *gorm.DB, where map[string]interface{}, newRecord *T) error {
	var existing T
	if err := db.Where(where).First(&existing).Error; err == gorm.ErrRecordNotFound {
		return db.Create(newRecord).Error
	} else {
		return nil
	}
}
