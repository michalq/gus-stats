package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	env "github.com/Netflix/go-env"
	"github.com/michalq/gus-stats/internal/subject"
	gus "github.com/michalq/gus-stats/pkg/client_gus"
	"github.com/michalq/gus-stats/pkg/limiter"
	"github.com/michalq/gus-stats/pkg/tree"
)

type Config struct {
	Gus struct {
		Client string `env:"GUS_CLIENT"`
	}
}

func main() {
	ctx := context.Background()
	var config Config
	if _, err := env.UnmarshalFromEnviron(&config); err != nil {
		log.Fatal(err)
	}

	gusConfig := gus.NewConfiguration()
	gusConfig.Debug = false
	gusConfig.Servers = gus.ServerConfigurations{
		{
			URL:         "https://bdl.stat.gov.pl/api/v1",
			Description: "No description provided",
		},
	}
	gusConfig.DefaultHeader["X-ClientId"] = config.Gus.Client
	gusClient := gus.NewAPIClient(gusConfig)

	subjectDownloader := tree.NewDownloader[*subject.Subject](
		subject.NewFinder(gusClient.SubjectsApi),
		registeredApiLimits(true),
	)
	subjectTree, err := subjectDownloader.Tree(ctx)
	if err != nil {
		log.Fatal(err)
	}
	root := subjectTree.Root()
	res, err := json.Marshal(root)
	fmt.Printf("err: %+v\nres: %s\n", err, res)
	// json marshal -> save to file
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
