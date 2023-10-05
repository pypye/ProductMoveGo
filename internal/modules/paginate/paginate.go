package paginate

import (
	"gorm.io/gorm"
	"net/http"
	"product_move/internal/infrastructure"
	"strconv"
)

type Paging[T any] struct {
	Page          int `json:"page"`
	Size          int `json:"size"`
	TotalElements int `json:"total_elements"`
	TotalPages    int `json:"total_pages"`
	Content       T   `json:"content"`
}

func getSize(r *http.Request) int {
	q := r.URL.Query()
	size, _ := strconv.Atoi(q.Get("size"))
	switch {
	case size > 50:
		size = 50
	case size <= 0:
		size = 10
	}
	return size
}

func getPage(r *http.Request) int {
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get("page"))
	if page < 0 {
		page = 0
	}
	return page
}

func Paginate(r *http.Request) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page := getPage(r)
		pageSize := getSize(r)
		offset := page * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func getTotalElements[T any](content T) int64 {
	var TotalElements int64
	infrastructure.GetDB().Get().Model(&content).Count(&TotalElements)
	return TotalElements
}

func GetPagination[T any](r *http.Request, content T) Paging[T] {
	totalElements := int(getTotalElements(content))
	return Paging[T]{
		Page:          getPage(r),
		Size:          getSize(r),
		TotalElements: totalElements,
		TotalPages:    (totalElements-1)/getSize(r) + 1,
		Content:       content,
	}
}
