package main

import "fmt"

type Tasks struct{
	title    string
	finished  bool
}

var taskList []Tasks


func main(){
	
	choise := 0
	newtask := Tasks{
		title: "1st task",
		finished: false,
	}
	taskList = append(taskList, newtask)
	taskList = append(taskList, Tasks{
		title: "2nd task",
		finished: true,
	})
	taskList = append(taskList, Tasks{
		title: "3rd task",
		finished: false,
	})
	taskList = append(taskList, Tasks{
		title: "4th task",
		finished: false,
	})
	taskList = append(taskList, Tasks{
		title: "5th task",
		finished: true,
	})

	fmt.Println(taskList)
	
	for choise != 5{
		fmt.Println("========= TODO CLI =========")

		// show menu
		fmt.Printf(
			"1. Add a Task\n2. View all Task\n3. Complete a Task\n4. Delete a Task\n5. Exit\n")

		// choose an option
		fmt.Print("Choose an option: ")
		fmt.Scan(&choise)

		switch choise {
		case 1:
			viewAllTask()
			continue
		case 2:
			fmt.Println("chose 2")
			continue
		case 3:
			fmt.Println("chose 3")
			continue
		case 4:
			fmt.Println("chose 4")
			continue
		case 5:
			fmt.Println("chose 5")
			continue
		}
	}

}


func viewAllTask(){
	for i := 0; i < len(taskList); i++ {
		status := "❌"
		if taskList[i].finished {
			status = "✅"
		}
		fmt.Println(taskList[i].title + " ", status)
	}
}