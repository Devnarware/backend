package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"io"
)	

type Tasks struct {
	title    string
	finished bool
}

var taskList []Tasks
var scanner = bufio.NewScanner(os.Stdin)

func main() {

	// newtask := Tasks{
	// 	title:    "1st task",
	// 	finished: false,
	// }
	// taskList = append(taskList, newtask)
	// taskList = append(taskList, Tasks{
	// 	title:    "2nd task",
	// 	finished: true,
	// })
	// taskList = append(taskList, Tasks{
	// 	title:    "3rd task",
	// 	finished: false,
	// })
	// taskList = append(taskList, Tasks{
	// 	title:    "4th task",
	// 	finished: false,
	// })
	// taskList = append(taskList, Tasks{
	// 	title:    "5th task",
	// 	finished: true,
	// })

	start()
	fmt.Println("")

}

func start(){
	choise := 0

	for choise != 5 {
		fmt.Println("")
		fmt.Println("========= TODO CLI =========")
		fmt.Println("")
		
		// show menu
		fmt.Printf("1. Add a Task\n2. View all Task\n3. Complete a Task\n4. Delete a Task\n5. Exit\n")
		fmt.Println("")
		
		// choose an option
		input, err := readIntput("Choose an option: ")
		choise = input
		fmt.Println("")

		if err == io.EOF{
			break
		}

		if err != nil || choise < 1 || choise > 5{
			fmt.Println("------------- Ivalid input -----------")
			continue
		}

		switch choise {
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

func addTask(){

	newtask := Tasks{
		title: "",
		finished: false,
	}


	fmt.Print("Enter the task tile: ")
		scanner.Scan()
        newtask.title = scanner.Text()

		 // 🔴 fix: here i dint check for empty input or extra white spaces

		if newtask.title == ""{
			fmt.Println("")
			fmt.Println("------------- Task cant be empty ---------")
			return
		}

	taskList = append(taskList, newtask)
	fmt.Println("")

	viewAllTask()
}

func viewAllTask() {

	if len(taskList) > 0 {
		fmt.Println("All the task are below:")
		fmt.Println("")

		for i := 0; i < len(taskList); i++ {
			status := "❌"
			if taskList[i].finished {
				status = "✅"
			}
			fmt.Printf("%d. %s  %s \n", i+1, taskList[i].title, status)
		}
	}else{
		fmt.Println("No task to show!!")
	}
}

func completeTask(){
	viewAllTask()

	// completed := 0
	fmt.Println("")
	completed, err := readIntput("Enter the number of task you completed or 0 to discard: ")
	fmt.Println("")

	if err != nil {
		fmt.Println("------- Ivalid Input -------")
		return
	}

	if completed <= 0 || completed > len(taskList) {
		fmt.Println("No changes have been made")
		return
	}


	taskList[completed-1].finished = true

	viewAllTask()
}

func deleteTask(){
	viewAllTask()

	// del := 0
	fmt.Println("")
	del, err := readIntput("Enter the number of task you want to delete or 0 to discard: ")
	fmt.Println("")


	if err != nil {
		fmt.Println("------- Ivalid Input -------")
		return
	}

	if del <= 0 || del > len(taskList) {
		fmt.Println("No task is removed")
		return
	}

	taskList = slices.Delete(taskList, del-1, del)

	viewAllTask()
}


// =========.  ERROR HANDLING .=========

func readIntput(msg string) (int, error){
	fmt.Print(msg)

	var input string

	if scanner.Scan() {
		input = strings.TrimSpace(scanner.Text())
	}else{
		return -1, io.EOF
	}

	// 🔴 fix: if the error caused due to the scanner.scan -> i am still returning io.EOF

	return strconv.Atoi(input)
}

// 	1. program should only fiinish when i press 5 ✅
// 	2. fix the input issues ✅
// 	3. empty input error ✅
// 	4. data storing  
// 	5. editing the task, and its status