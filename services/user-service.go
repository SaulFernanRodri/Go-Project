package services

import (
	"errors"
	"time"

	"myproject/models"
	"myproject/repositories"
	"myproject/utils"

	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("my_secret_key")

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

type UserServiceInterface interface {
	GetAllUsers() ([]models.User, error)
	CreateUser(user *models.User) (*models.User, error)
	UpdateUser(id uint64, user *models.User) (*models.User, error)
	DeleteUser(id uint64) error
	Authenticate(email, password string) (string, error)
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

func (s *UserService) CreateUser(user *models.User) (*models.User, error) {
	s.repo.Create(user)

	csvPath, error := utils.GenerateMilsymbol(user.Milsymbol)
	if error != nil {
		return user, error
	}

	user.CSV = csvPath
	return s.repo.Update(user.ID, user)
}

func (s *UserService) UpdateUser(id uint64, user *models.User) (*models.User, error) {
	return s.repo.Update(id, user)
}

func (s *UserService) DeleteUser(id uint64) error {
	return s.repo.Delete(id)
}

func (s *UserService) Authenticate(email, password string) (string, error) {
	var user models.User

	// Buscar el usuario por email
	if err := s.repo.FindByEmail(email, &user); err != nil {
		return "", errors.New("user not found")
	}

	// Verificar la contraseña (aquí puedes agregar el uso de bcrypt)
	if user.Password != password {
		return "", errors.New("invalid credentials")
	}

	// Generar un token JWT
	expirationTime := time.Now().Add(15 * time.Minute)
	claims := &Claims{
		Username: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
