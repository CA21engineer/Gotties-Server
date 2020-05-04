package form


type Article struct {
	Title string `validate:"required"`
	Before string `validate:"required,base64"`
	After string `validate:"required,base64"`
	Body string `validate:"required"`
	CategoryName string `validate:"required"`
	UserId string
}

func NewArticle(title, before, after, body, categoryName, userId string) *Article {
	return &Article{
		Title: title,
		Before: before,
		After: after,
		Body: body,
		CategoryName: categoryName,
		UserId: userId,
	}
}