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
		v, ok := f.Tag.Lookup("urlenc")
		if ok {
			parts := strings.Split(v, ",")
			for _, part := range parts {
				fmt.Println(part)
			}
		}
	}
}
