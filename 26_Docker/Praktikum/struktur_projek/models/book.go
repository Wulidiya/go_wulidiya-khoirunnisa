package models

import (
	"gorm.io/gorm"
)

type Book struct {
	gorm.Model
	Title     string `json:"title" form:"title"`
	Creator   string `json:"creator" form:"creator"`
	Publisher string `json:"publisher" form:"publisher"`
}
