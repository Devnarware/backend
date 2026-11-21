package main

import(
	"fmt"
	"log"
	"net/http"
)

type Movie struct{
	Title string `json:"title"`
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var Movies []Movie

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
	http.ServeFile(w, r, "./static/index.html")	
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