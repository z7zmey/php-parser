package node

type Visitor interface {
	EnterNode(node Node) bool
	GetChildrenVisitor(key string) Visitor
	LeaveNode(node Node)
}
