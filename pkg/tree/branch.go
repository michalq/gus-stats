package tree

type Branch[T any] interface {
	Id() string
	Parent() Branch[T]
	IsRoot() bool
	HasChildren() bool
	SetCorrupted()
	UnsetCorrupted()
	Value() T
}

type BranchValue interface {
	Id() string
}

type SubjectBranch[T BranchValue] struct {
	value       T
	corrupted   bool
	isRoot      bool
	hasChildren bool
	parent      Branch[T]
}

func NewSubjectBranch[T BranchValue](value T, parent Branch[T], hasChildren bool) Branch[T] {
	return &SubjectBranch[T]{
		value:       value,
		parent:      parent,
		corrupted:   false,
		isRoot:      false,
		hasChildren: hasChildren,
	}
}

// NewRootSubjectBranch creates dummy root branch, in situation where there are no specific root branch.
func NewRootSubjectBranch[T BranchValue](value T) *SubjectBranch[T] {
	return &SubjectBranch[T]{value: value, corrupted: false, isRoot: true, hasChildren: true}
}

func (s *SubjectBranch[T]) Id() string {
	return s.value.Id()
}

func (s *SubjectBranch[T]) IsRoot() bool {
	return s.isRoot
}

func (s *SubjectBranch[T]) Parent() Branch[T] {
	return s.parent
}

func (s *SubjectBranch[T]) SetCorrupted() {
	s.corrupted = true
}

func (s *SubjectBranch[T]) UnsetCorrupted() {
	s.corrupted = false
}

func (s *SubjectBranch[T]) HasChildren() bool {
	return s.hasChildren
}

func (s *SubjectBranch[T]) Value() T {
	return s.value
}
