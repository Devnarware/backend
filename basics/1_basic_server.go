package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	http.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("welcome to the server"))
	})

	fmt.Println("Server is starting at port :8000")

	err := http.ListenAndServe(":8000", nil)

	if err != nil {
		log.Fatal(err)
	}
}