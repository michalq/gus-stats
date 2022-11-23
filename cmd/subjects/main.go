package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/michalq/gus-stats/internal/config"
	gusClient "github.com/michalq/gus-stats/internal/gus"
	"github.com/michalq/gus-stats/internal/subject"

	"github.com/michalq/gus-stats/pkg/tree"
)

func main() {
	ctx := context.Background()
	cfg := config.LoadConfig()
	client := gusClient.NewClient(cfg)

	subjectDownloader := tree.NewDownloader[*subject.Subject](
		subject.NewFinder(client.SubjectsApi),
		gusClient.RegisteredApiLimits(true),
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
