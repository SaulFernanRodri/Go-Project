package repositories

import (
	"myproject/models"

	"gorm.io/gorm"
)

type UserRepoInterface interface {
	GetAll() ([]models.User, error)
	Create(user *models.User) (*models.User, error)
	Update(id uint64, user *models.User) (*models.User, error)
	Delete(id uint64) error
	GetByUsername(username string) ([]models.User, error)
}

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

func (r *UserRepo) Create(user *models.User) (*models.User, error) {
	err := r.db.Create(user).Error
	return user, err
}

func (r *UserRepo) Update(id uint64, user *models.User) (*models.User, error) {
	var existingUser models.User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return nil, err
	}

	existingUser.Name = user.Name

	if err := r.db.Save(&existingUser).Error; err != nil {
		return nil, err
	}

	return &existingUser, nil
}

func (r *UserRepo) Delete(id uint64) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *UserRepo) GetByUsername(username string) ([]models.User, error) {
	var users []models.User
	err := r.db.Where("auth_username = ?", username).Find(&users).Error
	return users, err
}
