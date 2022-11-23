package aggregate

import (
	"context"
	"fmt"
	"log"

	gus "github.com/michalq/gus-stats/pkg/client_gus"
)

func PrintAggregates(ctx context.Context, aggregatesApi gus.AggregatesApi) {
	aggrs, _, err := aggregatesApi.AggregatesGet(ctx).Execute()
	if err != nil {
		log.Fatal(err)
	}
	for _, aggregate := range aggrs.Results {
		fmt.Printf("| %d | %s (%s) | \n", aggregate.Id, string(*aggregate.Name), string(*aggregate.Level))
	}
}
