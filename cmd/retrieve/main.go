package main

import (
	"context"
	"fmt"

	gus "github.com/michalq/gus-stats/pkg/client_gus"
)

func main() {
	ctx := context.Background()
	gusConfig := gus.NewConfiguration()
	gusConfig.Debug = true
	gusConfig.Servers = gus.ServerConfigurations{
		{
			URL:         "https://bdl.stat.gov.pl/api/v1",
			Description: "No description provided",
		},
	}
	// gusConfig.DefaultHeader["X-ClientId"] = "xyz"
	gusClient := gus.NewAPIClient(gusConfig)
	aggr, _, err := gusClient.AggregatesApi.AggregatesGet(ctx).Execute()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%+v", aggr)
	}
	fmt.Println("Hello world!")
}
