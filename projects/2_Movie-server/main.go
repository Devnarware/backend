package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Movie struct {
	Title    string    `json:"title"`
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var Movies []Movie

func main() {

	Movies = append(Movies, Movie{
		Title: "Oddyssey",
		ID:    "1",
		Isbn:  "443512",
		Director: &Director{Firstname: "Christopher",
			Lastname: "Nolan",
		},
	})
	Movies = append(Movies, Movie{
		Title: "Spider Man",
		ID:    "2",
		Isbn:  "445212",
		Director: &Director{Firstname: "Jhon",
			Lastname: "Doe",
		},
	})

	mux := http.NewServeMux()

	mux.HandleFunc("GET /", home)
	mux.HandleFunc("GET /movies", getMovies)
	mux.HandleFunc("GET /movies/{id}", getMovie)
	mux.HandleFunc("POST /movies", createMovie)
	mux.HandleFunc("PUT /movies/{id}", updateMovie)
	mux.HandleFunc("DELETE /movies/{id}", deleteMovie)

	fmt.Println("Server is starting at the port 8000")

	if err := http.ListenAndServe(":8000", mux); err != nil {
		log.Fatal(err)
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "./static/index.html")
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	for idx, movie := range Movies {
		if movie.ID == id {
			Movies = append(Movies[:idx], Movies[idx+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(Movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	for _, movie := range Movies {
		if movie.ID == id {
			json.NewEncoder(w).Encode(movie)
			break
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000))
	Movies = append(Movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := r.PathValue("id")

	for idx, movie := range Movies {
		if movie.ID == id {
			Movies = append(Movies[:idx], Movies[idx+1:]...)

			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = id
			Movies = append(Movies, updatedMovie)
			json.NewEncoder(w).Encode(updatedMovie)
			return
		}
	}
}
