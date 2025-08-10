package main

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

var DB *gorm.DB

func initDB() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=postgres dbname=users port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматическая миграция схемы
	db.AutoMigrate(&User{})

	return db, nil
}
func main() {
	var err error
	DB, err = initDB()
	if err != nil {
		log.Fatal("Не удалось подключиться к базе данных:", err)
	}

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

	// Валидация
	if user.Name == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "имя пользователя обязательно"})
		return
	}

	if user.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email обязателен"})
		return
	}

	if !strings.Contains(user.Email, "@") {
		c.JSON(http.StatusBadRequest, gin.H{"error": "неверный формат email"})
		return
	}

	if len(user.Password) < 6 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "пароль должен содержать минимум 6 символов"})
		return
	}

	// Создание пользователя в БД
	if err := DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "не удалось создать пользователя"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": user})
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"user": user})
}

func updateUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	DB.Save(&user)
	c.JSON(http.StatusOK, gin.H{"user": user})
}

func deleteUser(c *gin.Context) {
	id := c.Param("id")
	var user User

	if err := DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "пользователь не найден"})
		return
	}

	DB.Delete(&user)
	c.JSON(http.StatusOK, gin.H{"message": "пользователь удален"})
}
