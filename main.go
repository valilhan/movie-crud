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
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	ur := r.URL
	fmt.Println(ur)
	fmt.Println(allMovie)
}

func getByIdMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars["id"])
	id := vars["id"]
	for i := 0; i < len(allMovie); i += 1 {
		if allMovie[i].Id == id {
			fmt.Println(allMovie[i])
			break
		}
	}
}

func postMovie(w http.ResponseWriter, r *http.Request) {
	var movie Movie
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Fatalln("There was an error decoding the request body into the struct")
	}
	ur := r.URL
	fmt.Println(ur)
	allMovie = append(allMovie, movie)
	fmt.Println(movie)
}

func putByIdMovie(w http.ResponseWriter, r *http.Request) {

}

func deleteById(w http.ResponseWriter, r *http.Request) {

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
