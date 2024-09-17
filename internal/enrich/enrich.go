package enrich

import (
	"popsa_tech_test/internal/model"
	"time"
)

var (
	// stored as a config/env variable?

	APIKEY     = "DnyeNV12aWOOPYbDZz25EtrLnyDbTNuz6Vzyz4SOTx8"
	url        = "https://revgeocode.search.hereapi.com/v1/revgeocode?at={lat&long}&lang=en-US&apiKey=" + APIKEY
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
	metaData.Cities = make(map[string]model.CityData)

	client := NewClient(url)
	for _, v := range album {

		// Find Country and City info
		country, city := client.reverseGeocode(v.Lat, v.Long)
		if country == "" && city == "" {
			//TODO
			//print error
			continue
		}

		// Save cities into map
		cityData, found := metaData.Cities[city]
		if found {
			if v.Taken.Before(cityData.Start) {
				cityData.Start = v.Taken
				cityData.Weather = weatherData(cityData.Start)
				metaData.Cities[city] = cityData
			}

			if v.Taken.After(cityData.End) {
				cityData.End = v.Taken
				metaData.Cities[city] = cityData
			}
		} else {
			metaData.Cities[city] = model.CityData{
				Start:   v.Taken,
				End:     v.Taken,
				Weather: weatherData(v.Taken),
			}
			metaData.CityKeys = append(metaData.CityKeys, city)
		}

		if metaData.Country == "" {
			metaData.Country = country
		}

		// Time of first and last picture
		if v.Taken.After(metaData.LastPic) {
			metaData.LastPic = v.Taken
		}

		if v.Taken.Before(metaData.FirstPic) {
			metaData.FirstPic = v.Taken
		}
	}

	return metaData
}

func weatherData(start time.Time) string {
	w, found := weatherMap[int(start.Month())]
	if !found {
		return ""
	}

	return w
}
