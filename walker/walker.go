package walker

// Walkable interface
type Walkable interface {
	Walk(v Visitor)
}

// Visitor interface
type Visitor interface {
	EnterNode(w Walkable) bool             // EnterNode invoked for each node encountered by Walk.
	GetChildrenVisitor(Key string) Visitor // GetChildrenVisitor returns visitor for children nodes
	LeaveNode(w Walkable)                  // LeaveNode invoked after process node
}
