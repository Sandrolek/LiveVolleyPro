package services

import (
	"errors"
	"fmt"
	"regexp"
	"strings"

	"volley/internal/models/orm"

	"gorm.io/gorm"
)

type SetActionService struct {
	DB *gorm.DB
}

func NewSetActionService(db *gorm.DB) *SetActionService {
	return &SetActionService{DB: db}
}

var actionAbbrMap = map[string]string{
	"A": "Атака",
	"B": "Блок",
	"C": "Защита",
	"S": "Подача",
	"R": "Прием",
}

func (s *SetActionService) RecordAction(setID int, record string) error {
	// Parse format like "12A++" or "5R-"
	re := regexp.MustCompile(`^(\d+)([ABC SR])([+_\-]{1,2}|\+\-)$`)
	matches := re.FindStringSubmatch(record)
	if len(matches) != 4 {
		return fmt.Errorf("invalid record format: %s", record)
	}

	playerNumber := matches[1]
	abbr := strings.TrimSpace(matches[2])
	rate := matches[3]

	actionName, ok := actionAbbrMap[abbr]
	if !ok {
		return fmt.Errorf("unknown action abbreviation: %s", abbr)
	}

	// Find Player by number
	var player orm.Player
	if err := s.DB.Where("number = ?", playerNumber).First(&player).Error; err != nil {
		return errors.New("player not found")
	}

	// Find Action by name
	var action orm.Action
	if err := s.DB.Where("name = ?", actionName).First(&action).Error; err != nil {
		return errors.New("action not found")
	}

	// Find ActionRate by ActionID + signature
	var rateObj orm.ActionRate
	if err := s.DB.Where("action_id = ? AND signature = ?", action.ActionID, rate).First(&rateObj).Error; err != nil {
		return errors.New("action rate not found")
	}

	// Create SetAction record
	setAction := orm.SetAction{
		SetID:        setID,
		PlayerID:     player.PlayerID,
		ActionRateID: rateObj.ActionRateID,
	}

	if err := s.DB.Create(&setAction).Error; err != nil {
		return fmt.Errorf("failed to create set action: %w", err)
	}

	return nil
}
