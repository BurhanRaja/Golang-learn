package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Director *Director `json:"director"`
	Title string `json:"title"`
	Description string `json:"description"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
} 

type ResponseMessage struct{
	Success bool `json:"success"`
	Message string `json:"message"`
}

var movies []Movie

func main() {
	
	movies = append(movies, Movie{ID: "1", Isbn: "hjkwei7i3", Title: "Mission Impossible: Dead Reconing", Description: "Ethan Hunt and his IMF team must track down a dangerous weapon before it falls into the wrong hands.", Director: &Director{Firstname: "Christopher", Lastname: "McQuarrie"}})

	movies = append(movies, Movie{ID: "2", Isbn: "qnwno83ll1m", Title: "Phir Hera Pheri", Description: "Baburao, Raju and Shyam are living happily after having risen from rags to riches. Still, money brings the joy of riches and with it the greed to make more money. And so, with a don as an unknowing investor, Raju initiates a new game.", Director: &Director{Firstname: "Neeraj", Lastname: "Vora"}})
	
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/movies", getAllMovieHandler).Methods("GET")
	r.HandleFunc("/movies", addMovieHandler).Methods("POST")
	r.HandleFunc("/movies/{id}", getMovieHandler).Methods("GET")
	r.HandleFunc("/movies/{id}", updateMovieHandler).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovieHandler).Methods("DELETE")
	fmt.Print("Server is listening to http://localhost:8000")
	http.ListenAndServe(":8000", r)
}

// Home API
func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var response ResponseMessage
	response.Success = true
	response.Message = "Welcome to the movie API!"
	json.NewEncoder(w).Encode(response)
}

// Get All Movie
func getAllMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

// Add Movie
func addMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var newMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&newMovie)
	newMovie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, newMovie)
	json.NewEncoder(w).Encode(newMovie)	
}

func getMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	for _, movie := range movies {
		if (movie.ID == id) {
			json.NewEncoder(w).Encode(movie)
			return
		}		
	}
}

func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]
	var updatedMovie Movie

	for index, movie := range movies {
		if (movie.ID == id){
			movies = append(movies[:index], movies[index+1:]...)
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			movies = append(movies, updatedMovie)
			break
		}
	}
	json.NewEncoder(w).Encode(updatedMovie)
}


func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	id := mux.Vars(r)["id"]

	for index, movie := range movies {
		if (movie.ID == id) {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	var response ResponseMessage
	response.Success = true
	response.Message = "Movie Deleted successfully.!"
	json.NewEncoder(w).Encode(response)
}