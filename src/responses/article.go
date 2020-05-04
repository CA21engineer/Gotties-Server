package responses

import (
	"Gotties-Server/src/models"
	"time"
)

type Article struct {
	ID uint `json:"id,string"`
	Title string `json:"title"`
	Before string `json:"before"`
	After string `json:"after"`
	Body string `json:"body"`
	Category *Category`json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Articles []Article

func (a *Article) ResponseArticle(article *models.Article) *Article{
	//TODO: 誰か綺麗なやり方教えてw
	return &Article{
		ID:     article.ID,
		Title:  article.Title,
		Before: article.Before,
		After:  article.After,
		Body:   article.Body,
		Category: new(Category).ResponseCategory(&article.Category),
		CreatedAt: article.CreatedAt,
		UpdatedAt: article.UpdatedAt,
	}
}