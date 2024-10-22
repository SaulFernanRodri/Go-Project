package services

import (
	"myproject/models"
	"myproject/repositories"
	"myproject/utils"
)

type UserService struct {
	repo *repositories.UserRepo
}

func NewUserService(repo *repositories.UserRepo) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	s.repo.Create(user)

	csvPath, error := utils.GenerateMilsymbol(user.Milsymbol)
	if error != nil {
		return user, error
	}

	user.CSV = csvPath
	return s.repo.Update(user.ID, user)
}

func (s *UserService) UpdateUser(id string, user *models.User) (*models.User, error) {
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id string) error {
	return s.repo.Delete(id)
}
