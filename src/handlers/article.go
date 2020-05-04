package handlers

import (
	"Gotties-Server/src/models"
	"Gotties-Server/src/responses"
	"github.com/gin-gonic/gin"
)

type Article struct{
	Title string
	Before string
	After string
	Body string
	//Category *handler.Category
	UserId string
}

type Articles []Article

type ArticleImpl interface {
	GetArticles(c *gin.Context)
	GetArticle(c *gin.Context)
	CreateArticle(c *gin.Context)
}

func NewArticle() ArticleImpl{
	return &Article{}
}


func (a *Article) GetArticles(c *gin.Context) {
	articles, err := models.NewArticle().All()

	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
	}


	c.JSON(200, gin.H{
		"article": articles,
	})
}


func (a *Article) GetArticle(c *gin.Context) {
	var article models.Article
	id := c.Param("id")
	if err := models.DbConnect.Where("id = ?", id).First(&article).Error; err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
	}

	c.JSON(200, gin.H{
		"article": article,
	})
}


func (a *Article) CreateArticle(c *gin.Context) {
	article := &Article{}
	models.DbConnect.NewRecord(article) // => 主キーが空の場合に `true` を返します。

	models.DbConnect.Create(&article)

	models.DbConnect.NewRecord(article) // => `user` が作られた後に `false` を返します。


	c.JSON(204, gin.H{
		"article": article,
	})
}