package main

import (
	"fmt"
	"reflect"
)

type Point struct {
	X int `urlenc:"x,omitempty"`
	Y int `urlenc:"y,omitempty"`
}

func main() {
	t := reflect.TypeOf(Point{X: 1, Y: 2})
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		if f.PkgPath != "" {
			continue
		}
		fmt.Println(f.Tag.Get("urlenc"))
	}
}
