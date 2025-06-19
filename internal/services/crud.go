package services

import (
	"gorm.io/gorm"
)

type CRUDService[T any] struct {
	DB *gorm.DB
}

func NewCRUDService[T any](db *gorm.DB) *CRUDService[T] {
	return &CRUDService[T]{DB: db}
}

func (s *CRUDService[T]) Create(entity *T, preloadFields ...string) error {
	if err := s.DB.Create(entity).Error; err != nil {
		return err
	}

	// Reload the entity with preloaded fields
	tx := s.DB
	for _, field := range preloadFields {
		tx = tx.Preload(field)
	}
	return tx.First(entity).Error
}

func (s *CRUDService[T]) GetWhere(out interface{}, condition string, args any, order string, preload ...string) error {
	db := s.DB

	// Применяем preload для указанных связей
	for _, p := range preload {
		db = db.Preload(p)
	}

	return db.Where(condition, args).Order(order).Find(out).Error
}

func (s *CRUDService[T]) GetById(id any, out *T, fields ...string) error {
	tx := s.DB
	for _, field := range fields {
		tx = tx.Preload(field)
	}
	result := tx.First(out, id)
	return result.Error
}

func (s *CRUDService[T]) GetAll(out *[]T, orderBy string, fields ...string) error {
	tx := s.DB.Order(orderBy)
	for _, field := range fields {
		tx = tx.Preload(field)
	}
	return tx.Find(out).Error
}

func (s *CRUDService[T]) Update(id any, updates map[string]interface{}) error {
	var entity T
	result := s.DB.First(&entity, id)
	if result.Error != nil {
		return result.Error
	}
	return s.DB.Model(&entity).Updates(updates).Error
}

func (s *CRUDService[T]) Delete(id any) (bool, error) {
	result := s.DB.Delete(new(T), id)
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected > 0, nil
}
