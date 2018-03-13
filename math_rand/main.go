package main

import (
	"fmt"
	"math/rand"
)

func main() {
	rand.Seed(42)
	n := rand.Intn(100)
	fmt.Println(n)
}
