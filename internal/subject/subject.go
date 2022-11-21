package subject

type Subject struct {
	Id        string
	Name      string
	Variables bool
	Children  []Subject
}

type SubjectBranch struct {
	subject     *Subject
	corrupted   bool
	isRoot      bool
	hasChildren bool
	parent      Branch
}

func NewSubjectBranch(subject *Subject, parent Branch, hasChildren bool) *SubjectBranch {
	return &SubjectBranch{
		subject:     subject,
		corrupted:   false,
		isRoot:      false,
		hasChildren: hasChildren,
	}
}

// NewRootSubjectBranch creates dummy root branch, in situation where there are no specific root branch.
func NewRootSubjectBranch() *SubjectBranch {
	return &SubjectBranch{subject: &Subject{}, corrupted: false, isRoot: true, hasChildren: true}
}

func (s *SubjectBranch) Id() string {
	return s.subject.Id
}

func (s *SubjectBranch) IsRoot() bool {
	return s.isRoot
}

func (s *SubjectBranch) Parent() Branch {
	return s.parent
}

func (s *SubjectBranch) SetCorrupted() {
	s.corrupted = true
}

func (s *SubjectBranch) UnsetCorrupted() {
	s.corrupted = false
}

func (s *SubjectBranch) HasChildren() bool {
	return s.hasChildren
}

type Branch interface {
	Id() string
	Parent() Branch
	IsRoot() bool
	HasChildren() bool
	SetCorrupted()
	UnsetCorrupted()
}
