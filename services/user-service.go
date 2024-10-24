package services

import (
	"myproject/models"
	"myproject/repositories"
	"myproject/utils"
)

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(id uint64, user *models.User) (*models.User, error)
	DeleteUser(id uint64) error
	GetByUsername(username string) ([]models.User, error)
}

type UserService struct {
	repo *repositories.UserRepo
}

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(UserRequest *models.UserRequest) (*models.UserRequest, error) {

	csvPath, error := utils.GenerateMilsymbol(UserRequest.Milsymbol)
	if error != nil {
		return UserRequest, error
	}
	user := &models.User{
		Name:         UserRequest.Name,
		CSV:          csvPath,
		AuthUsername: UserRequest.AuthUsername,
	}

	if err := s.repo.Create(user); err != nil {
		return UserRequest, err
	}

	return UserRequest, nil
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
