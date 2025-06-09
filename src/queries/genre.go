package queries

import (
	//"buffio"
	//"fmt"
	"io"
	"net/http"
	"MovieReccomendation/key"
	"MovieReccomendation/data"
	//"strings"
)

// genres
//todo: test for loop for getting genres or no genres, 

func GenreMovies() []byte{
	
	//anyGenre := "ANY"
	updatesGenres := getGenres()
	//fmt.Println(string(updatesGenres))
	jsonHelper.AddToJson(updatesGenres)
	/*
	chosen := false
	userGenres = new([]string)
	
	for (chosen){
		fmt.Println("Do you want to watch any genre specific movies(If any type any)?")
		input, _ = reader.ReadString('\n')
		input = strings.TrimSpace(input)
		input = strings.ToUpper(input)
		chosen = strings.Split(input, " ")
	}
	*/
	return updatesGenres
}

func getGenres() []byte{
	url := "https://api.themoviedb.org/3/genre/movie/list?language=en"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("accept", "application/json")
	tokenAuth := "Bearer " + key.Token
	req.Header.Add("Authorization", tokenAuth)

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)
	//fmt.Println(string(body))
	return body
}
