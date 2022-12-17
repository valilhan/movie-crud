package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"reflect"

	"github.com/gorilla/mux"
)

type Movie struct {
	Id     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var allMovie []Movie

func getAllMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(allMovie)
	if err != nil {
		log.Fatalln("There was an error in encoding ")
	}
}

func getByIdMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var movieById *Movie
	for i := 0; i < len(allMovie); i += 1 {
		if allMovie[i].Id == id {
			movieById = &allMovie[i]
			break
		}
	}
	if movieById != nil {
		err := json.NewEncoder(w).Encode(movieById)
		if err != nil {
			log.Fatalln("There was an error in encoding ")
		}
	}
}

func postMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie *Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	if movie != nil {
		allMovie = append(allMovie, *movie)
		err := json.NewEncoder(w).Encode(movie)
		if err != nil {
			log.Fatalln("There was an error in encoding ")
		}
	}
}

func putByIdMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var movie *Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	var movieById *Movie
	for i := 0; i < len(allMovie); i += 1 {
		if allMovie[i].Id == id {
			allMovie[i] = *movie
			movieById = movie
			break
		}
	}
	if movieById != nil {
		err := json.NewEncoder(w).Encode(movieById)
		if err != nil {
			log.Fatalln("There was an error in encoding ")
		}
	}
}

func deleteById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	id := vars["id"]
	var movieById *Movie
	for i := 0; i < len(allMovie); i += 1 {
		if allMovie[i].Id == id {
			movieById = &allMovie[i]
			allMovie = append(allMovie[:i], allMovie[i+1:]...)
			break
		}
	}
	if movieById != nil {
		err := json.NewEncoder(w).Encode(movieById)
		if err != nil {
			log.Fatalln("There was an error in encoding ")
		}
	}

}

func main() {
	var r *mux.Router = mux.NewRouter()
	fmt.Println(reflect.TypeOf(r))

	first := r.HandleFunc("/movies", getAllMovie)
	first.Methods("GET")
	second := r.HandleFunc("/movies/{id}", getByIdMovie)
	second.Methods("GET")

	third := r.HandleFunc("/movies", postMovie)
	third.Methods("POST")

	fourth := r.HandleFunc("/movies/{id}", putByIdMovie)
	fourth.Methods("PUT")

	fifth := r.HandleFunc("/movies/{id}", deleteById)
	fifth.Methods("DELETE")

	fmt.Printf("Server starting")
	log.Fatal(http.ListenAndServe(":8000", r))

}
