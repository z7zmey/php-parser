package traverser

type Traverser interface {
	Traverse(v Visitor)
}
