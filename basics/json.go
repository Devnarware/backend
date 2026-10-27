package main

import (
	"encoding/json"
	"fmt"
)

func dataHandling() {

	type Person struct {
		Name   string
		Age    int
		Gender string
	}

	person1 := Person{
		Name: "Devendra",
		Age: 20,
		Gender: "Male",
	}

	data, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error: ", err)
	}

	fmt.Println(string(data))

}
