package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"
)

type config struct {
	APIKey   string `json:"api_key"`
	MaxCount string `json:"max_count"`
}

func loadConfig() (*config, error) {
	u, error := user.Current()
	if error != nil {
		return nil, error
	}
	fname := filepath.Join(u.HomeDir, ".config", "my-app", "config.json")
	f, err := os.Open(fname)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var cfg config
	err = json.NewDecoder(f).Decode(&cfg)
	return &cfg, err
}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(config)
}
