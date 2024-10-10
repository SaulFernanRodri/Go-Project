package domain

import (
	"errors"
	"time"
)

// User representa la entidad principal de un usuario.
type User struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// NewUser es una función para crear un nuevo usuario aplicando las reglas de negocio.
func NewUser(name, email, password string) (*User, error) {
	if len(password) < 6 {
		return nil, errors.New("password must be at least 6 characters")
	}
	return &User{
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}, nil
}
