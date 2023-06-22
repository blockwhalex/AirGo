package struct_map_plugin

import (
	"fmt"
	"reflect"
	"strings"
)

// struct转map
func StructToMap(data interface{}) map[string]interface{} {

	m := make(map[string]interface{})

	v := reflect.ValueOf(data)

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return m
	}
	t := v.Type()

	for i := 0; i < t.NumField(); i++ {
		name := t.Field(i).Name
		tag := t.Field(i).Tag.Get("json")
		if tag == "-" || name == "-" {
			continue
		}
		if tag != "" {
			index := strings.Index(tag, ",")
			if index == -1 {
				name = tag
			} else {
				name = tag[:index]
			}
		}
		m[name] = v.Field(i).Interface()
	}
	return m
}

// 利用反射将结构体转化为map
func StructToMap1(obj interface{}) map[string]interface{} {
	obj1 := reflect.TypeOf(obj)
	obj2 := reflect.ValueOf(obj)

	data := make(map[string]interface{})
	for i := 0; i < obj1.NumField(); i++ {
		if obj1.Field(i).Tag.Get("mapstructure") != "" {
			data[obj1.Field(i).Tag.Get("mapstructure")] = obj2.Field(i).Interface()
		} else {
			data[obj1.Field(i).Name] = obj2.Field(i).Interface()
		}
	}
	return data
}

// golang struct 动态创建
func RegisterType(elem ...interface{}) map[string]reflect.Type {
	var typeRegistry = make(map[string]reflect.Type)
	for i := 0; i < len(elem); i++ {
		t := reflect.TypeOf(elem[i])
		typeRegistry[t.Name()] = t
	}
	return typeRegistry
}
func NewStruct(name string, typeRegistry map[string]reflect.Type) (interface{}, bool) {
	elem, ok := typeRegistry[name]
	fmt.Println("elem", elem)
	if !ok {
		return nil, false
	}
	return reflect.New(elem).Elem().Interface(), true

}
