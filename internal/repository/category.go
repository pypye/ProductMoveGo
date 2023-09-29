package repository

import (
	"database/sql"
	"product_move/internal/common/types"
)

type CategoryInterface interface {
	FindAll() (types.Categories, error)
	FindById(id int) (types.Category, error)
}

type CategoryRepository struct {
	Db *sql.DB
}

func (c *CategoryRepository) FindAll() (types.Categories, error) {
	var categories types.Categories
	rows, err := c.Db.Query("SELECT * FROM category")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		var category types.Category
		err := rows.Scan(&category.Id, &category.CreatedTime, &category.LastUpdatedTime, &category.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}

func (c *CategoryRepository) FindById(id int) (types.Category, error) {
	var category types.Category
	rows, err := c.Db.Query("SELECT * FROM category WHERE id = ?", id)
	if err != nil {
		return category, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {

		}
	}(rows)

	for rows.Next() {
		err := rows.Scan(&category.Id, &category.CreatedTime, &category.LastUpdatedTime, &category.Name)
		if err != nil {
			return category, err
		}
	}
	return category, nil
}
