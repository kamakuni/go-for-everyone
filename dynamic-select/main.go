package main

import (
	"os"
)

func main() {
	//filename := "data.txt"
	//tail(filename)
	var filenames []string
	filenames = append(filenames, "data1.txt")
	filenames = append(filenames, "data2.txt")
	filenames = append(filenames, "data3.txt")
	tripleTail(filenames)
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

func tripleTail(filenames []string) {
	chs := make([]chan []byte, 3)
	for i, filename := range filenames {
		f, err := os.Open(filename)
		if err != nil {
			return
		}
		defer f.Close()

		go func() {
			defer close(chs[i])
			buf := make([]byte, 4096)
			for {
				n, err := f.Read(buf)
				if err != nil {
					return
				}
				chs[i] <- buf[:n]
			}
		}()
	}
	var data []byte
	for {
		select {
		case data = <-chs[0]:
		case data = <-chs[1]:
		case data = <-chs[2]:
		}
		os.Stdout.Write(data)
	}
}
