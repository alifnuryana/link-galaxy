package services

import (
	"github.com/alifnuryana/link-galaxy/database"
	"github.com/alifnuryana/link-galaxy/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User

	tx := database.DB.Find(&users)
	if tx.Error != nil {
		return []models.User{}, tx.Error
	}

	return users, nil
}

func GetUser(id uint) (models.User, error) {
	var user models.User
	tx := database.DB.Where("id = ?", id).First(&user)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	return user, nil
}

func AddUser(userInput models.User) (models.User, error) {
	tx := database.DB.Create(&userInput)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	user, err := GetUser(userInput.ID)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func UpdateUser(userInput models.User, id uint) (models.User, error) {
	tx := database.DB.Where("id = ?", id).Updates(userInput)
	if tx.Error != nil {
		return models.User{}, tx.Error
	}
	user, err := GetUser(id)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}

func DeleteUser(id uint) error {
	tx := database.DB.Delete(&models.User{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
