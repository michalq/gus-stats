package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/michalq/gus-stats/internal/subject"
	gus "github.com/michalq/gus-stats/pkg/client_gus"
	"github.com/michalq/gus-stats/pkg/limiter"
	"github.com/michalq/gus-stats/pkg/tree"
)

func main() {
	ctx := context.Background()
	gusConfig := gus.NewConfiguration()
	gusConfig.Debug = false
	gusConfig.Servers = gus.ServerConfigurations{
		{
			URL:         "https://bdl.stat.gov.pl/api/v1",
			Description: "No description provided",
		},
	}
	// gusConfig.DefaultHeader["X-ClientId"] = "xyz"
	gusClient := gus.NewAPIClient(gusConfig)

	subjectDownloader := tree.NewDownloader[*subject.Subject](
		subject.NewFinder(gusClient.SubjectsApi),
		baseApiLimits(true),
	)
	subjectBranches, err := subjectDownloader.FindAllNodes(ctx)
	if err != nil {
		log.Fatal(err)
	}

	root := makeTree(subjectBranches)
	res, err := json.Marshal(root)
	fmt.Printf("err: %+v\nres: %s\n", err, res)
	// json marshal -> save to file
	fmt.Println("Hello world!")
}

func makeTree(subjectBranches []tree.Branch[*subject.Subject]) *subject.Subject {
	for _, branch := range subjectBranches {
		if branch.Parent() != nil {
			parent := branch.Parent().Value()
			parent.Children = append(parent.Children, branch.Value())
		}
	}
	return subjectBranches[0].Value()
}

func printAggregates(ctx context.Context, aggregatesApi gus.AggregatesApi) {
	aggrs, _, err := aggregatesApi.AggregatesGet(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}
	for _, aggregate := range aggrs.Results {
		fmt.Printf("| %d | %s (%s) | \n", aggregate.Id, string(*aggregate.Name), string(*aggregate.Level))
	}
}

func printVariables(ctx context.Context, dataApi gus.DataApi) {
	variables, _, err := dataApi.DataByUnitGet(ctx, "unitId").Execute()
	if err != nil {
		log.Fatal(err)
	}
	for _, variable := range variables.Results {
		fmt.Printf("| %d | %+v\n", variable.Id, *variable.MeasureUnitId)
	}
}

func baseApiLimits(debug bool) limiter.Limiters {
	return limiter.NewLimiters([]limiter.Definition{
		{Limit: 5, Duration: 1 * time.Second},
		{Limit: 100, Duration: 15 * time.Minute},
		{Limit: 1000, Duration: 12 * time.Hour},
		{Limit: 10000, Duration: 7 * 24 * time.Hour},
	}, debug)
}

func registeredApiLimits(debug bool) limiter.Limiters {
	return limiter.NewLimiters([]limiter.Definition{
		{Limit: 10, Duration: 1 * time.Second},
		{Limit: 500, Duration: 15 * time.Minute},
		{Limit: 5000, Duration: 12 * time.Hour},
		{Limit: 50000, Duration: 7 * 24 * time.Hour},
	}, debug)
}
