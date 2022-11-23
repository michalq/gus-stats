package subject

import "github.com/michalq/gus-stats/pkg/tree"

type Subject struct {
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	Variables   bool       `json:"variables"`
	ChildrenQty int32      `json:"childrenQuantity"`
	Children    []*Subject `json:"children"`
}

func (s *Subject) Id() string {
	return s.ID
}

func (s *Subject) AppendChild(child tree.BranchValueInterface) {
	s.Children = append(s.Children, child.(*Subject))
}
