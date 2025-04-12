package Reflection

import (
	"errors"
	"reflect"
)

func ValidateRequiredFields(input interface{}) error {
	val := reflect.ValueOf(input)

	if val.Kind() != reflect.Struct {
		return errors.New("expected a struct")
	}

	typ := val.Type()
	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("required")

		if tag == "true" {
			if isZeroValue(field) {
				return errors.New("missing required field: " + typ.Field(i).Name)
			}
		}
	}

	return nil
}

func isZeroValue(v reflect.Value) bool {
	return reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface())
}
