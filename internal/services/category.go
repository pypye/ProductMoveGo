package services

import (
	"net/http"
	"product_move/internal/domains"
	"product_move/internal/modules/pa"
	"product_move/internal/repositories"
)

type CategoryService struct {
	CategoryRep *repositories.CategoryRepository
}

func NewCategoryService() *CategoryService {
	return &CategoryService{CategoryRep: repositories.NewCategoryRepository()}
}

func (c *CategoryService) FindAll(r *http.Request) (pa.Response[domains.Category], error) {
	return c.CategoryRep.FindAll(r)
}

func (c *CategoryService) FindById(id int) (domains.Category, error) {
	return c.CategoryRep.FindById(id)
}
