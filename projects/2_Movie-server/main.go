package main

import(
	"fmt"
	"log"
	"net/http"
)

func main(){
	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /movies", getMovies)
	mux.HandleFunc("GET /movies/{id}", getMovie)
	mux.HandleFunc("POST /movies", createMovie)
	mux.HandleFunc("PUT /movies", updateMovie)
	mux.HandleFunc("DELETE /movies", deleteMovie)

	fmt.Println("Server is starting at the port 8000")

	if err := http.ListenAndServe(":8000", mux); err != nil{
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request){
	fileserver := http.FileServer("./statc")
}


func getMovies(w http.ResponseWriter, r *http.Request){

}


func getMovie(w http.ResponseWriter, r *http.Request){

}


func createMovie(w http.ResponseWriter, r *http.Request){

}


func updateMovie(w http.ResponseWriter, r *http.Request){

}


func deleteMovie(w http.ResponseWriter, r *http.Request){

}