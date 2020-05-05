package aws

import (
	"bytes"
	"encoding/base64"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"strconv"
	"time"
)

//https://cloud.google.com/storage/docs/uploading-objects#storage-upload-object-go
type Image struct {
	Base64 string
}

func NewImage(base64 string) *Image {
	return &Image{base64}
}


//画像をアップロードする
func (i *Image) UploadImage()(string, error) {
	// sessionの作成
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("ap-northeast-1")},
	)
	if err != nil {
		return "", err
	}

	//正確にcledentialsに接続できたかどうかを確認する
	_, err = sess.Config.Credentials.Get()
	if err != nil {
		return "", err
	}


	bucketName := os.Getenv("BUCKET_NAME")
	fileExtension := "png"

	uploader := s3manager.NewUploader(sess)

	data, _ := base64.StdEncoding.DecodeString(i.Base64)
	wb := new(bytes.Buffer)
	wb.Write(data)

	t := strconv.FormatInt(time.Now().Unix(), 10)

	res, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String("/images/" + t + "." + fileExtension),
		Body:   wb,
		ContentType: aws.String("image/" + fileExtension),
	})

	if err != nil {
		if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
			return "", err
		} else {
			return "", err
		}
	}

	return aws.StringValue(&res.Location), nil
}

//
//func (s S3Service) UploadToS3(i.base64 string, fileExtension string) error {
//	// 環境変数からS3Credential周りの設定を取得
//	var env Env
//	bucketName := "バケット名"
//	_ = envconfig.Process("", &env)
//
//	sess := session.Must(session.NewSession(&aws.Config{
//		Credentials: credentials.NewStaticCredentials(env.S3AK, env.S3SK, ""),
//		Region:      aws.String("ap-northeast-1"),
//	}))
//
//	uploader := s3manager.NewUploader(sess)
//
//	data, _ := base64.StdEncoding.DecodeString(imageBase64)
//	wb := new(bytes.Buffer)
//	wb.Write(data)
//
//	res, err := uploader.Upload(&s3manager.UploadInput{
//		Bucket: aws.String(bucketName),
//		Key:    aws.String("ファイル名" + "." + fileExtension),
//		Body:   wb,
//		ContentType: aws.String("image/" + fileExtension),
//	})
//
//	if err != nil {
//		fmt.Println(res)
//		if err, ok := err.(awserr.Error); ok && err.Code() == request.CanceledErrorCode {
//			return RaiseError(400, "Upload TimuOut", nil)
//		} else {
//			return RaiseError(400, "Upload Failed", nil)
//		}
//	}
//
//	return nil
//}
