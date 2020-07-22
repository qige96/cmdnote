package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path"

	"github.com/blevesearch/bleve"
)

// build index for a note, or otherwise it cannot
// be found during search
func IndexNote(noteTitle string) {
	index, err := bleve.Open(BlevePath)
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	notePath := path.Join(CONF.LocalRepoDir, noteTitle)
	data, err := ioutil.ReadFile(notePath)
	if err != nil {
		log.Fatal(err)
	}

	index.Index(noteTitle, noteTitle+" "+string(data))
}

func DeleteIndex(noteTitle string) {
	index, err := bleve.Open(BlevePath)
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	// notePath := path.Join(CONF.LocalRepoDir, noteTitle)
	err = index.Delete(noteTitle)
	if err != nil {
		log.Fatalf("When deleting index for %s: %s", noteTitle, err.Error())
	}
}

func noteTitlesBySearch(keywords string) []string {
	index, err := bleve.Open(BlevePath)
	if err != nil {
		log.Fatal(err)
	}
	defer index.Close()

	query := bleve.NewQueryStringQuery(keywords)
	searchRequest := bleve.NewSearchRequest(query)
	searchResult, _ := index.Search(searchRequest)

	noteTitles := []string{}
	for _, doc := range searchResult.Hits {
		noteTitles = append(noteTitles, doc.ID)
	}
	return noteTitles
}

// (full text) search notes by keywords
func searchNotes(keywords string) {
	noteTitles := noteTitlesBySearch(keywords)
	for _, title := range noteTitles {
		fmt.Println(title)
	}
}

// (full text) search notes by keywords with interactive inspection
func searchNotesInteractive(keywords string) {

	noteTitles := noteTitlesBySearch(keywords)
	for i, title := range noteTitles {
		fmt.Printf("%5d) %s\n", i, title)
	}

	interactiveSession(noteTitles)
}
