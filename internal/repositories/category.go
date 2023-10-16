package repositories

import (
	"net/http"
	"product_move/internal/domains"
	"product_move/internal/infrastructure"
	"product_move/internal/modules/pa"
)

type CategoryInterface interface {
	FindAll(r *http.Request) (domains.Category, error)
	FindById(id int) (domains.Category, error)
}

type CategoryRepository struct {
}

func NewCategoryRepository() *CategoryRepository {
	return &CategoryRepository{}
}

func (c *CategoryRepository) FindAll(r *http.Request) (pa.Response[domains.Category], error) {
	var categories []domains.Category
	infrastructure.GetDB().Scopes(
		pa.Filtering(r, "name"),
		pa.Sort(r),
		pa.Paging(r),
	).Find(&categories)
	return pa.GetResponse(r, categories), nil
}

func (c *CategoryRepository) FindById(id int) (domains.Category, error) {
	var category domains.Category
	infrastructure.GetDB().First(&category, id)
	return category, nil
}
