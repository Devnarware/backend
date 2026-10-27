package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func dataHandling() {

	type Person struct {
		Name   string `json:"name"`
		Age    int `json:"age"`
		Gender string `json:"gender"`
	}

	type Employee struct {
		Name    string
		Age     int
		Gender string
		Salary  int `json:"age"`
	}

	person1 := Person{
		Name:   "Devendra",
		Age:    20,
		Gender: "Male",
	}
	// -------------- marshling-------------

	data, err := json.Marshal(person1)
	if err != nil {
		log.Fatal("Error: ", err)
	}

	fmt.Println(string(data))

	// --------------- unmarshaling --------------

	var employee Employee

	err = json.Unmarshal(data, &employee)

	if err != nil {
		log.Fatal("Error: ", err)
	}
	fmt.Println((employee))

	// we can add `json:"keyname"` -> this help to change the key name in json file
	// ---> when we marshal the data in our code we can use the key which we declare in the struct but for the json file we can use this key
	// ---> same for demarshling the json file should have the different key name but in our code we can use the key which we declare in the struct.

	// if we didnt declare the `json:keyname` then it will use the struct key name as the key for the json file to marshling the data, but for unmarshaling if the struct key dont have have the same key as json file then it will not unmarshal the data, and if we have some key in common in both, so they will be unmarshal but the rest will be ignored and set the zero value for the particular key in the struct.

	// and the go lang is pretty smart, even if you set the struct key as name Name and the json file key as name as "name", still jsnon try to match the character in upper case lower case and if something match then it will unmarshall the data

	// `json:keyname` is supirior even if the key name is salary but if we set the `json:keyname` as age then it will unmarshall the data in the struct key salary but it will take the value from the json file key age, while we have the struct key age as well
}
