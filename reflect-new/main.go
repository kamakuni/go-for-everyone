package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	t := MakeTime()
	fmt.Println(reflect.TypeOf(t))
}

var timeT = reflect.TypeOf(time.Time{})

func MakeTime() *time.Time {
	rv := reflect.New(timeT)
	return rv.Interface().(*time.Time)
}
