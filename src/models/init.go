package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var DbConnect *gorm.DB


func init() {
	var err error
	//connectPath := os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+ "@tcp("+ os.Getenv("HOST") + ":3306)/" + os.Getenv("DB") + "?charset=utf8&parseTime=True&loc=Local"
	//DbConnect, err := gorm.Open("mysql", connectPath)
	DbConnect, err = gorm.Open("mysql", "root:password@tcp(db:3306)/gotties?charset=utf8&parseTime=True&loc=Local")

	//DbConnect, err = gorm.Open("mysql", "root:@/gotest?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	//logの有効化
	DbConnect.LogMode(true)

	DbConnect.AutoMigrate(&Article{}, &Category{})


	category, err := NewCategory("category").FindByNameORCreate()
	article := NewArticle("title1", "https://dummyimage.com/600x400/000/fff&text=before1", "https://dummyimage.com/600x400/000/fff&text=after1", "texttexttexttexttexttexttexttexttexttext", "uuid", category)
	DbConnect.Create(&article)
	article = NewArticle("title2", "https://dummyimage.com/600x400/000/fff&text=before2", "https://dummyimage.com/600x400/000/fff&text=after2", "texttexttexttexttexttexttexttexttexttext", "uuid", category)
	DbConnect.Create(&article)
}
