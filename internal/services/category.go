package services

import (
	"net/http"
	"product_move/internal/domains"
	"product_move/internal/modules/paginate"
	"product_move/internal/repositories"
)

type CategoryService struct {
	CategoryRep repositories.CategoryRepository
}

func (c *CategoryService) FindAll(r *http.Request) (paginate.Paging[[]domains.Category], error) {
	return c.CategoryRep.FindAll(r)
}

func (c *CategoryService) FindById(id int) (domains.Category, error) {
	return c.CategoryRep.FindById(id)
}
