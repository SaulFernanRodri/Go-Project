package services_test

import (
	"myproject/mocks"
	"myproject/models"
	"myproject/services"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllUsers(t *testing.T) {
	mockRepo := mocks.NewUserRepoInterface(t)
	service := services.NewUserService(mockRepo)

	expectedUsers := []models.User{
		{Name: "John"},
		{Name: "Jane"},
	}

	mockRepo.On("GetAll").Return(expectedUsers, nil)

	users, err := service.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}

func TestCreateUser(t *testing.T) {
	mockRepo := mocks.NewUserRepoInterface(t)
	service := services.NewUserService(mockRepo)

	userRequest := &models.User{
		AuthUsername: "john",
		Name:         "John Doe",
		BirthDate:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Email:        "john.doe@example.com",
		Address:      "123 Main St",
		Phone:        "123456789",
	}

	expectedUser := &models.User{
		AuthUsername: "john",
		Name:         "John Doe",
		BirthDate:    time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC),
		Email:        "john.doe@example.com",
		Address:      "123 Main St",
		Phone:        "123456789",
	}

	// Mock the Create method on the repository to return nil error
	mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	// Call the CreateUser method on the service with the user request
	createdUser, err := service.CreateUser(userRequest)

	// Assertions to verify no errors and the created user matches the expected user
	assert.NoError(t, err)
	assert.Equal(t, expectedUser.AuthUsername, createdUser.AuthUsername)
	assert.Equal(t, expectedUser.Name, createdUser.Name)
	assert.Equal(t, expectedUser.BirthDate, createdUser.BirthDate)
	assert.Equal(t, expectedUser.Email, createdUser.Email)
	assert.Equal(t, expectedUser.Address, createdUser.Address)
	assert.Equal(t, expectedUser.Phone, createdUser.Phone)
}

func TestUpdateUser(t *testing.T) {
	mockRepo := mocks.NewUserRepoInterface(t)
	service := services.NewUserService(mockRepo)

	userToUpdate := &models.User{
		Name: "UpdatedName",
	}

	mockRepo.On("Update", uint64(1), userToUpdate).Return(userToUpdate, nil)

	updatedUser, err := service.UpdateUser(1, userToUpdate)

	assert.NoError(t, err)
	assert.Equal(t, userToUpdate, updatedUser)
}

func TestDeleteUser(t *testing.T) {
	mockRepo := mocks.NewUserRepoInterface(t)
	service := services.NewUserService(mockRepo)

	mockRepo.On("Delete", uint64(1)).Return(nil)

	err := service.DeleteUser(1)

	assert.NoError(t, err)
}

func TestGetByUsername(t *testing.T) {
	mockRepo := mocks.NewUserRepoInterface(t)
	service := services.NewUserService(mockRepo)

	expectedUsers := []models.User{
		{Name: "John"},
	}

	mockRepo.On("GetByUsername", "John").Return(expectedUsers, nil)

	users, err := service.GetByUsername("John")

	assert.NoError(t, err)
	assert.Equal(t, expectedUsers, users)
}
