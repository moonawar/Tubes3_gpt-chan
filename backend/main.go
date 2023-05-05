package main

import (
	"database/sql"
	"fmt"
	"gpt-chan/api"
	db "gpt-chan/database/models"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var query_obj *db.Queries

func main() {
	// config, err := util.LoadConfig(".")
	// if err != nil {
	// 	log.Fatal("cannot load config:", err)
	// }
	dbSocketDir := os.Getenv("DB_SOCKET_DIR")
	conn, err := sql.Open("postgres", fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", dbSocketDir, os.Getenv("DB_USER"), os.Getenv("DB_PASS"), os.Getenv("DB_NAME")))
	if err != nil {
		log.Fatal("cannot connect to database:", err)
	}

	query_obj = db.New(conn)
	server := api.NewServer(query_obj)

	if err := server.Start("0.0.0.0:8080"); err != nil {
		log.Fatal("cannot start server:", err)
	}
	// alg := algorithm.New()
	// res := alg.LevenshteinDistance("Kapan tubes ini selesak", "Kapan ttubes ini selesai")
	// fmt.Println(res)
}
