package tree

import "context"

type NodeDiscover[T BranchValueInterface] interface {
	FindRoot() BranchInterface[T]
	FindChildren(ctx context.Context, parent BranchInterface[T]) []BranchInterface[T]
}
