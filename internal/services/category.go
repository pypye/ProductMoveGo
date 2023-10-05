package services

import (
	"product_move/internal/domains"
	"product_move/internal/repositories"
)

type CategoryService struct {
	CategoryRep repositories.CategoryRepository
}

func (c *CategoryService) FindAll() ([]domains.Category, error) {
	return c.CategoryRep.FindAll()
}

func (c *CategoryService) FindById(id int) (domains.Category, error) {
	return c.CategoryRep.FindById(id)
}
