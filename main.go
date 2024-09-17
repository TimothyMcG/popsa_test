package main

import (
	"fmt"
	"popsa_tech_test/internal/csv"
	"popsa_tech_test/internal/enrich"
	"popsa_tech_test/internal/generate"
	"popsa_tech_test/internal/model"
)

func main() {

	c := make(chan []model.RawAlbumData, 5)
	go csv.ReadCSV(c)

	for {
		album, ok := <-c
		if !ok {
			//TODO
			// Proper log here
			fmt.Println("finished reading files")
			break
		}

		enrichedAlbum := enrich.EnrichAlbumMetaData(album)
		titles := generate.GenerateTitles(enrichedAlbum)
		fmt.Printf("We have generated titles for album %s:\n %s\n", enrichedAlbum.FileName, titles)

	}

}
