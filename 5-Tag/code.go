package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Address string `json:"address,omitempty"`
}

func main() {
	p1 := Person{
		Name: "Jack",
		Age:  16,
	}

	p2 := Person{
		Name:    "yxv",
		Age:     18,
		Address: "China",
	}

	data1, err := json.Marshal(p1)
	data2, err := json.Marshal(p2)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", data1)
	fmt.Printf("%s", data2)
}
