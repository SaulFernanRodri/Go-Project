package services

import (
	"myproject/models"
	"myproject/repositories"
	"myproject/utils"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.UserRequest) (*models.UserRequest, error)
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

func (s *UserService) CreateUser(userRequest *models.UserRequest) (*models.UserRequest, error) {

	csvPath, error := utils.GenerateMilsymbol(userRequest.Milsymbol)
	if error != nil {
		return userRequest, error
	}
	user := &models.User{
		Name:         userRequest.Name,
		CSV:          csvPath,
		AuthUsername: userRequest.AuthUsername,
	}

	if err := s.repo.Create(user); err != nil {
		return userRequest, err
	}

	return userRequest, nil
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
