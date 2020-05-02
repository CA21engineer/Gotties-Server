package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DB ...
type DB struct {
	Host     string
	Username string
	Password string
	DBName   string
	Connect  *gorm.DB
}

// NewDB new
func NewDB() *DB {
	db := &DB{
		Host:     "mysql",
		Username: "root",
		Password: "password",
		DBName:   "sample",
	}

	db.Connect = db.connect()

	return db
}

func (d *DB) connect() *gorm.DB {
	db, err := gorm.Open("mysql", d.Username+":"+d.Password+"@tcp("+d.Host+":3306)/"+d.DBName+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err.Error())
	}
	return db
}

func main() {
	NewDB()

	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "hello",
		})
	})
	r.Run(":8080")
}
