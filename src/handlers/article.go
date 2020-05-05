package handlers

import (
	"Gotties-Server/src/form"
	"Gotties-Server/src/lib/firebase"
	"Gotties-Server/src/models"
	"Gotties-Server/src/responses"
	"fmt"
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
		return
	}
	var response responses.Articles

	for _, article := range *articles {
		responseArticle := new(responses.Article).ResponseArticle(&article)
		response = append(response, *responseArticle)
	}


	c.JSON(200, gin.H{
		"article":  response,
	})
}


func (a *Article) GetArticle(c *gin.Context) {
	id := c.Param("id")

	article, err := new(models.Article).Find(id)
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
		return
	}



	c.JSON(200, gin.H{
		"article":  new(responses.Article).ResponseArticle(article),
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
		responses.HTTPResponseBadRequest(c)
		return
	}

	//user_idの検証
	userId, err := firebase.NewAuth(c.PostForm("user_id")).IsUserLogedIn()
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
		return
	}

	//imageのアップロード
	beforeImg, err := firebase.NewImage(c.PostForm("before")).UploadImage()
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
		return
	}

	afterImg, err := firebase.NewImage(c.PostForm("after")).UploadImage()
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
		return
	}

	//categoryを検索し、もしcategoryがなかったら作成
	category, err := models.NewCategory(c.PostForm("category_name")).FindByNameORCreate()
	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
		return
	}

	article := models.NewArticle(
		f.Title,
		beforeImg,
		afterImg,
		f.Body,
		userId,
		category,
	)

	fmt.Println(article)

	if err = article.Create(); err!= nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message":  "success",

	})
}