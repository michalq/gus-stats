package tree

type Node interface {
	Id() string
	ParentId() string
}
