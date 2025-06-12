package jsonHelper

import (
	"encoding/json"
	"fmt"
	"os"
)

// Todo: 

type genre struct{
	Id int `json:"id"`
	NameG string `json:"name"`
}
type genreTag struct{
	G []genre `json:"genres"`
}

func AddToJson (body []byte)[]genre{
	file, err := os.Open("data/genres.json")
	var empty []genre 
	if err!= nil{
		fmt.Println("Error writing to file: ", err)
		return empty
	}
	defer file.Close()
	jsonData:= body
	if err!= nil{
		fmt.Println("Error writing to file: ", err)
		return empty
	}

	if err:= os.WriteFile("data/genres.json", jsonData, 0644);err!= nil{
		fmt.Println("Error writing to file: ", err)
	}


	var Genres genreTag
	
	error := json.Unmarshal(jsonData, &Genres)
	if error!=nil{
		fmt.Println("Unable to decode JSON: ", error)
	}
	



	for _, values := range(Genres.G){
		fmt.Printf("ID: %d, Name:%s ", values.Id, values.NameG)
	}

	return Genres.G


}