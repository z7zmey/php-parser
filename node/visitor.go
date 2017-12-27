package node

type Visitor interface {
	Visit(node Node) bool
	Children(key string) Visitor
	Scalar(key string, value interface{})
}
