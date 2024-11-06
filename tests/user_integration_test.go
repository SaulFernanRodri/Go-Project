package controllers_test

import (
	"myproject/controllers"
	"myproject/models"
	"myproject/repositories"
	"myproject/services"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

var db *gorm.DB
var userRepo *repositories.UserRepo
var userService *services.UserService
var userController *controllers.UserController

func setupTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	db.AutoMigrate(&models.User{})
	return db
}

func setupRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestCreateUserIntegration(t *testing.T) {
	db = setupTestDB()
	userRepo = repositories.NewUserRepo(db)
	userService = services.NewUserService(userRepo)
	userController = controllers.NewUserController(userService)

	router := setupRouter()
	router.POST("/users", userController.CreateUser)

	userRequest := `{
		"auth_username": "testuser",
		"name": "John Doe",
		"birth_date": "1990-01-01T00:00:00Z",
		"email": "john.doe@example.com",
		"address": "123 Main St",
		"phone": "123456789"
	}`

	req, _ := http.NewRequest("POST", "/users", strings.NewReader(userRequest))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Username", "testuser")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")

	var user models.User
	result := db.First(&user)
	assert.NoError(t, result.Error)
	assert.Equal(t, "John Doe", user.Name)
	assert.Equal(t, "testuser", user.AuthUsername)
	assert.Equal(t, "john.doe@example.com", user.Email)
	assert.Equal(t, "123456789", user.Phone)
}

func TestGetAllUsersIntegration(t *testing.T) {
	db = setupTestDB()
	userRepo = repositories.NewUserRepo(db)
	userService = services.NewUserService(userRepo)
	userController = controllers.NewUserController(userService)

	router := setupRouter()
	router.GET("/users", userController.GetAllUsers)

	db.Create(&models.User{Name: "John Doe", AuthUsername: "testuser", Email: "aaa@email.com", Phone: "123456789"})
	db.Create(&models.User{Name: "Jane Doe 2", AuthUsername: "testuser", Email: "a@email.com", Phone: "123456789"})

	req, _ := http.NewRequest("GET", "/users", nil)
	req.Header.Set("X-Auth-Username", "testuser")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "John Doe")
	assert.Contains(t, w.Body.String(), "Jane Doe 2")
}

func TestUpdateUserIntegration(t *testing.T) {
	db = setupTestDB()
	userRepo = repositories.NewUserRepo(db)
	userService = services.NewUserService(userRepo)
	userController = controllers.NewUserController(userService)

	router := setupRouter()
	router.PUT("/users/:id", userController.UpdateUser)

	user := models.User{Name: "John Doe", AuthUsername: "testuser"}
	db.Create(&user)

	userUpdate := `{
		"name": "Updated Name",
		"birth_date": "1990-01-01T00:00:00Z",
		"email": "aaa@gmail.com",
		"address": "123 Main St",
		"phone": "123456789"
	}`
	req, _ := http.NewRequest("PUT", "/users/"+strconv.Itoa(int(user.ID)), strings.NewReader(userUpdate))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-Username", "testuser")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "Updated Name")

	var updatedUser models.User
	db.First(&updatedUser, user.ID)
	assert.Equal(t, "Updated Name", updatedUser.Name)
}

func TestDeleteUserIntegration(t *testing.T) {
	db = setupTestDB()
	userRepo = repositories.NewUserRepo(db)
	userService = services.NewUserService(userRepo)
	userController = controllers.NewUserController(userService)

	router := setupRouter()
	router.DELETE("/users/:id", userController.DeleteUser)

	user := models.User{Name: "John Doe", AuthUsername: "testuser"}
	db.Create(&user)

	req, _ := http.NewRequest("DELETE", "/users/"+strconv.Itoa(int(user.ID)), nil)
	req.Header.Set("X-Auth-Username", "testuser")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Contains(t, w.Body.String(), "User deleted")

	var deletedUser models.User
	result := db.First(&deletedUser, user.ID)
	assert.Error(t, result.Error)
	assert.Equal(t, gorm.ErrRecordNotFound, result.Error)
}
