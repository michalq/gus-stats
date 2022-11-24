package crawler

import (
	"context"

	"github.com/michalq/gus-stats/internal/subject"
	"github.com/michalq/gus-stats/internal/variable"
	"github.com/michalq/gus-stats/pkg/tree"
)

type Crawler struct {
	subjectDownloader *tree.Walker[*subject.Subject]
	variableFinder    *variable.Finder
}

func NewCrawler(subjectDownloader *tree.Walker[*subject.Subject], variableFinder *variable.Finder) *Crawler {
	return &Crawler{subjectDownloader, variableFinder}
}

func (c *Crawler) DownloadSubjectsTree(ctx context.Context) (*subject.Subject, error) {
	subjectTree, err := c.subjectDownloader.Walk(ctx)
	if err != nil {
		panic(err)
	}
	return subjectTree.Root(), nil
}

func (c *Crawler) DownloadVariables(ctx context.Context) ([]*variable.Variable, error) {
	return c.variableFinder.FindAll(ctx)
}
