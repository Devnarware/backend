package main

import (
	"fmt"
	"log"
	"net/http"
)

func BasicServer(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("POST /post", setData)	
	mux.HandleFunc("PUT /put", updateData)
	mux.HandleFunc("DELETE /delete", deleteData)


	fmt.Println("starting the server at port 8000")

	err := http.ListenAndServe(":8000", mux)

	if err != nil {
		log.Fatal(err)
	}

}

func home(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Welcome to the server"))
}

func setData(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("Created some data"))
}

func updateData(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("update the data successfully ")) 
}

func deleteData(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("deleted successfully"))
}