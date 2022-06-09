package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/unrealandychan/golang_api_full/api"
	db "github.com/unrealandychan/golang_api_full/db/sqlc"
	"github.com/unrealandychan/golang_api_full/util"
	"log"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("Cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Can not start server", err)
	}
}
