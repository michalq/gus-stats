package cli

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/michalq/gus-stats/internal/crawler"
	"github.com/michalq/gus-stats/internal/subject"
)

func SubjectsHandler(ctx context.Context, dataCrawler *crawler.Crawler, subjectFinder *subject.Finder) error {
	tree, err := dataCrawler.DownloadSubjectsTree(ctx)
	if err != nil {
		return err
	}
	subjectsJson, err := json.Marshal(tree)
	if err != nil {
		return err
	}
	log.Printf("Max elmnts %d", subjectFinder.MaxElementsPerPage())
	if subjectFinder.MaxElementsPerPage() > 100 {
		log.Println("NEEDED PAGINATION, NOT ALL ELEMENTS ARE DOWNLOADED!")
	}
	return os.WriteFile("data/subjects.json", subjectsJson, 0644)
}
