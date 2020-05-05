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
	Category Category `gorm:"-"`
	CategoryId uint
}


// Articles list
type Articles []Article

func NewArticle(title, before, after, body, userId string, category *Category )*Article  {
	return &Article{
		Title: title,
		Before: before,
		After: after,
		Body: body,
		CategoryId: category.ID,
		UserId: userId,
	}
}


func (a *Article) All() (*Articles, error){
	var articles Articles
	if err := DbConnect.Find(&articles).Error; err != nil {
		return nil, err
	}

	for i, s := range articles {
		var category Category
		DbConnect.Where("id = ?", s.CategoryId).Find(&category)
		articles[i].Category = category
	}

	return &articles, nil
}

func (a *Article) Search(keyword string) (*Articles, error){
	var categories Categories
	if err := DbConnect.Where("name LIKE ?", "%"+keyword+"%").Find(&categories).Error; err != nil {
		return nil, err
	}
	var categoryIds []uint
	for _, v := range categories{
		categoryIds = append(categoryIds, v.ID)
	}

	var articles Articles
	if err := DbConnect.Where("category_id IN (?)", categoryIds).Find(&articles).Error; err != nil {
		return nil, err
	}

	for i, s := range articles {
		var category Category
		DbConnect.Where("id = ?", s.CategoryId).Find(&category)
		articles[i].Category = category
	}

	return &articles, nil
}

func (a *Article) Find(id string)(*Article, error){
	var article Article
	if err := DbConnect.Where("id = ?", id).First(&article).Error; err != nil {
		return nil, err
	}

	var category Category
	DbConnect.Where("id = ?", article.CategoryId).Find(&category)
	article.Category = category

	return &article, nil
}



func (a *Article) Create() error{

	if err := DbConnect.Create(a).Error; err != nil {
		return err
	}

	return nil
}