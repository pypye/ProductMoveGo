package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"product_move/internal/repository"
)

type CategoryController struct {
	categoryRep repository.CategoryRepository
}

func NewCategoryController(db *sql.DB) *CategoryController {
	return &CategoryController{
		categoryRep: repository.CategoryRepository{
			Db: db,
		},
	}
}

func (c *CategoryController) IndexHandler(w http.ResponseWriter, r *http.Request) {
	categories, err := c.categoryRep.FindAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(categories)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte(err.Error()))
		return
	}
}
