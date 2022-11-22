package tree

import (
	"context"
	"fmt"

	"github.com/michalq/gus-stats/pkg/limiter"
)

type NodeDiscover[T BranchValue] interface {
	FindRoot() Branch[T]
	FindChildren(ctx context.Context, parent Branch[T]) []Branch[T]
}

type Downloader[T BranchValue] struct {
	nodeDiscover NodeDiscover[T]
	apiLimits    limiter.Limiters
}

func NewDownloader[T BranchValue](nodeDiscover NodeDiscover[T], apiLimits limiter.Limiters) *Downloader[T] {
	return &Downloader[T]{nodeDiscover, apiLimits}
}

func (d *Downloader[T]) FindAllNodes(ctx context.Context) ([]Branch[T], error) {
	ctx, done := context.WithCancel(ctx)
	branches := make([]Branch[T], 0)
	branchesChan := make(chan Branch[T], 100)
	branchesChan <- d.nodeDiscover.FindRoot()
	i := 0
	for branch := range branchesChan {
		d.apiLimits.Wait(ctx)
		i++
		fmt.Printf("[%d] Processing children of %s\n", i, branch.Id())

		branches = append(branches, branch)
		go d.findChildren(ctx, branch, branchesChan)
		if i > 100 { // TODO only test
			done()
			return branches, nil
		} // TODO end only test
	}
	done()
	return branches, nil
}

func (d *Downloader[T]) findChildren(
	ctx context.Context,
	parent Branch[T],
	branchesChan chan<- Branch[T],
) {
	if !parent.HasChildren() {
		return
	}
	for _, child := range d.nodeDiscover.FindChildren(ctx, parent) {
		branchesChan <- child
	}
}
