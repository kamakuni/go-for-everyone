package main

import (
	"reflect"
	"time"
)

func main() {

}

var timeT = reflect.TypeOf(time.Time{})

func MakeTime() *time.Time {
	rv := reflect.New(timeT)
	return rv.Interface().(*time.Time)
}
