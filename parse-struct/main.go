package main

import (
	"fmt"
	"reflect"
	"strings"
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
		parts := strings.Split(f.Tag.Get("urlenc"), ",")
		fmt.Println(parts)
	}
}
