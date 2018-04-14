package main

import (
	"fmt"
	"reflect"
)

type Data struct {
	value string
}

func main() {
	data := &Data{value: "value"}
	rv := reflect.ValueOf(data)
	if rv.Kind() == reflect.Ptr {
		rv = rv.Elem()
		fmt.Println(rv)
	}
}
