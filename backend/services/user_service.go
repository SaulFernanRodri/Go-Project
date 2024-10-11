package services

import (
	"myproject/models"
	"myproject/repositories"

	"gorm.io/gorm"
)

type UserService struct {
	repo *repositories.UserRepository
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		repo: repositories.NewUserRepository(db),
	}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user *models.User) error {
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(id string, user *models.User) (*models.User, error) {
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
