// main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	InitDB()

	r := mux.NewRouter()

	r.HandleFunc("/movies/all", GetMovies).Methods("GET")
	r.HandleFunc("/movies/create", CreateMovie).Methods("POST")
	r.HandleFunc("/movies/{id}/update", UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}/delete", DeleteMovie).Methods("DELETE")

	// mettre après toutes les routes statiques
	r.HandleFunc("/movies/{id}", GetMovie).Methods("GET")

	fmt.Println("Server starting on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
