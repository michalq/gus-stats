package subject

import (
	"context"
	"fmt"

	gus "github.com/michalq/gus-stats/pkg/client_gus"
	"github.com/michalq/gus-stats/pkg/tree"
)

type Finder struct {
	subjectsApi gus.SubjectsApi
}

func NewFinder(subjectsApi gus.SubjectsApi) *Finder {
	return &Finder{subjectsApi: subjectsApi}
}

func (*Finder) FindRoot() tree.Branch[*Subject] {
	return tree.NewRootSubjectBranch(&Subject{})
}

func (f *Finder) FindChildren(ctx context.Context, parent tree.Branch[*Subject]) []tree.Branch[*Subject] {
	subjectsRequest := f.subjectsApi.SubjectsGet(ctx)
	if !parent.IsRoot() {
		subjectsRequest = f.subjectsApi.SubjectsGet(ctx).ParentId(parent.Id())
	}
	subjects, _, err := subjectsRequest.Execute()
	if err != nil {
		// TODO send to err chan?
		// TODO if "429 Too Many Requests" happen then lock?
		// TODO also maybe put parent in queue again?
		fmt.Println("ERROR!!! ", err)
		parent.SetCorrupted()
		return make([]tree.Branch[*Subject], 0)
	}
	children := make([]tree.Branch[*Subject], 0, len(subjects.Results))
	for _, apiSubject := range subjects.Results {
		children = append(children, tree.NewSubjectBranch(&Subject{
			ID:        *apiSubject.Id,
			Name:      *apiSubject.Name,
			Variables: *apiSubject.HasVariables, // TODO download variable?
		}, parent, len(apiSubject.Children) > 0))
	}
	return children
}
