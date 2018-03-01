package main

import (
	"bufio"
	"encoding/json"
	"fmt"
)

func parseEvents() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	b, _ := reader.Peek(1)
	if string(b) == "[" {
		var evs ConsulEvents
		dec := json.NewDecoder(reader)
		if err := dec.Decode(&evs); err != nil {
			return "", err
		}
		ev := &evs[len(evs)-1]
		return string(ev.Payload), nil
	} else {
		line, err := reader.ReadString("\n")
		return line, err
	}
}

func main() {

}
