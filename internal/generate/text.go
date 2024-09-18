package generate

import (
	"fmt"
	"popsa_tech_test/internal/model"
	"strings"
	"time"
)

var (
	generalTitles = []string{"A magical {time}\n", "A {time} exploring new places\n"}
)

func GenerateTitles(data model.AlbumMetaData) []string {
	var result []string
	t := getTimeContext(data.FirstPic, data.LastPic)
	general := generateGeneralTitles(t)
	loc := generateLocationTitles(data, t)
	result = append(result, general...)
	result = append(result, loc...)
	return result

}

// Create general titles with time
func generateGeneralTitles(t string) []string {
	var result []string
	for _, v := range generalTitles {
		result = append(result, strings.Replace(v, "{time}", t, -1))
	}

	return result
}

func generateLocationTitles(album model.AlbumMetaData, t string) []string {
	var titles []string
	// just visited one city
	if len(album.CityKeys) == 1 {
		city := album.CityKeys[0]
		cityData := album.Cities[city]
		if city == "New York" {
			titles = append(titles, "Taking a bite out of the big apple\n")
		}

		switch cityData.Weather {
		case "rain":
			titles = append(titles, fmt.Sprintf("Rain won't dampen the memories made in %s\n", city))
		case "sun":
			titles = append(titles, fmt.Sprintf("Under the sun in %s\n", city))
		case "cold":
			titles = append(titles, fmt.Sprintf("Shivering our way around %s\n", city))
		case "snow":
			titles = append(titles, fmt.Sprintf("Escaping to %s a winter wonderland\n", city))

		}
		titles = append(titles, fmt.Sprintf("You stole my heart %s\n", album.CityKeys[0]))
		return titles
	}

	if len(album.CityKeys) == 2 {
		titles = append(titles, fmt.Sprintf("A tale of two halves %s and %s\n", album.CityKeys[0], album.CityKeys[1]))
		return titles
	}

	buildString := ""
	for i, city := range album.CityKeys {
		if i == len(album.CityKeys)-1 {
			buildString += city + "."
			continue
		}
		buildString += city + " to "
	}
	buildString += " A " + t + " travelling through " + album.Country + "\n"
	titles = append(titles, buildString)
	return titles
}

func getTimeContext(start, end time.Time) string {
	duration := end.Sub(start)
	switch {
	case duration.Hours() <= 24:
		return "day"
	case duration.Hours() <= 72:
		if (start.Weekday().String() == "Friday" || start.Weekday().String() == "Saturday") && (end.Weekday().String() == "Sunday" || end.Weekday().String() == "Monday") {
			return "weekend"
		}
		return "couple of days"
	case duration.Hours() <= 7*24:
		return "week"
	}

	return ""
}
