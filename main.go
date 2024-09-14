package main

import (
	"bufio"
	"cdc-incremental-pipeline/postgres"
	"fmt"
	"log"
	"os"

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

	student := postgres.Student{Name: "lek", Email: "lek@dinossauro.com", Age: 100}
	err = postgres.InsertRowIntoTable(db, &student)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Pressione Enter para continuar...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	fmt.Println("Continuando...")

	updateData := postgres.UpdateTableData{
		TableName:   "students",
		ColumnName:  "email",
		ColumnValue: "'loucao@bla.co'",
		Id:          2,
	}
	err = postgres.UpdateTableRow(db, &updateData)
	if err != nil {
		log.Fatal(err)
	}
}
