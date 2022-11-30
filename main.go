package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Sugaml/mrc-payment/api/config"
	"github.com/Sugaml/mrc-payment/api/controller"
	"github.com/Sugaml/mrc-payment/api/db/postgres"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("error in load env file %v", err)
	} else {
		fmt.Println("Loaded env files...")
	}
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	conn, err := gorm.Open(conf.DBDriver, conf.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db ", err)
	}
	pdb := postgres.NewDB(conn)
	server, err := controller.NewServer(pdb.DB)
	if err != nil {
		log.Fatal(err)
		return
	}
	port := ":" + os.Getenv("PORT")
	server.Run(port)
}
