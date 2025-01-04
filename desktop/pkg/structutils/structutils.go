package structutils

import (
	"fmt"
	"reflect"
)

// Get value of field from provided interface.
// If it possible return value.
// Else throw 'error field doesn't exist'
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

func GetStructValues(i interface{}) []any {
	values := []any{}
	testStruct := i
	v := reflect.ValueOf(testStruct)

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		values = append(values, field)
		// fmt.Printf("Поле: %s, Тип: %v\n", field.Name, field.Type)
	}

	return values
}
