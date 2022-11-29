package subject

import (
	"context"
	"fmt"

	gus "github.com/michalq/gus-stats/pkg/client_gus"
	"github.com/michalq/gus-stats/pkg/tree"
)

const pageSize = 100 // Should be enough the biggest parent was with 70 nodes. 100 is also the maximum number for API.

type Finder struct {
	subjectsApi gus.SubjectsApi
}

func NewFinder(subjectsApi gus.SubjectsApi) *Finder {
	return &Finder{subjectsApi: subjectsApi}
}

func (*Finder) FindRoot() tree.BranchInterface[*Subject] {
	return tree.NewRootBranch(&Subject{})
}

func (f *Finder) FindChildren(ctx context.Context, parent tree.BranchInterface[*Subject]) ([]tree.BranchInterface[*Subject], error) {
	subjectsRequest := f.subjectsApi.SubjectsGet(ctx).PageSize(pageSize)
	if !parent.IsRoot() {
		subjectsRequest = f.subjectsApi.SubjectsGet(ctx).ParentId(parent.Id())
	}
	subjects, _, err := subjectsRequest.Execute()
	if err != nil {
		return make([]tree.BranchInterface[*Subject], 0), err
	}

	children := make([]tree.BranchInterface[*Subject], 0, len(subjects.Results))
	for _, apiSubject := range subjects.Results {
		children = append(children, tree.NewBranch(&Subject{
			ID:        *apiSubject.Id,
			Name:      *apiSubject.Name,
			Variables: *apiSubject.HasVariables, // TODO download variable?
		}, parent, len(apiSubject.Children) > 0))
	}

	return children, nil
}

func (f *Finder) HandleError(parent tree.BranchInterface[*Subject], err error) tree.HandleErrorDecision {
	// TODO if "429 Too Many Requests" happen then lock?
	fmt.Println("ERROR!!! ", err)

	return tree.HandleErrorIgnore
}
