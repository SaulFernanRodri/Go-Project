package services

import (
	"myproject/models"
	"myproject/repositories"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(id uint64, user *models.User) (*models.User, error)
	DeleteUser(id uint64) error
	GetByUsername(username string) ([]models.User, error)
}

type UserService struct {
	repo repositories.UserRepoInterface
}

func NewUserService(repo repositories.UserRepoInterface) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	return s.repo.Create(user)
}

func (s *UserService) UpdateUser(id uint64, user *models.User) (*models.User, error) {
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id uint64) error {
	return s.repo.Delete(id)
}

func (s *UserService) GetByUsername(username string) ([]models.User, error) {
	return s.repo.GetByUsername(username)
}
