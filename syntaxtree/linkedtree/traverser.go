package linkedtree

type Traverser interface {
	traverse(v Visitor)
}
