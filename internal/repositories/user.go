package repositories

import (
	"volley/internal/models/orm" // замените на актуальный путь

	"gorm.io/gorm"
)

type UserRepository interface {
	GetByID(id int) (*orm.User, error)
	GetAll() ([]orm.User, error)
	Create(user *orm.User) error
	Update(user *orm.User) error
	Delete(id int) error
	FindByEmail(email string) (*orm.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetByID(id int) (*orm.User, error) {
	var user orm.User
	if err := r.db.Preload("Teams").First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]orm.User, error) {
	var users []orm.User
	if err := r.db.Preload("Teams").Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Create(user *orm.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) Update(user *orm.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id int) error {
	return r.db.Delete(&orm.User{}, id).Error
}

func (r *userRepository) FindByEmail(email string) (*orm.User, error) {
	var user orm.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
