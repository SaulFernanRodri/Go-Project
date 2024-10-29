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

	userRequest := `{
		"name": "John Doe",
		"mil_symbol": {
			"symbolcode": "SFG-UCI---D",
			"size": 32,
			"frame": true,
			"fill": "#0000FF",
			"info_fields": {
				"uniqueDesignation": "Unit-1",
				"higherFormation": "Division-X",
				"staffComments": "No comments",
				"speed": "20km/h"
			},
			"quantity": 5,
			"direction": 90,
			"status": "active"
		}
	}`

	mockService.On("CreateUser", mock.AnythingOfType("*models.UserRequest")).Return(&models.UserRequest{
		Name: "John Doe",
		Milsymbol: models.Milsymbol{
			SymbolCode: "SFG-UCI---D",
			Size:       32,
			Frame:      true,
			Fill:       "#0000FF",
			InfoFields: models.InfoFields{
				UniqueDesignation: "Unit-1",
				HigherFormation:   "Division-X",
				StaffComments:     "No comments",
				Speed:             "20km/h",
			},
			Quantity:  5,
			Direction: 90,
			Status:    "active",
		},
		AuthUsername: "john",
	}, nil)

	router := setupRouter()
	router.POST("/users", controller.CreateUser)

	req, _ := http.NewRequest("POST", "/users", strings.NewReader(userRequest))
	req.Header.Set("X-Auth-Username", "john")
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "SFG-UCI---D")
}

func TestUpdateUser(t *testing.T) {
	mockService := mocks.NewUserServiceInterface(t)
	controller := controllers.NewUserController(mockService)

	id := 1
	mockService.On("GetByUsername", "john").Return([]models.User{{ID: uint64(id)}}, nil)
	mockService.On("UpdateUser", uint64(id), mock.AnythingOfType("*models.User")).Return(&models.User{Name: "Updated Name"}, nil)

	router := setupRouter()
	router.PUT("/users/:id", controller.UpdateUser)

	userUpdate := `{"name": "Updated Name"}`
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
