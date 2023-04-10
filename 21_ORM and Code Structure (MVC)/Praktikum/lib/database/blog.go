package database

import (
	"project_structure1/config"
	"project_structure1/models"
)

func CreateBlog(blog *models.Blog) error {
	if arr := config.DB.Create(blog).Error; arr != nil {
		return arr
	}
	return nil
}
func GetBlogs() (blogs []models.Blog, err error) {
	if err = config.DB.Find(&blogs).Error; err != nil {
		return
	}
	return
}

func GetBlog(id uint) (blog models.Blog, err error) {
	blog.ID = id
	if err = config.DB.First(&blog).Error; err != nil {
		return
	}
	return
}

func UpdateBlog(blog *models.Blog) error {
	if err := config.DB.Updates(blog).Error; err != nil {
		return err

	}
	return nil
}
func DeleteBlog(blog *models.Blog) error {
	if err := config.DB.Delete(blog).Error; err != nil {
		return err

	}
	return nil
}
