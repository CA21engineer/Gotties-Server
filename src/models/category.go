package models

import (
	"github.com/jinzhu/gorm"
)

type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
	Articles []Article `gorm:"foreignkey:CategoryId"`

}

type Categories []Category

func NewCategory(name string) *Category {
	return &Category{Name: name}
}

func (c *Category) All()(*Categories, error){
	var categories Categories
	if err := DbConnect.Find(&categories).Error; err != nil {
		return nil, err
	}

	return &categories, nil
}

func (c *Category) Find(id int)(*Category, error){
	var category Category
	if err := DbConnect.Where("id = ?", id).First(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c *Category) FindByNameORCreate()(*Category, error){
	var category Category
	if err := DbConnect.FirstOrCreate(&category, Category{Name: c.Name}).Error; err != nil {
			return nil, err
	}
	if err := DbConnect.Where("name = ?", c.Name).First(&category).Error; err != nil {
		return nil, err
	}


	return &category, nil
}

func (c *Category) Create()error{
	if err := DbConnect.Create(c).Error; err != nil {
		return err
	}

	return nil
}