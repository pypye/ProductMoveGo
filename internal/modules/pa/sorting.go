package pa

import (
	"gorm.io/gorm"
	"net/http"
	"strings"
)

type SortType string

const (
	ASC  SortType = "asc"
	DESC SortType = "desc"
)

type SortObject struct {
	Attribute string
	Order     SortType
}

func AsSortType(val any) SortType {
	if str, ok := val.(string); ok {
		str = strings.ToLower(str)
		str = strings.TrimSpace(str)
		switch str {
		case "asc", "ascending":
			return ASC
		case "desc", "descending":
			return DESC
		default:
			return DESC
		}
	}
	return DESC
}

func getSort(r *http.Request) SortObject {
	q := r.URL.Query()
	s := strings.Split(q.Get("sort"), ",")
	if len(s) < 2 {
		return SortObject{}
	}
	return SortObject{s[0], AsSortType(s[1])}
}

func Sort(r *http.Request) func(db *gorm.DB) *gorm.DB {
	s := getSort(r)
	return func(db *gorm.DB) *gorm.DB {
		if s.Attribute != "" && s.Order != "" {
			return db.Order(s.Attribute + " " + string(s.Order))
		}
		return db
	}
}
