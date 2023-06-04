package database

import (
	"struktur_projek/config"
	"struktur_projek/models"
)

func CreateUser(user *models.User) error {
	if arr := config.DB.Create(user).Error; arr != nil {
		return arr

	}
	return nil
}
func GetUsers() (users []models.User, err error) {
	if err = config.DB.Model(&models.User{}).Preload("Blogs").Find(&users).Error; err != nil {
		return
	}
	return
}

func GetUser(id uint) (user models.User, err error) {
	user.ID = id
	if err = config.DB.Model(&models.User{}).Preload("Blogs").First(&user).Error; err != nil {
		return
	}
	return
}

func UpdateUser(user *models.User) error {
	if err := config.DB.Updates(user).Error; err != nil {
		return err

	}
	return nil
}
func DeleteUser(user *models.User) error {
	if err := config.DB.Delete(user).Error; err != nil {
		return err

	}
	return nil
}
