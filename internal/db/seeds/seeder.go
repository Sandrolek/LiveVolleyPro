package seeds

import (
	"log"

	"gorm.io/gorm"
)

// SeedAll is the main entrypoint for seeding essential data.
func SeedAll(db *gorm.DB) {
	log.Println("🔄 Seeding core data...")

	SeedActions(db)     // must come before seeding rates
	SeedActionRates(db) // depends on existing Action rows
	SeedAmpluas(db)

	log.Println("✅ Seeding completed")
}

// SeedOnce checks if a record exists and creates it if not.
func SeedOnce[T any](db *gorm.DB, where map[string]interface{}, newRecord *T) error {
	var existing T
	if err := db.Where(where).First(&existing).Error; err == gorm.ErrRecordNotFound {
		return db.Create(newRecord).Error
	} else {
		return nil // already exists — skip
	}
}
