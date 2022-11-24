package main

import (
	"context"

	"github.com/michalq/gus-stats/internal/config"
	"github.com/michalq/gus-stats/internal/crawler"
	gusClient "github.com/michalq/gus-stats/internal/gus"
	"github.com/michalq/gus-stats/internal/subject"
	"github.com/michalq/gus-stats/internal/variable"

	"github.com/michalq/gus-stats/pkg/tree"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()
	client := gusClient.NewClient(cfg)

	subjectDownloader := tree.NewWalker[*subject.Subject](
		subject.NewFinder(client.SubjectsApi),
		gusClient.RegisteredApiLimits(true),
	)
	variableFinder := variable.NewFinder(
		client.VariablesApi,
		gusClient.RegisteredApiLimits(true),
	)

	dataCrawler := crawler.NewCrawler(subjectDownloader, variableFinder)
	// TODO add cli params, invoke specific method by cli arg
	// TODO add s3 repository and save the result
	_, err := dataCrawler.DownloadSubjectsTree(ctx)
	if err != nil {
		panic(err)
	}
	_, err = dataCrawler.DownloadVariables(ctx)
	if err != nil {
		panic(err)
	}
}
