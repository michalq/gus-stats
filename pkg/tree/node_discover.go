package tree

import "context"

type NodeDiscover[T BranchValue] interface {
	FindRoot() Branch[T]
	FindChildren(ctx context.Context, parent Branch[T]) []Branch[T]
}
