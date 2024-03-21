package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	db "simple_bank/db/sqlc"
	"simple_bank/internal/service/api"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://postgres:v8hlDV0yMAHHlIurYupj@localhost:5434/simplebank?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect db:", err)
	}
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server")
	}
}
