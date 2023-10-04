package repositories

import (
	"database/sql"
	"product_move/internal/domains"
	"product_move/internal/infrastructure"
)

type CategoryInterface interface {
	FindAll() (*domains.Categories, error)
	FindById(id int) (*domains.Category, error)
}

type CategoryRepository struct {
}

func (c *CategoryRepository) FindAll() (domains.Categories, error) {
	rows, err := infrastructure.GetDB().Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var categories domains.Categories
	for rows.Next() {
		var category domains.Category
		err := rows.Scan(&category.Id, &category.CreatedTime, &category.LastUpdatedTime, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *CategoryRepository) FindById(id int) (domains.Category, error) {
	rows, err := infrastructure.GetDB().Query("SELECT * FROM category WHERE id = ?", id)
	if err != nil {
		return domains.Category{}, err
	}
	defer func(rows *sql.Rows) {
		_ = rows.Close()
	}(rows)

	var category domains.Category
	for rows.Next() {
		err := rows.Scan(&category.Id, &category.CreatedTime, &category.LastUpdatedTime, &category.Name)
		if err != nil {
			return domains.Category{}, err
		}
	}
	return category, nil
}
