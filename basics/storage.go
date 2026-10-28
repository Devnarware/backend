package main


import "fmt"

type Student struct{
	Name string
	Roll int
}

func storage() {
	student1 := Student{
		Name: "dev",
		Roll: 1234,
	}

	fmt.Println(student1)	
}