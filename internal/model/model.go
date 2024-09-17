package model

import "time"

type AlbumMetaData struct {
	FileName  string
	FirstPic  time.Time
	LastPic   time.Time
	Cities    []string
	Countries []string
	Weather   string
}

// RawAlbum data pulled from CSVs
type RawAlbumData struct {
	FileName string
	Taken    time.Time
	Lat      string
	Long     string
}

type GeoLocate struct {
	Items []GeoLocateData `json:"items"`
}

type GeoLocateData struct {
	Address Address `josn:"address"`
}

type Address struct {
	City        string `json:"city"`
	CountryName string `json:"countryName"`
}
