package seeds

import (
	"log"
	"volley/internal/models/orm"

	"gorm.io/gorm"
)

func SeedActions(db *gorm.DB) {
	actions := []string{"Подача", "Блок", "Атака", "Защита", "Прием"}

	for _, name := range actions {
		action := &orm.Action{Name: name}
		if err := SeedOnce(db, map[string]interface{}{"name": name}, action); err != nil {
			log.Printf("❌ Failed to seed action '%s': %v", name, err)
		} else {
			log.Printf("✅ Seeded or found action: %s", name)
		}
	}
}
