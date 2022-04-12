package structx

import (
	"reflect"
)

// Struct2Map 结构体转 map
func Struct2Map(obj interface{}) map[string]interface{} {
	t := reflect.TypeOf(obj)
	v := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < t.NumField(); i++ {
		jsonTagName := t.Field(i).Tag.Get("json")
		val := v.Field(i).Interface()
		if jsonTagName != "" {
			data[jsonTagName] = val
		} else {
			data[t.Field(i).Name] = val
		}
	}

	return data
}
