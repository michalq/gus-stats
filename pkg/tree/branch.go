package tree

type BranchInterface[T any] interface {
	Id() string
	Parent() BranchInterface[T]
	IsRoot() bool
	HasChildren() bool
	setCorrupted()
	unsetCorrupted()
	Value() T
}

type Branch[T BranchValueInterface] struct {
	value       T
	corrupted   bool
	isRoot      bool
	hasChildren bool
	parent      BranchInterface[T]
}

func NewBranch[T BranchValueInterface](value T, parent BranchInterface[T], hasChildren bool) *Branch[T] {
	return &Branch[T]{
		value:       value,
		parent:      parent,
		corrupted:   false,
		isRoot:      false,
		hasChildren: hasChildren,
	}
}

// NewRootBranch creates dummy root branch, in situation where there are no specific root branch.
func NewRootBranch[T BranchValueInterface](value T) *Branch[T] {
	return &Branch[T]{value: value, corrupted: false, isRoot: true, hasChildren: true}
}

func (s *Branch[T]) Id() string {
	return s.value.Id()
}

func (s *Branch[T]) IsRoot() bool {
	return s.isRoot
}

func (s *Branch[T]) Parent() BranchInterface[T] {
	return s.parent
}

func (s *Branch[T]) setCorrupted() {
	s.corrupted = true
}

func (s *Branch[T]) unsetCorrupted() {
	s.corrupted = false
}

func (s *Branch[T]) HasChildren() bool {
	return s.hasChildren
}

func (s *Branch[T]) Value() T {
	return s.value
}
