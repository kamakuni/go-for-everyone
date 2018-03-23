package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type target struct {
	Name      string `json:"name"`
	Threshold int    `json:"threshold"`
}

type config struct {
	Addr   string   `json:"addr"`
	Target []target `json:"target"`
}

func main() {
	/*t1 := &target{
		Name:      "foo",
		Threshold: 3,
	}
	t2 := &target{
		Name:      "bar",
		Threshold: 4,
	}*/
	cfg := config{
		Addr: ":8080",
		Target: []target{
			target{
				Name:      "foo",
				Threshold: 3,
			},
			target{
				Name:      "bar",
				Threshold: 4,
			},
		},
	}
	b, err := json.Marshal(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
	b, err = json.MarshalIndent(&cfg, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(b))
}
