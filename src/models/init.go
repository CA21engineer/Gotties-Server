package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)


var DbConnect *gorm.DB


func init() {
	var err error
	//db, err := gorm.Open("mysql", "root:password@tcp(mysql:3306)/sample?charset=utf8&parseTime=True&loc=Local")
	DbConnect, err = gorm.Open("mysql", "root:@/gotties_server?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}


	//logの有効化
	DbConnect.LogMode(true)

	//
	DbConnect.AutoMigrate(&Article{}, &Category{})
}