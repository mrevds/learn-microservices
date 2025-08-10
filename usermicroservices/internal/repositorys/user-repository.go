package repository

import (
	"microservices-learn/usermicroservices/database"
	"microservices-learn/usermicroservices/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (ur *UserRepository) CreateUser(user *models.User) error {
	result := database.DB.Create(user)
	return result.Error
}

func (ur *UserRepository) GetUserByID(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("email = ?", email).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}

func (ur *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func (ur *UserRepository) UpdateUser(user *models.User) error {
	result := database.DB.Save(user)
	return result.Error
}

func (ur *UserRepository) DeleteUser(id uint) error {
	result := database.DB.Delete(&models.User{}, id)
	return result.Error
}

func (ur *UserRepository) UserExistsByEmail(email string) (bool, error) {
	var count int64
	result := database.DB.Model(&models.User{}).Where("email = ?", email).Count(&count)
	return count > 0, result.Error
}
