package generate

import (
	"fmt"
	"popsa_tech_test/internal/model"
	"strings"
	"time"
)

var (
	titleStore = map[string][]string{
		"general":  {"The best times with the best people\n", "Take me back", "Exploring New Horizons\n", "Frozen in Time\n"},
		"location": {"Remebering fun times in {location}\n", "{location} was the best\n"},
		"time":     {"A {time} away\n", "A {time} in paradise\n"},
		"New York": {"taking a bite out of the big apple\n", "seeing the city that never sleeps\n"},
	}
)

func GenerateTitles(data model.AlbumMetaData) []string {

	t := getTimeContext(data.FirstPic, data.LastPic)

	//generate different titles
	titles := generateGeneralTitles()
	titles = generateLocationSpecificTitles(data.Cities, data.Countries, titles)
	titles = generateTimeSpecificTitles(t, titles)

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

func generateGeneralTitles() []string {
	v, found := titleStore["general"]
	if !found {
		//TODO
		//log
		return []string{}
	}

	return v
}

func generateLocationSpecificTitles(cities, countries, titles []string) []string {
	fmt.Println("got cities: ", cities)
	genTitles, found := titleStore["location"]
	if !found {
		//TODO
		//log correctly
		return titles
	}

	for _, v := range genTitles {
		if len(cities) > 1 {
			v = strings.Replace(v, "{location}", countries[0], -1)
		} else {
			v = strings.Replace(v, "{location}", cities[0], -1)
		}
		titles = append(titles, v)
	}

	// check for city specific titles
	citySpecficTitles, found := titleStore[cities[0]]
	if !found {
		//TODO
		//log correctly
		fmt.Printf("no titles for city: %s\n", cities[0])
		return titles
	}

	titles = append(titles, citySpecficTitles...)
	return titles
}

func generateTimeSpecificTitles(t string, titles []string) []string {
	genTitles, found := titleStore["time"]
	if !found {
		//TODO
		// log here
		fmt.Println("n")
		return titles
	}
	for i := range genTitles {
		genTitles[i] = strings.Replace(genTitles[i], "{time}", t, -1)
	}

	titles = append(titles, genTitles...)
	return titles
}
