package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "root"
	password = "password"
	dbname   = "university"
)

func connect() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := connect()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	content, err := os.ReadFile("resources/sql/create_table.sql")
	if err != nil {
		panic(err)
	}

	create_table := string(content)
	result, err := db.Exec(create_table)
	if err != nil {
		panic(err)
	}
	fmt.Printf("result: %v\n", result)
}
