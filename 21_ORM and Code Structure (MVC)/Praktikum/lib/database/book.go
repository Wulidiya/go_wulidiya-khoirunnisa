package database

import (
	"project_structure1/config"
	"project_structure1/models"
)

func CreateBook(book *models.Book) error {
	if arr := config.DB.Create(book).Error; arr != nil {
		return arr
	}
	return nil
}
func GetBooks() (books []models.Book, err error) {
	if err = config.DB.Find(&books).Error; err != nil {
		return
	}
	return
}

func GetBook(id uint) (book models.Book, err error) {
	book.ID = id
	if err = config.DB.First(&book).Error; err != nil {
		return
	}
	return
}

func UpdateBook(book *models.Book) error {
	if err := config.DB.Updates(book).Error; err != nil {
		return err
	}
	return nil
}
func DeleteBook(book *models.Book) error {
	if err := config.DB.Delete(book).Error; err != nil {
		return err
	}
	return nil
}
