package main

import (
	"fmt"
	"github.com/kamakuni/go-for-everyone/reflect-package/myapp"
	"reflect"
)

func main() {
	p := myapp.Person{}
	rt := reflect.TypeOf(p)
	rv := reflect.ValueOf(p)
	for i := 0; i < rv.NumField(); i++ {
		ft := rt.Field(i)
		fv := rv.Field(i)
		if ft.PkgPath == "" {
			fmt.Printf("ft(i) -> %#v\n", i, ft)
			fmt.Printf("fv(i) -> %#v\n", i, fv.Interface())
		}
	}
}
