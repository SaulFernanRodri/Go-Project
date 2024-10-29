package services_test

import (
	"myproject/mocks"
	"myproject/models"
	"myproject/services"
	"testing"

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

	userRequest := &models.UserRequest{
		Name: "John",
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
	}

	expectedUser := &models.UserRequest{
		Name: "John",
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
	}
	mockRepo.On("Create", mock.AnythingOfType("*models.User")).Return(nil)

	createdUser, err := service.CreateUser(userRequest)

	assert.NoError(t, err)
	assert.Equal(t, expectedUser, createdUser)
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
