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

func CompileTemplate(rawTemplate *string, values interface{}) string {
	if !isPointerToStruct(values) {
		panic(errors.New("values must be a struct pointer"))
	}

	tmpl, err := template.New("").Parse(*rawTemplate)
	if err != nil {
		panic(err)
	}

	dereferentiatedPointer := reflect.ValueOf(values).Elem()
	var result bytes.Buffer
	err = tmpl.Execute(&result, dereferentiatedPointer)
	if err != nil {
		panic(err)
	}

	return result.String()
}
