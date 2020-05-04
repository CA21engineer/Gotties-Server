package responses

import (
	"Gotties-Server/src/models"
	"time"
)

type Category struct {
	ID uint `json:"id,string"`
	Name string `json:"name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
type Categories []Category

func (c *Category) ResponseCategory(category *models.Category) *Category{
	return &Category{
		ID:     category.ID,
		Name:  category.Name,
		CreatedAt: category.CreatedAt,
		UpdatedAt: category.UpdatedAt,
	}
}