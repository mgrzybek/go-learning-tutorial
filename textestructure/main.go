package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Message struct {
	Name string
	Age  int
}

func main() {
	const family = `[
		{"Name": "Sam", "Age": 10},
		{"Name": "Joe", "Age": 13}
	]`

	// Decode
	fmt.Println(family)
	dec := json.NewDecoder(strings.NewReader(family))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v\n", m.Name, m.Age)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// Encode
	m := Message{Name: "Jack", Age: 15}
	b, err := json.Marshal(m)
	if err != nil {
		panic("json.marshal")
	}

	fmt.Printf("%v\n", string(b))
}
