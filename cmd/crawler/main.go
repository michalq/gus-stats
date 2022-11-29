package main

import (
	"context"
	"log"

	"github.com/michalq/gus-stats/internal/cli"
	"github.com/michalq/gus-stats/internal/config"
	"github.com/michalq/gus-stats/internal/domain/crawler"
	gusClient "github.com/michalq/gus-stats/internal/domain/gus"
	"github.com/michalq/gus-stats/internal/domain/subject"
	"github.com/michalq/gus-stats/internal/domain/variable"

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

	subjectsFinder := subject.NewFinder(client.SubjectsApi)
	subjectDownloader := tree.NewWalker[*subject.Subject](
		subjectsFinder,
		gusClient.RegisteredApiLimits(true),
	)
	variableFinder := variable.NewFinder(
		client.VariablesApi,
		gusClient.RegisteredApiLimits(true),
	)

	dataCrawler := crawler.NewCrawler(subjectDownloader, variableFinder)

	flag.Parse()
	var err error
	// TODO add cli params, invoke specific method by cli arg
	// TODO add s3 repository and save the result?
	switch *resource {
	case "subjects":
		err = cli.SubjectsHandler(ctx, dataCrawler, subjectsFinder)
	case "variables":
		err = cli.VariablesHandler(ctx, dataCrawler)
	default:
		log.Fatalf("Resource not found %s, possible resources: %v", *resource, resources)
	}
	if err != nil {
		panic(err)
	}
}
