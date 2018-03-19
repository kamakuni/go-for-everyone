package main

import (
	"fmt"
	"time"
)

func doMain() {
	defer fmt.Println("done infinite loop")
	for {
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go doMain()
}
