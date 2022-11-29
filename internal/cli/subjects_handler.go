package cli

import (
	"context"
	"encoding/json"
	"os"

	"github.com/michalq/gus-stats/internal/domain/crawler"
	"github.com/michalq/gus-stats/internal/domain/subject"
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
	return os.WriteFile("data/subjects.json", subjectsJson, 0644)
}
