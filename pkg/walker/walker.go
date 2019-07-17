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
	EnterNode(w Walkable) bool
	LeaveNode(w Walkable)

	EnterChildNode(key string, w Walkable)
	LeaveChildNode(key string, w Walkable)

	EnterChildList(key string, w Walkable)
	LeaveChildList(key string, w Walkable)
}
