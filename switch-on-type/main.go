package main

import (
	"errors"
	"fmt"
	"reflect"
)

func SwitchOnType(v interface{}) (string, error) {
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Map:
		for _, key := range rv.MapKeys() {
			mv := rv.MapIndex(key)
			if mv.Kind() == reflect.String {
				return "Map value string", nil
			} else if mv.Kind() == reflect.Int64 {
				return "Map value int64", nil
			}
		}
		return "Map", nil
	case reflect.Struct:
		return "Struct", nil
	default:
		return "", errors.New("unsupported type:" + rv.Type().String())
	}
}

func main() {
	msg, e := SwitchOnType(map[string]string{"key": "value"})
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(msg)
	}
}
