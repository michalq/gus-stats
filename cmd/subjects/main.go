package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/michalq/gus-stats/internal/config"
	gusClient "github.com/michalq/gus-stats/internal/gus"
	"github.com/michalq/gus-stats/internal/subject"

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
	subjectTree, err := subjectDownloader.Walk(ctx)
	if err != nil {
		panic(err)
	}
	root := subjectTree.Root()
	res, err := json.Marshal(root)
	if err != nil {
		panic(err)
	}
	if err := os.WriteFile("data/subjects.json", res, 0644); err != nil {
		panic(err)
	}
	fmt.Println("All done!\nSee results in data/subjects.json")
}
