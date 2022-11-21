package main

import (
	"context"
	"fmt"
	"log"

	"github.com/michalq/gus-stats/internal/limiter"
	"github.com/michalq/gus-stats/internal/subject"
	gus "github.com/michalq/gus-stats/pkg/client_gus"
)

func main() {
	limiter.Xyz()
	return

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
	subjectDownloader := subject.NewDownloader(gusClient.SubjectsApi)
	branches, err := subjectDownloader.FindAllSubjects(ctx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", branches)
	// printAggregates(ctx, gusClient.AggregatesApi)
	// printSubjects(ctx, gusClient.SubjectsApi)
	// variables, _, err := gusClient.DataApi.DataByUnitGet(ctx).Execute()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// for _, variable := range variables.Results {
	// 	fmt.Printf("| %d | %+v\n", variable.Id, *variable.MeasureUnitId)
	// }
	fmt.Println("Hello world!")
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

func printSubjects(ctx context.Context, subjectsApi gus.SubjectsApi) {

}
