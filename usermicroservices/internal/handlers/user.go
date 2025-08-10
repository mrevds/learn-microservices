package handlers

import (
	"microservices-learn/usermicroservices/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func createUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username is empty"})
		return
	}
	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email is empty"})
		return
	}
	if len(user.Password) < 8 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be at least 8 characters"})
		return
	}
	if !strings.Contains(user.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email must contain @"})
		return
	}

	// Создаем экземпляр репозитория
	userRepo := repository.NewUserRepository()

	// Проверяем, существует ли пользователь с таким email
	exists, err := userRepo.UserExistsByEmail(user.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}
	if exists {
		c.JSON(http.StatusConflict, gin.H{"error": "User with this email already exists"})
		return
	}

	// Создаем пользователя
	if err := userRepo.CreateUser(&user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully", "user": user})
}
