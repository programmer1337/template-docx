package structutils

import (
	"fmt"
	"reflect"
)

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
