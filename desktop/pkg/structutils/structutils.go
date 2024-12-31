package structutils

import (
	"fmt"
	"reflect"
)

// Get value of field from provided interface
// if it possible return value
// else throw 'error field doesn't exist'
func GetFieldValue(s interface{}, fieldName string) (*reflect.Value, error) {
	v := reflect.ValueOf(s).Elem()
	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		return nil, fmt.Errorf("поле %s не найдено", fieldName)
	}

	return &field, nil
}

// Set value to the provided interface if it possible
// else throw error
func SetFieldValue(s interface{}, fieldName string, value interface{}) error {
	v := reflect.ValueOf(s).Elem()
	field := v.FieldByName(fieldName)

	if !field.IsValid() {
		return fmt.Errorf("поле %s не найдено", fieldName)
	}

	if !field.CanSet() {
		return fmt.Errorf("поле %s не может быть изменено", fieldName)
	}

	field.Set(reflect.ValueOf(value))
	return nil
}

// Get all fieldnames of provided object
func GetStructFieldNames(i interface{}) []string {
	var fieldNames []string
	val := reflect.ValueOf(i)

	if val.Kind() == reflect.Struct {
		for i := 0; i < val.NumField(); i++ {
			fieldNames = append(fieldNames, val.Type().Field(i).Name)
		}
	}

	return fieldNames
}
