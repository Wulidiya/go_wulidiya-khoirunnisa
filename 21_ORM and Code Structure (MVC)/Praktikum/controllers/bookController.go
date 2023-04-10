package controllers

import (
	"net/http"
	"strconv"

	"project_structure1/usecase"

	"project_structure1/config"
	"project_structure1/lib/database"
	"project_structure1/models"

	"github.com/labstack/echo"
)

// get all books
func GetBooksController(c echo.Context) error {
	var books []models.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"users":   books,
	})
}

// get book by id
func GetBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book, err := database.GetBook(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error get book",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  book,
	})
}

//create new book

func CreateBookController(c echo.Context) error {
	book := models.Book{}
	c.Bind(&book)

	if err := usecase.CreateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error create book",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new book",
		"users":    book,
	})
}

// delete book by id
func DeleteBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteBook(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete book",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succses delete book",
	})
}

// update book by id
func UpdateBookController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	book := models.Book{}
	c.Bind(&book)
	book.ID = uint(id)
	if err := usecase.UpdateBook(&book); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error update book",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succsess update book",
	})
}
