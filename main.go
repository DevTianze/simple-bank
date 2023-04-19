package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/DevTianze/simple-bank/api"
	db "github.com/DevTianze/simple-bank/db/sqlc"
	"github.com/DevTianze/simple-bank/util"
	_ "github.com/lib/pq"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	fmt.Print("server started at ", config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
