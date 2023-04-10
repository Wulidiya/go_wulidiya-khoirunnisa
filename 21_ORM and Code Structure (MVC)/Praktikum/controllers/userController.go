package controllers

import (
	"net/http"
	"strconv"

	"project_structure1/config"
	"project_structure1/lib/database"
	"project_structure1/models"
	"project_structure1/usecase"

	"github.com/labstack/echo"
)

// get all users
func GetUsersController(c echo.Context) error {
	var users []models.User

	if err := config.DB.Find(&users).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all users",
		"users":   users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user, err := database.GetUser(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error get user",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  user,
	})
}

//create new user

func CreateUserController(c echo.Context) error {
	user := models.User{}
	c.Bind(&user)

	if err := usecase.CreateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error create user",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new user",
		"users":    user,
	})
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteUser(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete user",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succses delete user",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	user := models.User{}
	c.Bind(&user)
	user.ID = uint(id)
	if err := usecase.UpdateUser(&user); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error update user",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succsess update user",
	})
}
