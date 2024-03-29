package controller

import (
	"net/http"
	"remed/model"
	"remed/usecase"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetItemsController(c echo.Context) error {
	items, e := usecase.GetListItems()

	if e != nil {
		return echo.NewHTTPError(http.StatusBadRequest, e.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"items":  items,
	})
}

func GetItemController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	item, err := usecase.GetItem(uint(id))

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error get item",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"items":  item,
	})
}

// create new item
func CreateItemController(c echo.Context) error {
	item := model.Item{}
	c.Bind(&item)

	if err := usecase.CreateItem(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error create item",
			"errorDescription": err,
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new item",
		"item":    item,
	})
}

// delete item by id
func DeleteItemController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	if err := usecase.DeleteItem(uint(id)); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error delete item",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf alat tidak dapat di hapus",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success delete item",
	})
}

// update item by id
func UpdateItemController(c echo.Context) error {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 32)
	item := model.Item{}
	c.Bind(&item)
	item.ID = uint(id)
	if err := usecase.UpdateItem(&item); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages":         "error update item",
			"errorDescription": err,
			"errorMessage":     "Mohon Maaf item tidak dapat di ubah",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success update item",
	})
}
