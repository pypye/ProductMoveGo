package infrastructure

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(mysql.Open("root:password@(localhost:3306)/product_move"))
	if err != nil {
		panic(err)
	}
	return &Database{db: db}
}

func (d *Database) Get() *gorm.DB {
	return d.db
}
