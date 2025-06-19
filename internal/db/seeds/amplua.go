package seeds

import (
	"log"
	"volley/internal/models/orm"

	"gorm.io/gorm"
)

func SeedAmpluas(db *gorm.DB) {
	ampluaNames := []string{
		"Доигровщик",
		"Центральный блокирующий",
		"Либеро",
		"Связующий",
		"Диагональный",
	}

	for _, name := range ampluaNames {
		amplua := &orm.Amplua{Name: name}
		where := map[string]interface{}{"name": name}

		if err := SeedOnce(db, where, amplua); err != nil {
			log.Printf("❌ Failed to seed amplua '%s': %v", name, err)
		} else {
			log.Printf("✅ Seeded amplua: %s", name)
		}
	}
}
