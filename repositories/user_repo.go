package repositories

import (
	"myproject/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepo) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepo) Update(id string, user *models.User) (*models.User, error) {
	var existingUser models.User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.Name = user.Name
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	existingUser.Milsymbol = user.Milsymbol
	existingUser.CSV = user.CSV

	if err := r.db.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (r *UserRepo) Delete(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}
