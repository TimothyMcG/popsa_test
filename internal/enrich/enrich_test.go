package enrich

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"popsa_tech_test/internal/model"
	"testing"
	"time"
)

var (
	expected = `{"items":[{"title":"Konduga, Borno, Nigeria","id":"here:cm:namedplace:23700684","resultType":"locality","localityType":"city","address":{"label":"Konduga, Borno, Nigeria","countryCode":"NGA","countryName":"Nigeria","state":"Borno","county":"Konduga","city":"Konduga"},"position":{"lat":11.65331,"lng":13.41104},"distance":0,"mapView":{"west":12.61552,"south":11.1925,"east":13.75925,"north":12.19968}}]}`
)

func TestEnrichAlbumData(t *testing.T) {

	svr := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expected)
	}))
	defer svr.Close()
	url = svr.URL

	testTime := time.Now()

	data := []model.RawAlbumData{
		{FileName: "test", Taken: testTime, Lat: "11.65331", Long: "13.41104"},
	}
	actual := EnrichAlbumMetaData(data)

	expectedCountry := "Nigeria"
	expectedCity := "Konduga"
	if actual.Country != expectedCountry {
		t.Errorf("expected country: %s got country: %s", expectedCountry, actual.Country)
	}

	if len(actual.Cities) != 1 {
		t.Errorf("expected one city. expected: %s", expectedCity)
	}

	if actual.CityKeys[0] != expectedCity {
		t.Errorf("expected city: %s got city %s", expectedCity, actual.CityKeys[0])
	}

	if actual.FirstPic != testTime || actual.LastPic != testTime {
		t.Errorf("expected test time: %v. got start:%v, end:%v", testTime, actual.FirstPic, actual.LastPic)
	}

	if actual.FileName != data[0].FileName {
		t.Errorf("expected filename: %s, got: %s", data[0].FileName, actual.FileName)
	}
}
