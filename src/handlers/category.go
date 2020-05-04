package handlers

import (
	"Gotties-Server/src/models"
	"Gotties-Server/src/responses"
	"github.com/gin-gonic/gin"
)

type Category struct{
	Name string
}

func NewCategory()*Category {
	return &Category{}
}

func (a *Category) GetCategories(c *gin.Context) {
	categories, err := new(models.Category).All()

	if err != nil {
		responses.HTTPResponseInternalServerError(c, err.Error())
	}


	c.JSON(200, gin.H{
		"categorie": categories,
	})
}


