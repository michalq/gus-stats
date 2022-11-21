package subject

import (
	"context"
	"fmt"
	"time"

	gus "github.com/michalq/gus-stats/pkg/client_gus"
	"golang.org/x/time/rate"
)

type Downloader struct {
	subjectsApi gus.SubjectsApi
}

func NewDownloader(subjectsApi gus.SubjectsApi) *Downloader {
	return &Downloader{subjectsApi}
}

func (d *Downloader) FindAllSubjects(ctx context.Context) ([]Branch, error) {
	// <-ctx.Done()
	branches := make([]Branch, 0)
	branchesChan := make(chan Branch, 100)
	branchesChan <- NewRootSubjectBranch()
	limiter1s := rate.NewLimiter(rate.Every(1*time.Second), 5)
	limiter15m := rate.NewLimiter(rate.Every(15*time.Minute), 100)
	limiter12h := rate.NewLimiter(rate.Every(12*time.Hour), 1000)
	limiter7d := rate.NewLimiter(rate.Every(7*24*time.Hour), 10000)
	for branch := range branchesChan {
		limiter1s.Wait(ctx)
		limiter15m.Wait(ctx)
		limiter12h.Wait(ctx)
		limiter7d.Wait(ctx)

		fmt.Println("Processing children")

		time.Sleep(2 * time.Second)
		branches = append(branches, branch)
		go d.findChildren(ctx, branch, branchesChan)
	}
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
	// TODO Stop if no children more
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
