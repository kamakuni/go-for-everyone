package main

import (
	"fmt"
)

var processing bool

func doSomething() {
	processing = false
	defer func() {
		processing = true
	}()
}

func main() {
	fmt.Println(processing)
	doSomething()
	fmt.Println(processing)
}
