package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)

func main() {
	u, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dir := filepath.Join(u.HomeDir, ".config", "myapp")
	err = os.MkdirAll(dir, 0775)
	if err != nil {
		log.Fatal(err)
	}
}
