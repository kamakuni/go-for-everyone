package main

import (
	"reflect"
)

type Point struct {
	X int
	Y int
}

func main() {
	p := Point{X: 5, Y: 10}
	rv := reflect.ValueOf(p)
	rv.Field(0).SetInt(10)
}
