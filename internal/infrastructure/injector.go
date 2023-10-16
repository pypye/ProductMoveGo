package infrastructure

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Injector struct {
	Db     *Database
	Server *Server
}

var instance *Injector

func NewInjector() *Injector {
	if instance == nil {
		instance = &Injector{}
	}
	return instance
}

func InitDB() {
	instance.Db = NewDatabase()
}

func InitServer() {
	instance.Server = NewServer()
}

func GetServer() *gin.Engine {
	return instance.Server.Get()
}

func GetDB() *gorm.DB {
	return instance.Db.Get()
}
