package pa

import (
	"gorm.io/gorm"
	"net/http"
	"product_move/internal/infrastructure"
	"strconv"
)

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

func Paging(r *http.Request) func(db *gorm.DB) *gorm.DB {
	page := getPage(r)
	pageSize := getSize(r)
	offset := page * pageSize
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(pageSize)
	}
}

func getTotalElements[T any](content T) int64 {
	var TotalElements int64
	infrastructure.GetDB().Model(&content).Count(&TotalElements)
	return TotalElements
}
