package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"
	"popsa_tech_test/internal/model"
	"time"
)

var (
	// stored as a config/env variable?
	APIKEY = "DnyeNV12aWOOPYbDZz25EtrLnyDbTNuz6Vzyz4SOTx8"

	weatherMap = map[int]string{
		1:  "rain",
		2:  "cold",
		3:  "snow",
		4:  "sun",
		5:  "sun",
		11: "cold",
		12: "snow",
	}
)

func EnrichAlbumMetaData(album []model.RawAlbumData) model.AlbumMetaData {

	var metaData model.AlbumMetaData
	metaData.FirstPic = time.Now()
	metaData.FileName = album[0].FileName

	// store countries and cities if unique
	var countries []string
	var cities []string
	seen := make(map[string]bool)

	for _, v := range album {

		// Find Country and City info
		country, city := reverseGeocode(v.Lat, v.Long)
		if country == "" && city == "" {
			//TODO
			//print error
			continue
		}

		if !seen[country] {
			seen[country] = true
			countries = append(countries, country)
		}

		if !seen[city] {
			seen[city] = true
			cities = append(cities, city)
		}

		// Time of first and last picture
		if v.Taken.After(metaData.LastPic) {
			metaData.LastPic = v.Taken
		}

		if v.Taken.Before(metaData.FirstPic) {
			metaData.FirstPic = v.Taken
		}
	}

	// generate weather "report" we can imagine this
	// would be a call to an external api that returns
	// weather data
	weather := weatherData(metaData.FirstPic)
	if weather != "" {
		metaData.Weather = weather
	}

	// Save cities and countries
	metaData.Cities = cities
	metaData.Countries = countries
	return metaData
}

func reverseGeocode(lat, long string) (string, string) {

	latlong := lat + "," + long
	url := "https://revgeocode.search.hereapi.com/v1/revgeocode?at=" + latlong + "&lang=en-US&apiKey=" + APIKEY

	reqeust, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//TODO
		//Handle Error
		fmt.Println("err creating GET req: ", err)
	}

	res, err := http.DefaultClient.Do(reqeust)
	if err != nil {
		//TODO
		//Handle error
		fmt.Println("err with call to google API: ", err)
	}

	var data model.GeoLocate
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		//TODO
		//Handle error
		fmt.Println("couldnt decode response: ", err)
	}

	if len(data.Items) == 0 {
		return "", ""
	}
	// Decoding the city
	return data.Items[0].Address.CountryName, data.Items[0].Address.City
}

func weatherData(start time.Time) string {
	w, found := weatherMap[int(start.Month())]
	if !found {
		return ""
	}

	return w
}
