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

// get all blogs
func GetBlogsController(c echo.Context) error {
	var blogs []models.Blog

	if err := config.DB.Find(&blogs).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all blogs",
		"users":   blogs,
	})
}

// get blog by id
func GetBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	blog, err := database.GetBlog(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error get blog",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"blogs":  blog,
	})
}

//create new blog

func CreateBlogController(c echo.Context) error {
	blog := models.Blog{}
	c.Bind(&blog)

	if err := usecase.CreateBlog(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         " error create blog",
			"errorDescription": err,
		})

	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create new blog",
		"users":    blog,
	})
}

// delete blog by id
func DeleteBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteBlog(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete blog",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succses delete blog",
	})
}

// update book by id
func UpdateBlogController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	blog := models.Blog{}
	c.Bind(&blog)
	blog.ID = uint(id)
	if err := usecase.UpdateBlog(&blog); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message":          "error update blog",
			"errorDescription": err,
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "succsess update blog",
	})
}
