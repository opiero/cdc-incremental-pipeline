package utils

import (
	"bytes"
	"errors"
	"reflect"
	"text/template"
)

func isPointerToStruct(value interface{}) bool {
	v := reflect.ValueOf(value)

	if v.Kind() == reflect.Ptr {
		if v.Elem().Kind() == reflect.Struct {
			return true
		}
	}

	return false
}

func CompileTemplate(rawTemplate *string, values *interface{}) (string, error) {
	if reflect.TypeOf(*values).Kind() != reflect.Struct {
		return "", errors.New("values must be a struct")
	}

	tmpl, err := template.New("").Parse(*rawTemplate)
	if err != nil {
		panic(err)
	}

	var result bytes.Buffer
	err = tmpl.Execute(&result, *values)
	if err != nil {
		panic(err)
	}

	return result.String(), nil
}
