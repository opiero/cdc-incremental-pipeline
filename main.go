package main

import (
	"cdc-incremental-pipeline/postgres"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "password"
	dbName   = "university"
	sqlPath  = "resources/sql"
)

func main() {
	db, err := postgres.Connect(host, port, user, password, dbName)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = postgres.ExecSqlScript(db, fmt.Sprintf("%s/create_table.sql", sqlPath))
	if err != nil {
		log.Fatal(err)
	}
}
