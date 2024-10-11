package repositories

import (
	"myproject/models"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *UserRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *UserRepository) Update(id string, user *models.User) (*models.User, error) {
	var existingUser models.User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return nil, err
	}
	existingUser.Name = user.Name
	existingUser.Email = user.Email
	r.db.Save(&existingUser)
	return &existingUser, nil
}

func (r *UserRepository) Delete(id string) error {
	return r.db.Delete(&models.User{}, id).Error
}
