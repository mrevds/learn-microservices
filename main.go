package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID       string `json:"id"`
	name     string `json:"name"`
	email    string `json:"email"`
	password string `json:"password"`
}

func main() {
	r := gin.Default()

	r.POST("/users", createUser)
	r.GET("/users/:id", getUser)
	r.PUT("/users/:id", updateUser)
	r.DELETE("/users/:id", deleteUser)

	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Валидация полей
	if user.name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "имя пользователя обязательно"})
		return
	}

	if user.email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email обязателен"})
		return
	}

	// Простая проверка формата email
	if !strings.Contains(user.email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат email"})
		return
	}

	if len(user.password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "пароль должен содержать минимум 6 символов"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "name": "John Doe"})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "User updated"})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{"id": id, "message": "User deleted"})
}
