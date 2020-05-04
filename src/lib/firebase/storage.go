package firebase


//https://cloud.google.com/storage/docs/uploading-objects#storage-upload-object-go
type Image struct {
	Base64 string
}

func NewImage(base64 string) *Image {
	return &Image{base64}
}


//画像をアップロードする
func (i *Image) UploadImage()(string, error) {
	//TODO:アップロード処理を書くよ


	return "https://dummyimage.com/600x400/000/fff&text=img", nil
}
