package queries

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"MovieReccomendation/key"
	"MovieReccomendation/data"
	"strings"
	"os"
)

// genres
//todo: adding genres by number to array not working 
// panic: runtime error: index out of range [1] with length 1

/* goroutine 1 [running]:
MovieReccomendation/queries.GenreMovies()
        C:/Users/Horizon/Documents/Github/MovieReccomend/src/queries/genre.go:50 +0x4c6
main.main()
        C:/Users/Horizon/Documents/Github/MovieReccomend/src/main.go:17 +0x13
exit status 2 */

func GenreMovies() []int{
	
	//anyGenre := "ANY"
	updatesGenres := getGenres()
	//fmt.Println(string(updatesGenres))
	genreList := jsonHelper.AddToJson(updatesGenres)
	
	//chosen := true
	var input string 
	
	//for (chosen){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Do you want to watch any genre specific movies(If any type any)?")
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
		//input = string.ToUpper(input)
		
		/*
		if input == anyGenre || input == ""{
			chosen = false
		}
		*/
	//}
	userGenres := strings.Split(input, " ")
	for i:=0; i < len(userGenres); i ++{
		userGenres[i] = capitalizeFirstLetter(userGenres[i])
	}
	userGenres = checkSciFiTV(userGenres)

	fmt.Println(userGenres)
	var genreNums []int
	count := 0
	for i := 0; i < len(genreList); i ++{
		for j :=0; j<=len(userGenres); j++{
			if genreList[i].NameG == userGenres[j]{
				//fmt.Println()
				genreNums = append(genreNums, genreList[i].Id)
				count++
			}
		}
		
		if len(userGenres) == len(genreNums){
			fmt.Println("Broken")
			break
			
		}
	}
	
	return genreNums
}
func capitalizeFirstLetter(s string) string {
	if len(s) == 0 {
		return s
	}
	return strings.ToUpper(string(s[0])) + s[1:]
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

// Science Fiction and TV Movie should be the only 2 word genres so need to check for it
func checkSciFiTV(g []string) []string{
	var ret []string = g
	
	for i := 0; i < len(g); i ++{
		if g[i] == "Science"{
			//ret[i] = "Science Fiction"
			ret = append(ret[:i], ret[i+2:]...)
			ret = append(ret, "Science Fiction")
			//ret = remove(ret, i+1)
		}
		if g[i] == "Tv"{
			//ret[i] = "TV Movie"
			ret = append(ret[:i], ret[i+2:]...)
			ret = append(ret, "TV Movie")
			//ret = remove(ret, i+1)
		}
	}


	return ret
}
func remove(s []string, i int)[]string{
	s[i] = s[len(s)-1]
    return s[:len(s)-1]
}

