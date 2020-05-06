package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func init()  {
	//envファイルを読み込み
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	fmt.Println("=========conf=====")
	os.Getenv("DB_PASS")
	fmt.Println("=========conf=====")
}