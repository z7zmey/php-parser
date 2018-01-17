package walker

// Walker interface
type Walker interface {
	Walk(v Visitor)
}

// Visitor interface
type Visitor interface {
	EnterNode(w Walker) bool               // EnterNode invoked for each node encountered by Walk.
	GetChildrenVisitor(Key string) Visitor // GetChildrenVisitor returns visitor for children nodes
	LeaveNode(w Walker)                    // LeaveNode invoked after process node
}
