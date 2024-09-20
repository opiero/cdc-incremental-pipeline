package postgres

import (
	"cdc-incremental-pipeline/utils"
	"database/sql"
	"errors"
	"os"
)

const (
	sqlPath                       = "./postgres/sql"
	templatesPath                 = sqlPath + "/templates"
	createTablePath               = sqlPath + "/create_table.sql"
	insertIntoTableTemplatePath   = templatesPath + "/insert_into_table.sql"
	updateRowTemplatePath         = templatesPath + "/update_line_from_table.sql"
	deleteRowTemplatePath         = templatesPath + "/delete_from_table.sql"
	createReplicatorUserPath      = sqlPath + "/create_replicator_role.sql"
	createStudentsPublicationPath = sqlPath + "/create_publication.sql"
)

type DeleteTableData struct {
	TableName string
	Id        int
}

type UpdateTableData struct {
	TableName   string
	ColumnName  string
	ColumnValue any
	Id          int
}

func (updateData *UpdateTableData) isValidColumnValue() bool {
	columnValue := updateData.ColumnValue

	_, isInt := columnValue.(int)
	_, isFloat := columnValue.(float64)
	_, isString := columnValue.(string)

	return isInt || isFloat || isString
}

func readFile(path string) string {
	content, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return string(content)
}

func runCommand(db *sql.DB, path string) error {
	command := readFile(path)

	_, err := db.Exec(command)

	return err
}

func runTemplateCommand(db *sql.DB, values any, path string) error {
	commandTemplate := readFile(path)
	command := utils.CompileTemplate(&commandTemplate, values)

	_, err := db.Exec(command)

	return err
}

func CreateTable(db *sql.DB) error {
	return runCommand(db, createTablePath)
}

func InsertRowIntoTable(db *sql.DB, values any) error {
	return runTemplateCommand(db, values, insertIntoTableTemplatePath)
}

func UpdateTableRow(db *sql.DB, data *UpdateTableData) error {
	if !data.isValidColumnValue() {
		return errors.New("UpdateTableData.ColumnValue must be an int, float64 or string.")
	}

	return runTemplateCommand(db, data, updateRowTemplatePath)
}

func DeleteRowFromTable(db *sql.DB, data *DeleteTableData) error {
	return runTemplateCommand(db, data, deleteRowTemplatePath)
}

func CreateReplicatorUser(db *sql.DB) error {
	return runCommand(db, createReplicatorUserPath)
}

func CreatePublication(db *sql.DB) error {
	return runCommand(db, createStudentsPublicationPath)
}

func CreateSubscription(db *sql.DB) error {
	return nil
}
