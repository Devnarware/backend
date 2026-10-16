package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)	

type Tasks struct {
	title    string
	finished bool
}

var taskList []Tasks

func main() {

	choise := 0
	newtask := Tasks{
		title:    "1st task",
		finished: false,
	}
	taskList = append(taskList, newtask)
	taskList = append(taskList, Tasks{
		title:    "2nd task",
		finished: true,
	})
	taskList = append(taskList, Tasks{
		title:    "3rd task",
		finished: false,
	})
	taskList = append(taskList, Tasks{
		title:    "4th task",
		finished: false,
	})
	taskList = append(taskList, Tasks{
		title:    "5th task",
		finished: true,
	})

	for choise != 5 {
		fmt.Println("")
		fmt.Println("========= TODO CLI =========")
		
		// show menu
		fmt.Printf("1. Add a Task\n2. View all Task\n3. Complete a Task\n4. Delete a Task\n5. Exit\n")
		fmt.Println("")
		
		// choose an option
		fmt.Print("Choose an option: ")
		fmt.Scan(&choise)
		fmt.Println("")

		switch choise {
		case 1:
			addTask()
			continue
		case 2:
			viewAllTask()
			continue
		case 3:
			completeTask()
			continue
		case 4:
			deleteTask()
			continue
		case 5:
			fmt.Println("chose 5")
			continue
		}
	}
	fmt.Println("")

}

func viewAllTask() {

	if len(taskList) > 0 {
		fmt.Println("All the task are below:")

		for i := 0; i < len(taskList); i++ {
			status := "❌"
			if taskList[i].finished {
				status = "✅"
			}
			fmt.Println(i+1, taskList[i].title + " ", status)
		}
	}else{
		fmt.Println("No task to show!!")
	}
}



func addTask(){

	scanner := bufio.NewScanner(os.Stdin)

	newtask := Tasks{
		title: "",
		finished: false,
	}


	fmt.Print("Enter the task tile: ")
	 if scanner.Scan() {
        newtask.title = scanner.Text()
    }

	taskList = append(taskList, newtask)
	fmt.Println("")

	viewAllTask()
}

func completeTask(){
	viewAllTask()

	del := 0
	fmt.Println("Enter the number of task you want to delete or 0 to discard: ")
	fmt.Scan(&del)

	if del == 0 || del > len(taskList) {
		fmt.Println("No task is removed")
		return
	}

	taskList = slices.Delete(taskList, del-1, del)

	viewAllTask()
}

func deleteTask(){
	viewAllTask()

	del := 0
	fmt.Println("Enter the number of task you want to delete or 0 to discard: ")
	fmt.Scan(&del)

	if del == 0 || del > len(taskList) {
		fmt.Println("No task is removed")
		return
	}

	taskList = slices.Delete(taskList, del-1, del)

	viewAllTask()
}
