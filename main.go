package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"Juno"`
}

func main() {
	person := Person{Name: "Juno"}
	jsonData, _ := json.Marshal(person)
	fmt.Println(string(jsonData))
}
