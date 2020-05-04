package models

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	gorm.Model
	Title string `gorm:"not null"`
	Before string `gorm:"not null"`
	After string `gorm:"not null"`
	Body string `gorm:"not null"`
	Category *Category `gorm:"-"`
	UserId string `gorm:"not null"`
}


// Articles list
type Articles []Article

func NewArticle()*Article  {
	return &Article{}
}


func (a *Article) All() (*Articles, error){
	var articles Articles
	if err := DbConnect.Find(&articles).Error; err != nil {
		return nil, err
	}

	return &articles, nil
}