package main

import (
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("my-app")
	b, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	b, err = japanese.ShiftJIS.NewDecoder().Bytes(b)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print(string(b))
}
