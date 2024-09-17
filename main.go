package main

import (
	"fmt"
	"popsa_tech_test/internal/csv"
	"popsa_tech_test/internal/enrich"
	"popsa_tech_test/internal/generate"
	"popsa_tech_test/internal/model"
	"sync"
)

func main() {

	c := make(chan []model.RawAlbumData, 5)
	go csv.ReadCSV(c)

	var wg sync.WaitGroup
	for {
		album, ok := <-c
		if !ok {
			//TODO
			// Proper log here
			fmt.Println("finished reading files")
			break
		}

		wg.Add(1)
		go func(album []model.RawAlbumData, wg *sync.WaitGroup) {
			enrichedAlbum := enrich.EnrichAlbumMetaData(album)
			titles := generate.GenerateTitles(enrichedAlbum)
			fmt.Printf("Generated titles for album %s:\n %s\n", enrichedAlbum.FileName, titles)
			wg.Done()
		}(album, &wg)
	}
	wg.Wait()

}
