package csv

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"popsa_tech_test/internal/model"
	"strings"
	"time"
)

// ReadCSV returns a slice of all the albums
func ReadCSV(c chan []model.RawAlbumData) {

	//Find the names of all albums in /data
	fileNames, err := os.ReadDir("./data/")
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range fileNames {
		fmt.Println("got filenames ", fileName)
		file, err := os.Open("./data/" + fileName.Name())
		if err != nil {
			//TODO
			//Handle error
			fmt.Println("error 1, ", err)
		}

		defer file.Close()

		reader := csv.NewReader(file)
		data, err := reader.ReadAll()
		if err != nil {
			//TODO
			//Handle error
			fmt.Println("error 2, ", err)
		}

		// index
		// 0 = time photo taken
		// 1 = latitude
		// 2 = longitude
		album := make([]model.RawAlbumData, len(data))
		for index, metaData := range data {
			timeTaken := formatTime(metaData[0])
			raw := model.RawAlbumData{
				FileName: fileName.Name(),
				Taken:    timeTaken,
				Lat:      metaData[1],
				Long:     metaData[2],
			}
			album[index] = raw
		}
		c <- album
	}

	close(c)
}

func formatTime(t string) time.Time {
	// remove any space at start or end of the time string
	t = strings.TrimRight(t, " ")
	t = strings.TrimLeft(t, " ")

	// if space instead of T add T in
	t = strings.Replace(t, " ", "T", -1)

	// append Z if not present for UTC
	if t[len(t)-1] != 'Z' {
		t = t + "Z"
	}

	timeTaken, err := time.Parse("2006-01-02T15:04:05Z", t)
	if err != nil {
		//TODO
		//Handle error
		fmt.Println("got err: ", err)
	}

	return timeTaken
}
