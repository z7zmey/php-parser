package graph

type Traverser interface {
	traverse(v Visitor)
}
