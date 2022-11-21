package subject

import (
	"context"
	"fmt"
	"time"

	"github.com/michalq/gus-stats/internal/limiter"
	gus "github.com/michalq/gus-stats/pkg/client_gus"
)

type Downloader struct {
	subjectsApi gus.SubjectsApi
}

func NewDownloader(subjectsApi gus.SubjectsApi) *Downloader {
	return &Downloader{subjectsApi}
}

func (d *Downloader) FindAllSubjects(ctx context.Context) ([]Branch, error) {
	ctx, done := context.WithCancel(ctx)
	branches := make([]Branch, 0)
	branchesChan := make(chan Branch, 100)
	branchesChan <- NewRootSubjectBranch()

	limit1s := limiter.NewLimit(5, 1*time.Second)
	limit15m := limiter.NewLimit(100, 15*time.Minute)
	limit12h := limiter.NewLimit(1000, 12*time.Hour)
	limit7d := limiter.NewLimit(10000, 7*24*time.Hour)

	for branch := range branchesChan {
		limit1s.Wait(ctx)
		limit15m.Wait(ctx)
		limit12h.Wait(ctx)
		limit7d.Wait(ctx)

		fmt.Println("Processing children")

		branches = append(branches, branch)
		go d.findChildren(ctx, branch, branchesChan)
	}
	done()
	return branches, nil
}

func (d *Downloader) findChildren(
	ctx context.Context,
	parent Branch,
	branchesChan chan Branch,
) {
	if !parent.HasChildren() {
		return
	}
	subjectsRequest := d.subjectsApi.SubjectsGet(ctx)
	if parent.IsRoot() {
		subjectsRequest = d.subjectsApi.SubjectsGet(ctx).ParentId(parent.Id())
	}
	subjects, _, err := subjectsRequest.Execute()
	if err != nil {
		// TODO send to err chan
		parent.SetCorrupted()
		return
	}
	for _, apiSubject := range subjects.Results {
		fmt.Println("Send to channel")
		branchesChan <- NewSubjectBranch(&Subject{
			Id:        *apiSubject.Id,
			Name:      *apiSubject.Name,
			Variables: *apiSubject.HasVariables, // TODO download variable?
		}, parent, len(apiSubject.Children) > 0)
	}
}
