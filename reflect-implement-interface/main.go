package main

import (
	"fmt"
	"os"
	"reflect"
)

type Data struct {
	value string
}

//func (data *Data) Read
func (data *Data) Read(b []byte) (n int, err error) {
	return 1, nil
}

func main() {
	v := &Data{value: ""}
	rv := reflect.ValueOf(v)
	iv := reflect.TypeOf(os.Stdout).Elem()
	fmt.Println(iv)
	if rv.Type().Implements(iv) {
		fmt.Println("this object is implemented a Read interface")
	}
}
