package main

import (
	"fmt"
	"time"
)

type AsyncResult struct {
	r   string
	err error
}

func doSomething() (string, error) {
	time.Sleep(5 * time.Minute)
	return "result", nil
}

func doSomethingWithTimeout(timer *time.Timer) (string, error) {
	done := make(chan AsyncResult)

	go func() {
		r, err := doSomething()
		done <- AsyncResult{r, err}
	}()
	select {
	case <-timer.C:
		return "", fmt.Errorf("timeout reached")
	case result := <-done:
		if result.err != nil {
			return "", result.err
		}
		return result.r, nil
	}
}

func main() {
	timer := time.NewTimer(10 * time.Second)
	r, err := doSomethingWithTimeout(timer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r)
}
