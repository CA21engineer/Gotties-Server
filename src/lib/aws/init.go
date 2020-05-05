package aws

import (
    "os"
)

type Config struct {
   Aws struct {
       S3 struct {
           Region string
           Bucket string
           AccessKeyID string
           SecretAccessKey string
       }
   }
}

func NewConfig() *Config {

   c := new(Config)

   // ex) アジアパシフィック (東京): ap-northeast-1
   c.Aws.S3.Region = "ap-northeast-1"
   c.Aws.S3.Bucket = "gotties"
   c.Aws.S3.AccessKeyID = os.Getenv("AccessKeyID")
   c.Aws.S3.SecretAccessKey = os.Getenv("SecretAccessKey")

   return c
}

//func init()  {
//    // sessionの作成
//    sess := session.Must(session.NewSessionWithOptions(session.Options{
//        SharedConfigState: session.SharedConfigEnable,
//    }))
//    sess, err := session.NewSession(&aws.Config{
//        Region: aws.String("ap-northeast-1")},
//    )
//
//}