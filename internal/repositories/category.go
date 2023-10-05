package repositories

import (
	"net/http"
	"product_move/internal/domains"
	"product_move/internal/infrastructure"
	"product_move/internal/modules/paginate"
)

type CategoryInterface interface {
	FindAll(r *http.Request) ([]domains.Category, error)
	FindById(id int) (domains.Category, error)
}

type CategoryRepository struct {
}

func (c *CategoryRepository) FindAll(r *http.Request) (paginate.Paging[[]domains.Category], error) {
	var categories []domains.Category
	infrastructure.GetDB().Get().Scopes(paginate.Paginate(r)).Find(&categories)
	return paginate.GetPagination(r, categories), nil
}

func (c *CategoryRepository) FindById(id int) (domains.Category, error) {
	var category domains.Category
	infrastructure.GetDB().Get().First(&category, id)
	return category, nil
}
