package enrich

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Client struct {
	httpClient *http.Client
	url        string
}

func NewClient(url string) Client {
	return Client{
		httpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		url: url,
	}
}

func (c Client) reverseGeocode(lat, long string) (string, string) {

	type Address struct {
		City        string `json:"city"`
		CountryName string `json:"countryName"`
	}

	type GeoLocateData struct {
		Address Address `josn:"address"`
	}

	type GeoLocate struct {
		Items []GeoLocateData `json:"items"`
	}

	latlong := lat + "," + long
	url := strings.Replace(c.url, "{lat&long}", latlong, -1)
	reqeust, err := http.NewRequest("GET", url, nil)
	if err != nil {
		//TODO
		//Handle Error
		fmt.Println("err creating GET req: ", err)
	}

	res, err := c.httpClient.Do(reqeust)
	if err != nil {
		//TODO
		//Handle error
		fmt.Println("err with call to google API: ", err)
	}

	var data GeoLocate
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
