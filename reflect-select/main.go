package main

import (
	"fmt"
	"os"
)

func readFromFile(ch chan []byte, f *os.File) {
	defer close(ch)
	defer f.Close()

	buf := make([]byte, 4096)
	for {
		if n, err := f.Read(buf); err == nil {
			ch <- buf[:n]
		}
	}
}

func main() {
	fmt.Println("Nothing")
}
