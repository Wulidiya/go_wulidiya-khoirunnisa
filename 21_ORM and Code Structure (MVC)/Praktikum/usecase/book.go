package usecase

import (
	"errors"
	"project_structure1/lib/database"
	"project_structure1/models"
)

func CreateBook(book *models.Book) error {
	if book.Title == "" {
		return errors.New("book title cannot be empty")
	}
	if book.Creator == "" {
		return errors.New("book creator cannot be empty")
	}
	err := database.CreateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func GetBook(id uint) (book models.Book, err error) {
	book, err = database.GetBook(id)
	return
}

func GetListBooks() (books []models.Book, err error) {
	books, err = database.GetBooks()
	return
}

func UpdateBook(book *models.Book) error {
	err := database.UpdateBook(book)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBook(id uint) error {
	book := models.Book{}
	book.ID = id
	err := database.DeleteBook(&book)
	if err != nil {
		return err
	}
	return nil
}
