package pa

import (
	"net/http"
)

type Response[T any] struct {
	Page          int `json:"page"`
	Size          int `json:"size"`
	TotalElements int `json:"total_elements"`
	TotalPages    int `json:"total_pages"`
	Content       []T `json:"content"`
}

func GetResponse[T any](r *http.Request, content []T) Response[T] {
	totalElements := int(getTotalElements(content))
	return Response[T]{
		Page:          getPage(r),
		Size:          getSize(r),
		TotalElements: totalElements,
		TotalPages:    (totalElements-1)/getSize(r) + 1,
		Content:       content,
	}
}
