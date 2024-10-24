package controllers

import (
	"myproject/models"
	"myproject/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserControllerInterface interface {
	ShowAllUsers(c *gin.Context)
	GetAllUsers(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
}

type UserController struct {
	service *services.UserService
}

func NewUserController(service *services.UserService) *UserController {
	return &UserController{service: service}
}

func (ctrl *UserController) ShowAllUsers(c *gin.Context) {
	users, err := ctrl.service.GetAllUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.HTML(http.StatusOK, "users.html", gin.H{"users": users})
}

func (ctrl *UserController) GetAllUsers(c *gin.Context) {
	username := c.GetHeader("X-Auth-Username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Auth-Username header is required"})
		return
	}

	users, err := ctrl.service.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}

func (ctrl *UserController) CreateUser(c *gin.Context) {
	username := c.GetHeader("X-Auth-Username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Auth-Username header is required"})
		return
	}

	var user models.UserRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.AuthUsername = username

	createdUser, err := ctrl.service.CreateUser(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdUser)
}

func (ctrl *UserController) UpdateUser(c *gin.Context) {
	username := c.GetHeader("X-Auth-Username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Auth-Username header is required"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Verificar que el usuario pertenece al username
	existingUsers, err := ctrl.service.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	found := false
	for _, u := range existingUsers {
		if u.ID == id {
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to update this user"})
		return
	}

	updatedUser, err := ctrl.service.UpdateUser(id, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedUser)
}

func (ctrl *UserController) DeleteUser(c *gin.Context) {
	username := c.GetHeader("X-Auth-Username")
	if username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "X-Auth-Username header is required"})
		return
	}

	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	// Verificar que el usuario pertenece al username
	existingUsers, err := ctrl.service.GetByUsername(username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	found := false
	for _, u := range existingUsers {
		if u.ID == id {
			found = true
			break
		}
	}
	if !found {
		c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to delete this user"})
		return
	}

	if err := ctrl.service.DeleteUser(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}
