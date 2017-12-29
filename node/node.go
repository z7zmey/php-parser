package node

type Node interface {
	Name() string
	Attributes() map[string]interface{}
	Walk(v Visitor)
}
