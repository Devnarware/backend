package main

import (
	"fmt"
	"log"
	"net/http"
)

func newServer(){
	mux := http.NewServeMux()
	fileserve :=  http.FileServer(http.Dir("./static"))


	mux.Handle("GET /", fileserve)
	mux.HandleFunc("GET /hello", helloFunc)

	fmt.Println("Starting the server at port :8080")

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

func helloFunc(w http.ResponseWriter, r *http.Request){
	_, err := fmt.Fprintf(w, "Just printing Hello")
		// w.Write([]byte("just printing hello"))  --> both are equivalent to eact other

	if err != nil {
		log.Fatal(err)
	}


}