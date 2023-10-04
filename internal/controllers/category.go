package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"product_move/internal/helpers"
	"product_move/internal/infrastructure"
	"product_move/internal/repositories"
	"product_move/internal/services"
	"strconv"
)

type CategoryController struct {
	categoryService services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		categoryService: services.CategoryService{
			CategoryRep: repositories.CategoryRepository{},
		},
	}
}

func (c *CategoryController) Build() {
	infrastructure.GetMapping("/", c.IndexHandler)
	infrastructure.GetMapping("/:id", c.FindByIdHandler)
}

func (c *CategoryController) IndexHandler(ctx *gin.Context) {
	categories, err := c.categoryService.FindAll()
	if err != nil {
		helpers.Write(ctx, http.StatusInternalServerError, err)
		return
	}
	helpers.Write(ctx, http.StatusOK, categories)
}

func (c *CategoryController) FindByIdHandler(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		helpers.WriteError(ctx, http.StatusBadRequest, err)
		return
	}
	category, err := c.categoryService.FindById(id)
	if err != nil {
		helpers.WriteError(ctx, http.StatusInternalServerError, err)
		return
	}
	helpers.Write(ctx, http.StatusOK, category)
}
