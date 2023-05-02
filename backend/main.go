package main

import (
	"gpt-chan/api"
	db "gpt-chan/database/models"
	"gpt-chan/util"

	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var query_obj *db.Queries

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	query_obj = db.New(conn)
	server := api.NewServer(query_obj)

	if err := server.Start(config.ServerAddress); err != nil {
		log.Fatal("cannot start server:", err)
	}
}
