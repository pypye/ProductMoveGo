package infrastructure

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	db *gorm.DB
}

func NewDatabase() *Database {
	db, err := gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@(%s:%d)/%s",
		viper.Get("database.username"),
		viper.Get("database.password"),
		viper.Get("database.host"),
		viper.Get("database.port"),
		viper.Get("database.schema"),
	)))
	if err != nil {
		panic(err)
	}
	return &Database{db: db}
}

func (d *Database) Get() *gorm.DB {
	return d.db
}
