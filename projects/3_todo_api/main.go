package main

import (
	"fmt"
	"log"
	"net/http"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	IsCompleted bool   `json:"isCompleted"`
}

var TaskList = make([]Task, 0)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /tasks", taskList)
	mux.HandleFunc("GET /tasks/{id}", getTask)
	mux.HandleFunc("POST /tasks", createTask)
	mux.HandleFunc("PUT /tasks/{id}", updateTask)
	mux.HandleFunc("DELETE /tasks/{id}", deleteTask)

	fmt.Println("starting the server at port :5500")
	err := http.ListenAndServe(":5500", mux)

	if err != nil {
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("we are in home"))
}

func taskList(w http.ResponseWriter, r *http.Request) {
	if len(TaskList) == 0 {
		w.Write([]byte("No task to show"))
	}
}
func getTask(w http.ResponseWriter, r *http.Request) {
	msg := r.PathValue("id")
	// fmt.Fprintf("get the task with id %d", msg)
	fmt.Println(msg)
	w.Write([]byte("1 task"))
}
func createTask(w http.ResponseWriter, r *http.Request) {
	
}
func updateTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("update the task"))
}
func deleteTask(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("to see all the tasks"))
}
