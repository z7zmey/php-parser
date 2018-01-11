package node

// Node interface
type Node interface {
	Attributes() map[string]interface{} // Attributes returns node attributes as map
	Walk(v Visitor)                     // Walk traverses nodes
}

// Visitor interface
type Visitor interface {
	EnterNode(node Node) bool              // EnterNode invoked for each node encountered by Walk.
	GetChildrenVisitor(Key string) Visitor // GetChildrenVisitor returns visitor for children nodes
	LeaveNode(node Node)                   // LeaveNode invoked after process node
}
