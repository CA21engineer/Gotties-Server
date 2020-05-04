package form


type Article struct {
	Title string `validate:"required`
	Before string `validate:"required, Base64`
	After string `validate:"required, Base64`
	Body string `validate:"required`
	CategoryName string `validate:"required`
	UserId string `validate:"required`
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