package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Tasks struct {
	Title   string
	Finished bool
}

var Tasklist []Tasks
var scanner = bufio.NewScanner(os.Stdin)

func main() {

	getData()

	start()
	fmt.Println("")

	storeData()

}

func start() {
	choice := 0

	for choice != 5 {
		fmt.Println("")
		fmt.Println("========= TODO CLI =========")
		fmt.Println("")

		// show menu
		fmt.Printf("1. Add a Task\n2. View all Task\n3. Complete a Task\n4. Delete a Task\n5. Exit\n")
		fmt.Println("")

		// choose an option
		input, err := readInput("Choose an option: ")
		choice = input
		fmt.Println("")

		if err != nil{
			fmt.Println("------------ Error: ", err, "-------------")
			return
		}else if choice < 1 || choice > 5 {
			fmt.Println("------------- Invalid input -----------")
			continue
		}

		switch choice {
		case 1:
			addTask()
		case 2:
			viewAllTask()
		case 3:
			completeTask()
		case 4:
			deleteTask()
		case 5:
			fmt.Println("Exiting the program")
		}
	}
}

func addTask() {

	newtask := Tasks{
		Title:    "",
		Finished: false,
	}

	fmt.Print("Enter the task title: ")

	if scanner.Scan() {
		newtask.Title = strings.TrimSpace(scanner.Text())
	} else {

		err := scanner.Err()
		if err != nil {
			fmt.Println("------------ Error: ", err, "-------------")
			return
		}

		fmt.Println("")
		fmt.Println("------------- EOF ------------")
		return
	}

	// 🔴 fix: here i dint check for empty input or extra white spaces

	if newtask.Title == "" {
		fmt.Println("")
		fmt.Println("------------- Task cant be empty ---------")
		return
	}

	Tasklist = append(Tasklist, newtask)
	fmt.Println("")

	viewAllTask()
}

func viewAllTask() {

	if len(Tasklist) > 0 {
		fmt.Println("All the task are below:")
		fmt.Println("")

		for i := 0; i < len(Tasklist); i++ {
			status := "❌"
			if Tasklist[i].Finished {
				status = "✅"
			}
			fmt.Printf("%d. %s  %s \n", i+1, Tasklist[i].Title, status)
		}
	} else {
		fmt.Println("No task to show!!")
	}
}

func completeTask() {
	viewAllTask()

	// completed := 0
	fmt.Println("")
	completed, err := readInput("Enter the S.No of task you completed or 0 to discard: ")
	fmt.Println("")

	if err != nil {
		fmt.Println("------------ Error: ", err, "-------------")
		return
	}

	if completed <= 0 || completed > len(Tasklist) {
		fmt.Println("No changes have been made")
		return
	}

	Tasklist[completed-1].Finished = !Tasklist[completed-1].Finished

	viewAllTask()
}

func deleteTask() {
	viewAllTask()

	// del := 0
	fmt.Println("")
	del, err := readInput("Enter the number of task you want to delete or 0 to discard: ")
	fmt.Println("")

	if err != nil {
		fmt.Println("------------ Error: ", err, "-------------")
		return
	}

	if del <= 0 || del > len(Tasklist) {
		fmt.Println("No task is removed")
		return
	}

	Tasklist = slices.Delete(Tasklist, del-1, del)

	viewAllTask()
}

// =========.  ERROR HANDLING .=========

func readInput(msg string) (int, error) {
	fmt.Print(msg)

	var input string

	if scanner.Scan() {
		input = strings.TrimSpace(scanner.Text())
	} else {

		err := scanner.Err()
		if err != nil {
			return -1, err
		}

		return -1, io.EOF
	}

	// 🔴 fix: if the error caused due to the scanner.scan -> i am still returning io.EOF

	return strconv.Atoi(input)
}

func getData(){
	// getting the data
	prevTodos, err := os.ReadFile("todo.json")

	if err != nil {
		log.Fatal(err)
	}

	err = json.Unmarshal(prevTodos, &Tasklist)
}

func storeData(){
	// storing the data
	todos, err := json.MarshalIndent(Tasklist, "", "\t")

	if err != nil {
		log.Fatal("Error: ", err)
	}

	os.WriteFile("todo.json", todos, 0644)
}

// 	1. program should only fiinish when i press 5 ✅
// 	2. fix the input issues ✅
// 	3. empty input error ✅
// 	4. data storing ✅
// 	5. editing the task, and its status
