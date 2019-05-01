package linear

type Traverser interface {
	traverse(v Visitor)
}
