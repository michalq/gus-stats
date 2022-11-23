package tree

import (
	"context"
	"fmt"
	"sync"

	"github.com/michalq/gus-stats/pkg/limiter"
)

type Downloader[T BranchValueInterface] struct {
	wg           sync.WaitGroup
	nodeDiscover NodeDiscover[T]
	apiLimits    limiter.Limiters
}

func NewDownloader[T BranchValueInterface](nodeDiscover NodeDiscover[T], apiLimits limiter.Limiters) *Downloader[T] {
	return &Downloader[T]{nodeDiscover: nodeDiscover, apiLimits: apiLimits}
}

func (d *Downloader[T]) findAllNodes(ctx context.Context) ([]BranchInterface[T], error) {
	ctx, done := context.WithCancel(ctx)
	branches := make([]BranchInterface[T], 0)
	branchesChan := make(chan BranchInterface[T], 100)
	branchesChan <- d.nodeDiscover.FindRoot()
	d.wg.Add(1)
	i := 0
	go (func() {
		d.wg.Wait()
		close(branchesChan)
	})()
	for branch := range branchesChan {
		d.apiLimits.Wait(ctx)
		i++
		fmt.Printf("[%d/%d] Processing children of %s\n", i, len(branchesChan), branch.Id())

		branches = append(branches, branch)
		go d.findChildren(ctx, branch, branchesChan)
	}
	done()
	return branches, nil
}

func (d *Downloader[T]) Tree(ctx context.Context) (*Tree[T], error) {
	nodes, err := d.findAllNodes(ctx)
	if err != nil {
		return nil, err
	}
	for _, branch := range nodes {
		if branch.Parent() != nil {
			parent := branch.Parent().Value()
			parent.AppendChild(branch.Value())
		}
	}
	return &Tree[T]{root: nodes[0].Value()}, nil
}

func (d *Downloader[T]) findChildren(
	ctx context.Context,
	parent BranchInterface[T],
	branchesChan chan<- BranchInterface[T],
) {
	defer d.wg.Done()
	if !parent.HasChildren() {
		return
	}
	for _, child := range d.nodeDiscover.FindChildren(ctx, parent) {
		d.wg.Add(1)
		branchesChan <- child
	}
}
