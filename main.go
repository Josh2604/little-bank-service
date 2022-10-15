package main

import (
	"database/sql"
	"log"

	"github.com/Josh2604/master-class/api"
	db "github.com/Josh2604/master-class/db/sqlc"
	"github.com/Josh2604/master-class/utils"
	_ "github.com/lib/pq"
)

func main() {
	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config: %v", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("can not start server: ", err)
	}
}
