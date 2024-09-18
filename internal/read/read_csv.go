package read

import (
	"encoding/csv"
	"log"
	"os"
	"popsa_tech_test/internal/model"
	"strings"
	"time"
)

var (
	dir = "./data/"
)

// ReadCSV returns a slice of all the albums
func ReadCSV(c chan []model.RawAlbumData) {

	defer close(c)
	//Find the names of all albums in /data
	fileNames, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, fileName := range fileNames {
		log.Println("got filenames ", fileName)
		file, err := os.Open(dir + fileName.Name())
		if err != nil {
			log.Println("ERROR: failed to open files, err: ", err)
			continue
		}

		defer file.Close()

		reader := csv.NewReader(file)
		data, err := reader.ReadAll()
		if err != nil {
			log.Println("ERROR: failed to read files, err: ", err)
			continue
		}

		// index
		// 0 = time photo taken
		// 1 = latitude
		// 2 = longitude
		album := make([]model.RawAlbumData, len(data))
		for index, metaData := range data {
			timeTaken, err := formatTime(metaData[0])
			if err != nil {
				//TODO
				// do we want to store incorrect meta data?
				log.Printf("ERROR: failed to format time: %s, with err: %v", timeTaken, err)
				continue
			}
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
}

func formatTime(t string) (time.Time, error) {
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
		return time.Time{}, err
	}

	return timeTaken, nil
}
