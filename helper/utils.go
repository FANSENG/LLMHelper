package helper

import "reflect"

func StructToMap(obj interface{}) map[string]interface{} {
	val := reflect.ValueOf(obj)
	typ := reflect.TypeOf(obj)

	if val.Kind() == reflect.Ptr {
		val = val.Elem()
		typ = typ.Elem()
	}
	if val.Kind() != reflect.Struct {
		panic("structToMap only accepts structs")
	}

	result := make(map[string]interface{})

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		fieldName := typ.Field(i).Tag.Get("json")
		if fieldName == "" {
			fieldName = typ.Field(i).Name
		}
		if fieldName != "-" {
			result[fieldName] = field.Interface()
		}
	}

	return result
}
