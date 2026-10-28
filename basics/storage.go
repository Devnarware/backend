package main

import (
	"encoding/json"
	// "fmt"
	"log"
	"os"
)

type Student struct{
	Name string
	Roll int
}

func storage() {
	data, err := encode()

	if err != nil {
		log.Fatal(err)
	}

	os.WriteFile("storage.json", data, 0644)
}

func encode() ([]byte, error){
	student1 := Student  {
		Name: "dev",
		Roll: 1234,
	}
	student2 := Student  {
		Name: "name2",
		Roll: 1234,
	}
	student3 := Student  {
		Name: "naaam",
		Roll: 1234,
	}

	list := []Student{student1, student2, student3}
	data, err := json.MarshalIndent(list, "", "\t")

	return data, err
}