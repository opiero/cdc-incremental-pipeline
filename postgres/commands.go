package postgres

import (
	"cdc-incremental-pipeline/utils"
	"database/sql"
	"errors"
	"os"
)

const (
	sqlPath                     = "./postgres/sql"
	templatesPath               = sqlPath + "/templates"
	createTablePath             = sqlPath + "/create_table.sql"
	insertIntoTableTemplatePath = templatesPath + "/insert_into_table.sql"
	updateRowTemplatePath       = templatesPath + "/update_line_from_table.sql"
)

type UpdateTableData struct {
	TableName   string
	ColumnName  string
	ColumnValue any
	Id          int
}

func (updateData *UpdateTableData) IsValidColumnValue() bool {
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

func CreateTable(db *sql.DB) error {
	command := readFile(createTablePath)
	_, err := db.Exec(command)

	return err
}

func InsertRowIntoTable(db *sql.DB, values any) error {
	commandTemplate := readFile(insertIntoTableTemplatePath)
	command := utils.CompileTemplate(&commandTemplate, values)

	_, err := db.Exec(command)

	return err
}

func UpdateTableRow(db *sql.DB, data *UpdateTableData) error {
	if !data.IsValidColumnValue() {
		return errors.New("UpdateTableData.ColumnValue must be an in, float64 or string.")
	}

	commandTemplate := readFile(updateRowTemplatePath)
	command := utils.CompileTemplate(&commandTemplate, data)

	_, err := db.Exec(command)

	return err
}
