package utils

import (
	"errors"
	"reflect"
)

func CompileTemplate(rawTemplate *string, values *interface{}) (string, error) {
	if reflect.TypeOf(*values).Kind() != reflect.Struct {
		return "", errors.New("values must be a struct")
	}
	return "", nil
}
