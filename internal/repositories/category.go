package repositories

import (
	"product_move/internal/domains"
	"product_move/internal/infrastructure"
)

type CategoryInterface interface {
	FindAll() ([]domains.Category, error)
	FindById(id int) (domains.Category, error)
}

type CategoryRepository struct {
}

func (c *CategoryRepository) FindAll() ([]domains.Category, error) {
	var categories []domains.Category
	infrastructure.GetDB().Get().Find(&categories)
	return categories, nil
}

func (c *CategoryRepository) FindById(id int) (domains.Category, error) {
	var category domains.Category
	infrastructure.GetDB().Get().First(&category, id)
	return category, nil
}
