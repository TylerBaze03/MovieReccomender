package discoverGen

import(
	"strings"
)


func BuildLink(genres []string, person string, timePeriod string) string{
	//base api link
	url := "https://api.themoviedb.org/3/discover/movie?include_adult=false&include_video=false&language=en-US&page=1&"
	// primary_release_date.gte=1980-01-01&
	if timePeriod != "ANY"{
		pRelease := "primary_release_date.gte=" + timePeriod + "-01-01"
		url += pRelease
	}
	url+= "sort_by=popularity.asc"

	// &with_genres=TV%20Movie%7CScience%20Fciton
	if genres != nil{
		url += "&with_genres="
		for i:=0; i < len(genres); i ++{
			if genres[i] == "Science Fiction" || genres[i] == "TV Movie"{
				genres[i] = strings.ReplaceAll(genres[i], " ", "%20")
			}
			url+=genres[i]
			if i != len(genres) {
				url += "%7"
			}
		}
	}
	// &with_people=brad%20pitt"
	if person != "ANY"{
		person = strings.ReplaceAll(person, " ", "%20")
		url+= "&with_people=" + person
	}

	return url
}
