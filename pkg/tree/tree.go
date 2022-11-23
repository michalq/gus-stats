package tree

type Tree[T BranchValue] struct {
	root T
}

func (t *Tree[T]) Root() T {
	return t.root
}
