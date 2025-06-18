package main

import (
	//"fmt"
	//"net/http"
	//"io"
	"MovieReccomendation/queries"
	"fmt"
)

// test file for making sure all these functions work
//todo: finish testing genres

func main(){
	// slice of genres interested in as first query, is a string if any is selected
	genreIDs := queries.GenreMovies()
	// fucntion to find movies based off of director, returned 0 if no specfic actor or director or producer, etc is selected
	person := queries.PersonSpec()
	//query for new(past 10 years), older(30 years), or very old (50+), string
	newOld := queries.MovieDate()

	// generate slice of movies from database, 5 of the top movies from the time and 5 not as well known(revenue)
	//movieRecSlice := Movies()
	fmt.Println((genreIDs))
	fmt.Println((person))
	fmt.Println((newOld))



}
