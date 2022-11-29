package cli

import (
	"context"
	"encoding/json"
	"os"

	"github.com/michalq/gus-stats/internal/domain/crawler"
)

func VariablesHandler(ctx context.Context, dataCrawler *crawler.Crawler) error {
	variables, err := dataCrawler.DownloadVariables(ctx)
	if err != nil {
		panic(err)
	}
	variablesJson, err := json.Marshal(variables)
	if err != nil {
		return err
	}
	err = os.WriteFile("data/variables.json", variablesJson, 0644)
	return err
}
