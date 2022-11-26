package tree

import "context"

type HandleErrorDecision string

const (
	HandleErrorRequeue HandleErrorDecision = "requeue"
	HandleErrorIgnore  HandleErrorDecision = "ignore"
)

// NodeDiscover implement this interface to be able to get a tree.
type NodeDiscover[T BranchValueInterface] interface {
	// FindRoot should return only one value of root branch.
	// In case there is more roots, you can return one dummy root, that as a children will store them.
	FindRoot() BranchInterface[T]

	// FindChildren should return children found by parent.
	FindChildren(ctx context.Context, parent BranchInterface[T]) ([]BranchInterface[T], error)

	// HandleError is responsible for handling error while processing children of parent
	HandleError(parent BranchInterface[T], err error) HandleErrorDecision
}
