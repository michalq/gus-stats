package cli

import (
	"context"
	"encoding/json"
	"os"

	"github.com/michalq/gus-stats/internal/crawler"
)

func SubjectsHandler(ctx context.Context, dataCrawler *crawler.Crawler) error {
	tree, err := dataCrawler.DownloadSubjectsTree(ctx)
	if err != nil {
		return err
	}
	subjectsJson, err := json.Marshal(tree)
	if err != nil {
		return err
	}
	err = os.WriteFile("data/subjects.json", subjectsJson, 0644)
	return err
}
