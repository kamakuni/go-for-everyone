package main

import (
	"fmt"
	"os"
)

func main() {
	f, _ := os.Open("file.txt")
	HandleData(f)
}

func HandleData(x interface{}) {
	f, ok := x.(*os.File)
	if ok {
		fmt.Println("file is", f.Name())
	}
}
