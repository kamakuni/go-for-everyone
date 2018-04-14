package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 5, Y: 10}
	rv := reflect.ValueOf(p)
	if !rv.Field(0).CanSet() {
		fmt.Println("This field can't set")
	}
	//rv.Field(0).SetInt(10)
}
