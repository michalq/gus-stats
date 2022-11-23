package tree

type Tree[T BranchValueInterface] struct {
	root T
}

func (t *Tree[T]) Root() T {
	return t.root
}
