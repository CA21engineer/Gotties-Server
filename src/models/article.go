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
	UserId string `gorm:"not null"`
	Category *Category `gorm:"foreignkey:CategoryRefer"`
}


// Articles list
type Articles []Article

func NewArticle(title, before, after, body, userId string, category *Category )*Article  {
	return &Article{
		Title: title,
		Before: before,
		After: after,
		Body: body,
		Category: category,
		UserId: userId,
	}
}


func (a *Article) All() (*Articles, error){
	var articles Articles
	if err := DbConnect.Find(&articles).Error; err != nil {
		return nil, err
	}

	return &articles, nil
}

func (a *Article) Create() error{

	if err := DbConnect.Create(a).Error; err != nil {
		return err
	}

	return nil
}