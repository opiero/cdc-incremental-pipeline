package main

import (
	"bytes"
	"cdc-incremental-pipeline/postgres"
	"fmt"
	"log"
	"os"
	"text/template"

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

	// err = postgres.ExecSqlScript(db, fmt.Sprintf("%s/create_table.sql", sqlPath))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	content, err := os.ReadFile(fmt.Sprintf("%s/templates/insert_into_table.sql", sqlPath))
	if err != nil {
		log.Fatal(err)
	}

	student := &postgres.Student{Name: "fodao", Email: "fod@o.com", Age: 22}
	tmpl, err := template.New("insert").Parse(string(content))
	var result bytes.Buffer
	tmpl.Execute(&result, student)
	command := result.String()
	fmt.Println(command)
	// db.Exec(command)
}
