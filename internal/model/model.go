package model

import (
	"time"
)

type AlbumMetaData struct {
	FileName string
	FirstPic time.Time
	LastPic  time.Time
	Cities   map[string]CityData
	CityKeys []string
	Country  string
}

// RawAlbum data pulled from CSVs
type RawAlbumData struct {
	FileName string
	Taken    time.Time
	Lat      string
	Long     string
}

type CityData struct {
	Start   time.Time
	End     time.Time
	Weather string
}
