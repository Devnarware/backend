package main

import (
	"encoding/json"
	"fmt"
	"log"
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
	// -------------- marshling-------------
	
	data, err := json.Marshal(person1)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println(string(data))

	// --------------- unmarshling --------------

	var person2 Person

	err = json.Unmarshal(data, &person2)

	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println((person2))
}
