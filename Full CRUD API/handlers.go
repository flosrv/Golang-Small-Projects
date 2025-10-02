// handlers.go
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/gorilla/mux"
)

// Create a json of fake movie data using the gofakeit library

// GET /movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movies []Movie
	db.Find(&movies)
	json.NewEncoder(w).Encode(movies)
}

// GET /movies/{id}
func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])
	var movie Movie
	if result := db.First(&movie, id); result.Error != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(movie)
}

// POST /movies
func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	db.Create(&movie)
	json.NewEncoder(w).Encode(movie)
}

// PUT /movies/{id}
func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var movie Movie
	if result := db.First(&movie, id); result.Error != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	var updatedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
	updatedMovie.ID = movie.ID
	db.Save(&updatedMovie)
	json.NewEncoder(w).Encode(updatedMovie)
}

// DELETE /movies/{id}
func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	var movie Movie
	if result := db.First(&movie, id); result.Error != nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	db.Delete(&movie)
	fmt.Fprintf(w, "Movie deleted")
}

// CreateFakeMovieData génère un film aléatoire et le retourne en tant que Movie

func CreateFakeMovies(count int) ([]byte, error) {
	if count <= 0 {
		count = 1
	}

	gofakeit.Seed(time.Now().UnixNano())

	var movies []interface{} // slice d’interface{} pour MongoDB

	for i := 0; i < count; i++ {
		movie := Movie{
			Title: gofakeit.MovieName(),
			Isbn:  gofakeit.DigitN(8),
			Director: Director{
				Firstname: gofakeit.FirstName(),
				Lastname:  gofakeit.LastName(),
			},
		}
		movies = append(movies, movie)
	}

	// Insérer dans MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := MoviesCollection.InsertMany(ctx, movies)
	if err != nil {
		return nil, err
	}

	// Retourner en JSON les films créés
	jsonData, _ := json.Marshal(movies)
	return jsonData, nil
}
