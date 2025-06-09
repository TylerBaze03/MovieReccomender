package jsonHelper

import (
	"encoding/json"
	"fmt"
	"os"
)

// Todo: exporting json out of the file cannot put it into the struct find fix, possibly with marshalling

type genre struct{
	Id int `json:"id"`
	NameG string `json:"name`
}
type genreTag struct{
	G []genre `json:"genres"`
}

func AddToJson (body []byte){
	file, err := os.Open("data/genres.json")
	if err!= nil{
		fmt.Println("Error writing to file: ", err)
		return
	}
	defer file.Close()
	jsonData:= body
	if err!= nil{
		fmt.Println("Error writing to file: ", err)
		return
	}

	if err:= os.WriteFile("data/genres.json", jsonData, 0644);err!= nil{
		fmt.Println("Error writing to file: ", err)
	}


	var Genres genreTag
	
	error := json.Unmarshal(jsonData, &Genres)
	if error!=nil{
		fmt.Println("Unable to decode JSON: ", error)
	}
	

	fmt.Println(Genres.G)


	for _, values := range(Genres.G){
		fmt.Printf("ID: %d, Name:%s", values.Id, values.NameG)
	}




}