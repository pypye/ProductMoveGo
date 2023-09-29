package main

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	dbhandler2 "product_move/internal/adapter/db_handler"
	"product_move/internal/adapter/server_handler"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	db, err := dbhandler2.NewMySqlStore()
	if err != nil {
		log.Fatal(err)
	}
	server, err := server_handler.NewAPIServer(8080, db)
	if err != nil {
		log.Fatal(err)
	}

	err = server.Run()
	if err != nil {
		log.Fatal(err)
	}
}
