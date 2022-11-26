package main

import (
	"context"
	"log"

	"github.com/michalq/gus-stats/internal/config"
	"github.com/michalq/gus-stats/internal/crawler"
	gusClient "github.com/michalq/gus-stats/internal/gus"
	"github.com/michalq/gus-stats/internal/subject"
	"github.com/michalq/gus-stats/internal/variable"

	"flag"

	"github.com/michalq/gus-stats/pkg/tree"
)

var (
	resource *string

	resources = []string{"subjects", "variables"}
)

func init() {
	resource = flag.String("resource", "", "resource to crawl [subjects, variables]")
}

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

	flag.Parse()
	// TODO add cli params, invoke specific method by cli arg
	// TODO add s3 repository and save the result?
	switch *resource {
	case "subjects":
		_, err := dataCrawler.DownloadSubjectsTree(ctx)
		if err != nil {
			panic(err)
		}
	case "variables":
		_, err := dataCrawler.DownloadVariables(ctx)
		if err != nil {
			panic(err)
		}
	default:
		log.Fatalf("Resource not found %s, possible resources: %v", *resource, resources)
	}

}
