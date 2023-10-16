package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"log"
	"product_move/internal/controllers"
	"product_move/internal/infrastructure"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile | log.Lmsgprefix)
	viper.SetConfigFile("configs/config.yml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	infrastructure.NewInjector()
	infrastructure.InitDB()
	infrastructure.InitServer()
	controllers.BuildController()
	err = infrastructure.GetServer().Run(fmt.Sprintf(":%d", viper.Get("server.port")))
	if err != nil {
		log.Fatal(err)
	}
}
