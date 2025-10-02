package main

// import db file content
import (
	"fmt"

	_ "gorm.io/driver/postgres"
)

// check movies table content
func CheckMovieTable() {
	var movies []Movie
	db.Find(&movies)

	for _, movie := range movies {
		fmt.Printf("ID: %d, ISBN: %s, Title: %s, Director: %s %s\n",
			movie.ID,
			movie.Isbn,
			movie.Title,
			movie.Director.Firstname,
			movie.Director.Lastname,
		)
	}
}

// call this function in main.go after InitDB() if you want to see the content of the movies table
