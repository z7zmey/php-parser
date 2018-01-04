package node

type Visitor interface {
	EnterNode(node Node) bool
	GetChildrenVisitor(Key string) Visitor
	LeaveNode(node Node)
}
