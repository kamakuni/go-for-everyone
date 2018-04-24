package main

import (
	"errors"
	"fmt"
	"os"
	"reflect"
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

func makeChannelsFromFiles(files ...string) ([]reflect.Value, error) {
	cs := make([]reflect.Value, len(files))

	for i, fn := range files {
		ch := make(chan []byte)
		f, err := os.Open(fn)
		if err != nil {
			return nil, err
		}
		go readFromFile(ch, f)

		cs[i] = reflect.ValueOf(ch)
	}
	return cs, nil
}

func makeSelectCases(cs ...reflect.Value) ([]reflect.SelectCase, error) {
	cases := make([]reflect.SelectCase, len(cs))
	for i, ch := range cs {
		if ch.Kind() != reflect.Chan {
			return nil, errors.New("argument must be a channel")
		}
		cases[i] = reflect.SelectCase{
			Chan: ch,
			Dir:  reflect.SelectRecv,
		}
	}
	return cases, nil
}

func main() {
	fmt.Println("Nothing")
}
