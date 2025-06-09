package main

// test file for making sure all these functions work

func main(){
	// slice of genres interested in as first query, is a string if any is selected
	genres := GenreMovies()
	// fucntion to find movies based off of director, returned 0 if no specfic actor or director or producer, etc is selected
	person := PersonSpec()
	//query for new(past 10 years), older(30 years), or very old (50+), string
	newOld := MovieDate()

	// generate slice of movies from database, 5 of the top movies from the time and 5 not as well known(revenue)
	movieRecSlice := Movies()



}
