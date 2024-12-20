// Code generated by mockery v2.46.3. DO NOT EDIT.

package mocks

import (
	models "myproject/models"

	mock "github.com/stretchr/testify/mock"
)

// UserRepoInterface is an autogenerated mock type for the UserRepoInterface type
type UserRepoInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UserRepoInterface) Create(user *models.User) (*models.User, error){
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*models.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return user, r0
}

// Delete provides a mock function with given fields: id
func (_m *UserRepoInterface) Delete(id uint64) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAll provides a mock function with given fields:
func (_m *UserRepoInterface) GetAll() ([]models.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAll")
	}

	var r0 []models.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]models.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []models.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByUsername provides a mock function with given fields: username
func (_m *UserRepoInterface) GetByUsername(username string) ([]models.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetByUsername")
	}

	var r0 []models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) ([]models.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) []models.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: id, user
func (_m *UserRepoInterface) Update(id uint64, user *models.User) (*models.User, error) {
	ret := _m.Called(id, user)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *models.User
	var r1 error
	if rf, ok := ret.Get(0).(func(uint64, *models.User) (*models.User, error)); ok {
		return rf(id, user)
	}
	if rf, ok := ret.Get(0).(func(uint64, *models.User) *models.User); ok {
		r0 = rf(id, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*models.User)
		}
	}

	if rf, ok := ret.Get(1).(func(uint64, *models.User) error); ok {
		r1 = rf(id, user)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewUserRepoInterface creates a new instance of UserRepoInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepoInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepoInterface {
	mock := &UserRepoInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
