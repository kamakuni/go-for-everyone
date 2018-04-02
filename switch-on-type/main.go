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
		for _, key := range rv.mapkeys() {
		}
		return "Map", nil
	case reflect.Struct:
		return "Struct", nil
	default:
		return "", errors.New("unsupported type:" + rv.Type().String())
	}
}

func main() {
	msg, e := SwitchOnType(make(map[string]string))
	if e != nil {
		fmt.Println(e)
	} else {
		fmt.Println(msg)
	}
}
