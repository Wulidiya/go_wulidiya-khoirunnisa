package model

import "github.com/google/uuid"

type Item struct {
	ID          uuid.UUID `gorm:"primarykey" json:"id" `
	UserID      uint      `json:"userID" form:"UserID"`
	Name        string    `json:"name" gorm:"unique"`
	Description string    `json:"description" form:"Description"`
	Stock       uint      `json:"stock" form:"Stock"`
	Price       uint      `json:"price" form:"Price"`
	CategoryID  uint      `json:"category_id" form:"Category_id"`
	Category    Category  `json:"-"`
}
