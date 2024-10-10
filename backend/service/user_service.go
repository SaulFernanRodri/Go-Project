package service

import (
	"my-backend/domain"
	"my-backend/repository"
)

type UserService struct {
	userRepository repository.UserRepository
}

// NewUserService crea un nuevo servicio de usuarios.
func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{userRepository: repo}
}

// Crear un nuevo usuario aplicando validaciones y reglas de negocio.
func (s *UserService) CreateUser(name, email, password string) (*domain.User, error) {
	user, err := domain.NewUser(name, email, password)
	if err != nil {
		return nil, err
	}
	return s.userRepository.Create(user)
}

// Obtener todos los usuarios.
func (s *UserService) GetAllUsers() ([]*domain.User, error) {
	return s.userRepository.FindAll()
}
