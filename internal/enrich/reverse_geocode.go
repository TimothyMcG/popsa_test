package enrich

import (
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"golang.org/x/time/rate"
)

type Client struct {
	httpClient *http.Client
	url        string
	rl         rate.Limiter
}

func NewClient(url string) *Client {
	return &Client{
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost: 20,
			},
			Timeout: 10 * time.Second,
		},
		url: url,
		rl:  *rate.NewLimiter(rate.Every(150*time.Millisecond), 1),
	}
}

func (c *Client) reverseGeocode(lat, long string) (string, string, error) {
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
		return "", "", err
	}

	ctx := context.Background()
	err = c.rl.Wait(ctx) // This is a blocking call. Honors the rate limit
	if err != nil {
		return "", "", err
	}

	res, err := c.httpClient.Do(reqeust)
	if err != nil {
		return "", "", err
	}

	defer res.Body.Close()

	var data GeoLocate
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return "", "", err
	}

	if len(data.Items) == 0 {
		return "", "", nil
	}
	// Decoding the city
	return data.Items[0].Address.CountryName, data.Items[0].Address.City, nil
}
