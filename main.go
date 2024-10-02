package main

import (
	"cdc-incremental-pipeline/debezium"
	"cdc-incremental-pipeline/postgres"

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

	debezium.ConnectToDebezium()
}
