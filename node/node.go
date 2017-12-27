package node

type Node interface {
	Name() string
	Walk(v Visitor)
}
