package node

type Visitor interface {
	EnterNode(node Node) bool
	GetChildrenVisitor(key string) Visitor
	Scalar(key string, value interface{})
	LeaveNode(node Node)
}
