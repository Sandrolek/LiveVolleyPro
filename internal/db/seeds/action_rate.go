package seeds

import (
	"log"

	"volley/internal/models/orm"

	"gorm.io/gorm"
)

func SeedActionRates(db *gorm.DB) {
	rateData := map[string][]struct {
		Signature string
		HelpText  string
	}{
		"Подача": {
			{Signature: "++", HelpText: "Эйс"},
			{Signature: "_", HelpText: "Обычная подача"},
			{Signature: "-", HelpText: "Ошибка на подаче"},
		},
		"Блок": {
			{Signature: "++", HelpText: "Заработали очко на блоке"},
			{Signature: "-", HelpText: "Ничего/Ошибка на блоке"},
		},
		"Атака": {
			{Signature: "++", HelpText: "Очко на атаке"},
			{Signature: "+-", HelpText: "Обратно переводят без атаки"},
			{Signature: "_", HelpText: "Обратно переводят с атакой"},
			{Signature: "-", HelpText: "Ошибка на подаче"},
		},
		"Защита": {
			{Signature: "++", HelpText: "Поднял в защите"},
			{Signature: "-", HelpText: "Ошибка в защите"},
		},
		"Прием": {
			{Signature: "++", HelpText: "Идеальный прием между 2 и 3 зоной"},
			{Signature: "+", HelpText: "Прием в пределах 3-хметровой линии"},
			{Signature: "_", HelpText: "Прием вверх"},
			{Signature: "-", HelpText: "Ошибка"},
		},
	}

	for actionName, rates := range rateData {
		var action orm.Action
		err := db.Where("name = ?", actionName).First(&action).Error
		if err != nil {
			log.Printf("❌ Action '%s' not found, skipping rates", actionName)
			continue
		}

		for _, r := range rates {
			rate := &orm.ActionRate{
				ActionID:  action.ActionID,
				Signature: r.Signature,
				HelpText:  r.HelpText,
			}
			where := map[string]interface{}{
				"action_id": action.ActionID,
				"signature": r.Signature,
			}
			if err := SeedOnce(db, where, rate); err != nil {
				log.Printf("❌ Failed to seed rate '%s' for action '%s': %v", r.Signature, actionName, err)
			} else {
				log.Printf("✅ Seeded rate '%s' for action '%s'", r.Signature, actionName)
			}
		}
	}
}
