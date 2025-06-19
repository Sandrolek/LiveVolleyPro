package services

import (
	"errors"

	"volley/internal/models/dto"
	"volley/internal/models/orm"

	"gorm.io/gorm"
)

type StatsService struct {
	DB *gorm.DB
}

func NewStatsService(db *gorm.DB) *StatsService {
	return &StatsService{DB: db}
}

func (s *StatsService) GetPlayerStats(input dto.PlayerStatsRequestDTO) (map[string]map[string]int, error) {
	if input.GameID == nil && input.SetID == nil {
		return nil, errors.New("either set_id or game_id must be provided")
	}

	query := s.DB.Model(&orm.SetAction{}).
		Select("actions.name AS action_name, action_rates.signature, COUNT(*) as total").
		Joins("JOIN action_rates ON action_rates.action_rate_id = set_actions.action_rate_id").
		Joins("JOIN actions ON actions.action_id = action_rates.action_id").
		Where("set_actions.player_id = ?", input.PlayerID)

	if input.SetID != nil {
		query = query.Where("set_actions.set_id = ?", *input.SetID)
	}

	if input.GameID != nil {
		query = query.Joins("JOIN sets ON sets.set_id = set_actions.set_id").
			Where("sets.game_id = ?", *input.GameID)
	}

	type row struct {
		ActionName string
		Signature  string
		Total      int
	}

	var rows []row
	if err := query.Group("actions.name, action_rates.signature").Scan(&rows).Error; err != nil {
		return nil, err
	}

	result := make(map[string]map[string]int)
	for _, r := range rows {
		if _, ok := result[r.ActionName]; !ok {
			result[r.ActionName] = make(map[string]int)
		}
		result[r.ActionName][r.Signature] = r.Total
	}

	return result, nil
}
