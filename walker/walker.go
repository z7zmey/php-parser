// Package walker declares walking behavior
package walker

// Walkable interface
//
// Every node must implement this interface
type Walkable interface {
	Walk(v Visitor)
}

// Visitor interface
type Visitor interface {
	EnterNode(w Walkable) bool             // EnterNode is invoked for each node encountered by Walk.
	GetChildrenVisitor(Key string) Visitor // GetChildrenVisitor is invoked at every node parameter that contains children nodes
	LeaveNode(w Walkable)                  // LeaveNode is invoked after node processed
}
