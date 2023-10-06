package pa

import (
	"gorm.io/gorm"
	"net/http"
)

func getFilter(r *http.Request, query string) string {
	q := r.URL.Query()
	filter := q.Get(query)
	return filter
}

func Filtering(r *http.Request, query string) func(db *gorm.DB) *gorm.DB {
	q := getFilter(r, query)
	return func(db *gorm.DB) *gorm.DB {
		if q != "" {
			return db.Where(query+" LIKE ?", "%"+q+"%")
		} else {
			return db
		}
	}
}

func FilteringExactly(r *http.Request, query string) func(db *gorm.DB) *gorm.DB {
	q := getFilter(r, query)
	return func(db *gorm.DB) *gorm.DB {
		if q != "" {
			return db.Where(query+" = ?", q)
		} else {
			return db
		}
	}
}

func FilteringDate(r *http.Request, queryFrom, queryTo string) func(db *gorm.DB) *gorm.DB {
	qf := getFilter(r, queryFrom)
	qt := getFilter(r, queryTo)
	return func(db *gorm.DB) *gorm.DB {
		if qf != "" && qt != "" {
			return db.Where(queryFrom+" >= ? AND "+queryTo+" <= ?", qf, qt)
		} else {
			return db
		}
	}
}
