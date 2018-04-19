package main

import (
	"os"
)

func main() {
	filename := "data.txt"
	tail(filename)
}

func tail(filename string) {
	f, err := os.Open(filename)
	if err != nil {
		return
	}
	defer f.Close()

	ch := make(chan []byte)
	go func() {
		defer close(ch)
		buf := make([]byte, 4096)
		for {
			n, err := f.Read(buf)
			if err != nil {
				return
			}
			ch <- buf[:n]
		}
	}()
	for in := range ch {
		os.Stdout.Write(in)
	}
}

func multipleTail() {
	buf := make([]byte, 4096)
	os.Stdout.Write(buf)
}
