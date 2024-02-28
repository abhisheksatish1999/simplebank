package main

import (
	"database/sql"
	"log"

	"github.com/abhisheksatish1999/simplebank/api"
	db "github.com/abhisheksatish1999/simplebank/db/sqlc"
	"github.com/abhisheksatish1999/simplebank/util"
	_ "github.com/lib/pq"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("unable to load config file:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(store, config)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("failed to start server: ", err)
	}
}
