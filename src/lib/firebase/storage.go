package firebase


//https://cloud.google.com/storage/docs/uploading-objects#storage-upload-object-go
type Image struct {
	Base64 string
}

type Storage struct {
	Base64 string
}

type StrageImpl interface {
	UploadImage()(string error)
	GetImage(url string)(string error)
}

func NewStrage() StrageImpl {
	return &Storage{}
}


//画像を取得する
func (s *Storage) GetImage(url string)(string error) {

}

//画像をアップロードする
func (s *Storage) UploadImage()(string error) {

}
