package node

// Node interface
type Node interface {
	Attributes() map[string]interface{}
	Walk(v Visitor)
}
