package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string `json:"Juno"`
}

func main() {
	p := Person{
		Name: "Juno",
	}

	data, err := json.Marshal(p)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(data))
}
