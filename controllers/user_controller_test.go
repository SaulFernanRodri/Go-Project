package controllers_test

import (
	"myproject/controllers"
	"myproject/mocks"
	"myproject/models"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Configurar un router Gin para pruebas
func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetAllUsers(t *testing.T) {
	mockService := mocks.NewUserServiceInterface(t)
	controller := controllers.NewUserController(mockService)

	mockUsers := []models.User{
		{ID: 1, Name: "John Doe"},
		{ID: 2, Name: "Jane Doe"},
	}
	mockService.On("GetByUsername", "john").Return(mockUsers, nil)

	router := setupRouter()
	router.GET("/users", controller.GetAllUsers)

	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("X-Auth-Username", "john")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "Jane Doe")
}

func TestCreateUser(t *testing.T) {
	mockService := mocks.NewUserServiceInterface(t)
	controller := controllers.NewUserController(mockService)

	user := `{
		"auth_username": "john",
		"name": "John Doe",
		"birth_date": "1990-01-01T00:00:00Z",
		"email": "john.doe@example.com",
		"address": "123 Main St",
		"phone": "123456789"
	}`

	mockService.On("CreateUser", mock.AnythingOfType("*models.User")).Return(&models.User{
		ID:           1,
		AuthUsername: "john",
		Name:         "John Doe",
		BirthDate:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Email:        "john.doe@example.com",
		Address:      "123 Main St",
		Phone:        "123456789",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}, nil)

	router := setupRouter()
	router.POST("/users", controller.CreateUser)

	req, _ := http.NewRequest("POST", "/users", strings.NewReader(user))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Username", "john")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "john.doe@example.com")
	assert.Contains(t, w.Body.String(), "123 Main St")
	assert.Contains(t, w.Body.String(), "123456789")
}

func TestUpdateUser(t *testing.T) {
	mockService := mocks.NewUserServiceInterface(t)
	controller := controllers.NewUserController(mockService)

	id := 1
	mockService.On("GetByUsername", "john").Return([]models.User{{ID: uint64(id)}}, nil)
	mockService.On("UpdateUser", uint64(id), mock.AnythingOfType("*models.User")).Return(&models.User{
		ID:        uint64(id),
		Name:      "Updated Name",
		BirthDate: time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Email:     "updated.email@example.com",
		Phone:     "987654321",
		UpdatedAt: time.Now(),
	}, nil)
	router := setupRouter()
	router.PUT("/users/:id", controller.UpdateUser)

	userUpdate := `{
		"name": "Updated Name",
		"birth_date": "1990-01-01T00:00:00Z",
		"email": "updated.email@example.com",
		"phone": "987654321"
	}`
	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(id), strings.NewReader(userUpdate))
	req.Header.Set("X-Auth-Username", "john")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Name")
}

func TestDeleteUser(t *testing.T) {
	mockService := mocks.NewUserServiceInterface(t)
	controller := controllers.NewUserController(mockService)

	id := 1
	mockService.On("GetByUsername", "john").Return([]models.User{{ID: uint64(id)}}, nil)
	mockService.On("DeleteUser", uint64(id)).Return(nil)

	router := setupRouter()
	router.DELETE("/users/:id", controller.DeleteUser)

	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(id), nil)
	req.Header.Set("X-Auth-Username", "john")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User deleted")
}
