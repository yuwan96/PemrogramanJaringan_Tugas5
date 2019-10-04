package main

import (
	"encoding/json"
	"fmt"
)

//Person a
type Person struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

func main() {
	bytes, err := json.Marshal(Person{
		Firstname: "John",
		Lastname:  "Dow",
	})
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))
}