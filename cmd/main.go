package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"product_move/internal/controllers"
	"product_move/internal/infrastructure"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	infrastructure.NewInjector()
	db, err := infrastructure.NewDatabase()
	if err != nil {
		log.Fatal(err)
	}
	infrastructure.InitDB(db)
	infrastructure.InitServer(infrastructure.NewServer())
	controllers.NewCategoryController().Build()
	controllers.NewAuthController().Build()
	err = infrastructure.GetServer().Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
