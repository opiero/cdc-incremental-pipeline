package postgres

import (
	"database/sql"
	"fmt"
	"os"
)

type Student struct {
	Id    int
	Name  string
	Email string
	Age   int
}

func Connect(host string, port int, user string, password string, dbName string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbName)

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

func ExecSqlScript(db *sql.DB, filepath string) error {
	content, err := os.ReadFile(filepath)
	if err != nil {
		panic(err)
	}

	command := string(content)
	_, err = db.Exec(command)
	return err
}
