package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/michalq/gus-stats/internal/config"
	gusClient "github.com/michalq/gus-stats/internal/gus"
	"github.com/michalq/gus-stats/internal/variable"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()
	client := gusClient.NewClient(cfg)
	variableFinder := variable.NewFinder(
		client.VariablesApi,
		gusClient.RegisteredApiLimits(true),
	)
	variables, err := variableFinder.FindAll(ctx)
	if err != nil {
		panic(err)
	}
	res, err := json.Marshal(variables)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("data/variables.json", res, 0644); err != nil {
		panic(err)
	}
	fmt.Println("All done!\nSee results in data/variables.json")
}
