package main

import (
	crand "crypto/rand"
	"encoding/binary"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var s int64
	if err := binary.Read(crand.Reader, binary.LittleEndian, &s); err != nil {
		s = time.Now().UnixNano()
	}
	rand.Seed(s)
	n := rand.Intn(100)
	fmt.Println(n)
}
