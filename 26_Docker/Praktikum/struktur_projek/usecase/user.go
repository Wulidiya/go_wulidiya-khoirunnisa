package usecase

import (
	"errors"
	"struktur_projek/lib/database"
	"struktur_projek/models"
)

func CreateUser(user *models.User) error {
	if user.Name == "" {
		return errors.New("user name cannot be empty")
	}
	if user.Email == "" {
		return errors.New("user email cannot be empty")
	}
	err := database.CreateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func GetUser(id uint) (user models.User, err error) {
	user, err = database.GetUser(id)
	return
}

func GetListUsers() (users []models.User, err error) {
	users, err = database.GetUsers()
	return
}

func UpdateUser(user *models.User) error {
	err := database.UpdateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func DeleteUser(id uint) error {
	user := models.User{}
	user.ID = id
	err := database.DeleteUser(&user)
	if err != nil {
		return err
	}
	return nil
}
