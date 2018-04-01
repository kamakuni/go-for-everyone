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
		return "Map", nil
	case reflect.Struct:
		return "Struct", nil
	default:
		return "", errors.New("unsupported type:" + rv.Type().String())
	}
}

func main() {
	msg, e := SwitchOnType("test")
	if e != nil {
		fmt.Println(msg)
	} else {
		fmt.Println(e)
	}
}
