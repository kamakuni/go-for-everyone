package main

import (
	"fmt"
	"reflect"
)

type Data struct {
}

func (d *Data) PrtFoo() string {
	return "nothing"
}

func (d Data) Foo() string {
	return "nothing"
}

func main() {
	val := reflect.TypeOf(Data{})
	fmt.Printf("Name = %s\n", val.Method(0).Name)
	val = reflect.TypeOf(&Data{})
	fmt.Printf("Name = %s\n", val.Method(0).Name)
}
