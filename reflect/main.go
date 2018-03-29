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
	p := &Point{X: 10, Y: 5}
	rv := reflect.ValueOf(p)
	fmt.Printf("rv.Type = %v\n", rv.Type())
	fmt.Printf("rv.Kind = %v\n", rv.Kind())
	fmt.Printf("rv.Interface = %v\n", rv.Interface())
	/*xv := rv.Field(0)
	fmt.Printf("xv = %d\n", xv.Int())
	xv.SetInt(100)
	fmt.Printf("xv (after SetInt) = %d\n", xv.Int())*/
	rv1 := reflect.ValueOf(1)
	/*rv2 := reflect.ValueOf("Hello World")
	rv3 := reflect.ValueOf([]byte{0xa, 0xd})
	rv4 := reflect.ValueOf(make(chan struct{}))*/
	fmt.Printf("rv1 = %d\n", rv1.Int())
	var num int64
	if rv1.Kind() == reflect.Int {
		num = rv1.Int()
	}
	fmt.Printf("num = %d\n", num)
	/*fmt.Printf("rv2 = %d\n", rv2.Int())
	fmt.Printf("rv3 = %d\n", rv3.Int())
	fmt.Printf("rv4 = %d\n", rv4.Int())*/
	rvm := reflect.ValueOf(map[string]int{"foo": 1})
	value := rvm.MapIndex(reflect.ValueOf("foo"))
	fmt.Println(value)
	rvm.SetMapIndex(reflect.ValueOf("bar"), reflect.ValueOf(2))
	value = rvm.MapIndex(reflect.ValueOf("bar"))
	fmt.Println(value)
	//rvm.Int()
	//reflect.Type("Point")
	fmt.Println(reflect.TypeOf(Point{}))
}
