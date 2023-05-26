package usecase

import (
	"errors"
	"fmt"
	"remed/model"
	"remed/repository/database"
)

func CreateItem(item *model.Item) error {

	// check tool_id
	if item.Name == "" {
		return errors.New("item tool_id cannot be empty")
	}

	err := database.CreateItem(item)
	if err != nil {
		return err
	}

	return nil
}

func GetItem(id uint) (item model.Item, err error) {
	item, err = database.GetItem(id)
	if err != nil {
		fmt.Println("GetItem: Error getting item from database")
		return
	}
	return
}

func GetListItems() (items []model.Item, err error) {
	items, err = database.GetItems()
	if err != nil {
		fmt.Println("GetListItems: Error getting items from database")
		return
	}
	return
}

func UpdateItem(item *model.Item) (err error) {
	err = database.UpdateItem(item)
	if err != nil {
		fmt.Println("UpdateItem : Error updating item, err: ", err)
		return
	}

	return
}

func DeleteItem(id uint) (err error) {
	item := model.Item{}
	item.ID = id
	err = database.DeleteItem(&item)
	if err != nil {
		fmt.Println("DeleteItem : error deleting item, err: ", err)
		return
	}

	return
}
