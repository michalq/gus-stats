package tree

// BranchValueInterface indicates what method should be implemented for branch value.
type BranchValueInterface interface {
	Id() string
	AppendChild(BranchValueInterface)
}
