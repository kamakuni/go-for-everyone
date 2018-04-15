package main

import (
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	/*p := Point{X: 5, Y: 10}
	rv := reflect.ValueOf(p)
	if !rv.Field(0).CanSet() {
		fmt.Println("This field can't set")
	}
	rv.Field(0).SetInt(10)*/
	p := Point{X: 10, Y: 5}
	rv := reflect.ValueOf(&p)
	if f := rv.Field(0); f.CanSet() {
		f.SetInt(100)
	}
}
