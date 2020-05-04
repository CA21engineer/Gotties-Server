package handlers

import (
	"Gotties-Server/src/form"
	"Gotties-Server/src/lib/firebase"
	"Gotties-Server/src/models"
	"Gotties-Server/src/responses"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type Article struct{}

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
	articles, err := new(models.Article).All()

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
	f := form.NewArticle(
		c.PostForm("title"),
		c.PostForm("before"),
		c.PostForm("after"),
		c.PostForm("body"),
		c.PostForm("category_name"),
		c.PostForm("user_id"),
	)

	//バリデーション検証
	if errors := validator.New().Struct(f); errors != nil {
		responses.HTTPResponseInternalServerError(c, errors.Error())
	}

	//useridの検証
	userId, err := firebase.NewAuth(c.PostForm("userId")).IsUserLogedIn()
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
	}

	//imageのアップロード
	beforeImg, err := firebase.NewImage(c.PostForm("before")).UploadImage()
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
	}

	afterImg, err := firebase.NewImage(c.PostForm("after")).UploadImage()
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
	}

	//categoryを検索し、もしcategoryがなかったら作成
	category, err := models.NewCategory(c.PostForm("category_name")).FindByNameORCreate()
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
	}

	article := models.NewArticle(
		f.Title,
		beforeImg,
		afterImg,
		f.Body,
		userId,
		category,
	)

	if err = article.Create(); err!= nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
	}


	c.JSON(204, gin.H{
		"article": article,
	})
}