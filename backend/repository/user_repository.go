package repository

import "my-backend/domain"

// UserRepository es la interfaz que define las operaciones para gestionar usuarios en la base de datos.
type UserRepository interface {
	Create(user *domain.User) (*domain.User, error)
	FindByID(id uint) (*domain.User, error)
	FindAll() ([]*domain.User, error)
	Update(user *domain.User) (*domain.User, error)
	Delete(id uint) error
}
