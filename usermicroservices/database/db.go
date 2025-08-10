package database

import (
	"microservices-learn/usermicroservices/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() (*gorm.DB, error) {
	dsn := "host=postgres user=postgres password=postgres dbname=users port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Автоматическая миграция схемы
	db.AutoMigrate(&models.User{})

	return db, nil
}
