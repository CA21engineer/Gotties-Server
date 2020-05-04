package models

import "github.com/jinzhu/gorm"

type Category struct {
	gorm.Model
	Name string `gorm:"unique;not null"`
}

type Categories []Category


func (a *Category) All() (*Categories, error){
	var articles Categories
	if err := DbConnect.Find(&articles).Error; err != nil {
		return nil, err
	}

	return &articles, nil
}