package repositories

import "volley/internal/models"

type UserRepository interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user models.User) (models.User, error)
}

type userRepository struct {
	// db *sql.DB
}

func NewUserRepository( /*db *sql.DB*/ ) UserRepository {
	return &userRepository{
		// db: db,
	}
}

func (r *userRepository) GetAllUsers() ([]models.User, error) {
	// Чтение из базы
	return []models.User{}, nil
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	// Запись в базу
	return user, nil
}
