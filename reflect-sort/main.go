package main

import (
	"reflect"
)

type Sortwrap struct {
	value    reflect.Value
	lessfunc func(int, int) bool
}

func NewSortwrap(v interface{}, lessfunc func(int, int) bool) *Sortwrap {
	return &Sortwrap{
		value:    reflect.ValueOf(v),
		lessfunc: lessfunc,
	}
}
func (s *Sortwrap) Len() int {
	return s.value.Len()
}

func main() {

}
