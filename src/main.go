package main

import (
	"Gotties-Server/src/handlers"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
	"os"
	//"github.com/aws/aws-sdk-go/service/s3"
)



func main() {
	//envファイルを読み込み
	err := godotenv.Load(fmt.Sprintf("../%s.env", os.Getenv("GO_ENV")))
	if err != nil {
		panic(err)
	}

	r := gin.Default()

	article := handlers.NewArticle()
	category := handlers.NewCategory()



	r.GET("/articles", article.GetArticles)
	r.GET("/articles/:id", article.GetArticle)
	r.POST("/articles", article.CreateArticle)

	r.GET("/categories", category.GetCategories)


	r.Run(":8080")
}
