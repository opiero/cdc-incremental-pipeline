package main

import (
	"cdc-incremental-pipeline/postgres"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "password"
	dbName   = "university"
)

func main() {
	db, err := postgres.Connect(host, port, user, password, dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = postgres.CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}
}
