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
	// mux.HandleFunc("GET /form", func(w http.ResponseWriter, r *http.Request){
	// 	http.ServeFile(w, r, "./static/form.html")
	// })
	mux.HandleFunc("POST /form", details)

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

func details(w http.ResponseWriter, r *http.Request){

	if err := r.ParseForm(); err != nil{
		log.Fatal(err)
	}

	fmt.Fprintf(w, "Post request successfull")
	fmt.Println("\n")

	name := r.FormValue("name")
	address := r.FormValue("address")

	fmt.Printf("Name: %s, Address: %s", name, address)
}
