package controllers

import (
	"net/http"
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

func (c *CategoryController) IndexHandler(ctx *infrastructure.ServerCtx) {
	categories, err := c.categoryService.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, categories)
}

func (c *CategoryController) FindByIdHandler(ctx *infrastructure.ServerCtx) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	category, err := c.categoryService.FindById(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}
	ctx.JSON(http.StatusOK, category)
}
