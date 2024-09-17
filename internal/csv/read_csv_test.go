package csv

import (
	"popsa_tech_test/internal/model"
	"testing"
	"time"
)

func TestReadCSV(t *testing.T) {

	testCase := []struct {
		expectedLat  string
		expectedLong string
	}{
		{expectedLat: "40.000000", expectedLong: "-70.000000"},
		{expectedLat: "41.000000", expectedLong: "-71.000000"},
		{expectedLat: "42.000000", expectedLong: "-72.000000"},
	}

	dir = "../../data_test/"
	testChan := make(chan []model.RawAlbumData)
	go ReadCSV(testChan)
	actual := <-testChan

	if len(actual) != 3 {
		t.Errorf("expected 3 results got: %v", len(actual))
	}

	for i, expected := range testCase {
		if actual[i].Lat != expected.expectedLat {
			t.Errorf("expected latitude of: %v, got: %v", expected.expectedLat, actual[i].Lat)
		}
		if actual[i].Long != expected.expectedLong {
			t.Errorf("expected longitude of: %v got: %v", expected.expectedLong, actual[i].Long)
		}
	}
}

func TestFormatTime(t *testing.T) {
	expected, err := time.Parse("2006-01-02T15:04:05Z", "2019-10-31T10:50:16Z")
	if err != nil {
		t.Errorf("couldn't parse expected time: %v", err)
	}

	testCases := []string{"2019-10-31T10:50:16Z", "2019-10-31 10:50:16Z", "2019-10-31T10:50:16", "2019-10-31 10:50:16"}

	for _, test := range testCases {
		actual, err := formatTime(test)
		if err != nil {
			t.Errorf("error running formatTime, err: %v", err)
		}
		if actual != expected {
			t.Errorf("expected: %s got: %s", expected, actual)
		}
	}
}
